package main

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"atomicgo.dev/cursor"
)

func execute_et(command []string, runkeys map[string]string) {
	if len(command) >= 1 {
		result := strings.Join(command, " ")
		if strings.Contains(result, "run") {
			arg := strings.Split(result, " ")
			for key, value := range runkeys {
				if key == "RUN_"+strings.ToUpper(arg[1]) {
					arg[1] = strings.ReplaceAll(value, "/", "\\")
					result = strings.Join(arg[1:], " ")
					break
				}
			}
		}
		output := exec.Command("cmd.exe", "/c", result)
		output.Stdout = os.Stdout
		output.Stdin = os.Stdin
		output.Stderr = os.Stderr
		cursor.Down(2)
		terminal_clock_start := time.Now()
		err := output.Run()
		if err != nil {
			terminalErrorMessage("Invalid Command. For Command Usage, Type [helpme]", "Et")
		} else {
			time_taken := time.Since(terminal_clock_start).String()
			terminalMessage("TOOK: " + time_taken)
		}
	} else {
		terminalErrorMessage("Too Low Arguments", "Et")
	}
}
