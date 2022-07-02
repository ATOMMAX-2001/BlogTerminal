package main

import (
	"os/exec"
	"strings"

	"atomicgo.dev/cursor"
)

func execute_spawn(command []string, runkeys map[string]string) {
	if len(command) == 1 {
		availableCommand := []string{
			"logout", "copy", "view", "write", "clear", "clean", "cp", "create", "remove", "rm", "rename", "ren", "drive", "find", "goto", "list", "ls", "new", "date", "setdate", "type", "storage", "tree", "time", "settime", "checkdisk", "eval", "poweroff", "restart", "filefmt", "cmd", "terminal", "exit", "env", "$", "\"", "tele", "setcolor", "run", "et",
		}
		foundTCommand := false
		for i := range availableCommand {
			if strings.EqualFold(availableCommand[i], command[0]) {
				terminalErrorMessage("Spawn is mostly used for applications and won't work for Terminal commands", "Spawn")
				foundTCommand = true
			}
		}
		if !foundTCommand {
			// terminalErrorMessage("Invalid Command, Type [helpme] For Commnad Usage", "Command")
			notApplicable := []string{"dir", "cd", "type", "Xcopy", "cls", "del"}
			foundCommand := false
			for i := range command {
				for j := range notApplicable {

					if strings.EqualFold(command[i], notApplicable[j]) {
						terminalErrorMessage("Invalid Command, Type [helpme] For Commnad Usage", command[i])
						foundCommand = true
					}
				}
			}
			if !foundCommand {
				if strings.Contains(command[0], "run") {
					arg := strings.Split(command[0], " ")
					for key, value := range runkeys {
						if key == "RUN_"+strings.ToUpper(arg[1]) {
							arg[1] = strings.ReplaceAll(value, "/", "\\")
							command[0] = strings.Join(arg[1:], " ")
							break
						}
					}
				}
				output := exec.Command("cmd.exe", "/c", strings.Join(command, " "))
				// output.Stdout = os.Stdout
				// output.Stdin = os.Stdin
				// output.Stderr = os.Stderr
				cursor.Down(2)
				go output.Run()
			}
		}
	} else {
		terminalErrorMessage("Too Much Arguments", "Spawn")
	}
}
