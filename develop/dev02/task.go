package main
import (
    "fmt"
    "unicode"
    "strconv"
)
/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func unpack(r []rune) ([]rune, error) {
    if len(r) == 0 {
	return []rune(""), nil
    }
    if unicode.IsDigit(r[0]) {
	return []rune(""), fmt.Errorf("unpacker: passed []rune starts with digit")
    }

    result := make([]rune, 0)
    escape := false
    var tmpch rune

    for i:=0; i < len(r); i++ {
	// slash processing
	if r[i] == '/' {
	    if escape {
		result = append(result, '/')
		escape = false
	    } else {
		escape = true
	    }
	// escape processing
	} else if escape {
	    result = append(result, r[i])
	    escape = false
	// digits processing
	} else if unicode.IsDigit(r[i]) {
	    count := make([]rune, 0)
	    for ; i < len(r) && unicode.IsDigit(r[i]); i++ {
		count = append(count, r[i])
	    }

	    n, _ := strconv.Atoi(string(count))
	    if n == 0 {
		result = result[:len(result)-1]
	    }
	    for j:=0; j < n-1; j++ {
		result = append(result, result[len(result)-1])
	    } 
	    i--
	} else {
	    tmpch = r[i]
	    result = append(result, tmpch)
	}
    }
    return result, nil
}

func main() {
    tr := []rune("ab3/4//c5")
    res, err := unpack(tr)
    if err != nil {
	panic(err)
    }
    fmt.Printf("%s\n", string(res))
}
