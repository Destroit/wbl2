package main
import "fmt"
/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
	============================================
	Паттерн "Фасад" - Структурный паттерн, предоставляющий унифицированный интерфейс
	над сложной подсистемой интерфейсов. 

	"+"
	- Упрощает использование подсистемы
	- Изоляция использованя и реализации

	"-"
	- Может быть лишней, если подсистема достаточно проста для прямого использованяи
	- Дополнительный слой абстракции иожет усложнить понимание кода
*/
type AlarmClock struct {}

func (ac *AlarmClock) MakeBeep() {
    fmt.Println("[AlarmClock]: BEEP! BEEP! BEEP!")
}

type CoffeeMachine struct{}

func (cm *CoffeeMachine) MakeCoffee() {
    fmt.Println("[CoffeeMachine]: Making coffee")
}

type LightController struct{}

func (lc *LightController) TurnOnLights() {
    fmt.Println("[LightController]: Lights on")
}

type Radio struct{}
func (rd *Radio) TuneStation() {
    fmt.Println("[Radio]: Tuning to 222.22 FM")
}

type SmartHomeFacade struct {
    alarmClock *AlarmClock
    lightController *LightController
    coffeeMachine *CoffeeMachine
    radio *Radio
}

func (sh *SmartHomeFacade) MorningScenario() {
    sh.alarmClock.MakeBeep()
    fmt.Println("[SmartHome]: Good morning!")
    sh.lightController.TurnOnLights()
    sh.coffeeMachine.MakeCoffee()
    sh.radio.TuneStation()
}

func main() {
    f := &SmartHomeFacade{&AlarmClock{}, &LightController{}, &CoffeeMachine{}, &Radio{}}
    f.MorningScenario()
}
