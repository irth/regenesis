package simple

import (
	"fmt"
)

type CommandDirective struct {
	Name  string
	Extra string
}

func (w CommandDirective) Render() (string, error) {
	return fmt.Sprintf("@%s %s\n", w.Name, w.Extra), nil
}
