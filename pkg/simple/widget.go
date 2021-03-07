package simple

import "fmt"

type Coordinate interface {
	Render() string
}

type Percent int

func (val Percent) Render() string { return fmt.Sprintf("%d%%", val) }

type Abs int

func (val Abs) Render() string { return fmt.Sprintf("%d", val) }

type relative bool

func (val relative) Render() string {
	if val {
		return "same"
	} else {
		return "step"
	}
}

var Same Coordinate = relative(true)
var Step Coordinate = relative(false)

type Position struct {
	X      Coordinate
	Y      Coordinate
	Width  Coordinate
	Height Coordinate
}

func Pos(x, y, w, h Coordinate) Position {
	return Position{x, y, w, h}
}

func (p Position) Render() string {
	return fmt.Sprintf("%s %s %s %s", p.X.Render(), p.Y.Render(), p.Width.Render(), p.Height.Render())
}

type Widget interface {
	Render() (string, error)
	Update(stdout Output) ([]BoundEventHandler, error)
}
