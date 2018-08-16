package validation

import (
	"errors"
	"strconv"
)

var (
	errFieldMustBeSet = errors.New("must not be zero value, empty, or nil")
)

const (
	fmtFieldIndexMustBeSet = "index %d must not be zero value, empty, or nil"
	fmtFieldIndexMustMatch = "index %d must match %+v: %+v"
	fmtFieldMustBeZero     = "must be zero value, empty, or nil: %+v"
	fmtFieldMustMatch      = "must match %+v: %+v"
	fmtFieldLenRange       = "length must range between %d and %d"
	fmtFieldLenMin         = "length must be minimum %d: %+v"
	fmtFieldLenMax         = "length exceeds maximum %d: %+v"
	fmtFieldLen            = "length must be exactly %d: %+v"
)

type Tree struct {
	name   string
	errors []error
	fields map[string]*Tree
}

func NewTree(name string) *Tree {
	return &Tree{
		name: name,
	}
}

func (t *Tree) error(err error) {
	t.errors = append(t.errors, err)
}

func (t *Tree) Result() error {
	if t.Valid() {
		return nil
	}
	return &Result{Tree: t}
}

func (t *Tree) Valid() bool {
	for _, err := range t.errors {
		if err != nil {
			return false
		}
	}
	for _, c := range t.fields {
		if c != nil && !c.Valid() {
			return false
		}
	}
	return true
}

func (t *Tree) Field(field string) *Tree {
	if t, ok := t.fields[field]; ok {
		return t
	}
	n := NewTree(field)
	t.fields[field] = n
	return n
}

func (t *Tree) Index(index int) *Tree {
	return t.Field(strconv.Itoa(index))
}
