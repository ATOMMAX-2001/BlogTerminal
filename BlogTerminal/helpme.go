package main

import "fmt"

func helpme(command []string) {
	if len(command) == 1 {
		fmt.Println("BlogTerminal from Blogcorp <HELP>")
		switch command[0] {
		case "copy", "cp":
			helpme_copy()
		case "clear", "clean":
			helpme_clear()
		case "ref", "refresh":
			helpme_refresh()
		case "spawn", "sp":
			helpme_spawn()
		case "exit", "logout":
			helpme_exit()
		case "list", "ls":
			helpme_list()
		case "goto":
			helpme_goto()
		case "drive":
			helpme_drive()
		case "goback", "gb":
			helpme_goback()
		case "create":
			helpme_create()
		case "remove", "rm":
			helpem_remove()
		case "view", "read":
			helpme_view()
		case "new":
			helpme_new()
		case "rename", "ren":
			helpme_rename()
		case "setcolor":
			helpme_setcolor()
		case "poweroff":
			helpme_poweroff()
		case "restart":
			helpme_restart()
		case "setdate":
			helpme_setdate()
		case "settime":
			helpme_settime()
		case "checkdisk":
			helpme_checkdisk()
		case "volinfo":
			helpme_volinfo()
		case "elevate":
			helpme_elevate()
		case "teleport", "tele":
			helpme_teleport()
		case "run":
			helpme_run()
		case "write":
			helpme_write()
		case "et":
			helpme_et()

		case "evaluate", "eval":
			helpme_eval()
		case "type":
			helpem_type()

		case "date":
			helpme_date()
		case "time":
			helpme_time()
		case "find":
			helpme_find()

		case "tree":
			helpme_tree()

		default:
			terminalErrorMessage("Invalid Command. Type [helpme] for command usage", "Helpme")

		}
	} else if len(command) > 1 {
		terminalErrorMessage("Too Much Arguments", "Helpme")
	} else {
		display_terminal_commands()
	}
}

func display_terminal_commands() {
	fmt.Println("\nBlogTerminal Commands:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\n\t\tcopy,cp         <Copy files>")
	fmt.Println("\t\tclear,clean      <Clear the console>")
	fmt.Println("\t\tcreate           <Create a new directory>")
	fmt.Println("\t\tremove,rm        <To Remove one or more File(s)/Folder(s)>")
	fmt.Println("\t\tdrive            <To set different drive path>")
	fmt.Println("\t\tfind             <Search a file in current working directory>")
	fmt.Println("\t\tgoto             <To set the specified path as the working directory>")
	fmt.Println("\t\tgoback,gb        <To set the specified parent directory as working directory>")
	fmt.Println("\t\tlist,ls          <List all the files in current working directory")
	fmt.Println("\t\tnew              <Create an empty file>")
	fmt.Println("\t\trename,ren       <Rename an existing file>")
	fmt.Println("\t\tview             <Display the content in a file>")
	fmt.Println("\t\twrite            <To write a file>")
	fmt.Println("\t\tdate             <Display the date>")
	fmt.Println("\t\tsetdate          <To set The Date>")
	fmt.Println("\t\ttime             <Display the time>")
	fmt.Println("\t\tsettime          <To set time>")
	fmt.Println("\t\ttype             <Display the filetype>")
	fmt.Println("\t\tvolinfo          <Display the volume of the drive>")
	fmt.Println("\t\ttree             <Display the folder structure>")
	fmt.Println("\t\tcheckdisk        <Check the disk And report the status(Administrator Required)>")
	fmt.Println("\t\teval             <Calculator>")
	fmt.Println("\t\tpoweroff         <Shutdown>")
	fmt.Println("\t\trestart          <Restart>")
	fmt.Println("\t\tcmd              <Open an instance of Cmd>")
	fmt.Println("\t\tbtermi           <Open an instance of BlogTerminal>")
	fmt.Println("\t\texit,logout      <Close The BlogTerminal>")
	fmt.Println("\t\tenv              <Display all the environment variable>")
	fmt.Println("\t\t$                <To Concat one or more Arguments>")
	fmt.Println("\t\t\"                <To Concat one or more Arguments>")
	fmt.Println("\t\tteleport,tele    <To set the working directory to a defined path>")
	fmt.Println("\t\tsetcolor         <To change the theme color>")
	fmt.Println("\t\trun              <To alias a path>")
	fmt.Println("\t\tet               <Display the execution time for an application>")
	fmt.Println("\t\tspawn,sp         <Spawn a new Thread to execute an Application>")
	fmt.Println("\t\televate          <To execute an application by the super-user or admin>")
	fmt.Println("\t\tfind             <To check if a file exist or not>")
}

func helpme_copy() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Copy:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tCopy command is used to copy a file/folder from one location to another")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tcopy <src> <dst>  |OR| cp <src> <dst>")
}
func helpme_clear() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Clear:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tClear command is used to clear the console")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tclear |OR| clean")
}

func helpme_refresh() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Refresh:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tRefresh command is used to update the config file")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\trefresh |OR| ref")
}
func helpme_spawn() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Spawn:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tSpawn command is used to run an application in different thread")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tspawn <application> |OR| sp <application>")
	fmt.Println("\t\tNOTE: Spawn can have only one arguments. Add [$] in prefix of your application to add more than one arguments")
}
func helpme_exit() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Exit:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo exit from BlogTerminal")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\texit |OR| logout")
}
func helpme_list() {
	set_default_terminal_theme("GREEN")
	fmt.Println("List:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo list all the files and folders in the working directory")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tlist |OR| ls |OR| ls/list <path>")
}
func helpme_goto() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Goto:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo set specified path as the working directory ")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tgoto <path>")
}
func helpme_drive() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Drive:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo set different drive path in the working directory")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tdrive <driveName>")
}
func helpme_goback() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Goback:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo set specified parent directory as a working directory ")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tgoback <parentDir> |OR| gb <parentPath>")
}
func helpme_create() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Create:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo create one or more folder(s)")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tcreate <name> ...")
}
func helpem_remove() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Remove:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo Remove one or more File(s)/Folder(s)")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tremove <file/folder> ... |OR| rm <file/folder> ...")
}
func helpme_view() {
	set_default_terminal_theme("GREEN")
	fmt.Println("View:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo view the content of a file")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tview <fileName> |OR| read <fileName>")
}
func helpme_new() {
	set_default_terminal_theme("GREEN")
	fmt.Println("New:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo create an empty file")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tnew <fileName>")
}
func helpme_rename() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Rename:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo rename an existing file")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\trename <old> <new> |OR| ren <old> <new>")
}
func helpme_setcolor() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Setcolor:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo set a theme")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tsetcolor |OR| setcolor <value/name>")
}
func helpme_poweroff() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Powerff:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo shutdown the computer")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tpoweroff")
}
func helpme_restart() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Restart:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo restart the computer")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\trestart")
}
func helpme_setdate() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Setdate:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo update the date")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tsetdate")
}

func helpme_date() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Date:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tDisplay the date")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tdate")
}

func helpme_settime() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Settime:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo update the time")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tsettime")
}
func helpme_time() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Time:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tDisplay the time")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\ttime")
}

func helpme_checkdisk() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Checkdisk:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tCheck and report the disk")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tcheckdisk")
	fmt.Println("NOTE: It needs an admin privilege to run the commnad. Add [elevate] in prefix to run the command ")
}
func helpme_volinfo() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Volinfo:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tDisplay the serialno of the disk")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tvolinfo")
}
func helpme_elevate() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Elevate:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo execute an application by the super-user or admin")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\televate <application>")
}
func helpme_teleport() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Teleport:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo set the working directory to a defined path")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tteleport <keyword> |OR| tele <keyword>")
	fmt.Println("NOTE: The user need to define the path for the keyword in the [terminal.config] file")
}
func helpme_run() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Run:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo alias a path")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\trun <keyword>")
	fmt.Println("NOTE: The user need to define the path for the keyword in the [terminal.config] file")
}
func helpme_write() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Write:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo write a file in the console")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\twrite <fileName>")
}
func helpme_et() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Et:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tDisplay the execution time for an application ")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tet <application>")
}
func helpme_eval() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Eval:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tA simple calculator")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\teval <expression>")
}
func helpem_type() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Type:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tDisplay the file extension association")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\ttype |OR| type <extension>")
}

func helpme_find() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Find:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo check a file exist or not")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tfile <fileName>")
}

func helpme_tree() {
	set_default_terminal_theme("GREEN")
	fmt.Println("Tree:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\tTo display the structure of a folder")
	set_default_terminal_theme("GREEN")
	fmt.Println("Syntax:")
	set_default_terminal_theme("YELLOW")
	fmt.Println("\t\ttree  |OR|  tree <path>")
}
