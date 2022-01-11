package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/wooseopkim/ck/v2/ui/main_page"
	"github.com/wooseopkim/ck/v2/usecases"
)

func main() {
	a := app.New()
	w := a.NewWindow("CK")

	a.Settings().SetTheme(theme.LightTheme())

	inferRemoteTime := usecases.NewInferRemoteTime()

	widget := main_page.NewWidget(
		inferRemoteTime,
	)
	w.SetContent(widget)
	w.Resize(fyne.NewSize(800.0, 100.0))
	w.ShowAndRun()
}
