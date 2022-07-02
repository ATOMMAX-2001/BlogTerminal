package main

import (
	"os"
	"os/exec"
	"strings"

	"atomicgo.dev/cursor"
)

func execute_eval(command []string) {
	path := os.Args[0][0 : len(os.Args[0])-16]
	result := path + "pyeval\\eval.exe" + " " + strings.Join(command, " ")
	output := exec.Command("cmd.exe", "/c", result)
	output.Stdout = os.Stdout
	output.Stdin = os.Stdin
	output.Stderr = os.Stderr
	cursor.Down(2)
	err := output.Run()
	if err != nil {
		terminalErrorMessage("Invalid Command. For Command Usage Type [helpme] ", "Eval")
	}
}
