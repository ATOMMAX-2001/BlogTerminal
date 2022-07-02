package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func execute_drive(command []string, Tuser Terminaluser) Terminaluser {
	if len(command) == 1 {
		_, err := ioutil.ReadDir(command[0] + ":")
		if err != nil {
			terminalErrorMessage(err.Error(), "Drive")
		} else {
			Tuser.cwd = strings.ToUpper(command[0]) + ":\\"
			os.Chdir(Tuser.cwd)
			return Tuser
		}

	} else if len(command) > 1 {
		terminalErrorMessage("Too Much Arguments", "Drive")
	} else {
		terminalErrorMessage("Too Low Argumets", "Drive")
	}
	return Tuser
}
