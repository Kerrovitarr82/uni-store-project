package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findSubstring(s string) (longestWord string) {
	str := strings.Split(s, " ")
	for _, v := range str {
		if len(v) > len(longestWord) {
			longestWord = v
		}
	}
	return longestWord
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите предложение: ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	longestWord := findSubstring(str)

	fmt.Printf("Самое длинное слово в предложении: %s\n", longestWord)
}
