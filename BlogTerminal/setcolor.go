package main

import "fmt"

func execute_setcolor(command []string, Tuser Terminaluser) Terminaluser {
	if len(command) == 0 {
		fmt.Println("\n These are the Available Color Theme:")
		terminalMessage("[1 => CYAN,2 => GREEN,3 => RED,4 => MAGENTA,5 => WHITE,6 => BLUE]")
	} else if len(command) == 1 {
		switch command[0] {
		case "1", "cyan":
			set_default_terminal_theme("CYAN")
			Tuser.theme = "CYAN"
		case "2", "green":
			set_default_terminal_theme("GREEN")
			Tuser.theme = "GREEN"
		case "3", "red":
			set_default_terminal_theme("RED")
			Tuser.theme = "RED"
		case "4", "magenta":
			set_default_terminal_theme("MAGENTA")
			Tuser.theme = "MAGENTA"
		case "5", "white":
			set_default_terminal_theme("WHITE")
			Tuser.theme = "WHITE"
		case "6", "blue":
			set_default_terminal_theme("BLUE")
			Tuser.theme = "BLUE"
		case "7", "yellow":
			set_default_terminal_theme("YELLOW")
		default:
			terminalErrorMessage("Invalid Color ", "Setcolor")
		}
	} else if len(command) > 1 {
		terminalErrorMessage("Too Much Arguments", "Setcolor")
	} else {
		terminalErrorMessage("Too Low Arguments", "Setcolor")
	}
	return Tuser
}
