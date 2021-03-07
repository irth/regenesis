package simple

import "fmt"

var _ Widget = Justify("")

type Justify string

const (
	Left   Justify = "left"
	Center Justify = "center"
	Right  Justify = "right"
)

func (j Justify) Render() (string, error) {
	return fmt.Sprintf("@justify %s", j), nil
}

func (j Justify) Update(out Output) ([]BoundEventHandler, error) {
	return nil, nil
}
