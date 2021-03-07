package simple

import "strings"

type WidgetList []Widget

func (w WidgetList) Render() (string, error) {
	children := []string{}
	for _, widget := range w {
		out, err := widget.Render()
		if err != nil {
			return "", err
		}
		children = append(children, out)
	}
	return strings.Join(children, "\n"), nil
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
	return nil, nil
}
