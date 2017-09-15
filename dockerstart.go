package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func exeCmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func main() {
	configPtr := flag.String("config", "../data/runner.json", "Tasks configuration file")
	flag.Parse()

	fmt.Println(*configPtr)
	// wg := new(sync.WaitGroup)
	// wg.Add(3)

	// x := []string{"echo newline >> foo.o", "echo newline >> f1.o", "echo newline >> f2.o"}
	// go exeCmd(x[0], wg)
	// go exeCmd(x[1], wg)
	// go exeCmd(x[2], wg)

	// wg.Wait()
}
