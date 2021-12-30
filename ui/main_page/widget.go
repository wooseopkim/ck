package main_page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/wooseopkim/ck/v2/usecases"
)

const (
	http  = "http://"
	https = "https://"
)

func NewWidget(
	inferRemoteTime *usecases.InferRemoteTime,
) fyne.CanvasObject {
	urlEntry := widget.NewEntry()
	protocolSelect := widget.NewSelect([]string{http, https}, func(s string) {})
	protocolSelect.SetSelected(protocolSelect.Options[0])
	inputContainer := fyne.NewContainerWithLayout(
		layout.NewFormLayout(),
		protocolSelect,
		urlEntry,
	)
	datetimeLabel := widget.NewLabel("")
	submitButton := widget.NewButton("Go", func() {})

	view := NewView(
		protocolSelect,
		urlEntry,
		datetimeLabel,
		submitButton,
	)
	presenter := NewPresenter(
		view,
		inferRemoteTime,
	)
	view.Attach(presenter)

	return container.NewVBox(
		inputContainer,
		datetimeLabel,
		layout.NewSpacer(),
		submitButton,
	)
}
