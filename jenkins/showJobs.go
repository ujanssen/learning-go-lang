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

func responseBody(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func JenkinsJobs(jenkins *string) ([]JobState, error) {
	contents, err := responseBody("http://" + *jenkins + "/api/json?pretty=true")
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s\n", string(contents))

	var data Jenkins
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}
	return data.Jobs, nil
}
func main() {
	jenkins := flag.String("jenkins", "127.0.0.1:8080", "Jenkins hostname")
	flag.Parse()

	jobs, err := JenkinsJobs(jenkins)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	for _, job := range jobs {
		fmt.Printf("Name: %s, URL: %s, Color: %s\n", job.Name, job.URL, job.Color)
	}
}
