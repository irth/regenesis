package simple

import "fmt"

var _ Widget = FontSize(10)

type FontSize int

func (f FontSize) Render() (string, error) {
	return fmt.Sprintf("@fontsize %d", f), nil
}

func (f FontSize) Update(out Output) ([]BoundEventHandler, error) {
	return nil, nil
}
