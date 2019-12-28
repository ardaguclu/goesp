package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ardaguclu/goesp/analysis"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatalf("package path is required\n")
		return
	}

	path := os.Args[1]

	cmd := exec.Command("go", "build", "-gcflags=-m", path)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("command execution failure: %s\n", err)
	}

	a := analysis.New()
	a.Start(stderr.String())

	fmt.Println(a)
}
