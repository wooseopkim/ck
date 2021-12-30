package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/wooseopkim/ck/v2/ui/main_page"
	"github.com/wooseopkim/ck/v2/usecases"
)

func main() {
	a := app.New()
	w := a.NewWindow("CK")

	inferRemoteTime := usecases.NewInferRemoteTime()

	widget := main_page.NewWidget(
		inferRemoteTime,
	)
	w.SetContent(widget)
	w.Resize(fyne.NewSize(800.0, 600.0))
	w.ShowAndRun()
}
