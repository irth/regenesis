package simple

var _ Widget = &Label{}

type Label struct {
	Text string
	Position
}

func NewLabel(pos Position, text string) *Label {
	return &Label{text, pos}
}

func (b *Label) Render() (string, error) {
	return CommandWidget{
		Name:     "label",
		Position: b.Position,
		Extra:    b.Text,
	}.Render()
}

func (b *Label) Update(out Output) ([]BoundEventHandler, error) {
	return nil, nil
}
