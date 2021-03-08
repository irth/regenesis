package simple

var _ Widget = &ButtonWidget{}

type ClickHandler func(a *App, b *ButtonWidget) error

type ButtonWidget struct {
	ID      string
	Name    string
	OnClick ClickHandler
	Position
}

func Button(id string, pos Position, name string, onClick ClickHandler) *ButtonWidget {
	return &ButtonWidget{id, name, onClick, pos}
}

func (b *ButtonWidget) Render() (string, error) {
	return CommandWidget{
		Name:     "button",
		ID:       b.ID,
		Position: b.Position,
		Extra:    b.Name,
	}.Render()
}

func (b *ButtonWidget) Update(out Output) ([]BoundEventHandler, error) {
	if !out.Selected(b.ID) || b.OnClick == nil {
		return nil, nil
	}
	return []BoundEventHandler{
		func(a *App) error {
			return b.OnClick(a, b)
		},
	}, nil
}
