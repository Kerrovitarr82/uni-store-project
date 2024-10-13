package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findSubstring(s, sub string) int {
	runesS := []rune(s)
	runesSub := []rune(sub)

	lenS := len(runesS)
	lenSub := len(runesSub)

	if lenSub == 0 || lenSub > lenS {
		return -1
	}

	for i := 0; i <= lenS-lenSub; i++ {
		match := true
		for j := 0; j < lenSub; j++ {
			if runesS[i+j] != runesSub[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите основную строку: ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	fmt.Print("Введите подстроку для поиска: ")
	substr, _ := reader.ReadString('\n')
	substr = strings.TrimSpace(substr)

	position := findSubstring(str, substr)

	if position != -1 {
		fmt.Printf("Подстрока найдена на позиции: %d\n", position)
	} else {
		fmt.Println("Подстрока не найдена")
	}
}
