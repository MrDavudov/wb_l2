package main

import (
	"fmt"
)

/*
Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн Builder относится к порождающим паттернам уровня объекта.

Паттерн Builder определяет процесс поэтапного построения сложного продукта. После того 
как будет построена последняя его часть, продукт можно использовать.

Плюсы
- Позволяет создавать действия последовательно
- Позволяет использовать один тот же код для различных обьектов
- Изолирует сложный код в сборке обьекта и его основной бизнес логики
Минусы
- Усложняет код из за введения дополнительных классов
- Клиент будет привязан конкретному обьекту строителся, так как в интерфейсе может не быть
	какого то метода и тогда ему придется добавить его
*/

// Создания основных методов
type CardInteface interface {
	SetCard()
	SetBalance()
	SetUser()
	GetCard()
}
// Определения какая карта будет выпускаться
func GetCard(card string) CardInteface {
	switch card {
	default:
		return nil
	case "debit":
		return &DebitCard{}
	case "credit":
		return &CreditCard{}
	}
	return nil
}

// Параметры карты
type Card struct {
	Card        string
	Balance		int
	User		string
}
// Метод для вывода полученной карты
func (card *Card) Print() {
	fmt.Printf("Card: [%s]\nBalance: [%d]\nUser: [%d]", 
	card.Card, 
	card.Balance, 
	card.User)
}

// Создания дебитовая карты
type DebitCard struct {
	Card        string
	Balance		int
	User		string
}
func (card *DebitCard) SetCard() {
	card.Card = "debit"
}
func (card *DebitCard) SetBalance() {
	card.Balance = 0
}
func (card *DebitCard) SetUser(user string) {
	card.User = user
}
func (card *DebitCard) GetCard() Card {
	return Card{
		Card: card.Card,
		Balance: card.Balance,
		User: card.User,
	}
}

// Создания кредитной карты
type CreditCard struct {
	Card        string
	Balance		int
	User		string
}
func (card *CreditCard) SetCard() {
	card.Card = "credit"
}
func (card *CreditCard) SetBalance() {
	card.Balance = 100000
}
func (card *CreditCard) SetUser(user string) {
	card.User = user
}
func (card *CreditCard) GetCard() Card {
	return Card{
		Card: card.Card,
		Balance: card.Balance,
		User: card.User,
	}
}

// Создания карты
type Factory struct {
	Factory CardInteface
}

func NewFactory(card CardInteface) *Factory {
	return &Factory{Factory: card}
}

func (factory Factory) SetCard(card CardInteface) {
	factory.Factory = card
}

func (factory Factory) CreateCard(user string) Card {
	factory.Factory.SetCard()
	factory.Factory.SetBalance()
	factory.Factory.SetUser(user)
	return factory.Factory.GetCard()
}


func main() {
	debit := GetCard("debit")
	credit := GetCard("credit")
	factory := NewFactory(debit)
	debitCard := factory.CreateCard("Name")
	debitCard.Print()

	factory.SetCard(credit)
	creditCard := factory.CreateCard("Name")
	creditCard.Print()
}