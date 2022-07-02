package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"atomicgo.dev/cursor"
)

func execute_list(command []string) {
	if len(command) == 1 {
		files, err := ioutil.ReadDir(command[0])
		if err != nil {
			terminalErrorMessage(err.Error(), "list")
		} else {
			dirCount, fileCount := 0, 0
			var memoryUsed int64
			cursor.Down(3)
			for _, file := range files {
				cursor.Right(3)
				if file.IsDir() {
					dirCount++
					name := file.Name()
					if len(name) > 40 {
						name = name[:40]
					}
					mode := file.Mode().String()
					mode = strings.ReplaceAll(mode, "-", " ")
					fmt.Printf("<N/A>\t <%s>\t [DIR]\t %s\n ", mode, name)
				} else {
					fileCount++
					name := file.Name()
					if len(name) > 40 {
						name = name[:40] + " ...."
					}
					mode := file.Mode().String()
					mode = strings.ReplaceAll(mode, "-", " ")
					size := file.Size()
					memoryUsed += int64(size)
					fsize := strconv.Itoa(int(size))
					if len(fsize) > 6 {
						fsize = fsize[:6] + "+"
					}
					fmt.Printf("<%s B>\t <%s>\t [FIL]\t %s\n ", fsize, mode, name)
				}
			}
			summary := "\tTotal File(s): " + strconv.Itoa(fileCount) + "\n\tTotal Dir(s): " + strconv.Itoa(dirCount) + "\n\tMemory Used: " + strconv.FormatInt(memoryUsed, 10) + " Bytes"
			terminalMessage(summary)
		}
	} else if len(command) == 0 {
		path, _ := os.Getwd()
		files, err := ioutil.ReadDir(path)
		if err != nil {
			terminalErrorMessage(err.Error(), "list")
		} else {
			dirCount, fileCount, memoryUsed := 0, 0, 0
			cursor.Down(3)
			for _, file := range files {
				cursor.Right(3)
				if file.IsDir() {
					dirCount++
					name := file.Name()
					if len(name) > 40 {
						name = name[:40]
					}
					mode := file.Mode().String()
					mode = strings.ReplaceAll(mode, "-", " ")
					fmt.Printf("<N/A>\t <%s>\t [DIR]\t %s\n ", mode, name)
				} else {
					fileCount++
					name := file.Name()
					if len(name) > 40 {
						name = name[:40] + " ...."
					}
					mode := file.Mode().String()
					mode = strings.ReplaceAll(mode, "-", " ")
					size := file.Size()
					memoryUsed += int(size)
					fsize := strconv.Itoa(int(size))
					if len(fsize) > 6 {
						fsize = fsize[:6] + "+"
					}
					fmt.Printf("<%s B>\t <%s>\t [FIL]\t %s\n ", fsize, mode, name)
				}
			}
			summary := "\tTotal File(s): " + strconv.Itoa(fileCount) + "\n\tTotal Dir(s): " + strconv.Itoa(dirCount) + "\n\tMemory Used: " + strconv.Itoa(memoryUsed) + " Bytes"
			terminalMessage(summary)

		}
	} else {
		terminalErrorMessage("Too Much Arguments", "list")
	}
}
