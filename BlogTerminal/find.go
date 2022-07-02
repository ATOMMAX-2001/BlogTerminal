package main

import (
	"os"
	"os/exec"
)

func execute_find(command []string) {
	if len(command) == 1 {
		output := exec.Command("cmd.exe", "/c", "dir", command[0])
		output.Stdout = os.Stdout
		output.Run()
	} else if len(command) > 1 {
		terminalErrorMessage("Too Much Arguments", "Find")
	} else {
		terminalErrorMessage("Too Low Arguments", "Find")
	}
}
