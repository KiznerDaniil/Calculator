package main

import (
	"bufio"
	"fmt"
	"go/types"
	"os"
	"strconv"
	"strings"
)

type NumRim[string]int

type Numbers_arab struct {
	num1 int
	num2 int
}

type Numbers_arab struct {
	num1 string
	num2 string
}

type I interface {
	arab_rim() string
	rim_arab() int
}

func (n Numbers) arab_rim() (A, B string) {
	m
}

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
			if mass[]
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
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	}
	return
}
