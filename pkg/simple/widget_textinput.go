package simple

var _ Widget = &TextInputWidget{}

type UpdateHandler func(a *App, t *TextInputWidget, value string) error

type TextInputWidget struct {
	ID       string
	Value    string
	OnUpdate UpdateHandler
	Position
}

func TextInput(id string, pos Position, value string, onUpdate UpdateHandler) *TextInputWidget {
	return &TextInputWidget{id, value, onUpdate, pos}
}

func (t *TextInputWidget) Render() (string, error) {
	return WidgetCommand{
		Name:     "textinput",
		ID:       t.ID,
		Position: t.Position,
		Extra:    t.Value,
	}.Render()
}

func (t *TextInputWidget) Update(out Output) ([]BoundEventHandler, error) {
	value, updated := out.Input(t.ID)
	if !updated || t.OnUpdate == nil {
		return nil, nil
	}

	return []BoundEventHandler{
		func(a *App) error {
			return t.OnUpdate(a, t, value)
		},
	}, nil
}
