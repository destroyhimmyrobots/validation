package validation

import (
	"fmt"
)

var (
	errMustBeFalse = fmt.Errorf("must be false")
	errMustBeTrue  = fmt.Errorf("must be true")
)

func (t *Tree) BoolFalse(value bool) *Tree {
	if value {
		t.error(errMustBeFalse)
	}
	return t
}

func (t *Tree) BoolTrue(value bool) *Tree {
	if !value {
		t.error(errMustBeTrue)
	}
	return t
}

func (t *Tree) BoolPtrFalse(value *bool) *Tree {
	if value != nil {
		return t.BoolFalse(*value)
	}
	t.error(errFieldMustBeSet)
	return t
}

func (t *Tree) BoolPtrTrue(value *bool) *Tree {
	if value != nil {
		return t.BoolTrue(*value)
	}
	t.error(errFieldMustBeSet)
	return t
}
