package validation

import (
	"fmt"
	"strings"
)

func (t *Tree) StringPtr(value *string) *Tree {
	if value == nil {
		t.error(errFieldMustBeSet)
	}
	return t
}

func (t *Tree) StringPtrZero(value *string) *Tree {
	if value != nil {
		t.error(fmt.Errorf(fmtFieldMustBeZero, *value))
	}
	return t
}

func (t *Tree) StringPtrSliceNotZero(value []*string) *Tree {
	if len(value) == 0 {
		t.error(errFieldMustBeSet)
	} else {
		for i, elem := range value {
			if elem == nil || *elem == "" {
				t.error(fmt.Errorf(fmtFieldIndexMustBeSet, i))
			}
		}
	}
	return t
}

func (t *Tree) StringSliceNotZero(value []string) *Tree {
	if len(value) == 0 {
		t.error(errFieldMustBeSet)
	} else {
		for i, elem := range value {
			if elem == "" {
				t.error(fmt.Errorf(fmtFieldIndexMustBeSet, i))
			}
		}
	}
	return t
}

func (t *Tree) StringSlice(value, expect []string) *Tree {
	switch n := len(value); {
	case n == 0:
		t.error(errFieldMustBeSet)
	case n != len(expect):
		t.error(fmt.Errorf(fmtFieldMustMatch, expect, value))
	default:
		for i := 0; i < n; i++ {
			vi := value[i]
			ei := expect[i]
			if vi != ei {
				t.error(fmt.Errorf(fmtFieldIndexMustMatch, i, ei, vi))
			}
		}
	}
	return t
}

func (t *Tree) StringPtrSliceZero(value []*string) *Tree {
	if value != nil {
		t.error(fmt.Errorf(fmtFieldMustBeZero, value))
	}
	return t
}

func (t *Tree) StringPtrNotZero(value *string) *Tree {
	if value != nil {
		return t.StringNotZero(*value)
	}
	t.error(errFieldMustBeSet)
	return t
}

func (t *Tree) StringNotZero(value string) *Tree {
	if value == "" {
		t.error(errFieldMustBeSet)
	}
	return t
}

func (t *Tree) StringLenRange(value string, min, max int) *Tree {
	if n := len(value); n < min || n > max {
		t.error(fmt.Errorf(fmtFieldLenRange, min, max))
	}
	return t
}

func (t *Tree) StringZero(value string) *Tree {
	if value != "" {
		t.error(fmt.Errorf(fmtFieldMustBeZero, value))
	}
	return t
}

func (t *Tree) String(value, expected string) *Tree {
	if value != expected {
		t.error(fmt.Errorf(fmtFieldMustMatch, expected, value))
	}
	return t
}

func (t *Tree) StringPtrMaxLenBytes(value *string, max int) *Tree {
	if value == nil {
		t.error(errFieldMustBeSet)
	} else if len(*value) > max {
		t.error(fmt.Errorf(fmtFieldLenMax, max, value))
	}
	return t
}

func (t *Tree) StringPtrListNotZero(value *string, sep string) *Tree {
	if value != nil {
		return t.StringListNotZero(*value, sep)
	}
	t.error(errFieldMustBeSet)
	return t
}

func (t *Tree) StringListNotZero(value, sep string) *Tree {
	if list := strings.Split(value, sep); len(list) == 0 {
		t.error(errFieldMustBeSet)
	} else {
		for i, elem := range list {
			if elem == "" {
				t.error(fmt.Errorf(fmtFieldIndexMustBeSet, i))
			}
		}
	}
	return t
}
