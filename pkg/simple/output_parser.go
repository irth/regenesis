package simple

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Output struct {
	selected map[string]bool
	updated  map[string]string
	Raw      []string
}

func ParseOutput(r io.Reader) (Output, error) {
	out := Output{
		selected: make(map[string]bool),
		updated:  make(map[string]string),
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		out.Raw = append(out.Raw, line)
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			continue
		}
		switch strings.TrimSpace(split[0]) {
		case "selected":
			out.selected[strings.TrimSpace(split[1])] = true
		case "input":
			split := strings.SplitN(line, ":", 2)
			if len(split) != 2 {
				continue
			}
			// should I really trim space from the value here?
			out.updated[strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return out, fmt.Errorf("while reading simple stdout: %w", err)
	}

	return out, nil
}

func (o Output) Selected(id string) bool {
	_, ok := o.selected[id]
	return ok
}

func (o Output) Input(id string) (string, bool) {
	u, ok := o.updated[id]
	return u, ok
}
