package main

import (
	"fmt"
	"laba1/pkg"
	"log"
	"math"
	"strconv"
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
		functions []*pkg.CoordinateFunction
		seqBinary [][]int
		seqF      [][]int
		seqTemp   []int
		/*test      = []int{
			3, 5, 15, 12, 8, 0, 4, 14, 10, 6, 1, 11, 9, 13, 2, 7,
		}*/
		zhigalkin []string
	)
	functions = make([]*pkg.CoordinateFunction, field)
	seqF = make([][]int, field)
	for i := 0; i < field; i++ {
		seqF[i] = make([]int, int(math.Pow(2, float64(field))))
	}

	log.Println("Генерируется последовательность.")
	seq := pkg.NewSequence(field)
	seq.Sattolo(items64)
	log.Println("Последовательность сгенерирована.")
	seq.Print()

	log.Println("Конвертируем исходную последовательность в двочиную систему.")
	fmt.Println("-------------------------------------------------------------------")
	seqBinary = ConvertToBinary(seq.Seq)
	log.Println("Последовательность конвертирована.")

	for j := 0; j < field; j++ {
		for i := 0; i < int(math.Pow(2, float64(field))); i++ {
			seqTemp = append(seqTemp, seqBinary[i][j])
		}
		functions[j] = pkg.NewCoordinateFunction(seqTemp, field)
		seqTemp = nil
	}

	for i := 0; i < field; i++ {
		fmt.Printf("   f%d", field-i)
	}
	fmt.Println("-------------------------------------------------------------------")

	for i := 0; i < int(math.Pow(2, float64(field))); i++ {
		fmt.Printf("%d:", seq.Seq[i])

		for j := 0; j < field; j++ {
			if seq.Seq[i] > 9 {
				fmt.Printf(" %d  ", seqBinary[i][j])
			} else {
				fmt.Printf("  %d ", seqBinary[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println("-------------------------------------------------------------------")
	log.Println("Посчитаем вес функции.")

	for i := 0; i < len(functions); i++ {
		fmt.Printf("Вес для f%d = %d\n", field-i, functions[i].GetWeight())
		functions[i].Polinom = TransformationToPolinom(functions[i].Function)
	}

	fmt.Println("-------------------------------------------------------------------")
	log.Println("Получим полином Жигалкина и фиктивные переменные для каждой функции.")

	for i := 0; i < len(functions); i++ {
		zhigalkin = functions[i].CreatePolinom()
		functions[i].PrintPolinom(i, zhigalkin)
	}

	fmt.Println("-------------------------------------------------------------------")
	log.Println("Найдем преобладание единиц.")
	for i := 0; i < len(functions); i++ {
		if functions[i].Predominance() {
			fmt.Printf("Для f%d преобладания единиц нет.\n", field-i)
		}
	}

	fmt.Println("-------------------------------------------------------------------")
	log.Println("Найдем запрeт для функций.")
	for i := 0; i < len(functions); i++ {
		if len(functions[i].ComputeZapret()) != 0 {
			fmt.Printf("Запрет найден для f%d.\n", field-i)
			fmt.Printf("Функция f%d не сильноравновероятная.\n", field-i)
		}
	}
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
