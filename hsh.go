package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

type HerokuCmd struct {
	subcommand string
	app        string
}

func (h *HerokuCmd) String() string {
	return "heroku " + h.subcommand + " -a " + h.app
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: hsh app-name")
		os.Exit(0)
	}
	app_name := os.Args[1]
	ps := app_name + " > "
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(ps)
	for scanner.Scan() {
		command := HerokuCmd{subcommand: scanner.Text(), app: app_name}
		switch command.subcommand {
		case ":exit":
			os.Exit(0)
		case "\n":
			continue
		case "console":
			command.subcommand = "run console"
		}
		cmd := exec.Command("/bin/bash", "-c", command.String())
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

    //Channel Setup for process execution
		done := make(chan error)
		sig := make(chan os.Signal, 1)
    // Setup kill notification
    signal.Notify(sig, os.Interrupt)

		go func() {
			done <- cmd.Run()
		}()
		select {
		case err := <-done:
			if err != nil {
				log.Print(err)
			}
			break
		case s := <-sig:
			err := cmd.Process.Signal(s)
      <-done // wait for the subprocess to finish
			if err != nil {
				log.Fatal(err)
			}
      break
		}
    fmt.Print(ps)
	}
}
