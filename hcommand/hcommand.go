package hcommand

import (
  "os/exec"
  "os"
)

type herokuCmd struct {
	subcommand string
	app        string
}

func (h *herokuCmd) String() string {
	return "heroku " + h.subcommand + " -a " + h.app
}

func HerokuCmd(subcommand string, app string) *exec.Cmd {
  command := herokuCmd{subcommand: subcommand, app: app}
  cmd := exec.Command("/bin/bash", "-c", command.String())

  // Rewire the new processes' iostreams to ours
  cmd.Stdout = os.Stdout
  cmd.Stdin = os.Stdin
  cmd.Stderr = os.Stderr
  return cmd
}

