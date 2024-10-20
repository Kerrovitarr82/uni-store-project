package main

import (
	"fmt"
	"math/rand"
)

func guessTheNumber(min, max, n int) (bool, error) {
	num := 0
	hiddenNum := rand.Intn(max-min+1) + min
	fmt.Println(hiddenNum)
	for i := 0; i < n; i++ {
		fmt.Println("---------------------------")
		fmt.Printf("%v-ая попытка: ", i+1)
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			return false, err
		}
		if num == hiddenNum {
			return true, err
		} else {
			fmt.Println("Неверно!")
		}
	}
	return false, nil
}

func main() {
	min := 1
	max := 100
	n := 5
	fmt.Printf("Угадай число от %v до %v! У тебя есть %v попыток!\n", min, max, n)
	guess, err := guessTheNumber(min, max, n)
	fmt.Println("---------------------------")
	if err != nil {
		fmt.Println(err)
	} else if guess {
		fmt.Println("Ты угадал!")
	} else {
		fmt.Println("К сожалению, ты не угадал.")
	}
}
