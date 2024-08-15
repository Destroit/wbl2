package main
import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "sort"
    "strconv"
    "strings"
)
/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func check(err error) {
    if err != nil {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
    }
}

func prepareInput(path string, uniq bool) ([]string, error) {
    lines := make([]string, 0)
    var ( 
	file *os.File
	err error
    )

    if len(path) != 0 {
	file, err = os.Open(path)
	if err != nil {
	    return nil, fmt.Errorf("Invalid path")
	}
    } else {
	file = os.Stdin
    }
    scanner := bufio.NewScanner(file)

    if uniq {
	set := make(map[string]struct{})
	for scanner.Scan() {
	    if _, ok := set[scanner.Text()]; !ok {
		set[scanner.Text()] = struct{}{}
		lines = append(lines, scanner.Text())
	    }
	}
    } else {
	for scanner.Scan() {
	    lines = append(lines, scanner.Text())
	}
    }
    return lines, scanner.Err()
}

func doSort(lines []string, key int, num bool) ([]string, error) {
    if num {
	sort.Slice(lines, func(i int , j int) bool {
	    numi, err := strconv.Atoi(lines[i])
	    check(err)
	    numj, err := strconv.Atoi(lines[j])
	    check(err)
	    return numi < numj
	})
    } else if key < 1 {
	    return nil, fmt.Errorf("Invalid key")
    } else if key == 1 {
	sort.Slice(lines, func(i int , j int) bool {
	    return lines[i] < lines[j]
	})
    } else {
	sort.Slice(lines, func(i int , j int) bool {
	    // Split into words(columns)
	    iField := strings.Fields(lines[i])
	    jField := strings.Fields(lines[j])
	    // Word count
	    iSize := len(iField)
	    jSize := len(jField)

	    if iSize >= key && jSize >= key {
		iCol := iField[key-1]
		jCol := jField[key-1]
		if iCol < jCol {
		    return true
		} else if iCol == jCol {
		    return lines[i] < lines[j]
		} else {
		    return false
		}
	    } else if iSize < key && jSize < key {
		return lines[i] < lines[j]
	    } else {
		return iSize < jSize
	    }
	})
    }
    return lines, nil
}

func writeOutput(path string, lines []string, reverse bool) error {
    var (
	writer *bufio.Writer 
	file *os.File
	err error
    )
    gotPath := (len(path) != 0)

    if gotPath {
	file, err = os.Create(path)
	if err != nil {
	    fmt.Println("os create")
	    return err
	}
	defer file.Close()
    } else {
	file = os.Stdout
    }
    writer = bufio.NewWriter(file)

    if reverse {
	for i:=len(lines)-1; i > 0; i-- {
	    _, _ = writer.WriteString(lines[i])
	    _, _ = writer.WriteString("\n")
	}
    } else {
	for _, v := range lines {
	    _, _ = writer.WriteString(v)
	    _, _ = writer.WriteString("\n")
	}
    }
    
    err = writer.Flush()
    if err != nil {
	return err
    }

    return nil
}

func main() {
    key := flag.Int("k", 1, "sort via a key")
    num := flag.Bool("n", false, "compare according to string numerical value")
    reverse := flag.Bool("r", false, "reverse the result of comparison")
    unique := flag.Bool("u", false, "output first of an equal run")
    input := flag.String("f", "", "read from FILE instead of standard input")
    output := flag.String("o", "", "write result to FILE instead of standard output")
    flag.Parse()
    fmt.Println("key:", *key)
    fmt.Println("num:", *num)
    fmt.Println("reverse:", *reverse)
    fmt.Println("unique:", *unique)
    fmt.Println("input:", *input)
    fmt.Println("output:", *output)

    data, err := prepareInput(*input, *unique)
    check(err)
    sorted, err := doSort(data, *key, *num)
    check(err)
    err = writeOutput(*output, sorted, *reverse)
    check(err)
   }
