package main
import "fmt"
/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
	=============================================
	Паттерн "Посетитель" - поведенческий паттерн проектирования, 
	позволяющий добавить операции над объектами,
	не прибегая к изменению их классов

	"+"
	- Упрощается добавление новых операций
	- Объединение родственных операций в Visitor
	- Расширяемость

	"-"
	- Код становится более усложнённым
*/

type Visitor interface {
    VisitTV(t *TV)
    VisitDisplay(d *Display)
}

type TV struct {
    model string
}

func (t *TV) Accept(v Visitor) {
    v.VisitTV(t)
}   

type Display struct {
    model string
}

func (d *Display) Accept(v Visitor) {
    v.VisitDisplay(d)
}   

type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) VisitTV(t *TV) {
    fmt.Println("Visited", t.model, "from TV")
}

func (cv *ConcreteVisitor) VisitDisplay(d *Display) {
    fmt.Println("Visited", d.model, "from Display")
}

func main() {
    tv := &TV{"DEXP"}
    disp := &Display{"Dell"}

    visitor := &ConcreteVisitor{}

    tv.Accept(visitor)
    disp.Accept(visitor)
}
