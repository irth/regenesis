package main

import "github.com/irth/regenesis/pkg/simple"

func Header() simple.Widget {
	return simple.WidgetList{
		simple.Left,
		simple.FontSize(100),
		simple.NewLabel(
			simple.Pos(simple.Abs(100), simple.Abs(100), simple.Percent(100), simple.Abs(50)),
			"LibGen",
		),
		simple.FontSize(32),
	}
}
