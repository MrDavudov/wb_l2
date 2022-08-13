package main

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Паттерн Factory Method относится к порождающим паттернам уровня класса и сфокусирован только на отношениях между классами.

Паттерн Factory Method полезен, когда система должна оставаться легко расширяемой путем добавления объектов новых типов.
Этот паттерн является основой для всех порождающих паттернов и может легко трансформироваться под нужды системы. По этому,
если перед разработчиком стоят не четкие требования для продукта или не ясен способ организации взаимодействия между продуктами,
то для начала можно воспользоваться паттерном Factory Method, пока полностью не сформируются все требования.

Паттерн Factory Method применяется для создания объектов с определенным интерфейсом, реализации которого предоставляются потомками.
Другими словами, есть базовый абстрактный класс фабрики, который говорит, что каждая его наследующая фабрика должна реализовать
такой-то метод для создания своих продуктов.


*/

import (
	"fmt"
	"log"
)

// action помогает клиентам узнать доступные действия.
type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

// Creator предоставляет фабричный интерфейс.
type Creator interface {
	CreateProduct(action action) Product // фабричный метод
}

// Product обеспечивает интерфейс продукта.
// Все продукты, возвращаемые заводом-изготовителем, должны иметь единый интерфейс.
type Product interface {
	Use() string // Каждый продукт должен быть полезным
}

// ConcreteCreator реализует интерфейс Creator.
type ConcreteCreator struct{}

// NewCreator является конструктором ConcreteCreator.
func NewCreator() Creator {
	return &ConcreteCreator{}
}

// CreateProduct является фабричным методом.
func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

// ConcreteProductA реализация продукта "A".
type ConcreteProductA struct {
	action string
}

// Используйте действие возврата продукта.
func (p *ConcreteProductA) Use() string {
	return p.action
}

// ConcreteProductB реализация продукта "B".
type ConcreteProductB struct {
	action string
}

// Используйте действие возврата продукта.
func (p *ConcreteProductB) Use() string {
	return p.action
}

func main() {
	factory := NewCreator()
	products := []Product{
		factory.CreateProduct(A),
		factory.CreateProduct(B),
	}

	for _, v := range products {
		fmt.Println(v)
	}

}