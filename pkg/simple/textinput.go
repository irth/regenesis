package simple

import "fmt"

var _ Widget = &TextInput{}

type UpdateHandler func(a *App, t *TextInput) error

type TextInput struct {
	ID       string
	Value    string
	OnUpdate UpdateHandler
	Position
}

func (t *TextInput) Render() (string, error) {
	return fmt.Sprintf("textinput:%s %s %s", t.ID, t.Position.Render(), t.Value), nil
}

func (t *TextInput) Update(out Output) ([]BoundEventHandler, error) {
	value, updated := out.Input(t.ID)
	if !updated {
		return nil, nil
	}

	t.Value = value

	return []BoundEventHandler{
		func(a *App) error {
			return t.OnUpdate(a, t)
		},
	}, nil
}
