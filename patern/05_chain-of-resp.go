package main

import "fmt"

/*
Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Паттерн Chain Of Responsibility(цепочка вызововов) относится к поведенческим паттернам уровня объекта.

Паттерн Chain Of Responsibility(цепочка вызововов) позволяет избежать привязки объекта-отправителя запроса к 
объекту-получателю запроса, при этом давая шанс обработать этот запрос нескольким объектам. Получатели связываются 
в цепочку, и запрос передается по цепочке, пока не будет обработан каким-то объектом.

По сути это цепочка обработчиков, которые по очереди получают запрос, а затем решают, обрабатывать его или нет. 
Если запрос не обработан, то он передается дальше по цепочке. Если же он обработан, то паттерн сам решает передавать 
его дальше или нет. Если запрос не обработан ни одним обработчиком, то он просто теряется.

Плюсы
- Уменьшает зависимость между клиентов и обработчиком(плэтому можно изменить его в дальнейшем)
- Реализует принцип единсвенной обязаности(каждый сервис выполняет свою роль)
- Реализует принцип закрытости и открытости
Минусы
- Запрос может остаться не обработаным из за нарушения логики
 */

// Handler предоставляет интерфейс обработчика.
type Handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerA реализация обработкчика "A".
type ConcreteHandlerA struct {
	next Handler
}

// SendRequest реализация отправки запроса.
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerB реализация обработкчика "B".
type ConcreteHandlerB struct {
	next Handler
}

// SendRequest реализация отправки запроса.
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerC реализация обработкчика "C".
type ConcreteHandlerC struct {
	next Handler
}

// SendRequest реализация отправки запроса.
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

func main() {
	// создания обекта обработчиков
	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}

	res := handlers.SendRequest(1)
	fmt.Println(res)

}