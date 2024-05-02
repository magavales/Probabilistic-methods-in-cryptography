package pkg

import (
	"fmt"
	"math"
	"math/rand"
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

func (seq *Sequence) Print() {
	for i := 0; i < int(math.Pow(2, float64(seq.Field))); i++ {
		fmt.Printf("%d ", seq.Seq[i])
	}
}
