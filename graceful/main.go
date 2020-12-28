package main

import (
	"log"
	"os"
	"os/exec"
	"net"
)

func main() {
	l, err := net.Listener
	f, _ := os.Open("/tmp/test.txt")

	netListener, _ := l
	file := netListener.File() // this returns a Dup()
	path := "/path/to/executable"
	args := []string{
		"-graceful",
	}

	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{file}

	err := cmd.Start()
	if err != nil {
		log.Fatalf("gracefulRestart: Failed to launch, error: %v", err)
	}
}