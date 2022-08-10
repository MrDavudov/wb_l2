package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

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