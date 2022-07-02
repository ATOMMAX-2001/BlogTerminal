package main

import (
	"io"
	"os"
	"strconv"

	cp "github.com/otiai10/copy"
)

func execute_copy(command []string) {
	if len(command) == 2 {
		soruceFile, err := os.Stat(command[0])
		if err != nil {
			terminalErrorMessage(err.Error(), "FileNotFound")
		} else {
			if soruceFile.IsDir() {
				//move folder
				err := cp.Copy(command[0], command[1])
				if err != nil {
					terminalErrorMessage(err.Error(), "DirCopy")
				} else {
					terminalMessage("All Files Copied from " + command[0] + " to " + command[1] + ", Successfuly")
				}
			} else if !soruceFile.Mode().IsRegular() {
				terminalErrorMessage("Terminal Dont Have the Permission To Read The File ["+soruceFile.Name()+"]", "FilePermission")
			} else {
				//move file
				source, err := os.Open(command[0])
				if err != nil {
					terminalErrorMessage(err.Error(), "FileOpen")
				} else {
					defer source.Close()
					destination, err := os.Create(command[1])
					if err != nil {
						terminalErrorMessage(err.Error(), "FileCreate")
					} else {
						defer destination.Close()
						fileCopyBytes, err := io.Copy(destination, source)
						if err != nil {
							terminalErrorMessage(err.Error(), "FileCopy")
						} else {
							terminalMessage("[" + soruceFile.Name() + "] Copied " + strconv.Itoa(int(fileCopyBytes)) + " Bytes to [" + destination.Name() + "], Successfully")
						}
					}
				}
			}
		}
	} else if len(command) > 2 {
		terminalErrorMessage("Too Much Arguments", "Copy")
	} else {
		terminalErrorMessage("Too Low Arguments", "Copy")
	}
}
