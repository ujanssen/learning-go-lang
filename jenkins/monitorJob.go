package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
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
	return contents, nil
}

func JenkinsJobData(jenkins, jobName *string) (Job, error) {
	url := "http://" + *jenkins + "/job/" + *jobName + "/api/json?pretty=true"
	// fmt.Printf("%s\n", url)
	contents, err := responseBody(url)
	if err != nil {
		return Job{}, err
	}
	// fmt.Printf("%s\n", string(contents))

	var data Job
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return Job{}, err
	}
	return data, nil
}

func JenkinsBuild(jenkins, jobName *string, buildNumber *int) (Build, error) {
	build := Build{}

	url := "http://" + *jenkins +
		"/job/" + *jobName + "/" +
		strconv.Itoa(*buildNumber) + "/api/json?pretty=true"

	// fmt.Printf("%s\n", url)
	contents, err := responseBody(url)
	if err != nil {
		return build, err
	}
	// fmt.Printf("%s\n", string(contents))

	err = json.Unmarshal(contents, &build)
	if err != nil {
		return build, err
	}
	return build, nil
}

type Job struct {
	Name            string    `json:"name"`
	URL             string    `json:"url"`
	Color           string    `json:"color"`
	NextBuildNumber int       `json:"nextBuildNumber"`
	InQueue         bool      `json:"inQueue"`
	LastBuild       LastBuild `json:"lastBuild"`
}
type LastBuild struct {
	Number int    `json:"number"`
	URL    string `json:"url"`
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

type JobState string

const (
	Unknown  JobState = "Unknown" // no contact to jenkins
	InQueue  JobState = "InQueue"
	Building JobState = "Building"
	Finished JobState = "Finished"
)

type JenkinsJob struct {
	Name     string   `json:"name"`
	JobState JobState `json:"state"`
}

func JenkinsJobState(jenkins, jobName *string) JenkinsJob {
	jenkinsJob := JenkinsJob{Name: *jobName}
	var buildNumber int

	// query /job while InQueue
	// set buildNumber
	job, err := JenkinsJobData(jenkins, jobName)
	if err != nil {
		fmt.Printf("%s", err)
		jenkinsJob.JobState = Unknown
		return jenkinsJob
	}
	fmt.Println("InQueue:", job.InQueue)

	if job.InQueue {
		jenkinsJob.JobState = InQueue
		return jenkinsJob
	} else {
		buildNumber = job.LastBuild.Number
	}

	// query /job/{{buildNo}} for Building
	build, err := JenkinsBuild(jenkins, jobName, &buildNumber)
	if err != nil {
		fmt.Printf("%s", err)
		jenkinsJob.JobState = Unknown
		return jenkinsJob
	}
	if build.Building {
		start := time.Unix(int64(build.Timestamp/1000), 0)
		estEnd := time.Unix(int64((build.Timestamp+build.EstimatedDuration)/1000), 0)

		buildDuration := int(time.Since(start).Seconds())
		buildCountdown := (build.EstimatedDuration / 1000) - buildDuration

		fmt.Println("Building:", build.Building)
		fmt.Println("Start:", start)
		fmt.Println("Build Duration: (s): ", buildDuration)
		fmt.Println("Build Countdown (s): ", buildCountdown)
		fmt.Println("Estimated End:", estEnd)
		fmt.Println("Estimated Duration (s): ", build.EstimatedDuration/1000)
		jenkinsJob.JobState = Building
	} else {
		// show build result
		fmt.Println("Duration:", build.Duration)
		fmt.Println("Result:", build.Result)
		jenkinsJob.JobState = Finished
	}
	return jenkinsJob
}

func main() {
	jenkins := flag.String("jenkins", "127.0.0.1:8080", "Jenkins hostname")
	jobName := flag.String("jobName", "test", "Jenkins job name")
	flag.Parse()
	oldState := JenkinsJobState(jenkins, jobName)
	fmt.Println("Time:", time.Now())
	fmt.Println("oldState:", oldState)
	for {
		time.Sleep(time.Second)
		newState := JenkinsJobState(jenkins, jobName)
		if newState.JobState != oldState.JobState || newState.JobState == Building {
			fmt.Println("Time:", time.Now())
			fmt.Println("newState:", newState)
			fmt.Printf("\n\n")
		}
		oldState = newState
	}

}
