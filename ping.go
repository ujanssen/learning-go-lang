package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	ICMP_ECHO_REQUEST = 8
	ICMP_ECHO_REPLY   = 0
)

func makePingRequest(id, seq, pktlen int, filler []byte) []byte {
	p := make([]byte, pktlen)
	copy(p[8:], bytes.Repeat(filler, (pktlen-8)/len(filler)+1))

	p[0] = ICMP_ECHO_REQUEST // type
	p[1] = 0                 // code
	p[2] = 0                 // cksum
	p[3] = 0                 // cksum
	p[4] = uint8(id >> 8)    // id
	p[5] = uint8(id & 0xff)  // id
	p[6] = uint8(seq >> 8)   // sequence
	p[7] = uint8(seq & 0xff) // sequence

	// calculate icmp checksum
	cklen := len(p)
	s := uint32(0)
	for i := 0; i < (cklen - 1); i += 2 {
		s += uint32(p[i+1])<<8 | uint32(p[i])
	}
	if cklen&1 == 1 {
		s += uint32(p[cklen-1])
	}
	s = (s >> 16) + (s & 0xffff)
	s = s + (s >> 16)

	// place checksum back in header; using ^= avoids the
	// assumption the checksum bytes are zero
	p[2] ^= uint8(^s & 0xff)
	p[3] ^= uint8(^s >> 8)

	return p
}

func parsePingReply(p []byte) (id, seq int) {
	id = int(p[4])<<8 | int(p[5])
	seq = int(p[6])<<8 | int(p[7])
	return
}

func main() {
	flag.Parse()
	host := flag.Arg(0)

	ipconn, err := net.Dial("ip4:icmp", host)
	if err != nil {
		log.Fatalf("net.Dial(ip4:icmp, nil, %v) = %v", ipconn, err)
	}

	sendid := os.Getpid() & 0xffff
	sendseq := 1
	pingpktlen := 64

	for {
		sendpkt := makePingRequest(sendid, sendseq, pingpktlen, []byte("Go Ping"))

		start := time.Now()

		n, err := ipconn.Write(sendpkt)
		if err != nil || n != pingpktlen {
			log.Fatalf("net.Write(...) = %v, %v", n, err)
		}

		milliseconds := 500
		ipconn.SetReadDeadline(start.Add(time.Duration(milliseconds)))

		resp := make([]byte, 1024)
		for {
			n, err := ipconn.Read(resp)
			elapsed := time.Since(start)
			fmt.Printf("%d bytes from %s: icmp_req=%d time=%v \n", n, host, sendseq, elapsed)

			if err != nil {
				log.Fatalf("Read(...) = %v, %v", n, err)
			}
			if resp[0] != ICMP_ECHO_REPLY {
				continue
			}
			rcvid, rcvseq := parsePingReply(resp)
			if rcvid != sendid || rcvseq != sendseq {
				log.Fatalf("Ping reply saw id,seq=0x%x,0x%x (expected 0x%x, 0x%x)", rcvid, rcvseq, sendid, sendseq)
			}
			break
		}

		sendseq++
		time.Sleep(1e9) // 1 second. like -i (interval) option
	}
}
