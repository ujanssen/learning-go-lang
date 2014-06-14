package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	jenkins := flag.String("jenkins", "127.0.0.1:8080", "Jenkins hostname")
	jobName := flag.String("jobName", "test", "Jenkins job name")
	flag.Parse()

	url := "http://" + *jenkins + "/job/" + *jobName + "/build"
	_, err := http.Post(url, "", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", url)
}
