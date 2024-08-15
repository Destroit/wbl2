Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

err содержит значение nil, поэтому выводится <nil>.
Интерфейс равен nil, если его тип и значение равны nil, поэтому err == nil выводит false

Обычный интерфейс использует следующую структуру:
type iface struct { // 16 bytes on a 64bit arch
    tab  *itab
    data unsafe.Pointer
}
гдe itab содержит информацию об интерфейсе, такую как: метаданные интерфейса, тип хранимого значения, хэш, и список методов, удовлетворяющих интерфейсу. data содержит значение.

Пустой же использует другую структуру:
type eface struct { // 16 bytes on a 64bit arch
    _type *_type
    data  unsafe.Pointer
}
вместо itab присутствует только информация о типе хранимого значения
```
