package simple

import (
	"fmt"
)

type DirectiveCommand struct {
	Name  string
	Extra string
}

func (w DirectiveCommand) Render() (string, error) {
	return fmt.Sprintf("@%s %s\n", w.Name, w.Extra), nil
}
