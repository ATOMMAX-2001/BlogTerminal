package main

import (
	"os"
	"strings"
)

func execute_remove(command []string) {
	if len(command) >= 1 {
		for i := 0; i < len(command); i++ {
			file, err := os.Stat(command[i])
			if err != nil {
				terminalErrorMessage(err.Error(), "Remove")
			} else {
				err := os.RemoveAll(command[i])
				if err != nil {
					error := strings.Replace(err.Error(), "remove", "", 1)
					terminalErrorMessage(error, "Remove")
				}
				if file.IsDir() {
					terminalMessage("[" + command[i] + "] Folder Removed Successfully")
				} else {
					terminalMessage("[" + command[i] + "] File Removed Successfully")
				}
			}
		}

	} else {
		terminalErrorMessage("Too Low Arguments", "Remove")
	}
}
