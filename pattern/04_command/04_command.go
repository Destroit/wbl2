package main
import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
	=============================================
	Паттерн "Комманда" - поведенческий паттерн проектирования, который инкапсулирует
	всю информацию для совершения запроса в объекте

	"+"
	- Поддержка операций отмены и повтора
	- Разъединение отправителя и получателя

	"-"
	- Приводит к созданию множества маленьких классов
*/

type Command interface {
    Execute()
}

type Car struct {
    ignition bool
}

func (c *Car) Ignite() {
    fmt.Println("Starting car")
    c.ignition = true
}

func (c *Car) Stop() {
    fmt.Println("Stopping car")
    c.ignition = false
}

type IgniteCommand struct {
    car *Car
}

func (ic *IgniteCommand) Execute() {
    ic.car.Ignite()
}

type StopCommand struct {
    car *Car
}

func (ic *StopCommand) Execute() {
    ic.car.Stop()
}

type RadioKey struct {
    command Command
}

func (rk *RadioKey) SetCommand(command Command) {
    rk.command = command
}

func (rk *RadioKey) Send() {
    rk.command.Execute()
}

func main() {
    car := &Car{}

    start := &IgniteCommand{car}
    stop := &StopCommand{car}

    rk := &RadioKey{}

    rk.SetCommand(start)
    rk.Send()
    
    rk.SetCommand(stop)
    rk.Send()

}
