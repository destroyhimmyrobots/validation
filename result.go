package validation

import (
	"strings"
)

type Result struct {
	*Tree
}

func formatError(t *Tree, sb *strings.Builder, depth int) {
	const indent = "    "
	indentation := strings.Repeat(indent, depth)

	for _, err := range t.errors {
		if err != nil {
			sb.WriteString(indentation)
			sb.WriteString(err.Error())
			sb.WriteString("\n")
		}
	}

	if len(t.fields) > 0 {
		depth++
		for _, f := range t.fields {
			if f != nil {
				formatError(f, sb, depth)
			}
		}
	}
}

func (r *Result) Error() string {
	var sb strings.Builder
	formatError(r.Tree, &sb, 0)
	return sb.String()
}
