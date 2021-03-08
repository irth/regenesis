package main

import (
	"fmt"

	libgen "github.com/irth/golibgen"
	ui "github.com/irth/regenesis/pkg/simple"
)

type BookScreen struct {
	r    *Regenesis
	Book libgen.Book
}

func NewBookScreen(r *Regenesis, b libgen.Book) ui.Scene {
	return &BookScreen{r, b}
}

func (b *BookScreen) Render() (ui.Widget, error) {
	return ui.WidgetList{
		Header(),
		ui.FontSize(64),
		ui.Label(
			ui.Pos(ui.Abs(100), ui.Abs(300), ui.Percent(100), ui.Abs(100)),
			"Book details",
		),
		ui.FontSize(48),
		ui.Label(
			ui.Pos(ui.Abs(150), ui.Abs(430), ui.Abs(0), ui.Abs(0)),
			b.Book.Title(),
		),
		ui.FontSize(32),
		ui.Label(
			ui.Pos(ui.Abs(150), ui.Abs(480), ui.Abs(0), ui.Abs(0)),
			b.Book.Author(),
		),
		ui.Label(
			ui.Pos(ui.Abs(150), ui.Abs(550), ui.Abs(0), ui.Abs(0)),
			fmt.Sprintf("%s, %s / %s", b.Book.Language(), b.Book.Size(), b.Book.Format()),
		),
		ui.FontSize(64),
		ui.Button(
			"download",
			ui.Pos(ui.Abs(150), ui.Abs(630), ui.Abs(325), ui.Abs(75)),
			"[download]",
			b.download,
		),
		ui.FontSize(38),
		ui.TextInput(
			"location",
			ui.Pos(ui.Abs(150), ui.Abs(730), ui.Abs(b.r.app.ScreenWidth()-2*150), ui.Abs(50)),
			b.r.BookLocation,
			b.updateLocation,
		),

		BackButton(b.r),
	}, nil
}

func (b *BookScreen) updateLocation(a *ui.App, t *ui.TextInputWidget, newValue string) error {
	b.r.BookLocation = newValue
	// TODO: store this somewhere permanently
	return nil
}

func (b *BookScreen) download(a *ui.App, t *ui.ButtonWidget) error {
	err := DownloadBook(b.Book, b.r.BookLocation)
	if err != nil {
		println(err)
	}
	// TODO: catch and display the returned error
	return nil
}
