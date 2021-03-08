package simple

import (
	"fmt"
	"os/exec"
)

type Scene interface {
	Render() (Widget, error)
}

type BoundEventHandler func(app *App) error

func RunScene(s Scene) ([]BoundEventHandler, error) {
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

	widgets, err := s.Render()
	if err != nil {
		return nil, fmt.Errorf("while rendering the scene: %w", err)
	}

	sas, err := widgets.Render()
	if err != nil {
		return nil, fmt.Errorf("while rendering a widget: %w", err)
	}

	stdin.Write([]byte(sas))
	stdin.Close()

	parsed, err := ParseOutput(stdout)
	if err != nil {
		return nil, fmt.Errorf("failed to read simple's output: %w", err)
	}

	err = cmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("error ocurred when runing simple: %w", err)
	}

	handlersToRun, err := widgets.Update(parsed)
	if err != nil {
		return nil, fmt.Errorf("while running update handlers: %w", err)
	}

	return handlersToRun, nil
}
