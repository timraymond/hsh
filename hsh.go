package main

import (
  "github.com/timraymond/hsh/hcommand"

	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: hsh app-name")
		os.Exit(0)
	}
	app_name := os.Args[1]
	ps := app_name + " > "
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(ps)

  // REPL
	for scanner.Scan() {
    input := scanner.Text()

    // Parse out special commands
		switch input {
		case ":exit":
			os.Exit(0)
		case "":
      fmt.Print(ps)
			continue
		case "console":
			input = "run console"
		}
    cmd := hcommand.HerokuCmd(input, app_name)

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
