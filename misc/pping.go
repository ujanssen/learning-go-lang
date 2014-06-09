package main

import "fmt"
import "os/exec"

func main() {
	lsCmd := exec.Command("ping", "-c1", "localhost")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(lsOut))
}
