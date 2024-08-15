package main
import "fmt"
/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
	==============================================
	Паттерн "Стратегия" - поведенческий паттерн, позволяющий выбирать поведение
	алгоритма во время выполнения программы. При его использовании выделяются схожие
	алгоритмы, после чего их реализация выносится в отдельные классы

	"+"
	- Замена алгоритмов во время выполнения программы
	- Изоляция кода и данных алгоритмов

	"-"
	- Усложняет программу дополнительныи классами
	- Клиент должен знать о стратегиях, чтобы выбрать подходящую
*/
type Strategy interface {
    Execute() string
}

type Context struct {
    strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
    c.strategy = s
}

func (c *Context) ExecuteStrategy() string {
    return c.strategy.Execute()
}

type FindRoof struct {}

func (fr *FindRoof) Execute() string {
    return "Found roof"
}

type OpenUmbrella struct {}

func(ou *OpenUmbrella) Execute() string {
    return "Opened umbrella"
}

func main() {
    ctx := &Context{}
    ctx.SetStrategy(&OpenUmbrella{})
    fmt.Println(ctx.ExecuteStrategy(), "to hide from the rain")
    ctx.SetStrategy(&FindRoof{})
    fmt.Println(ctx.ExecuteStrategy(), "to hide from the rain")
}
