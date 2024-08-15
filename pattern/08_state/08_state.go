package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
    Show()
}

type Lights struct {
    state State
}

type LightsOn struct{}

func (l *LightsOn) Show() {
    fmt.Println("Lights ON")
}

type LightsOff struct{}

func (l *LightsOff) Show() {
    fmt.Println("Lights OFF")
}

func main() {
    lrm := Lights{}
    lrm.state = &LightsOn{}
    lrm.state.Show()
    lrm.state = &LightsOff{}
    lrm.state.Show()
}
