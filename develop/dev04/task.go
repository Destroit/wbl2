package main
import (
    "strings"
    "slices"
    "sort"
    "fmt"
)
/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func searchAnagram(words []string) map[string][]string {
    anagrams := make(map[string][]string)
    wordsSet := make(map[string]string)
    for _, word := range words {
	if len(word) <= 1 {
	    continue
	}
	r := []rune(strings.ToLower(word))
	//sorting letters in words with same letters gives same result
	sort.Slice(r, func(i, j int) bool {
	    return r[i] < r[j]
	})
	if _, ok := wordsSet[string(r)]; !ok {
	    wordsSet[string(r)] = word
	    anagrams[word] = append(anagrams[word], word)
	} else {
	    key := wordsSet[string(r)]
	    if _, ok := slices.BinarySearch(anagrams[key], word); !ok {
		anagrams[key] = append(anagrams[key], word)
	    }
	}
    }
    for _, words := range anagrams {
	    sort.Strings(words)
    }
    return anagrams
}

func main() {
    m :=  searchAnagram([]string{"тяпка", "пятка", "пятак", "листок", "слиток", "слиток", "a", "a", ""})
    fmt.Println(m)
}
