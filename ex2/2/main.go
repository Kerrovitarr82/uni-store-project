package main

import (
	"fmt"
)

func isPol(str string) bool {
	revStr := ""
	for _, v := range str {
		revStr = string(v) + revStr
	}
	if str == revStr {
		return true
	}
	return false
}

func main() {
	str := ""
	fmt.Print("Введите слово: ")
	_, err := fmt.Scanf("%s", &str)
	if err != nil {
		return
	}
	if isPol(str) {
		fmt.Printf("%v - палиндром\n", str)
	} else {
		fmt.Printf("%v - не палиндром", str)
	}
}
