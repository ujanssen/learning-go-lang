package main

import (
	"flag"
	"fmt"

	"github.com/bndr/gojenkins"
)

var (
	host    = flag.String("host", "jenkins", "jenkins host name")
	proto   = flag.String("proto", "https", "protocol http or https")
	user    = flag.String("user", "jenkins", "jenkins user name")
	token   = flag.String("token", "***", "jenkins user api token")
	jobName = flag.String("job", "Package-and-deploy-frontend", "jenkins job name")
)

func main() {
	flag.Parse()

	jenkins, err := gojenkins.CreateJenkins(*proto+"://"+*host+"/", *user, *token).Init()

	if err != nil {
		panic("Something Went Wrong: " + err.Error())
	}
	fmt.Printf("jenkins: %+v\n", jenkins)

	job, err := jenkins.GetJob(*jobName)
	if err != nil {
		panic("Job Does Not Exist: " + err.Error())
	}

	build, err := job.GetLastSuccessfulBuild()
	if err != nil {
		panic("Build Does Not Exist: " + err.Error())
	}

	fmt.Printf("LastSuccessfulBuild: %+v\n", build)

	build.GetBuildNumber()
	fmt.Printf("GetBuildNumber: %+v\n", build.GetBuildNumber())
}
