package main
import (
    "flag"
    "fmt"
    "strconv"
    "bufio"
    "strings"
    "sort"
    "os"
)
/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
    fields := flag.String("f", "", "Select fields(columns)")
    delimiter := flag.String("d", "\t", "Select other separator")
    separated := flag.Bool("s", false, "Only lines with separator")
    flag.Parse()

    fieldSet := make(map[int]struct{})
    fieldNums := make([]int, 0)
    if *fields == "" {
	fmt.Fprintln(os.Stderr, "Fields must not be empty")
	os.Exit(1)
    }
    fieldSplit := strings.Split(*fields, ",")
    for _,v := range fieldSplit {
	num, err := strconv.Atoi(v)
	if err != nil {
	    fmt.Fprintln(os.Stderr, err)
	    os.Exit(1)
	}
	if num <= 0 {
	    fmt.Fprintln(os.Stderr, "Fields can't be less or equal zero")
	} else {
	    if _, ok := fieldSet[num]; !ok {
		fieldSet[num] = struct{}{}
		fieldNums = append(fieldNums, num)
	    }
	}
    }
    sort.Ints(fieldNums)

    var sb strings.Builder

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
	txt := scanner.Text()
	if !strings.Contains(txt, *delimiter) {
	    if !(*separated) {
		sb.WriteString(txt)
		sb.WriteString("\n")
	    }
	    continue
	}

	split := strings.Split(txt, *delimiter)
	for i:=0; i < len(fieldNums)-1; i++ {
	    col := fieldNums[i]
	    if col > len(split) {
		break
	    }
	    sb.WriteString(split[col-1])
	    sb.WriteString(*delimiter)
	}
	col := fieldNums[len(fieldNums)-1]
	if col <= len(split) {
	    sb.WriteString(split[col-1])
	}
	sb.WriteString("\n")
    }
    fmt.Println(strings.TrimSuffix(sb.String(), "\n"))
}
