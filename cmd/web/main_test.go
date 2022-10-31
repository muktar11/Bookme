package main

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	_, err := run()
	if err == nil{
		fmt.Println("main package passed test successfully")
	}
	if err != nil {
		t.Error("failed run(")
	}
}

//run wtih cd cmd/web  go test -v  //go test cover to get summary


//go test -coverprofile=coverage.out && go tool cover -html=coverage.out