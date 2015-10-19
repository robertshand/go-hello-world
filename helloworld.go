package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {

    UnamecmdArgs := []string{"-n"}
    UnamecmdExe := "uname"
	Uname := exe_cmd( UnamecmdExe, UnamecmdArgs )

	log.Printf("%s %s\n", req.Proto, req.URL)
	fmt.Fprintln(w, "*** Go - Hello World ! ***")
	fmt.Fprintln(w, "WELCOME_MSG : ", os.Getenv("WELCOME_MSG") )
	fmt.Fprint(w, "Hostname is : ", Uname)
	fmt.Fprint(w, "Process ID  : ", os.Getpid() )
}

func exe_cmd(cmdName string, cmdArgs []string) string {
    var (
        cmdOut []byte
        err    error
    )
    if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
        fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
        os.Exit(1)
    }
    // fmt.Println("output of command is:", string(cmdOut))
    return string(cmdOut)
}
