package main

import (
	"strings"
	"fmt"
	"strconv"
)

const num = "0123456789"

func main() {
	arr := []string { "a4bc2d5e", "abcd", "45", "", "3a3"}
	fmt.Println(unpacking(arr[4]))
	
}

func unpacking(str string) string {
	var res string
	if "" == str || strings.Contains(num, string(str[0])) {
		return fmt.Sprintf("Некоректное значение")
	}
	
	for i := range str {
		intNum, _ := strconv.Atoi(string(str[i]))
		if i >= 1 && strings.Contains(num, string(str[i])) {
			if strings.Contains(num, string(str[i])) == strings.Contains(num, string(str[i-1])) {
				return fmt.Sprintf("Некоректное значение")
			} 
			res += strings.Repeat(string(str[i-1]), intNum)
		} else {
			res += string(str[i])
		}
	}
	return res
}