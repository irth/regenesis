package simple

import "strconv"

var _ Widget = FontSize(10)

type FontSize int

func (f FontSize) Render() (string, error) {
	return DirectiveCommand{"fontsize", strconv.Itoa(int(f))}.Render()
}

func (f FontSize) Update(out Output) ([]BoundEventHandler, error) {
	return nil, nil
}
