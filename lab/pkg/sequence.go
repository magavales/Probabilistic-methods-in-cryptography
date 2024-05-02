package pkg

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

type Sequence struct {
	Seq   []int
	Field int
}

func NewSequence(field int) *Sequence {
	return &Sequence{
		Seq:   nil,
		Field: field,
	}
}

/*
*Алгоритм Саттоло
*Данный алгоритм генерирует последовательность
 */
func (seq *Sequence) Sattolo(items []int) {
	for i := len(items) - 1; i > 1; i-- {
		j := rand.Intn(i)
		items[j], items[i] = items[i], items[j]
	}

	seq.Seq = items
}

/**
*Данная функция выполняет преобразование десятичной последовательности в двоичную
 */
func (seq *Sequence) ConvertToBinary() [][]int {
	var (
		output    [][]int
		temp      string
		tempInt64 int64
	)
	output = make([][]int, int(math.Pow(2, float64(seq.Field))))
	for i := 0; i < int(math.Pow(2, float64(seq.Field))); i++ {
		output[i] = make([]int, seq.Field)
	}

	for i := 0; i < len(seq.Seq); i++ {
		temp = strconv.FormatInt(int64(seq.Seq[i]), 2)

		switch seq.Field - len(temp) {
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

func (seq *Sequence) Print() {
	for i := 0; i < int(math.Pow(2, float64(seq.Field))); i++ {
		fmt.Printf("%d ", seq.Seq[i])
	}
}
