package main
import (
    "bufio"
    "os"
    "os/exec"
    "fmt"
    "strings"
    "strconv"
    "github.com/mitchellh/go-ps"
)
/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func shellHandler(){
    scanner := bufio.NewScanner(os.Stdin)
    for { 
	fmt.Print("gosh ~> ")
	if scanner.Scan() {
	    request := strings.Fields(scanner.Text())
	    if len(request) == 0 {
		continue
	    }
	    cmd := request[0]
	    args := request[1:]
	    switch cmd {
	    case "cd":
		doCd(args)
	    case "pwd":
		doPwd(args)
	    case "echo":
		doEcho(args)
	    case "kill":
		doKill(args)
	    case "ps":
		doPs(args)
	    case "exec":
		doExec(args)
	    case "exit":
		os.Exit(0)
	    default:
		fmt.Fprintln(os.Stderr, "Unknown command: ", cmd)
	    }
	} else {
	    break
	}
    }
}

func errorHandler(err error) bool {
    if err != nil {
	fmt.Fprintln(os.Stderr, err)
	return true
    }
    return false
}

func doCd(args []string) {
    if len(args) == 0 {
	dir, err := os.UserHomeDir()
	if errorHandler(err) {
	    return
	}
	err = os.Chdir(dir)
	errorHandler(err)
    } else if len(args) > 1 {
	fmt.Fprintln(os.Stderr, "Too many arguments. Usage: cd [DIRECTORY]")
    } else {
	err := os.Chdir(args[0])
	errorHandler(err)
    }
}

func doPwd(args []string) {
    if len(args) != 0 {
	fmt.Fprintln(os.Stderr, "Too many arguments. Usage: pwd")
	return
    }
    wd, err := os.Getwd()
    if errorHandler(err) {
	return
    }
    fmt.Println(wd)
}

func doEcho(args []string) {
    fmt.Println(strings.Join(args, " "))
}

func doKill(args []string) {
    var pidsToKill []int
    if len(args) == 0 {
	fmt.Fprintln(os.Stderr, "Usage: kill <pid> [...]")
	return
    }

    for _,v := range args {
	pid, err := strconv.Atoi(v)
	if errorHandler(err) {
	    return
	}
	pidsToKill = append(pidsToKill, pid)
    }

    for _, pid := range pidsToKill {
	process, err := os.FindProcess(pid)
	if errorHandler(err) {
	    return
	}
	err = process.Kill()
	if errorHandler(err) {
	    return
	}
    }
}

func doPs(args []string) {
    if len(args) != 0 {
	fmt.Fprintln(os.Stderr, "Too many arguments. Usage: ps")
	return
    }

    list, err := ps.Processes()
    if errorHandler(err) {
	return
    }
    fmt.Println("PID\tCMD")
    for _,process := range list {
	fmt.Printf("%v\t%v\n", process.Pid(), process.Executable())
    }
}

func doExec(args []string) {
    if len(args) == 0 {
	fmt.Fprintln(os.Stderr, "Usage: exec COMMAND")
	return
    }
    newcmd := exec.Command(args[0], args[1:]...)
    newcmd.Stderr = os.Stderr
    newcmd.Stdout = os.Stdout
    newcmd.Stdin = os.Stdin
    err := newcmd.Run()
    errorHandler(err)
}

func main() {
    shellHandler()
}
