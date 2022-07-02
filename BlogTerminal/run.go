package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"atomicgo.dev/cursor"
)

func execute_run(command []string, runkeys map[string]string) {
	if len(command) >= 1 {
		if command[0] == "show" {
			execute_show_run_data(runkeys)
			return
		} else {
			for key, value := range runkeys {
				if key == "RUN_"+strings.ToUpper(command[0]) {
					value = strings.ReplaceAll(value, "/", "\\")
					var output *exec.Cmd
					if len(command) > 1 {
						result := value + " " + strings.Join(command[1:], " ")
						output = exec.Command("cmd.exe", "/c", result)
						fmt.Println(value + " " + strings.Join(command[1:], " "))
					} else {
						output = exec.Command("cmd.exe", "/c", value)
					}
					output.Stdout = os.Stdout
					output.Stdin = os.Stdin
					output.Stderr = os.Stderr
					cursor.Down(2)
					output.Run()
					return
				}
			}
		}
	} else {
		terminalErrorMessage("Too Low Arguments", "Run")
		return
	}
	terminalErrorMessage("No Such RUN Path Found", "Run")
}
func execute_show_run_data(runkeys map[string]string) {
	terminalMessage("\tNAME\t\t\tPATH")
	set_default_terminal_theme("YELLOW")
	for key, value := range runkeys {
		if strings.Contains(key, "RUN_") {
			key = strings.Replace(key, "RUN_", "", 1)
			value = strings.ReplaceAll(value, "/", "\\")
			fmt.Println("\t" + key + "\t               " + value)
		}
	}
}
