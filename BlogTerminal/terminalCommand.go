package main

import (
	"os"
	"os/exec"
	"strings"

	"atomicgo.dev/cursor"
)

func filter_and_split_the_command(command string) []string {
	var userData []string
	userCommand := strings.Split(command, " ")
	for i := 0; i < len(userCommand); i++ {
		if len(userCommand[i]) == 0 {
			continue
		} else if string(userCommand[i][0]) == "$" {
			userCommand[i] = userCommand[i][1:]
			temp := strings.Join(userCommand[i:], " ")
			userData = append(userData, temp)
			break
		} else if string(userCommand[i][0]) == "\"" {
			userCommand[i] = userCommand[i][1:]
			if string(userCommand[i][len(userCommand[i])-1]) == "\"" {
				userData = append(userData, userCommand[i][0:len(userCommand[i])-1])
			} else {
				var quoteFoundFlag bool = false
				for j := i + 1; j < len(userCommand); j++ {
					if string(userCommand[j][len(userCommand[j])-1]) == "\"" {
						quoteFoundFlag = true
						userCommand[j] = userCommand[j][0 : len(userCommand[j])-1]
						temp := strings.Join(userCommand[i:j+1], " ")
						userData = append(userData, temp)
						i = j
					}
				}
				if !quoteFoundFlag {
					terminalErrorMessage("Closing Quotes[\" \"] is not found", "Literal")
				}
			}
		} else {
			userData = append(userData, userCommand[i])
		}

	}
	return userData
}

func execute_terminal_command(command []string, Tuser Terminaluser) Terminaluser {
	var multipleCommandIndex int = 0
	for i := range command {
		if command[i] == "&&" {
			multipleCommandIndex = i
		}
	}
	if multipleCommandIndex > 0 {
		Tuser = execute_command(command[:multipleCommandIndex], Tuser)
		Tuser = execute_terminal_command(command[multipleCommandIndex+1:], Tuser)
		return Tuser
	}
	Tuser = execute_command(command, Tuser)
	return Tuser
}

func execute_command(command []string, Tuser Terminaluser) Terminaluser {
	switch command[0] {
	case "abilash", "s.abilash", "author":
		terminalMessage("BlogTerminal was designed & Developed by S.Abilash <github.com/ATOMMAX-2001>")
	case "copy", "cp":
		execute_copy(command[1:])

	case "clear", "clean":
		output := exec.Command("cmd.exe", "/c", "cls")
		output.Stdout = os.Stdout
		output.Run()

	case "ref", "refresh":
		Tuser = read_terminal_config(Tuser)
		Tuser = assign_theme(Tuser)
		Tuser = assign_username_and_cwd(Tuser)

	case "spawn", "sp":
		execute_spawn(command[1:], Tuser.config)

	case "exit", "logout":
		terminalMessage("Logging Out!!!")
		os.Exit(0)

	case "list", "ls":
		execute_list(command[1:])

	case "goto":
		Tuser = execute_goto(command[1:], Tuser)

	case "drive":
		Tuser = execute_drive(command[1:], Tuser)

	case "goback", "gb":
		Tuser = execute_goback(command[1:], Tuser)

	case "create":
		execute_create(command[1:])

	case "remove", "rm":
		execute_remove(command[1:])

	case "view", "read":
		execute_view(command[1:])
	case "new":
		execute_new(command[1:])

	case "rename", "ren":
		execute_rename(command[1:])

	case "setcolor":
		Tuser = execute_setcolor(command[1:], Tuser)

	case "poweroff":
		output := exec.Command("cmd.exe", "/c", "shutdown", "-s", "-t", "1", "-c", "Thank You, Have a Nice Day!!")
		output.Stdout = os.Stdout
		output.Run()

	case "restart":
		output := exec.Command("cmd.exe", "/c", "shutdown", "-r", "-c", "Taking a Quick Nap")
		output.Stdout = os.Stdout
		output.Run()

	case "date":
		output := exec.Command("cmd.exe", "/c", "date", "/T")
		output.Stdout = os.Stdout
		cursor.Down(1)
		output.Run()

	case "setdate":
		output := exec.Command("cmd.exe", "/c", "date")
		output.Stdout = os.Stdout
		output.Stdin = os.Stdin
		output.Stderr = os.Stderr
		output.Run()

	case "time":
		output := exec.Command("cmd.exe", "/c", "time", "/T")
		output.Stdout = os.Stdout
		cursor.Down(1)
		output.Run()

	case "settime":
		output := exec.Command("cmd.exe", "/c", "time")
		output.Stdout = os.Stdout
		output.Stdin = os.Stdin
		output.Stderr = os.Stderr
		output.Run()

	case "type":
		var output *exec.Cmd
		if len(command) > 1 {
			output = exec.Command("cmd.exe", "/c", "assoc", strings.Join(command[1:], " "))
		} else {
			cursor.Down(1)
			output = exec.Command("cmd.exe", "/c", "assoc")
		}
		output.Stdout = os.Stdout
		output.Run()
	case "find":
		execute_find(command[1:])

	case "env":
		output := exec.Command("cmd.exe", "/c", "set")
		output.Stdout = os.Stdout
		output.Run()

	case "checkdisk":
		output := exec.Command("cmd.exe", "/c", "chkdsk")
		output.Stdout = os.Stdout
		output.Stdin = os.Stdin
		output.Stderr = os.Stderr
		output.Run()

	case "volinfo":
		output := exec.Command("cmd.exe", "/c", "vol")
		output.Stdout = os.Stdout
		output.Stdin = os.Stdin
		output.Stderr = os.Stderr
		output.Run()

	case "elevate":
		output := exec.Command("cmd.exe", "/c", "powershell", "-Command", "Start-Process", "-Verb", "RunAs", command[1])
		output.Stdout = os.Stdout
		output.Stdin = os.Stdin
		output.Run()

	case "teleport", "tele":
		Tuser.cwd = execute_teleport(command[1:], Tuser.config, Tuser.cwd)

	case "run":
		execute_run(command[1:], Tuser.config)

	case "write":
		execute_write(command[1:])

	case "et":
		execute_et(command[1:], Tuser.config)

	case "helpme":
		helpme(command[1:])

	case "evaluate", "eval":
		execute_eval(command[1:])

	default: // terminalErrorMessage("Invalid Command, Type [helpme] For Commnad Usage", "Command")
		notApplicable := []string{"dir", "cd", "type", "Xcopy", "cls", "del", "date", "time", "chkdsk"}
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
			output := exec.Command("cmd.exe", "/c", strings.Join(command, " "))
			output.Stdout = os.Stdout
			output.Stdin = os.Stdin
			output.Stderr = os.Stderr
			cursor.Down(2)
			err := output.Run()
			if err != nil {
				terminalErrorMessage("Invalid Command. For Command Usage Type [helpme] ", strings.Join(command, " "))
			}
		}

	}
	return Tuser
}
