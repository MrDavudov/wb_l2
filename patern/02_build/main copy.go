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

// Создания акбстрактного класса с помощью интерфейса
type CardBuilder interface {
	SetCard()
	SetBalance()
	SetUser(str string)
	GetCard()
}

// Factory реализует менеджера
type Factory struct {
	Factory CardBuilder
}

// Construct сообщает строителю, что делать и в каком порядке.
func (f Factory) CreateCard() Card {
	f.Factory.SetCard()
	f.Factory.SetBalance()
	f.Factory.SetUser("Udfdf")
	return f.Factory.GetCard()
}

// DebitCard реализует интерфейс CardBuilder.
type DebitCard struct {
	Card 	string
	Balance	int
	User	string
}

// SetCard создает дебитовую карту.
func (d *DebitCard) SetCard() {
	d.Card = "debit"
}

// SetBalance указывает баланс карты.
func (d *DebitCard) SetBalance(str string) {
	d.Balance = 0
}

// SetUser указывает владельца карты.
func (d *DebitCard) SetUser(str string) {
	d.User = str
}

func (d *DebitCard) GetCard() Card {
	return Card{
		Card: d.Card,
		Balance: d.Balance,
		User: d.User,
	}
}

// Реализация параметры карты.
type Card struct {
	Card 	string
	Balance	int
	User	string
}

// Выводит информацию о выпушенной карте
func (card *Card) Show() {
	fmt.Printf("Card: [%s]\nBalance: [%d]\nUser: [%s]", card.Card, card.Balance, card.User)
}

func main() {
	debitcard := Factory.CreateCard()
	debitcard.Show()
}
