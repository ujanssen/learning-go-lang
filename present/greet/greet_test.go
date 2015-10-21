// greet/greet_test.go
package main

import "testing"

func Test(t *testing.T) {
	if got, want := greet(), "Hello World!"; got != want {
		t.Fatal(got, want)
	}
}
