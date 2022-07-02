package main

import (
	"os"
	"strings"
)

func execute_create(command []string) {
	if len(command) >= 1 {
		for i := 0; i < len(command); i++ {
			err := os.Mkdir(command[i], 0755)
			if err != nil {
				error := strings.Replace(err.Error(), "mkdir", "", 1)
				error = strings.ReplaceAll(error, "file", "folder")
				terminalErrorMessage(error, "Create")
			} else {
				terminalMessage("[" + command[i] + "] Folder Created Successfully")
			}
		}
	} else {
		terminalErrorMessage("Too Low Arguments", "Create")
	}
}
