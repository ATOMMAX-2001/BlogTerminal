package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"atomicgo.dev/cursor"
	ct "github.com/daviddengcn/go-colortext"
)

type Terminaluser struct {
	username string
	theme    string
	cwd      string
	config   map[string]string
}

func read_terminal_config(Tuser Terminaluser) Terminaluser {
	var path string
	if strings.Contains(os.Args[0], "BlogTerminal.exe") {
		path = os.Args[0][0 : len(os.Args[0])-16]
	} else if strings.Contains(os.Args[0], "btermi.exe") {
		path = os.Args[0][0 : len(os.Args[0])-10]
	}
	result := path + "terminal.config"
	configFile, err := os.Open(result)
	if err != nil {
		terminalErrorMessage("Can't able to read the [terminal.config] file", "Config")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	defer configFile.Close()
	collectionOfKeyAndValues := make(map[string]string)
	fileScanner := bufio.NewScanner(configFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		var configData string = strings.TrimSpace(fileScanner.Text())
		if len(configData) == 0 {
			continue
		}
		if string(configData[0]) == "#" {
			continue
		} else {
			var colonIndex int = strings.Index(configData, ":")
			if colonIndex != -1 {
				var configKey string = configData[0:colonIndex]
				var configValue string = configData[colonIndex+1:]
				collectionOfKeyAndValues[configKey] = configValue
			}
		}
	}
	Tuser.config = make(map[string]string)
	Tuser.config = collectionOfKeyAndValues

	return Tuser
}

func checkKeyExist(Tuser Terminaluser, configKey string) bool {
	for key := range Tuser.config {
		if key == configKey {
			return true
		}
	}
	return false
}
func getValue(Tuser Terminaluser, configKey string) string {
	for key, value := range Tuser.config {
		if key == configKey {
			return value
		}
	}
	return "NIL"
}

func set_default_terminal_theme(theme string) {
	switch theme {
	case "CYAN":
		ct.Foreground(ct.Cyan, true)
	case "GREEN":
		ct.Foreground(ct.Green, true)
	case "RED":
		ct.Foreground(ct.Red, true)
	case "MAGENTA":
		ct.Foreground(ct.Magenta, true)
	case "WHITE":
		ct.Foreground(ct.White, true)
	case "BLUE":
		ct.Foreground(ct.Blue, true)
	case "YELLOW":
		ct.Foreground(ct.Yellow, true)
	default:
		terminalErrorMessage("Invalid Theme Option. GOT: '"+theme+"', [Cyan] is considered as a default Theme", "Theme")
		ct.Foreground(ct.Cyan, true)
	}
}

func assign_theme(Tuser Terminaluser) Terminaluser {
	if checkKeyExist(Tuser, "DEFAULT_THEME") {
		Tuser.theme = Tuser.config["DEFAULT_THEME"]
		set_default_terminal_theme(Tuser.theme)
	} else {
		terminalErrorMessage("[DEFAULT_THEME] is not defined in [terminal.config] file, [Cyan] is Considered as the default Theme", "Theme")
		Tuser.theme = "CYAN"
	}
	return Tuser
}

func assign_username_and_cwd(Tuser Terminaluser) Terminaluser {
	if checkKeyExist(Tuser, "DEFAULT_USER") {
		userName := getValue(Tuser, "DEFAULT_USER")
		if userName != "NIL" {
			Tuser.username = userName
		} else {
			terminalErrorMessage("[DEFAULT_USER] is not defined in [terminal.config] file", "config")
		}
	} else {
		tempUserProfile := os.Getenv("USERPROFILE")
		var userProfile string
		for i := len(tempUserProfile) - 1; i >= 0; i-- {
			userProfile += string(tempUserProfile[i])
		}
		var userIndex int = strings.Index(userProfile, "\\")
		Tuser.username = tempUserProfile[len(tempUserProfile)-userIndex:]
	}

	//assigning cwd
	if checkKeyExist(Tuser, "DEFAULT_START_PATH") {
		userStartPath := getValue(Tuser, "DEFAULT_START_PATH")
		if userStartPath != "NIL" {
			Tuser.cwd = userStartPath
			os.Chdir(Tuser.cwd)
		} else {
			terminalErrorMessage("[DEFAULT_START_PATH] is not defined in terminal.config file", "config")
		}
	} else {
		userStartPath, err := os.Getwd()
		if err != nil || userStartPath+"\\BlogTerminal.exe" == Tuser.cwd {
			userStartPath = os.Getenv("USERPROFILE")
			Tuser.cwd = userStartPath
			os.Chdir(Tuser.cwd)
		}
		Tuser.cwd = userStartPath
		os.Chdir(Tuser.cwd)
	}
	return Tuser
}

func terminalErrorMessage(message string, short_message string) {
	ct.Foreground(ct.White, true)
	ct.Background(ct.Red, true)
	cursor.Right(10)
	cursor.Down(2)
	fmt.Println(" ERROR: " + short_message + " \n")
	cursor.Right(10)
	fmt.Print(" DESCRIPTION:")
	cursor.Right(3)
	ct.ResetColor()
	ct.Foreground(ct.Red, true)
	fmt.Println(message + "\n\n")
	ct.ResetColor()
}

func terminalMessage(message string) {
	ct.Foreground(ct.Green, true)
	cursor.Down(2)
	cursor.Right(1)
	fmt.Println(message)
	ct.ResetColor()
}
