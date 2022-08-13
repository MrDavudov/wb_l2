package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Паттерн State относится к поведенческим паттернам уровня объекта.

Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния и является
объектно-ориентированной реализацией конечного автомата. Поведение объекта изменяется настолько, что создается
впечатление, будто изменился класс объекта.
*/

// MobileAlertStater обеспечивает общий интерфейс для различных состояний.
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert реализует предупреждение в зависимости от его состояния.
type MobileAlert struct {
	state MobileAlertStater
}

// Alert возвращает строку предупреждения
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState изменяет состояние
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// NewMobileAlert является конструктором MobileAlert.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration реализует виброзвонок
type MobileAlertVibration struct {
}

// Alert возвращает строку предупреждения
func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr...\n"
}

// MobileAlertSong реализует звуковой сигнал
type MobileAlertSong struct {
}

// Alert возвращает строку предупреждения
func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы...\n"
}

func main() {
	mobile := NewMobileAlert()

	result := mobile.Alert()
	result += mobile.Alert()

	mobile.SetState(&MobileAlertSong{})

	result += mobile.Alert()
	fmt.Println(result)
}