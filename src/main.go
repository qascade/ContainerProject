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

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
