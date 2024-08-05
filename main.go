package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Printf("Process ID: %v ===>>> %v\n", syscall.Getpid(), os.Args)
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Usage: go run main.go run")
	}
}

func run() {
	cmd := exec.Command(os.Args[0], append([]string{"child"}, os.Args[2:]...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func child() {
	cmd := exec.Command(os.Args[2])
	err := syscall.Sethostname([]byte("docker"))
	if err != nil {
		panic(err)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
