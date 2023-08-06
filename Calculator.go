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
	var seq = "1"
	fmt.Println("Введите выражение:")
	seq, _ = reader.ReadString('\n')
	for {
		if seq == "0" {
			fmt.Println("Спасибо за использование данного калькулятора, до свидания!")
		} else {
			seq = strings.TrimSpace(seq)
			var mass []string // МАССИВ ДАННЫХ
			mass = strings.Split(seq, " ")
			a, _ := strconv.Atoi(mass[0])
			b, _ := strconv.Atoi(mass[2])
			c := mass[1]
			fmt.Println(operation(a, b, c))
			fmt.Println("Введите выражение:")
			seq, _ = reader.ReadString('\n')
		}
	}
}

func operation(x, y int, z string) (itog int) {
	switch z {
	case "+":
		return x + y
	case "-":
		itog = x - y
	case "*":
		itog = x * y
	case "/":
		itog = x / y
	}
	return
}
