package simple

import (
	"fmt"
	"os/exec"
)

type Scene struct {
	Widgets WidgetList
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

	data, err := s.Widgets.Render()
	if err != nil {
		return nil, fmt.Errorf("an error occured while drawing a widget: %w", err)
	}
	stdin.Write([]byte(data))
	stdin.Write([]byte("\n"))
	stdin.Close()

	parsed, err := ParseOutput(stdout)
	if err != nil {
		return nil, fmt.Errorf("failed to read simple's output: %w", err)
	}

	handlersToRun, err := s.Widgets.Update(parsed)
	if err != nil {
		return nil, fmt.Errorf("while running update handlers: %w", err)
	}

	return handlersToRun, nil
}
