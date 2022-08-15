package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Поиск анаграмм по словарю

Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
- Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
- Выходные данные: ссылка на мапу множеств анаграмм
- Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
- Массив должен быть отсортирован по возрастанию.
- Множества из одного элемента не должны попасть в результат. 
- Все слова должны быть приведены к нижнему регистру.
- В результате каждое слово должно встречаться только один раз.

*/

// сравнения двух строк является одна анаграмой другой
func сomparison(str1, str2 string) bool {
	if len(str1) != len(str2) || str1 == str2 {
		return false
	}

	// конвертируем строку в массив рун для сортировки по алфавиту
	arr := []rune(str1)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	arr2 := []rune(str2)
	sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })

	// сравниваем
	if string(arr) == string(arr2) {
		return true
	}

	return false
}

func toLower(data []string) []string {
	for _, v := range data {
		v = strings.ToLower(v)
	}

	return data
}

func anagram(data []string) map[string][]string {
	result := make(map[string][]string)
	exist := []string{}
	// функция для приведения всех слов в нижний регистр
	exist = toLower(data)
	
	for i, v := range exist {
		for j:=0;j<len(exist);j++ {
			if v == exist[j] {
				continue
			}
			if сomparison(v, exist[j]) {
				result[v] = append(result[v], exist[j])
				exist = append(exist[:j], exist[j+1:]...)
				j--
			}
		}
		if len(exist) == 1 {
			break
		} else {
			// сортировка по возростанию
			sort.Strings(result[v])
			exist = append(exist[:i], exist[i+1:]...)
		}
	}
	
	return result
}

func main() {
	data := []string{"пятак", "тяпка", "пятка", "листок", "слиток", "столик"}
	
	fmt.Println(anagram(data))
}