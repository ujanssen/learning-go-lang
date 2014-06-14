package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
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

func JenkinsBuild(jenkins, jobName *string, buildNumber *int) (Build, error) {
	build := Build{}

	url := "http://" + *jenkins +
		"/job/" + *jobName + "/" +
		strconv.Itoa(*buildNumber) + "/api/json?pretty=true"

	fmt.Printf("%s\n", url)
	contents, err := responseBody(url)
	if err != nil {
		return build, err
	}
	fmt.Printf("%s\n", string(contents))

	err = json.Unmarshal(contents, &build)
	if err != nil {
		return build, err
	}
	return build, nil
}

type Job struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Color           string `json:"color"`
	NextBuildNumber int    `json:"nextBuildNumber"`
	InQueue         bool   `json:"inQueue"`
}

type Build struct {
	Number            int    `json:"number"`
	Duration          int    `json:"duration"`
	EstimatedDuration int    `json:"estimatedDuration"`
	Timestamp         int    `json:"timestamp"`
	URL               string `json:"url"`
	Result            string `json:"result"`
	Building          bool   `json:"building"`
}

func main() {
	jenkins := flag.String("jenkins", "127.0.0.1:8080", "Jenkins hostname")
	jobName := flag.String("jobName", "test", "Jenkins job name")
	flag.Parse()

	job, err := JenkinsJob(jenkins, jobName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("job.NextBuildNumber: %d\n", job.NextBuildNumber)
	buildNumber := job.NextBuildNumber

	url := "http://" + *jenkins + "/job/" + *jobName + "/build"
	_, err = http.Post(url, "", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", url)

	// InQueue
	for {
		time.Sleep(time.Second)
		job, err := JenkinsJob(jenkins, jobName)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Println("Time:", time.Now())
		fmt.Println("Name:", job.Name)
		fmt.Println("URL:", job.URL)
		fmt.Println("Color:", job.Color)
		fmt.Println("InQueue:", job.InQueue)

		if !job.InQueue {
			break
		}
	}
	for {
		build, err := JenkinsBuild(jenkins, jobName, &buildNumber)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Println("Number:", build.Number)
		fmt.Println("URL:", build.URL)
		fmt.Println("Building:", build.Building)
		fmt.Println("Timestamp:", build.Timestamp)
		fmt.Println("Duration:", build.Duration)
		fmt.Println("EstimatedDuration:", build.EstimatedDuration)
		fmt.Println("Result:", build.Result)
		if len(build.Result) > 0 {
			break
		}
	}
}
