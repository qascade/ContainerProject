package main

import(
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main(){
	
	if os.Args[1] == "start" {
		start()
	}else if os.Args[1]=="child"{
		child()
	}else{
		panic("cant start.")
	}

}
func start(){
	fmt.Printf("Starting %v as pid %d\n",os.Args[2:],os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"},os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//namespace system call
	//CLONE_NEWUTS for Hostname and NIS DomainName
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	fmt.Println(cmd)
	checkErr(cmd.Run())
}
func child(){
	fmt.Printf("Starting %v as pid %d\n",os.Args[2:],os.Getpid())
	
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	checkErr(syscall.Sethostname([]byte("myContainer")))
	checkErr(cmd.Run())
}
func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
