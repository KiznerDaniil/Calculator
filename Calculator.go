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
	seq = strings.TrimSpace(seq)
	var mass []string                 // МАССИВ ДАННЫХ
	mass = strings.Split(seq, " ")
	a, _ := strconv.Atoi(mass[0])
	b, _ := strconv.Atoi(mass[2])
	fmt.Println(a + b)
}