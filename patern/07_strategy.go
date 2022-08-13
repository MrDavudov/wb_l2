package main

import (
	"fmt"
	"strconv"
)

/*
Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Паттерн Strategy относится к поведенческим паттернам уровня объекта.

Паттерн Strategy определяет набор алгоритмов схожих по роду деятельности, инкапсулирует их в отдельный
класс и делает их подменяемыми. Паттерн Strategy позволяет подменять алгоритмы без участия клиентов,
которые используют эти алгоритмы.
*/

// StrategySort предоставляет интерфейс для алгоритмов сортировки.
type StrategySort interface {
	Sort([]int)
}

// BubbleSort реализует алгоритм пузырьковой сортировки.
type BubbleSort struct {
}

// Sort сортировка данных.
func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// InsertionSort реализует алгоритм сортировки вставками.
type InsertionSort struct {
}

// Sort сортировка данных.
func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

// Context обеспечивает контекст для выполнения стратегии.
type Context struct {
	strategy StrategySort
}

// Algorithm заменяет стратегии.
func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

// Sort сортирует данные в соответствии с выбранной стратегией.
func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}

func main() {
	data1 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}

	ctx := new(Context)

	ctx.Algorithm(&BubbleSort{})

	ctx.Sort(data1)

	var result1 string
	for _, val := range data1 {
		result1 += strconv.Itoa(val) + ","
	}
	fmt.Println(result1)
}