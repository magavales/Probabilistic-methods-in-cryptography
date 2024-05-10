package pkg

import (
	"fmt"
	"math"
	"strings"
)

var (
	table = [][]int{
		{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 1, 0}, {0, 0, 0, 0, 1, 1}, {0, 0, 0, 1, 0, 0}, {0, 0, 0, 1, 0, 1}, {0, 0, 0, 1, 1, 0}, {0, 0, 0, 1, 1, 1},
		{0, 0, 1, 0, 0, 0}, {0, 0, 1, 0, 0, 1}, {0, 0, 1, 0, 1, 0}, {0, 0, 1, 0, 1, 1}, {0, 0, 1, 1, 0, 0}, {0, 0, 1, 1, 0, 1}, {0, 0, 1, 1, 1, 0}, {0, 0, 1, 1, 1, 1},
		{0, 1, 0, 0, 0, 0}, {0, 1, 0, 0, 0, 1}, {0, 1, 0, 0, 1, 0}, {0, 1, 0, 0, 1, 1}, {0, 1, 0, 1, 0, 0}, {0, 1, 0, 1, 0, 1}, {0, 1, 0, 1, 1, 0}, {0, 1, 0, 1, 1, 1},
		{0, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 0, 1}, {0, 1, 1, 0, 1, 0}, {0, 1, 1, 0, 1, 1}, {0, 1, 1, 1, 0, 0}, {0, 1, 1, 1, 0, 1}, {0, 1, 1, 1, 1, 0}, {0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 1}, {1, 0, 0, 0, 1, 0}, {1, 0, 0, 0, 1, 1}, {1, 0, 0, 1, 0, 0}, {1, 0, 0, 1, 0, 1}, {1, 0, 0, 1, 1, 0}, {1, 0, 0, 1, 1, 1},
		{1, 0, 1, 0, 0, 0}, {1, 0, 1, 0, 0, 1}, {1, 0, 1, 0, 1, 0}, {1, 0, 1, 0, 1, 1}, {1, 0, 1, 1, 0, 0}, {1, 0, 1, 1, 0, 1}, {1, 0, 1, 1, 1, 0}, {1, 0, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 1}, {1, 1, 0, 0, 1, 0}, {1, 1, 0, 0, 1, 1}, {1, 1, 0, 1, 0, 0}, {1, 1, 0, 1, 0, 1}, {1, 1, 0, 1, 1, 0}, {1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0}, {1, 1, 1, 0, 0, 1}, {1, 1, 1, 0, 1, 0}, {1, 1, 1, 0, 1, 1}, {1, 1, 1, 1, 0, 0}, {1, 1, 1, 1, 0, 1}, {1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1},
	}
)

type CoordinateFunction struct {
	Function []int
	Field    int
	weight   int
	Polinom  []int
	Ratios   []float64
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

func (function *CoordinateFunction) CorrelativeImmunity() int {
	var order int
	function.Ratios = make([]float64, len(function.Function))
	seqOut := fastFuries(function.Function)
	function.Ratios = walshAdamar(seqOut, function.Field)

	for i := 0; i < len(function.Ratios); i++ {
		if function.Ratios[i] != 0 && (i == 1 || i == 4 || i == 8 || i == 16 || i == 32) {
			order = 0
			break
		}
		if function.Ratios[i] != 0 && (i == 3 || i == 5 || i == 6 || i == 9 || i == 10 || i == 22) {
			order = 1
			break
		}
		if function.Ratios[i] != 0 && (i == 7 || i == 15 || i == 19) {
			order = 2
			break
		}
	}

	return order
}

func (function *CoordinateFunction) Elastic() int {
	var (
		count0 = 0
		count1 = 0
	)
	for i := 0; i < len(function.Function); i++ {
		if function.Function[i] == 1 {
			count1++
		} else {
			count0++
		}
	}

	if count0 == count1 {
		return 0
	}

	return 0
}

func (function *CoordinateFunction) GetSpectre() [][]int {
	var (
		spectre         []float64
		maximum         float64 = 0
		possibleVectors [][]int
	)
	spectre = make([]float64, len(function.Function))
	for idx, val := range function.Ratios {
		spectre[idx] = val * math.Pow(2, float64(function.Field-1))
	}

	for _, val := range spectre {
		if val > maximum {
			maximum = val
		}
	}

	for idx, val := range spectre {
		if val == maximum {
			possibleVectors = append(possibleVectors, table[idx])
		}
	}

	return possibleVectors
}

func (function *CoordinateFunction) ComputeAutocorrelationRatios() []float64 {
	var (
		idxX     int
		idxUxorX int
	)
	autocorrelationRatios := make([]float64, int(math.Pow(2, float64(function.Field))))
	uXorX := make([]int, function.Field)

	for i, uVector := range table {
		for _, xVector := range table {
			for j := 0; j < function.Field; j++ {
				uXorX[j] = uVector[j] ^ xVector[j]
			}
			for idx, v := range table {
				if equal(v, xVector) {
					idxX = idx
				}
				if equal(v, uXorX) {
					idxUxorX = idx
				}
			}
			autocorrelationRatios[i] += math.Pow(-1, float64(function.Function[idxX]^function.Function[idxUxorX]))
		}
		autocorrelationRatios[i] /= math.Pow(2, float64(function.Field))
	}

	return autocorrelationRatios
}

func (function *CoordinateFunction) GetBentStatus() bool {
	if function.Field%2 == 0 {
		return false
	}
	for i := 0; i < len(function.Ratios)-1; i++ {
		if function.Ratios[i] != function.Ratios[i+1] {
			return false
		}
	}
	return true
}

func equal(array1, array2 []int) bool {
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}

func walshAdamar(seq []int, field int) []float64 {
	ratios := make([]float64, len(seq))
	for i := 0; i < len(seq); i++ {
		ratios[i] = float64(seq[i]) / math.Pow(2, float64(field))
		if i == 0 {
			ratios[i] = 1 - 2*ratios[i]
		} else {
			ratios[i] = -2 * ratios[i]
		}
	}
	return ratios
}

func fastFuries(seq []int) []int {
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
		seqLeft[i] = seq[i] + seq[i+len(seq)/2]
		seqRight[i] = seq[i] - seq[i+len(seq)/2]
	}

	if len(seq) == 2 {
		seqOut[0] = seqLeft[0]
		seqOut[1] = seqRight[0]
		return seqOut
	}

	temp1 = fastFuries(seqLeft)
	temp2 = fastFuries(seqRight)

	for i := 0; i < len(seqOut)/2; i++ {
		seqOut[i] = temp1[i]
		seqOut[i+len(seqOut)/2] = temp2[i]
	}

	return seqOut
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
