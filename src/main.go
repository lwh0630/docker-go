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
		Run()
	case "init":
		Init()
	default:
		panic("Usage: go run main.go run")
	}
}

func Run() {
	cmd := exec.Command(os.Args[0], append([]string{"init"}, os.Args[2])...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getuid(), Size: 1},
		},
		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getgid(), Size: 1},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
	// 取消ns中的mount
	//err := syscall.Unmount("rootfs/proc", syscall.MNT_FORCE)
	//if err != nil {
	//	panic(err)
	//}
}

func Init() {
	fmt.Println(os.Getwd())
	err := syscall.Sethostname([]byte("docker"))
	if err != nil {
		panic(err)
	}
	err = syscall.Chroot("rootfs")
	err = syscall.Chdir("/")
	if err != nil {
		panic(err)
	}
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	err = syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	if err != nil {
		panic(err)
	}
	_ = syscall.Exec(os.Args[2], os.Args[2:], os.Environ())
	//if err != nil {
	//	panic(err)
	//}
	//err = syscall.Unmount("/proc", syscall.MNT_FORCE)
	//if err != nil {
	//	panic(err)
	//}
}
