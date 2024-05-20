package pkg

import "math"

type Tree struct {
	functions []*CoordinateFunction
	parent    *Tree
	c         []int
	field     int
	zero      *Tree
	one       *Tree
}

func NewTree(functions []*CoordinateFunction, field int) *Tree {
	return &Tree{
		functions: functions,
		parent:    nil,
		c:         make([]int, 0),
		field:     field,
	}
}

func (t *Tree) nextStep(function []int) {
	t.zero = NewTree([]*CoordinateFunction{}, t.field)
	t.one = NewTree([]*CoordinateFunction{}, t.field)
	t.zero.parent = t
	t.one.parent = t
	t.zero.c = append(t.c, 0)
	t.one.c = append(t.c, 1)

	for _, i := range t.functions {
		j := i.Function[len(i.Function)-t.field+1:]
		sum := 0
		for k := range j {
			sum += j[k] * int(math.Pow(2, float64(t.field-k-1)))
		}
		if function[sum] == 1 {
			i.Function = append(i.Function, 0)
			t.one.functions = append(t.one.functions, NewCoordinateFunction(i.Function, i.Field))
		} else {
			i.Function = append(i.Function, 0)
			t.zero.functions = append(t.zero.functions, NewCoordinateFunction(i.Function, i.Field))
		}

		sum++
		if function[sum] == 1 {
			i.Function = append(i.Function, 1)
			t.one.functions = append(t.one.functions, NewCoordinateFunction(i.Function, i.Field))
		} else {
			i.Function = append(i.Function, 1)
			t.zero.functions = append(t.zero.functions, NewCoordinateFunction(i.Function, i.Field))
		}
	}
}
