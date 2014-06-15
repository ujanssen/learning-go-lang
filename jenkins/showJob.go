package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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

func JenkinsJob(jenkins, jobName *string) (Job, error) {
	noJob := Job{}

	url := "http://" + *jenkins + "/job/" + *jobName + "/api/json?pretty=true"
	fmt.Printf("%s\n", url)
	contents, err := responseBody(url)
	if err != nil {
		return noJob, err
	}
	fmt.Printf("%s\n", string(contents))

	var data Job
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return noJob, err
	}
	return data, nil
}

type Build struct {
	Number int    `json:"number"`
	URL    string `json:"url"`
}

type Job struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Color     string `json:"color"`
	InQueue   bool   `json:"inQueue"`
	LastBuild Build  `json:"lastBuild"`
}

func main() {
	jenkins := flag.String("jenkins", "127.0.0.1:8080", "Jenkins hostname")
	jobName := flag.String("jobName", "test", "Jenkins job name")
	flag.Parse()

	job, err := JenkinsJob(jenkins, jobName)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Println("Name:", job.Name)
	fmt.Println("URL:", job.URL)
	fmt.Println("Color:", job.Color)
	fmt.Println("InQueue:", job.InQueue)
	fmt.Println("LastBuild number:", job.LastBuild.Number)
	fmt.Println("LastBuild URL:", job.LastBuild.URL)
}
