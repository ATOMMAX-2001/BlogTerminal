package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func execute_goto(command []string, Tuser Terminaluser) Terminaluser {
	if len(command) == 1 {
		_, err := ioutil.ReadDir(command[0])
		if err != nil {
			terminalErrorMessage(err.Error(), "Goto")
		} else {
			if command[0] == ".." {
				Tuser.cwd = execute_goto_parent(Tuser.cwd)
				os.Chdir(Tuser.cwd)
				return Tuser
			} else {
				path := strings.ReplaceAll(command[0], "/", "\\")
				if strings.Contains(path, "..") {
					index := strings.Index(path, "..")
					newpath := execute_goto_parent(path[:index-1])
					newpath += path[index+2:]
					path = newpath
				}
				Tuser.cwd += "\\" + path
				if string(Tuser.cwd[len(Tuser.cwd)-1]) == "\\" {
					Tuser.cwd = Tuser.cwd[:len(Tuser.cwd)-1]
				}

				Tuser.cwd = strings.ReplaceAll(Tuser.cwd, "\\\\", "\\")
				os.Chdir(Tuser.cwd)
				return Tuser
			}
		}

	} else if len(command) > 1 {
		terminalErrorMessage("Too Much Arguments", "Goto")
	} else {
		terminalErrorMessage("Too Low Arguments", "Goto")
	}
	return Tuser
}

func execute_goto_parent(path string) string {
	var index int
	for i := len(path) - 1; i >= 0; i-- {
		if string(path[i]) == "\\" {
			index = i
			break
		}
	}
	path = path[0:index]
	return path
}

func execute_goback(command []string, Tuser Terminaluser) Terminaluser {
	if len(command) == 1 {
		if strings.Contains(Tuser.cwd, command[0]) {
			index := strings.Index(Tuser.cwd, command[0])
			Tuser.cwd = Tuser.cwd[0 : index+len(command[0])]
			os.Chdir(Tuser.cwd)
			return Tuser
		} else {
			terminalErrorMessage("System Cannot Find The path", "Goback")
		}
	} else if len(command) > 1 {
		terminalErrorMessage("Too Much Arguments", "Goback")
	} else {
		terminalErrorMessage("Too Low Arguments", "Goback")
	}
	return Tuser
}
