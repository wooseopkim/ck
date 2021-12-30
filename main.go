package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/wooseopkim/ck/v2/ui/main_page"
	"github.com/wooseopkim/ck/v2/usecases"
	"github.com/wooseopkim/goclock"
)

func main() {
	a := app.New()
	w := a.NewWindow("CK")

	inferRemoteTime := usecases.NewInferRemoteTime(func(url string) (time.Duration, error) {
		gc, err := goclock.New(goclock.Request{URL: url})
		if err != nil {
			return 0, err
		}
		return gc.Offset, nil
	})

	w.SetContent(main_page.NewWidget(inferRemoteTime))
	w.Resize(fyne.NewSize(800.0, 600.0))
	w.ShowAndRun()
}
