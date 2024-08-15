package main
import (
    "flag"
    "fmt"
    "regexp"
    "bufio"
    "os"
)
/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func check(err error) {
    if err != nil {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
    }
}

func processArgs() (string, string, error) {
    args := flag.Args()
    argsLen := len(args)
    switch argsLen {
    case 0:
	return "", "", fmt.Errorf("Usage: gogrep [OPTION...] PATTERNS [FILE...]")
    case 1:
	pattern := args[0]
	return pattern, "", nil
    case 2:
	pattern := args[0]
	path := args[1]
	if _, err := os.Stat(path); os.IsNotExist(err) {
	    return "", "", fmt.Errorf("No such file: %v", path)
	}
	return pattern, path, nil
    default:
	return "", "", fmt.Errorf("Too many arguments. Usage: grep [OPTION...] PATTERNS [FILE...]")
    }
}

func prepareInput(path string) ([]string, error) {
    lines := make([]string, 0)
    var (
	file *os.File
	err error
    )

    if len(path) != 0 {
	file, err = os.Open(path)
	if err != nil {
	    return nil, err
	}
    } else {
	file = os.Stdin
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	    line := scanner.Text()
	    lines = append(lines, line)
    }
    return lines, scanner.Err()
}

func prepareExpression(pattern string, ignoreCase bool, fixed bool) (*regexp.Regexp, error) {
    if fixed {
	pattern = `(\Q` + pattern + `\E)`
    }
    if ignoreCase {
	pattern = "(?i)(" + pattern + ")"
    }
    rgx, err := regexp.Compile(pattern)
    if err != nil {
	return nil, err
    }
    return rgx, nil
}

func doCount(lines []string, rgx *regexp.Regexp) int {
    cnt := 0
    for _, line := range lines {
	match := rgx.Match([]byte(line))
	if match {
	    cnt++
	}
    }
    return cnt
}

func doSearchAndPrint(lines []string, rgx *regexp.Regexp, before int, after int, invert bool, number bool) {
    for idx, line := range lines {
	match := rgx.Match([]byte(line))
	if !invert && match {
	    doPrint(lines, idx, before, after, number)
	} else if invert && !match {
	    doPrint(lines, idx, before, after, number)
	}
    }
}

func doPrint(lines []string, idx int, before int, after int, number bool) {
    start := idx - before
    if start < 0 {
	start = 0
    }

    end := idx + after
    if end > len(lines) - 1 {
	end = len(lines) - 1
    }
    fmt.Println("---------------")
    if number {
	for i := start; i < idx; i++ {
	    fmt.Println(i+1, lines[i])
	}
	fmt.Println(idx+1, "\033[31m" + lines[idx] + "\033[0m")
	for i := idx+1; i <= end; i++ {
	    fmt.Println(i+1, lines[i])
	}
    } else {
	for i := start; i < idx; i++ {
	    fmt.Println(lines[i])
	}
	fmt.Println("\033[31m" + lines[idx] + "\033[0m")
	for i := idx+1; i <= end; i++ {
	    fmt.Println(lines[i])
	}

    }
}

func main() {
    after := flag.Int("A", 0, "Print NUM lines of trailing context after matching lines")
    before := flag.Int("B", 0, "Print NUM lines of trailing context before matching lines")
    context := flag.Int("C", 0, "Print NUM lines of output context")
    count := flag.Bool("c", false, "Supress normal output. Print a count of matching lines for each input file")
    ignoreCase := flag.Bool("i", false, "Ignore case distinctions in patterns and input")
    invert := flag.Bool("v", false, "Invert sense of matching, to select non-maching lines")
    fixed := flag.Bool("F", false, "Interpret pattern as fixed string")
    lineNum := flag.Bool("n", false, "Print numbers of lines")

    flag.Parse()

    pattern, path, err := processArgs()
    check(err)
    lines, err := prepareInput(path)
    check(err)
    rgx, err := prepareExpression(pattern, *ignoreCase, *fixed)
    check(err)

    if *count {
	cnt := doCount(lines, rgx)
	if *invert {
	    cnt = len(lines) - cnt
	}
	fmt.Println(cnt)
    } else if *after != 0 || *before != 0 {
	doSearchAndPrint(lines, rgx, *before, *after, *invert, *lineNum)
    } else {
	doSearchAndPrint(lines ,rgx, *context, *context, *invert, *lineNum)
    }
}
