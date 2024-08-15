package main
import (
    "fmt"
    "flag"
    "os"
    "os/signal"
    "net"
    "bufio"
    "time"
    "syscall"
)
/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/
func readConn(conn net.Conn) {
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
	fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
	conn.Close()
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
    }
}

func readStdin(conn net.Conn) {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
	_, err := fmt.Fprintln(conn, scanner.Text())
	if err != nil {
	    fmt.Fprintln(os.Stderr, err)
	    os.Exit(1)
	}
	if err := scanner.Err(); err != nil {
	    fmt.Fprintln(os.Stderr, err)
	    os.Exit(1)
	}
    }
}

func makeConnection(addr string, timeout *time.Duration, stop chan struct{}) error {
    fmt.Println(addr)
    conn, err := net.DialTimeout("tcp", addr, *timeout)
    if err != nil {
	return err
    }
    defer conn.Close()
    go readConn(conn)
    go readStdin(conn)

    <-stop
    return nil
}

func main() {
    timeout := flag.Duration("timeout", 10 * time.Second, "Set timeout for connection. Example: 10s")
    flag.Parse()
    args := flag.Args()
    lenArgs := len(args)
    if lenArgs == 0 {
	fmt.Println("Usage: telnet [--timeout] [HOST] [PORT]")
	os.Exit(1)
    } else if lenArgs < 2 {
	fmt.Fprintln(os.Stderr, "Missing PORT. Usage: telnet [--timeout] [HOST] [PORT]")
	os.Exit(1)
    } else if lenArgs > 2 {
	fmt.Fprintln(os.Stderr, "Too many arguments. Usage: telnet [--timeout] [HOST] [PORT]")
	os.Exit(1)
    }
    addr := args[0] + ":" + args[1]
    stop := make(chan struct{})
    err := makeConnection(addr, timeout, stop)
    if err != nil {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
    }
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    close(stop)
}
