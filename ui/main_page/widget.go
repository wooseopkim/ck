package main_page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/usecases"
)

const (
	http  = "http://"
	https = "https://"
)

func NewWidget(inferRemoteTime *usecases.InferRemoteTime) fyne.CanvasObject {
	var presenter adapters.Presenter

	urlEntry := widget.NewEntry()
	protocolSelect := widget.NewSelect([]string{http, https}, func(s string) {})
	datetimeLabel := widget.NewLabel("")
	submitButton := widget.NewButton("Go", func() {
		presenter.OnSubmit(protocolSelect.Selected + urlEntry.Text)
	})

	view := NewView(datetimeLabel)
	presenter = NewPresenter(view, inferRemoteTime)

	return container.NewVBox(
		container.NewHBox(
			protocolSelect,
			urlEntry,
		),
		datetimeLabel,
		submitButton,
	)
}
