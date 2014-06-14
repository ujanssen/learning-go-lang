package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type JobState struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Color string `json:"color"`
}

type Jenkins struct {
	NodeDescription string     `json:"nodeDescription"`
	Jobs            []JobState `json:"jobs"`
}

func main() {
	jenkins := flag.String("jenkins", "127.0.0.1:8080", "Jenkins hostname")
	flag.Parse()

	response, err := http.Get("http://" + *jenkins + "/api/json?pretty=true")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))

		var data Jenkins

		err = json.Unmarshal(contents, &data)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		fmt.Printf("%v", data.Jobs)
	}
}
