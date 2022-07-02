package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"atomicgo.dev/cursor"
	ct "github.com/daviddengcn/go-colortext"
)

func terminalSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	if <-sig != nil {
		signal.Stop(sig)
	}
}

func logo() {
	ct.Foreground(ct.Cyan, true)
	fmt.Println("\t\t\t\t___________                  .__.             .__.   ")
	fmt.Println("\t\t\t\t\\__    ___/__________  _____ |__| ____ _____  |  |  ")
	ct.Foreground(ct.Cyan, true)
	fmt.Println("\t\t\t\t  |    |_/ __ \\_  __ \\/     \\|  |/     \\__  \\ |  |  ")
	ct.Foreground(ct.Magenta, true)
	fmt.Print("\t\t\t\t  |    |\\  ___/|  | \\/")
	ct.Foreground(ct.White, true)
	fmt.Print("  Y Y")
	ct.Foreground(ct.Magenta, true)
	fmt.Println("  \\  |   |  \\/ __ \\|  |__")
	ct.Foreground(ct.Yellow, true)
	fmt.Println("\t\t\t\t  |____| \\___  >__|  |__|_|  /__|___|  (____  /____/")
	fmt.Println("\t\t\t\t             \\/            \\/        \\/     \\/      	")
	ct.Foreground(ct.White, false)
	cursor.Right(28)
	fmt.Println("| All Rights Reserved || Author: S.Abilash || from Blogcorp |")
	cursor.Right(50)
	ct.Foreground(ct.Green, false)
	fmt.Println("||  Version: 1.3.2  ||")
	cursor.Right(48)
	ct.Foreground(ct.Yellow, true)
	fmt.Println("||  ForHelpType: helpme ||")
	cursor.Down(3)
	ct.ResetColor()
}

func main() {
	//display logo
	logo()
	var Tuser = new(Terminaluser)
	Tuser.cwd = os.Args[0]
	*Tuser = read_terminal_config(*Tuser)
	*Tuser = assign_theme(*Tuser)
	*Tuser = assign_username_and_cwd(*Tuser)
	//get the user input buffer
	reader := bufio.NewReader(os.Stdin)
	for {
		//display username and cwd
		// cursor.Right(1)
		terminalMessage(Tuser.username + "@Terminal")
		cursor.Right(1)
		set_default_terminal_theme(Tuser.theme)
		fmt.Print(Tuser.cwd + " -:$ ")
		userCommand, err := reader.ReadString('\n')
		if err != nil {
			terminalErrorMessage("Can't able to Recognize the Keyboard buffer input ", "KeyboardBuffer")
		}
		//trim the input buffer
		userCommand = strings.TrimSpace(userCommand)
		if len(userCommand) != 0 {
			userData := filter_and_split_the_command(userCommand)
			go terminalSignal()
			*Tuser = execute_terminal_command(userData, *Tuser)
		}
	}
}
