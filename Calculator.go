package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение:")
	seq, _ := reader.ReadString('\n') // ВХОДЯЩИЕ ДАННЫЕ
	var mass []string                 // МАССИВ ДАННЫХ
	mass = strings.Split(seq, " ")
	a, _ := strconv.Atoi(mass[0])
	b := uint(mass[1])
	fmt.Println(mass, a, b)
	operation(a, b)
}

func operation(X, Y int) int {
	var c int
	c = X + Y
	fmt.Println(X + Y)
	return c
}
