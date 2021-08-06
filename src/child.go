package main

import(
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func child(){
	fmt.Printf("Starting %v as pid %d\n",os.Args[2:],os.Getpid())
	
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	checkErr(syscall.Sethostname([]byte("myContainer")))
	checkErr(cmd.Run())
}