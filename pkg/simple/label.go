package simple

import "fmt"

var _ Widget = &Label{}

type Label struct {
	Text string
	Position
}

func NewLabel(pos Position, text string) *Label {
	return &Label{text, pos}
}

func (b *Label) Render() (string, error) {
	return fmt.Sprintf("label %s %s", b.Position.Render(), b.Text), nil
}

func (b *Label) Update(out Output) ([]BoundEventHandler, error) {
	return nil, nil
}
