package main

import (
	"bufio"
	"fmt"
	"go/types"
	"os"
	"strconv"
	"strings"
)

var RimNum map[string]int = map[string]int{
	"I" : 1,
	"II" : 2,
	"III" : 3,
	"IV" : 4,
	"V" : 5,
	"VI" : 6,
	"VII" : 7,
	"VIII" : 8,
	"IX" : 9,
	"X" : 10,
	"L" : 50,
	"C" : 100,
}

type Rim struct {
	num1 string
	num2 string
}

type Arab struct {
	num1 int
}

type I interface {
	arab_rim() string
	rim_arab() int
}

func (r Rim) arab_rim(x int) (A string) {
	RimNum = make(map[string]int)
	var h int = x / 100 // СОТНИ
	var t int = x / 10 // ДЕСЯТКИ
	var ch int = x % t
	if h == 1 {
		h1 := "C"
	} else {
		switch {
		case t == 5:
			t1 := "L"
		case t > 5:
			for k, v := range RimNum {
				if v == t - 5 {
					t1 = "L" + k
				}
			}
		case t < 5:
			for k, v := range RimNum {
				if v == t {
					t1 = "L" + k
				}
			}
	}

	}
	for k, v := range RimNum {
		if v == x {

		}
	}
}

func (a Arab) rim_arab(x, y string) (A, B int) {
	RimNum = make(map[string]int)
	*a.num1 = x
	*a.num2 = y
	A = RimNum[a.num1]
	B = RimNum[a.num2]
	return A, B
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
