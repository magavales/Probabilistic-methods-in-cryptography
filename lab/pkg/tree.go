package pkg

import "math"

type Tree struct {
	functions []*CoordinateFunction
	parent    *Tree
	c         []int
	field     int
}

func NewTree(functions []*CoordinateFunction, field int) *Tree {
	return &Tree{
		functions: functions,
		parent:    nil,
		c:         make([]int, 0),
		field:     field,
	}
}

func (t *Tree) nextStep(vector map[int]bool) (*Tree, *Tree) {
	zero := NewTree([]*CoordinateFunction{}, t.field)
	one := NewTree([]*CoordinateFunction{}, t.field)
	zero.parent = t
	one.parent = t
	zero.c = append(t.c, 0)
	one.c = append(t.c, 1)

	for _, i := range t.functions {
		j := i.Function[len(i.Function)-t.field+1:]
		sum := 0
		for k := range j {
			sum += j[k] * int(math.Pow(2, float64(t.field-k-1)))
		}
		if vector[sum] {
			i.Function = append(i.Function, 0)
			one.functions = append(one.functions, NewCoordinateFunction(i.Function, i.Field))
		} else {
			i.Function = append(i.Function, 0)
			zero.functions = append(zero.functions, NewCoordinateFunction(i.Function, i.Field))
		}

		sum++
		if vector[sum] {
			i.Function = append(i.Function, 1)
			one.functions = append(one.functions, NewCoordinateFunction(i.Function, i.Field))
		} else {
			i.Function = append(i.Function, 1)
			zero.functions = append(zero.functions, NewCoordinateFunction(i.Function, i.Field))
		}
	}
	return zero, one
}
