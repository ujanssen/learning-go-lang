package md5_test

import (
	"github.com/ujanssen/learning-go-lang/md5"
	"testing"
)

func testMd5(in, want string, t *testing.T) {
	if got := md5.Hash(in); want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFranz(t *testing.T) {
	testMd5("Franz jagt im komplett verwahrlosten Taxi quer durch Bayern", "a3cca2b2aa1e3b5b3b5aad99a8529074", t)
}

func TestFrank(t *testing.T) {
	testMd5("Frank jagt im komplett verwahrlosten Taxi quer durch Bayern", "7e716d0e702df0505fc72e2b89467910", t)
}

func TestEmpty(t *testing.T) {
	testMd5("", "d41d8cd98f00b204e9800998ecf8427e", t)
}
