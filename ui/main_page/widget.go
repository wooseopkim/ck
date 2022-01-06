package main_page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/wooseopkim/ck/v2/usecases"
	"github.com/wooseopkim/ck/v2/widgets"
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
	panelLabel := widget.NewLabel("")
	submitButton := widget.NewButton("Go", func() {})
	layout := container.NewVBox(
		inputContainer,
		panelLabel,
		layout.NewSpacer(),
		submitButton,
	)

	vm := NewViewModel(inferRemoteTime)
	submitButton.OnTapped = func() {
		vm.OnSubmit(protocolSelect.Selected + urlEntry.Text)
	}
	inputEnabled := vm.InputEnabled()
	widgets.OnBoolChange(inputEnabled, func(value bool) {
		if value {
			submitButton.Enable()
		} else {
			submitButton.Disable()
		}
	})
	panel := vm.Panel()
	widgets.OnStringChange(panel, func(value string) {
		panelLabel.Text = value
		panelLabel.Refresh()
		submitButton.Refresh()
	})

	return layout
}
