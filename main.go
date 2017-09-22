package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type parse interface {
	parse() error
}

// config is struct for describe running commands
type config struct {
	Tasks []command
}

// command is struct for describe single command
type command struct {
	Name    string
	Command string
	Shell   string
}

func (c *config) parse(configFile string) error {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("Read File Error: %v", err)
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		return fmt.Errorf("JSON Parse Error: %v", err)
	}
	return nil
}

func parseCmd(cmd command) []string {
	if cmd.Shell != "" {
		return []string{cmd.Shell, "-c", cmd.Command}
	}
	return strings.Fields(cmd.Command)
}

func exeCmd(cmd command, wg *sync.WaitGroup) {
	command := parseCmd(cmd)

	fmt.Printf("<%s> Run %s\n", cmd.Name, command)

	executor := command[0]
	arguments := command[1:len(command)]

	out := exec.Command(executor, arguments...)
	stdout, err := out.StdoutPipe()
	if err != nil {
		fmt.Printf("<%s> %s\n", cmd.Name, err)
	}
	stderr, err := out.StderrPipe()
	if err != nil {
		fmt.Printf("<%s> %s\n", cmd.Name, err)
	}

	if err := out.Start(); err != nil {
		fmt.Printf("<%s> %s\n", cmd.Name, err)
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	out.Wait()
	wg.Done()
}

func main() {
	var cfg config
	configPtr := flag.String("config", "../../data/runner.json", "Tasks configuration file")
	flag.Parse()

	err := cfg.parse(*configPtr)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(cfg.Tasks))

	for _, c := range cfg.Tasks {
		go exeCmd(c, wg)
	}
	wg.Wait()
}
