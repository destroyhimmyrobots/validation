package validation

import (
	"fmt"
)

const (
	minFieldValueFmt = "minimum value of %d"
)

func (t *Tree) Int64Min(value int64, min int64) *Tree {
	if value < min {
		t.error(fmt.Errorf(minFieldValueFmt, min))
	}
	return t
}

func (t *Tree) Int64PtrMin(value *int64, min int64) *Tree {
	if value != nil {
		return t.Int64Min(*value, min)
	}
	t.error(errFieldMustBeSet)
	return t
}

func (t *Tree) Int64(value, expected int64) *Tree {
	if value != expected {
		t.error(fmt.Errorf(fmtFieldMustMatch, expected, value))
	}
	return t
}

func (t *Tree) Int64Zero(value int64) *Tree {
	if value != 0 {
		t.error(fmt.Errorf(fmtFieldMustBeZero, value))
	}
	return t
}
