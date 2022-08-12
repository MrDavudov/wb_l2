package main

/*
Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"errors"
	"fmt"
	"time"
)

// структуру покупателя
type User struct {
	Name	string
	Card	*Card
}
// запрос на получения баланса карты у покупателя
func (user User) GetBalance() float64 {
	return user.Card.Balance
}

// структура банковских карт
type Card struct {
	Name	string
	Balance	float64
	Bank	*Bank
}
// проверяем баланс карты
func (card Card) CheckBalance() error {
	fmt.Println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 500)

	return card.Bank.CheckBalance(card.Name)
}

// структура банк с именем и количеством карт
type Bank struct {
	Name	string
	Cards	[]Card
}
// получения от банка баланса по карте
func (bank Bank) CheckBalance(cardNum string) error {
	fmt.Printf("[Банк] Получения остатка по карте %s", cardNum)
	time.Sleep(time.Millisecond * 500)
	for _, card := range bank.Cards {
		if card.Name != cardNum {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}
	println("[Банк] Остаток положительный")
	return nil
}

// структура магазин
type Shop struct {
	Name		string
	Products	map[string]float64
}
// товары которые продаются в магазине
var product = map[string]float64 {
	"Mercedes-Benz": 87000,
	"BMW": 64000,
}
// продажа продукта 
func (shop Shop) Sell(user User, prod string) error {
	fmt.Println("[Магазин] Запрос пользователю, для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка может ли [%s] купить товар!\n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for name, price := range product {
		if name != prod {
			continue
		}
		if price >= user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара!")
		}
		fmt.Printf("[Магазин] Товар [%s] - куплен!", name)
	}
	return nil
}

// добавляем данные для проверки продажи товара покупателю
var (
	bank = Bank{
		Name:	"Банк",
		Cards:	[]Card{},
	}
	shop = Shop{
		Name: "Shop",
		Products: product,
	}
	card1 = Card{
		Name: "CRD-1",
		Balance: 70000,
		Bank: &bank,
	}
	user = User{
		Name: "Покупатель",
		Card: &card1,
	}
)


func main() {
	// добавим карту в структуру банка
	fmt.Println("[Банк] Выпуск карты")
	time.Sleep(time.Millisecond * 500)
	bank.Cards = append(bank.Cards, card1)
	fmt.Printf("[%s]\n", user.Name)

	// продажа товара Mercedes-Benz
	err := shop.Sell(user, "Mercedes-Benz")
	if err != nil {
		fmt.Println(err.Error())
	}
	// продажа товара BMW
	err = shop.Sell(user, "BMW")
	if err != nil {
		fmt.Println(err.Error())
	}
}