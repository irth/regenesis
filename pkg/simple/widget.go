package simple

import "fmt"

type Coordinate interface {
	Render() string
}

type Percent int

func (val Percent) Render() string { return fmt.Sprintf("%d%%", val) }

type Absolute int

func (val Absolute) Render() string { return fmt.Sprintf("%d", val) }

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

func (p Position) Render() string {
	return fmt.Sprintf("%s %s %s %s", p.X.Render(), p.Y.Render(), p.Width.Render(), p.Height.Render())
}

type Widget interface {
	Render() (string, error)
	Update(stdout string) ([]BoundEventHandler, error)
}
