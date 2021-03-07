package simple

import (
	"fmt"
	"os/exec"
)

type Scene struct {
	Widgets []Widget
}

type BoundEventHandler func(app *App) error

func (s *Scene) Render() ([]BoundEventHandler, error) {
	cmd := exec.Command("simple")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to simple's stdin: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to simple's stdout: %w", err)
	}

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("couldn't start simple: %w", err)
	}

	for _, widget := range s.Widgets {
		data, err := widget.Render()
		if err != nil {
			return nil, fmt.Errorf("an error occured while drawing the widget: %w", err)
		}
		stdin.Write([]byte(data))
		stdin.Write([]byte("\n"))
	}
	stdin.Close()

	parsed, err := ParseOutput(stdout)
	if err != nil {
		return nil, fmt.Errorf("failed to read simple's output: %w", err)
	}

	var handlersToRun []BoundEventHandler
	for _, widget := range s.Widgets {
		handlers, err := widget.Update(parsed)
		if err != nil {
			return nil, err
		}
		handlersToRun = append(handlersToRun, handlers...)
	}

	return handlersToRun, nil
}
