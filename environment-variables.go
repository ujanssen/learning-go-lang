package main

import "os"
import "fmt"

func main() {
	fmt.Println("HOME:", os.Getenv("HOME"))
}
