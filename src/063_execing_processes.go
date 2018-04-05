package main

import (
	"os/exec"
	"os"
	"syscall"
)

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	areg := []string{"ls","-a","-h","-l"}

	env := os.Environ()


	if execErr := syscall.Exec(binary,areg,env); execErr != nil {
		panic(execErr)
	}
}
