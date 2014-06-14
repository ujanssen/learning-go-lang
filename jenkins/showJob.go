package main

import (
	"encoding/json"
	"errors"
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
func JenkinsJob(jenkins, jobName *string) (JobState, error) {
	noJob := JobState{}

	url := "http://" + *jenkins + "/job/" + *jobName + "/api/json?pretty=true"
	fmt.Printf("%s\n", url)
	contents, err := responseBody(url)
	if err != nil {
		return noJob, err
	}
	fmt.Printf("%s\n", string(contents))

	// var data Jenkins
	// err = json.Unmarshal(contents, &data)
	// if err != nil {
	// 	return nil, err
	// }
	// return data.Jobs, nil
	return noJob, errors.New("Job not found.")
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
	fmt.Printf("Name: %s, URL: %s, Color: %s\n", job.Name, job.URL, job.Color)
}
