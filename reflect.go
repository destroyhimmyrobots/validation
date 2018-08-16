package validation

import (
	"fmt"
	"reflect"
)

func (t *Tree) SliceMinLen(value interface{}, min int) *Tree {
	if reflect.ValueOf(value).Len() < min {
		t.error(fmt.Errorf(fmtFieldLenMin, min, value))
	}
	return t
}

func (t *Tree) Len(value interface{}, len int) *Tree {
	if reflect.ValueOf(value).Len() != len {
		t.error(fmt.Errorf(fmtFieldLen, len, value))
	}
	return t
}

func (t *Tree) NotNil(value interface{}) *Tree {
	if reflect.ValueOf(value).IsNil() {
		t.error(errFieldMustBeSet)
	}
	return t
}

func (t *Tree) Nil(value interface{}) *Tree {
	if !reflect.ValueOf(value).IsNil() {
		t.error(fmt.Errorf(fmtFieldMustBeZero, value))
	}
	return t
}

func (t *Tree) Ignore() {}

func (t *Tree) WithinPtr(value interface{}, fn func(*Tree)) *Tree {
	if err := t.NotNil(fn); err == nil {
		fn(t)
	}
	return t
}

func (t *Tree) WithinSlice(value interface{}, fn func(*Tree)) *Tree {
	if err := t.SliceMinLen(value, 1); err == nil {
		fn(t)
	}
	return t
}
