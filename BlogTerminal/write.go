package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func execute_write(command []string) {
	if len(command) == 1 {
		editorScanner := bufio.NewReader(os.Stdin)
		var data string = ""
		terminalMessage("Type The Content for " + command[0] + ". Type [texit] in newline to exit from the editor\n")
		for {
			teditor, err := editorScanner.ReadString('\n')
			if err != nil {
				terminalErrorMessage(err.Error(), "Write")
				if len(data) > 0 {
					ioutil.WriteFile(command[0], []byte(data), 0644)
				}
				return
			} else {
				if strings.TrimSpace(teditor) == "texit" {
					break
				}
				data += teditor
			}
		}
		if len(data) > 0 {
			err := ioutil.WriteFile(command[0], []byte(data), 0644)
			if err != nil {
				terminalErrorMessage(err.Error(), "Write")
			} else {
				terminalMessage("Sucessfully Wrote " + strconv.Itoa(len(data)) + " bytes in " + command[0])
			}
		}
		return
	} else if len(command) > 1 {
		terminalErrorMessage("Too Much Arguments", "Write")
	} else {
		terminalErrorMessage("Too Low Arguments", "Write")
	}
}
