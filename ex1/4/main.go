package main

import (
	"fmt"
)

func concat(str [4]string) string {
	return str[0] + " " + str[1] + " " + str[2] + " " + str[3]

}

func main() {
	str := [4]string{}
	fmt.Println("Введите 4 слова через пробел")
	for i, _ := range str {
		_, err := fmt.Scanf("%s", &str[i])
		if err != nil {
			return
		}
	}

	fmt.Printf("Итоговая строка %v", concat(str))
}
