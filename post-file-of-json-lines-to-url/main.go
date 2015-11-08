package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

// read a file line by line and send each line to the channel
func readFile(fileName *string, jsonData chan string) {
	log.Println("Open File ", *fileName)

	file, err := os.Open(*fileName)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file)

	for {
		line, err := r.ReadString('\n')
		if err == nil {
			jsonData <- line
		} else {
			log.Println("got err: ", err.Error())
			break
		}
	}
	close(jsonData)
}

// read json data from channel and post it to a given url
func sendJSONData(url *string, jsonData chan string) {
	for s := range jsonData {
		log.Println("send line: ", s)
		reader := strings.NewReader(s)
		resp, err := http.Post(*url, "application/json", reader)
		resp.Body.Close()
		if err == nil {
			log.Printf("resp: %+v\n", resp)
		} else {
			log.Println("got err: ", err.Error())
			break
		}
	}
}

func main() {
	fileName := flag.String("file", "", "File name of a json file")
	url := flag.String("url", "", "URL to post the json data")
	flag.Parse()

	jsonData := make(chan string)

	go readFile(fileName, jsonData)

	sendJSONData(url, jsonData)
}
