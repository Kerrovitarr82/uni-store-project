package main

import (
	"fmt"
)

func strReverse(str string) (revStr string) {
	for _, v := range str {
		revStr = string(v) + revStr
	}
	return revStr
}

func main() {
	str := ""
	fmt.Print("Введите слово: ")
	_, err := fmt.Scanf("%s", &str)
	if err != nil {
		return
	}
	fmt.Printf("%v - перевернутая строка\n", strReverse(str))
}
