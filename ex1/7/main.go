package main

import (
	"fmt"
	"strconv"
)

func toBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	binary := toBinary(n)
	fmt.Printf("Число %d в двоичной системе: %s\n", n, binary)
}
