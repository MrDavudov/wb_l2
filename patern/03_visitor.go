package main

import "fmt"

/*
Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Паттерн Visitor относится к поведенческим паттернам уровня объекта.

Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными
интерфейсами, а также позволяет добавить новый метод в класс объекта, при этом,
не изменяя сам класс этого объекта.

Требуется для реализации:
Абстрактный класс Visitor, описывающий интерфейс визитера;
Класс ConcreteVisitor, реализующий конкретного визитера. Реализует методы для обхода конкретного элемента;
Класс ObjectStructure, реализующий структуру(коллекцию), в которой хранятся элементы для обхода;
Абстрактный класс Element, реализующий интерфейс элементов структуры;
Класс ElementA, реализующий элемент структуры;
Класс ElementB, реализующий элемент структуры.
*/

// Visitor предоставляет интерфейс посетителя.
type Visitor interface {
	VisitTinkoff(p *Tinkoff) string
	VisitSber(p *Sber) string
}

// CardInterface предоставляет интерфейс для места, которое должен посетить посетитель.
type Card interface {
	Accept(v Visitor) string
}

// People реализуют интерфейс посетителя.
type People struct {
}

// метод VisitTinkoff осуществляет посещение Tinkoff.
func (v *People) VisitTinkoff(p *Tinkoff) string {
	return p.BuyTinkoff()
}

// метод VisitSber осуществляет посещение Sber.
func (v *People) VisitSber(p *Sber) string {
	return p.BuySber()
}

// Bank реализует коллекцию мест для посещения банков.
type Bank struct {
	Cards []Card
}

// метод Add добавляет место в коллекцию.
func (c *Bank) Add(p Card) {
	c.Cards = append(c.Cards, p)
}

// метод Accept показывает какие карты были приобретены
func (c *Bank) Accept(v Visitor) string {
	var result string
	for _, p := range c.Cards {
		result += p.Accept(v)
	}
	return result
}

// Tinkoff реализует интерфейс Card.
type Tinkoff struct {
}

// Accept реализация визита в банк.
func (s *Tinkoff) Accept(v Visitor) string {
	return v.VisitTinkoff(s)
}

// BuyTinkoff реализация получения карты.
func (s *Tinkoff) BuyTinkoff() string {
	return "get a card tinkoff...\n"
}

// Sber реализует интерфейс Card.
type Sber struct {
}

// Accept реализация визита в банк.
func (p *Sber) Accept(v Visitor) string {
	return v.VisitSber(p)
}

// BuySber реализация получения карты.
func (p *Sber) BuySber() string {
	return "get a card sber...\n"
}


func main() {
	// создаем обьект city
	city := new(Bank)

	// добавляем карты которые можно получить
	city.Add(&Tinkoff{})
	city.Add(&Sber{})

	// перебираем 
	res := city.Accept(&People{})

	fmt.Println(res)
}