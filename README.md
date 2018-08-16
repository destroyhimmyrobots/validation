## Validation

Concept for a structure validation DSL.

All possible fields are validated, and a tree of errors may be obtained.

### Usage

```go
package delivery

import (
    "github.com/destroyhimmyrobots/validation"
)

type A struct { B *B }
type B struct { Description, Name string; C []*C }
type C struct { D bool }

func Validate(a *A) error {
    return validation.NewTree("A").WithinPtr(a, func(v *validation.Tree) {
        b := a.B
        v.Field("B").WithinPtr(b, func(v *validation.Tree) {
            v.Field("Description").StringNotZero(b.Description)
            v.Field("Name").StringLenRange(b.Name, 1, 128)

            c := b.C
            v.Field("C").SliceMinLen(c, 1).WithinSlice(c, func(v *validation.Tree) {
                for i, ci := range c {
                    v.Index(i).WithinPtr(ci, func(v *validation.Tree) {
                        v.Field("D").BoolTrue(ci.D)
                    })
                }
            })
        })
    }).Result()
}
```

