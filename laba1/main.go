package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

var field = 6

func main() {
	var (
		items64 = []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
			16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
			32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
			48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63,
		}
		seq       []int
		seqBinary [][]int
		seqF      [][]int
		seqTemp   []int
		/*test      = []int{
			3, 5, 15, 12, 8, 0, 4, 14, 10, 6, 1, 11, 9, 13, 2, 7,
		}*/
		members = []string{
			"1", "x1", "x2", "x1x2", "x3", "x1x3", "x2x3", "x1x2x3", "x4", "x1x4", "x2x4", "x1x2x4", "x3x4", "x1x3x4", "x2x3x4", "x1x2x3x4",
			"x5", "x1x5", "x2x5", "x1x2x5", "x3x5", "x1x3x5", "x2x3x5", "x1x2x3x5", "x4x5", "x1x4x5", "x2x4x5", "x1x2x4x5", "x3x4x5", "x1x3x4x5",
			"x2x3x4x5", "x1x2x3x4x5", "x6", "x1x6", "x2x6", "x1x2x6", "x3x6", "x1x3x6", "x2x3x6", "x1x2x3x6", "x4x6", "x1x4x6", "x2x4x6", "x1x2x4x6",
			"x3x4x6", "x1x3x4x6", "x2x3x4x6", "x1x2x3x4x6", "x5x6", "x1x5x6", "x2x5x6", "x1x2x5x6", "x3x5x6", "x1x3x5x6", "x2x3x5x6", "x1x2x3x5x6",
			"x4x5x6", "x1x4x5x6", "x2x4x5x6", "x1x2x4x5x6", "x3x4x5x6", "x1x3x4x5x6", "x2x3x4x5x6", "x1x2x3x4x5x6",
		}
		x = []string{
			"x1", "x2", "x3", "x4", "x5", "x6",
		}
		zhigalkin string
		weight    = 0
	)
	seqF = make([][]int, field)
	for i := 0; i < field; i++ {
		seqF[i] = make([]int, int(math.Pow(2, float64(field))))
	}

	log.Println("Генерируется последовательность.")
	seq = Sattolo(items64)
	log.Println("Последовательность сгенерирована.")
	for i := 0; i < int(math.Pow(2, float64(field))); i++ {
		fmt.Printf("%d ", seq[i])
	}

	log.Println("Конвертируем исходную последовательность в двочиную систему.")
	seqBinary = ConvertToBinary(seq)
	log.Println("Последовательность конвертирована.")

	for i := 0; i < field; i++ {
		fmt.Printf("   f%d", field-i)
	}
	fmt.Println()

	for i := 0; i < int(math.Pow(2, float64(field))); i++ {
		fmt.Printf("%d:", seq[i])

		for j := 0; j < field; j++ {
			if seq[i] > 9 {
				fmt.Printf(" %d  ", seqBinary[i][j])
			} else {
				fmt.Printf("  %d ", seqBinary[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println()
	log.Println("Посчитаем вес функции.")

	for j := 0; j < field; j++ {
		seqTemp = make([]int, int(math.Pow(2, float64(field))))
		for i := 0; i < int(math.Pow(2, float64(field))); i++ {
			seqTemp[i] = seqBinary[i][j]
			if seqTemp[i] == 1 {
				weight++
			}
		}
		fmt.Printf("Вес для f%d = %d\n", field-j, weight)
		seqF[j] = TransformationToPolinom(seqTemp)
		weight = 0
	}

	fmt.Println()
	log.Println("Получим полином Жигалкина и фиктивные переменные для каждой функции.")

	for i := 0; i < len(seqF); i++ {
		for j := 0; j < len(seqF[i]); j++ {
			if seqF[i][j] == 1 {
				zhigalkin = zhigalkin + members[j] + " "
			}
		}
		str := strings.Split(zhigalkin, " ")
		str = str[:len(str)-1]
		fmt.Printf("Полином Жигалкина для f%d: %s", field-i, strings.Join(str, "+"))

		for k := 0; k < len(x); k++ {
			if !strings.Contains(zhigalkin, x[k]) {
				fmt.Printf("Фиктивные переменные для f%d: %s\n", field-i, x[k])
			}
		}
		zhigalkin = ""
		fmt.Println()
	}
}

/*
*Алгоритм Саттоло
*Данный алгоритм генерирует последовательность
 */
func Sattolo(items []int) []int {
	for i := len(items) - 1; i > 1; i-- {
		j := rand.Intn(i)
		items[j], items[i] = items[i], items[j]
	}

	return items
}

/**
*Данная функция выполняет преобразование десятичной последовательности в двоичную
 */
func ConvertToBinary(seq []int) [][]int {
	var (
		output    [][]int
		temp      string
		tempInt64 int64
	)
	output = make([][]int, int(math.Pow(2, float64(field))))
	for i := 0; i < int(math.Pow(2, float64(field))); i++ {
		output[i] = make([]int, field)
	}

	for i := 0; i < len(seq); i++ {
		temp = strconv.FormatInt(int64(seq[i]), 2)

		switch field - len(temp) {
		case 1:
			temp = "0" + temp
		case 2:
			temp = "00" + temp
		case 3:
			temp = "000" + temp
		case 4:
			temp = "0000" + temp
		case 5:
			temp = "00000" + temp
		}

		for j, r := range temp {
			tempInt64, _ = strconv.ParseInt(string(r), 10, 64)
			output[i][j] = int(tempInt64)
		}
	}

	return output
}

func TransformationToPolinom(seq []int) []int {
	var (
		seqLeft  []int
		seqRight []int
		seqOut   []int
		temp1    []int
		temp2    []int
	)

	seqLeft = make([]int, len(seq)/2)
	seqRight = make([]int, len(seq)/2)
	seqOut = make([]int, len(seq))

	for i := 0; i < len(seq)/2; i++ {
		seqLeft[i] = seq[i]
		seqRight[i] = (seq[i] + seq[i+len(seq)/2]) % 2
	}

	if len(seq) == 2 {
		seqOut[0] = seqLeft[0]
		seqOut[1] = seqRight[0]
		return seqOut
	}

	temp1 = TransformationToPolinom(seqLeft)
	temp2 = TransformationToPolinom(seqRight)

	for i := 0; i < len(seqOut)/2; i++ {
		seqOut[i] = temp1[i]
		seqOut[i+len(seqOut)/2] = temp2[i]
	}

	return seqOut
}
