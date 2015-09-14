package main

import "fmt"

type teiler struct {
	zahl, teiler int
}

func main() {
	max := 100000
	c := make(chan *teiler)
	d := make(chan bool)
	go calculate(max, c, d)
	print(max, c, d)
}

func sendTeiler(num int, c chan *teiler, d chan bool) {
	c <- &teiler{num, 1}

	for t := 2; t <= num/2; t++ {
		if num%t == 0 {
			c <- &teiler{num, t}
		}
	}

	if num > 1 {
		c <- &teiler{num, num}
	}

	d <- true
}

func calculate(max int, c chan *teiler, d chan bool) {
	worker := 8
	for z := 1; z < max; z += worker {
		for w := 0; w < worker; w++ {
			if z+w < max {
				go sendTeiler(z+w, c, d)
			}
		}
	}
}

func print(max int, c chan *teiler, d chan bool) {
	info := make([][]int, max+1)
	for i := range info {
		info[i] = make([]int, 0)
	}
	done := 0
	for {
		select {
		case t := <-c:
			info[t.zahl] = append(info[t.zahl], t.teiler)

		case _ = <-d:
			done++
			if done == max-1 {
				for i := range info {
					if i > 0 && i < max {
						fmt.Printf("%3d %6d %v\n", len(info[i]), i, info[i])
					}
				}
				return
			}
		}
	}
}
