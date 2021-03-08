package simple

import "strings"

type WidgetList []Widget

func (w WidgetList) Render() (string, error) {
	var sb strings.Builder

	for _, widget := range w {
		if widget == nil {
			continue
		}

		out, err := widget.Render()
		if err != nil {
			return "", err
		}
		sb.WriteString(out)
	}
	return sb.String(), nil
}

func (w WidgetList) Update(out Output) ([]BoundEventHandler, error) {
	handlers := []BoundEventHandler{}
	for _, widget := range w {
		h, err := widget.Update(out)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, h...)
	}
	return handlers, nil
}
