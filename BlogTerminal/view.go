package main

import (
	"io/ioutil"
	"strings"
)

func execute_view(command []string) {
	if len(command) == 1 {
		file, err := ioutil.ReadFile(command[0])
		if err != nil {
			error := strings.Replace(err.Error(), "open", "", 1)
			terminalErrorMessage(error, "View")
		} else {
			terminalMessage(string(file))
		}
	} else if len(command) > 1 {
		terminalErrorMessage("Too Much Arguments", "View")
	} else {
		terminalErrorMessage("Too Low Arguments", "View")
	}
}
