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
	var ch int          // ЕДИНИЦЫ
	if t == 0 {
		ch = x
	} else {
		ch = x % (t * 10)
	}
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
	A = RimNum[x]
	B = RimNum[y]
	return A, B
}

func main() {
	a, b, c, T, err := input_and_check()
	if err == nil {
		I := operation(a, b, c)
		if T == "Arab" {
			fmt.Println(I)
		} else {
			if I <= 0 {
				err = fmt.Errorf("ошибка: итог вычислений между римскими числами не может быть меньше или равен нулю")
				fmt.Println(err)
			} else {
				fmt.Println(arab_rim(I))
			}
		}
	} else {
		fmt.Println(err)
	}

}

func operation(x, y int, z string) (itog int) {
	switch z {
	case "+":
		itog = x + y
	case "-":
		itog = x - y
	case "*":
		itog = x * y
	case "/":
		itog = x / y
	}
	return
}

func loop(x int) (y string) {
	for k, v := range RimNum {
		if v == x {
			y = k
		}
	}
	return
}

func input_and_check() (A, B int, C, T string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")
	seq, _ := reader.ReadString('\n')
	seq = strings.TrimSpace(seq)
	var mass []string // МАССИВ ДАННЫХ
	mass = strings.Split(seq, " ")
	switch {
	case len(mass) > 3:
		err = fmt.Errorf("ошибка: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	case len(mass) < 3:
		err = fmt.Errorf("ошибка: введённая строка не является математической операцией")
		return
	}
	A, err1 := strconv.Atoi(mass[0])
	B, err2 := strconv.Atoi(mass[2])
	C = mass[1]
	if err1 == nil && err2 == nil {
		T = "Arab"
		if A > 10 || B > 10 || A <= 0 || B <= 0 {
			err = fmt.Errorf("ошибка: введённые числа не принадлежат промежутку [1; 10]")
		} else {
			return A, B, C, T, err
		}
	} else {
		if err = check_rim(err1, err2); err == nil {
			T = "Rim"
			if err, A, B := check_rim_count(mass[0], mass[2]); err == nil {
				if A > 10 || B > 10 || A <= 0 || B <= 0 {
					err = fmt.Errorf("ошибка: введённые числа не принадлежать промежутку [1; 10]")
				} else {
					return A, B, C, T, err
				}
			} else {
				return A, B, C, T, err
			}
		}
	}
	return
}

func check_rim(err1, err2 error) (err error) {
	switch {
	case err1 != nil && err2 != nil:
		err = nil
		return err
	case (err1 == nil && err2 != nil) || (err1 != nil && err2 == nil):
		err = fmt.Errorf("ошибка: используются разные системы исчисления")
		break
		return err
	default:
		err = fmt.Errorf("ошибка: введены неизвестные значения, не являющиеся числами")
		return err
	}
	return
}

func check_rim_count(x, y string) (err error, num1, num2 int) {
	num1, err1 := RimNum[x]
	num2, err2 := RimNum[y]
	switch {
	case num1 > 10 || num2 > 10:
		err = fmt.Errorf("ошибка: для ввода доступны только цифры из арабской и римской системы исчисления в промежутке [1; 10]")
		return
	case err1 && err2:
		err = nil
		return
	default:
		err = fmt.Errorf("ошибка: для ввода доступны только цифры из арабской и римской системы исчисления в промежутке [1; 10]")
		return
	}
}
