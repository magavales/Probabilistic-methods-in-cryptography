package pkg

import (
	"fmt"
	"strings"
)

type CoordinateFunction struct {
	Function []int
	Field    int
	weight   int
	Polinom  []int
}

func NewCoordinateFunction(seq []int, field int) *CoordinateFunction {
	return &CoordinateFunction{
		Function: seq,
		Field:    field,
		weight:   calculateWeight(seq),
	}
}

func (function *CoordinateFunction) GetWeight() int {
	return function.weight
}

func (function *CoordinateFunction) CreatePolinom() []string {
	var (
		members = []string{
			"1", "x1", "x2", "x1x2", "x3", "x1x3", "x2x3", "x1x2x3", "x4", "x1x4", "x2x4", "x1x2x4", "x3x4", "x1x3x4", "x2x3x4", "x1x2x3x4",
			"x5", "x1x5", "x2x5", "x1x2x5", "x3x5", "x1x3x5", "x2x3x5", "x1x2x3x5", "x4x5", "x1x4x5", "x2x4x5", "x1x2x4x5", "x3x4x5", "x1x3x4x5",
			"x2x3x4x5", "x1x2x3x4x5", "x6", "x1x6", "x2x6", "x1x2x6", "x3x6", "x1x3x6", "x2x3x6", "x1x2x3x6", "x4x6", "x1x4x6", "x2x4x6", "x1x2x4x6",
			"x3x4x6", "x1x3x4x6", "x2x3x4x6", "x1x2x3x4x6", "x5x6", "x1x5x6", "x2x5x6", "x1x2x5x6", "x3x5x6", "x1x3x5x6", "x2x3x5x6", "x1x2x3x5x6",
			"x4x5x6", "x1x4x5x6", "x2x4x5x6", "x1x2x4x5x6", "x3x4x5x6", "x1x3x4x5x6", "x2x3x4x5x6", "x1x2x3x4x5x6",
		}
		str string
	)
	for i := 0; i < len(function.Function); i++ {
		if function.Function[i] == 1 {
			str = str + members[i] + " "
		}
	}
	zhigalkin := strings.Split(str, " ")
	zhigalkin = zhigalkin[:len(zhigalkin)-1]

	return zhigalkin
}

func (function *CoordinateFunction) PrintPolinom(i int, zhigalkin []string) {
	var x = []string{
		"x1", "x2", "x3", "x4", "x5", "x6",
	}
	fmt.Printf("Полином Жигалкина для f%d: %s\n", function.Field-i, strings.Join(zhigalkin, "+"))

	str := strings.Join(zhigalkin, "")

	for k := 0; k < len(x); k++ {
		if !strings.Contains(str, x[k]) {
			fmt.Printf("Фиктивные переменные для f%d: %s\n", function.Field-i, x[k])
		}
	}
}

func calculateWeight(seq []int) int {
	var (
		weight int
	)
	for i := 0; i < len(seq); i++ {
		if seq[i] == 1 {
			weight++
		}
	}

	return weight
}
