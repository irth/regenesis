package simple

import "fmt"

var _ Widget = Button{}

type ClickHandler func(b Button)

type Button struct {
	ID      string
	Name    string
	OnClick ClickHandler
	Position
}

func (b Button) Render() (string, error) {
	return fmt.Sprintf("button:%s %s %s", b.ID, b.Position.Render(), b.Name), nil
}

func (b Button) Update(stdout string) ([]BoundEventHandler, error) {
	// TODO: parse
	return []BoundEventHandler{
		func(a *App) error {
			fmt.Println("update called", a)
			return nil
		},
	}, nil
}
