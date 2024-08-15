package main
import "fmt"
/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
	=============================================
	Паттерн "Строитель" -  порождаюий паттерн проектирования, позволяющий поэтапно 
	создавать сложные объекты, и изолирует создание объекта от его представления

	"+"
	- Позволяет варьировать внутреннее представление объекта
	- Изолирует код для конструкции и представления
	- Позволяет контролировать процесс конструкции

	"-"
	- для каждого типа Product нужен свой отдельный ConcreteBuilder
*/

type Car struct {
    model string
    gearbox string
    engine string
    maxspeed int
    horsepower int
}

func (c *Car) Stats() {
    fmt.Println("[Model]:", c.model)
    fmt.Println("=========================")
    fmt.Println("(Engine):\t", c.engine)
    fmt.Println("(Gearbox):\t", c.gearbox)
    fmt.Println("{Max Speed}:\t", c.maxspeed)
    fmt.Println("{Horsepower}:\t", c.horsepower)
}

type CarBuilder struct {
    model string
    gearbox string
    engine string
    maxspeed int
    horsepower int
}

func (cb *CarBuilder) SetModel(model string) *CarBuilder {
    cb.model = model
    return cb
}

func (cb *CarBuilder) SetEngine(engine string) *CarBuilder {
    cb.engine = engine
    return cb
}

func (cb *CarBuilder) SetGearbox(gearbox string) *CarBuilder {
    cb.gearbox = gearbox
    return cb
}

func (cb *CarBuilder) SetMaxspeed(maxspeed int) *CarBuilder {
    cb.maxspeed = maxspeed
    return cb
}

func (cb *CarBuilder) SetHorsepower(horsepower int) *CarBuilder {
    cb.horsepower = horsepower
    return cb
}

func (cb *CarBuilder) Build() *Car {
    return &Car {
	model: cb.model,
	engine: cb.engine,
	gearbox: cb.gearbox,
	maxspeed: cb.maxspeed,
	horsepower: cb.horsepower,
    }
}

func main() {
    builder := &CarBuilder{}
    thiscar := builder.SetModel("Honda Civic Type R").SetEngine("B16B(Petrol)").SetGearbox("S4C(Manual)").SetHorsepower(185).SetMaxspeed(225).Build()
    thiscar.Stats()
}
