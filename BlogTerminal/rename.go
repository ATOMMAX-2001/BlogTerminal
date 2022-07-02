package main

import "os"

func execute_rename(command []string) {
	if len(command) == 2 {
		err := os.Rename(command[0], command[1])
		if err != nil {
			terminalErrorMessage(err.Error(), "Rename")
		} else {
			terminalMessage("[" + command[0] + "] Renamed To [" + command[1] + "] Successfully")
		}
	} else if len(command) > 2 {
		terminalErrorMessage("Too Much Arguments", "Rename")
	} else {
		terminalErrorMessage("Too Low Arguments", "Rename")
	}
}
