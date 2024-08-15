package main

import "fmt"
/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
	====================================================
	Паттерн "Фабричный метод" - это порождающий паттерн, который определяет 
	общий интерфейс создания объектов в родительском классе и позволяет 
	изменять создаваемые объекты в дочерних классах.

	"+"
	- Упрощает добавление новых продуктов в программу 
	- Избавляет слой создания объектов от конкретных классов продуктов

	"-"
	- Может привести к созданию больших параллельных иерархий классов, т.к. для
	каждого класса продукта нужно создать свой подкласс создателя
*/

type Gamepader interface {
    PrintName()
    PressButton()
    Vibrate()
}

type Gamepad struct {
    name string
}

func NewGamepad(objType string) (Gamepader, error) {
    switch objType {
    case "Dualshock":
	return NewDualshock(), nil
    case "XboxController":
	return NewXboxController(), nil
    default:
	return nil, fmt.Errorf("Unknown gamepad type")
    }
}


type Dualshock struct {
    Gamepad
}

func (d *Dualshock) PrintName() {
    fmt.Println("SONY", d.name)
}

func (d *Dualshock) PressButton() {
    fmt.Println("X O L1 R1")
}

func (d *Dualshock) Vibrate() {
    fmt.Println("Playstation brbrbrbrbrbr")
}

func NewDualshock() Gamepader {
    return &Dualshock{
	Gamepad: Gamepad{
	    name: "Dualshock",
	},
    }
}


type XboxController struct {
    Gamepad
}

func (x *XboxController) PrintName() {
    fmt.Println("MICROSOFT", x.name)
}

func (x *XboxController) PressButton() {
    fmt.Println("A B X Y LB RB")
}

func (x *XboxController) Vibrate() {
    fmt.Println("Xbox brrrrrbrrrrr")
}

func NewXboxController() Gamepader {
    return &XboxController{
	Gamepad: Gamepad{
	    name: "Xbox Controller",
	},
    }
}

func main() {
    gp1, _ := NewGamepad("Dualshock")
    gp1.PrintName()
    gp1.PressButton()
    gp1.Vibrate()
    fmt.Println("----")

    gp2, _ := NewGamepad("XboxController")
    gp2.PrintName()
    fmt.Println("----")

    _, err := NewGamepad("Cheese")
    fmt.Println(err)
}
