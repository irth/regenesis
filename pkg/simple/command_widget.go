package simple

import "strings"

type WidgetCommand struct {
	Name      string
	ID        string
	Position  Position
	Extra     string
	Multiline bool
}

func (w WidgetCommand) Render() (string, error) {
	var str strings.Builder
	if w.Multiline {
		str.WriteRune('[')
	}

	str.WriteString(w.Name)
	if w.ID != "" {
		str.WriteRune(':')
		str.WriteString(w.ID)
	}

	str.WriteRune(' ')
	str.WriteString(w.Position.Render())

	if w.Extra != "" {
		str.WriteRune(' ')
		str.WriteString(w.Extra)
	}

	if w.Multiline {
		str.WriteRune(']')
	}

	str.WriteRune('\n')

	return str.String(), nil
}
