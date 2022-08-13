package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Паттерн Command позволяет представить запрос в виде объекта. Из этого следует, что команда - это объект.
Такие запросы, например, можно ставить в очередь, отменять или возобновлять.

В этом паттерне мы оперируем следующими понятиями: Command - запрос в виде объекта на выполнение;
Receiver - объект-получатель запроса, который будет обрабатывать нашу команду; Invoker - объект-инициатор запроса.

Паттерн Command отделяет объект, инициирующий операцию, от объекта, который знает, как ее выполнить.
Единственное, что должен знать инициатор, это как отправить команду.

Требуется для реализации:
- Базовый абстрактный класс Command описывающий интерфейс команды;
- Класс ConcreteCommand, реализующий команду;
- Класс Invoker, реализующий инициатора, записывающий команду и провоцирующий её выполнение;
- Класс Receiver, реализующий получателя и имеющий набор действий, которые команда можем запрашивать;
*/

// предоставляет командный интерфейс.
type Command interface {
	Execute() string
}

// ToggleOnCommand реализация командного интерфейса
type ToggleOnCommand struct {
	receiver *Receiver
}
// метод Execute команда.
func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

// ToggleOffCommand реализация командного интерфейса
type ToggleOffCommand struct {
	receiver *Receiver
}
// метод Execute команда.
func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

// Receiver реализация пустой структуры.
type Receiver struct {
}
// ToggleOn реализация включения.
func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}
// ToggleOff реализация выключения.
func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

// Invoker список команд для реализации.
type Invoker struct {
	commands []Command
}
// метод StoreCommand добавления команды в список.
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// UnStoreCommand удаления последней команды из списка.
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Execute выполения всех команд в списке
func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}

func main() {
	// добавления команд
	invoker := &Invoker{}
	receiver := &Receiver{}

	// добавления команд в список
	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})

	// удаления команды
	invoker.UnStoreCommand()

	res := invoker.Execute()
	fmt.Println(res)
}