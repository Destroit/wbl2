package main
import "fmt"
/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
	Паттерн "Цепочка вызовов" - это поведенческий паттерн, который передаёт запросы
	последовательно по цепи обработчиков. При получении запроса, кажды1 обработчик
	принимает решение - обработать ли запрос или передать следующему обработчику

	"+"
	- Уменьшает зависимость между клиентами и разработчиком
	- Добавляет гибкости в назначе

	"-"
	- Запрос может остаться никем не обработанным
	- Может привести к созданию сложной цепочки
*/
const (
    LOW int = 0
    MID int = 1
    HIGH int = 2
)

type Event struct {
    info string
    level int
}

type Handler interface {
    SetNext(Handler) Handler
    Handle(Event)
}

type BaseHandler struct {
    next Handler
}

func (bh *BaseHandler) SetNext(next Handler) Handler{
    bh.next = next
    return next
}

type LowAlarmHandler struct {
    BaseHandler
}

func (lah *LowAlarmHandler) Handle(ev Event) {
    fmt.Println("[!]", ev.info)
    if ev.level > LOW && lah.next != nil {
	lah.next.Handle(ev)
    }
}

type MidAlarmHandler struct {
    BaseHandler
}

func (mah *MidAlarmHandler) Handle(ev Event) {
    fmt.Println("[!!]", ev.info)
    if ev.level > MID && mah.next != nil {
	mah.next.Handle(ev)
    }
}

type HighAlarmHandler struct {
    BaseHandler
}

func (hah *HighAlarmHandler) Handle(ev Event) {
    fmt.Println("[!!!]", ev.info)
    if ev.level > HIGH && hah.next != nil {
	hah.next.Handle(ev)
    }
}

func main() {
    lowh := &LowAlarmHandler{}
    midh := &MidAlarmHandler{}
    highh := &HighAlarmHandler{}

    noise := &Event{"Strange noise", LOW} 
    smoke := &Event{"Smoke from the hall", MID}
    explosion := &Event{"Big bright mushroom", HIGH} 

    lowh.SetNext(midh).SetNext(highh)

    lowh.Handle(*noise)
    lowh.Handle(*smoke)
    lowh.Handle(*explosion)
}


