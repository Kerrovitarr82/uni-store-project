package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func countWords(str string) (map[string]int, error) {
	if len(str) == 0 {
		return map[string]int{}, errors.New("пустая строка")
	}

	wordMap := make(map[string]int)
	words := strings.Split(str, " ")
	for _, word := range words {
		wordMap[word]++
	}
	return wordMap, nil
}

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	fmt.Println("---------------------------")
	wordMap, err := countWords(str)
	if err != nil {
		fmt.Println(err)
	} else {
		for key, value := range wordMap {
			fmt.Printf("Слово %s: %d повтор(ов)\n", key, value)
		}
	}

}
