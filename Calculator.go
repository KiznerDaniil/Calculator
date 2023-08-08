package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var RimNum map[string]int = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XX":   20,
	"XXX":  30,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
}

func arab_rim(x int) (A string) {
	var h int = x / 100 // СОТНИ
	var t int = x / 10  // ДЕСЯТКИ
	var ch int = x % t  // ЕДИНИЦЫ
	if h == 1 {
		A = "C"
		return A
	} else {
		t1 := loop(t * 10)
		ch1 := loop(ch)
		A = t1 + ch1
		return A
	}
}

func rim_arab(x, y string) (A, B int) {
	RimNum = make(map[string]int)
	A = RimNum[x]
	B = RimNum[y]
	return A, B
}

func main() {
	fmt.Println(operation(input_and_check()))
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

func loop(x int) (y string) {
	RimNum = make(map[string]int)
	for k, v := range RimNum {
		if v == x {
			y = k
		}
	}
	return
}

func input_and_check() (A, B int, C string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")
	for seq, _ := reader.ReadString('\n'); seq != "0"; seq, _ = reader.ReadString('\n') {
		seq = strings.TrimSpace(seq)
		var mass []string // МАССИВ ДАННЫХ
		mass = strings.Split(seq, " ")
		A, err1 := strconv.Atoi(mass[0])
		B, err2 := strconv.Atoi(mass[2])
		C = mass[1]
		if err1 == nil && err2 == nil {
			if A > 10 || B > 10 || A <= 0 || B <= 0 {
				err := fmt.Errorf("введённые числа не принадлежать промежутку [1; 10]")
				fmt.Println(err)
			} else {
				return A, B, C
			}
		} else {
			if check_rim(err1, err2) == nil {
				if v := check_rim_count(mass[0], mass[2]); v == nil {
					A, B = rim_arab(mass[0], mass[2])
					switch {
					case A > 10 || B > 10 || A <= 0 || B <= 0:
						err := fmt.Errorf("введённые числа не принадлежать промежутку [1; 10]")
						fmt.Println(err)
					case A-B <= 0:
						err := fmt.Errorf("итог вычислений в арабской системе исчисления не может быть меньше или равен нулю")
						fmt.Println(err)
					default:
						return A, B, C
					}
				} else {
					fmt.Print(v)
				}
			}
		}
		fmt.Println("Введите выражение:")
	}
	return
}

func check_rim(err1, err2 error) (err error) {
	switch {
	case err1 != nil && err2 != nil:
		err = nil
		return err
	case err1 != nil || err2 != nil:
		err = fmt.Errorf("ошибка: используются разные системы исчисления")
		return err
	}
	return
}

func check_rim_count(x, y string) (err error) {
	RimNum = make(map[string]int)
	_, err1 := RimNum[x]
	_, err2 := RimNum[y]
	if err1 && err2 {
		err = nil
		return err
	} else {
		err = fmt.Errorf("ошибка: для ввода доступны только цифры из арабской и римской системы исчисления в промежутке [1; 10]")
		return
	}
}
