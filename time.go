package validation

import (
	"fmt"
	"time"
)

func (t *Tree) TimePtrNil(value *time.Time) *Tree {
	if value != nil {
		t.error(fmt.Errorf(fmtFieldMustBeZero, value))
	}
	return t
}
