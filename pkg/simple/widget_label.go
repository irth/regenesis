package simple

var _ Widget = &LabelWidget{}

type LabelWidget struct {
	Text string
	Position
}

func Label(pos Position, text string) *LabelWidget {
	return &LabelWidget{text, pos}
}

func (b *LabelWidget) Render() (string, error) {
	return WidgetCommand{
		Name:     "label",
		Position: b.Position,
		Extra:    b.Text,
	}.Render()
}

func (b *LabelWidget) Update(out Output) ([]BoundEventHandler, error) {
	return nil, nil
}
