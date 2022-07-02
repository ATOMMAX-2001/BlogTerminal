package main

import (
	"fmt"
	"os"
	"strings"
)

func execute_teleport(command []string, telekey map[string]string, path string) string {
	if len(command) == 1 {
		if command[0] == "show" {
			execute_show_teleport_data(telekey)
			return path
		} else {
			for key, value := range telekey {
				if key == "TELE_"+strings.ToUpper(command[0]) {
					value = strings.ReplaceAll(value, "/", "\\")
					err := os.Chdir(value)
					if err != nil {
						terminalErrorMessage(err.Error(), "Teleport")
						return path
					} else {
						path = value
						return path
					}
				}
			}
		}
	} else if len(command) > 1 {
		terminalErrorMessage("Too Much Arguments", "Teleport")
		return path
	} else {
		terminalErrorMessage("Too Low Arguments", "Teleport")
		return path
	}
	terminalErrorMessage("No Such Teleport Path Found", "Teleport")
	return path
}
func execute_show_teleport_data(telekeys map[string]string) {
	terminalMessage("\tNAME\t\t\tPATH")
	set_default_terminal_theme("YELLOW")
	for key, value := range telekeys {
		if strings.Contains(key, "TELE_") {
			key = strings.Replace(key, "TELE_", "", 1)
			value = strings.ReplaceAll(value, "/", "\\")
			fmt.Println("\t" + key + "\t               " + value)
		}
	}
}
