package simple

var _ Widget = &Button{}

type ClickHandler func(a *App, b *Button) error

type Button struct {
	ID      string
	Name    string
	OnClick ClickHandler
	Position
}

func NewButton(id string, pos Position, name string, onClick ClickHandler) *Button {
	return &Button{id, name, onClick, pos}
}

func (b *Button) Render() (string, error) {
	return CommandWidget{
		Name:     "button",
		ID:       b.ID,
		Position: b.Position,
		Extra:    b.Name,
	}.Render()
}

func (b *Button) Update(out Output) ([]BoundEventHandler, error) {
	if !out.Selected(b.ID) || b.OnClick == nil {
		return nil, nil
	}
	return []BoundEventHandler{
		func(a *App) error {
			return b.OnClick(a, b)
		},
	}, nil
}
