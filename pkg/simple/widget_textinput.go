package simple

var _ Widget = &TextInput{}

type UpdateHandler func(a *App, t *TextInput, value string) error

type TextInput struct {
	ID       string
	Value    string
	OnUpdate UpdateHandler
	Position
}

func NewTextInput(id string, pos Position, value string, onUpdate UpdateHandler) *TextInput {
	return &TextInput{id, value, onUpdate, pos}
}

func (t *TextInput) Render() (string, error) {
	return CommandWidget{
		Name:     "textinput",
		ID:       t.ID,
		Position: t.Position,
		Extra:    t.Value,
	}.Render()
}

func (t *TextInput) Update(out Output) ([]BoundEventHandler, error) {
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
