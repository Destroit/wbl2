package main
import (
    "os"
    "io"
    "strings"
    "flag"
    "net/http"
    "net/url"
    "fmt"
)
/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func errorHandler(err error) {
    if err != nil {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
    }
}

func getURL(urlRaw string) (*url.URL, error) {
    u, err := url.ParseRequestURI(urlRaw)
    if err != nil || u.Scheme == "" || u.Host == "" {
	return nil, fmt.Errorf("invalid URL")
    }
    return u, nil
}

func makeFile(u *url.URL) (*os.File, error) {
    splitUrl := strings.Split(u.Path, "/")
    name := splitUrl[len(splitUrl)-1]
    if name == "" {
	name = "website"
    }
    file, err := os.Create(name)
    if err != nil {
	return nil, err
    }
    return file, nil
}

func downloadData(u *url.URL, file *os.File) error {
    defer file.Close()
    resp, err := http.Get(u.String())
    if err != nil {
	return err
    }
    defer resp.Body.Close()
    _, err = io.Copy(file, resp.Body)
    if err != nil {
	return err
    }
    return nil
}

func main() {
    flag.Parse()
    args := flag.Args()
    argLen := len(args)
    if argLen == 0 {
	fmt.Println("Missing URL. Usage: wget [URL]")
	os.Exit(1)
    } else if argLen > 1 {
	fmt.Fprintln(os.Stderr, "Too many arguments. Usage: wget [URL]")
	os.Exit(1)
    }
    urlRaw := args[0]

    u, err := getURL(urlRaw)
    errorHandler(err)

    file, err := makeFile(u)
    errorHandler(err)
    err = downloadData(u, file)
    errorHandler(err)
}
