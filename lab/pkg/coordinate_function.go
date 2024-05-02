package pkg

import (
	"fmt"
	"math"
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

func (function *CoordinateFunction) Predominance() bool {
	if 1-(function.weight)/(int(math.Pow(2, float64(function.Field)-1))) == 0 {
		return true
	} else {
		return false
	}
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
		str       string
		zhigalkin []string
	)
	for i := 0; i < len(function.Function); i++ {
		if function.Function[i] == 1 {
			str = str + members[i] + " "
		}
	}
	zhigalkin = strings.Split(str, " ")
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

func (function *CoordinateFunction) ComputeZapret() []int {
	var vector map[int]bool
	vec := make([]*CoordinateFunction, 0)
	for _, v := range product([]int{0, 1}, function.Field-1) {
		vec = append(vec, NewCoordinateFunction(v, function.Field))
	}
	t := NewTree(vec, function.Field)
	zapret := make([]int, function.Field)

	tmp := []*Tree{t}
	min := math.Inf(1)
	var next []*Tree
	count := 0
	for {
		for _, i := range tmp {
			zero, one := i.nextStep(vector)
			if float64(len(one.functions)) < min {
				min = float64(len(one.functions))
				next = []*Tree{one}
			} else if float64(len(one.functions)) == min {
				next = append(next, one)
			}

			if float64(len(zero.functions)) < min {
				min = float64(len(zero.functions))
				next = []*Tree{zero}
			} else if float64(len(zero.functions)) == min {
				next = append(next, zero)
			}
		}
		if min == 0 {
			break
		}
		tmp = next
		next = []*Tree{}
		count++
		if min == math.Pow(2, float64(function.Field)) && count > function.Field*4 {
			return []int{-1}
		}
	}
	for _, i := range tmp {
		zero, one := i.nextStep(vector)
		if len(one.functions) == 0 {
			return one.c
		}
		if len(zero.functions) == 0 {
			return zero.c
		}
	}

	return zapret
}

func product(arr []int, repeat int) [][]int {
	if repeat == 0 {
		return [][]int{{}}
	}

	var result [][]int
	subsets := product(arr, repeat-1)
	for _, subset := range subsets {
		for _, item := range arr {
			result = append(result, append(subset, item))
		}
	}
	return result
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
