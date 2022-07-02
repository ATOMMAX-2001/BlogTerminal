package main

import "os"

func execute_new(command []string) {
	if len(command) >= 1 {
		for i := 0; i < len(command); i++ {
			_, err := os.Stat(command[i])
			if os.IsNotExist(err) {
				file, err := os.Create(command[i])
				if err != nil {
					terminalErrorMessage(err.Error(), "New")
				} else {
					terminalMessage("[" + command[i] + "] File Created Successfully")
					file.Close()
				}
			} else {
				terminalErrorMessage(command[i]+": File Already Exist", "New")
			}
		}
	} else {
		terminalErrorMessage("Too Low Arguments", "New")
	}
}
