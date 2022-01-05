package main_page

import (
	"time"

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

type view struct {
	datetimeLabel *widget.Label
	submitButton  *widget.Button
}

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

	presenter := NewPresenter(&view{
		datetimeLabel: datetimeLabel,
		submitButton:  submitButton,
	}, inferRemoteTime)
	submitButton.OnTapped = func() {
		presenter.OnSubmit(protocolSelect.Selected + urlEntry.Text)
	}

	return container.NewVBox(
		inputContainer,
		datetimeLabel,
		layout.NewSpacer(),
		submitButton,
	)
}

func (v *view) DisplayTime(time time.Time) {
	v.datetimeLabel.SetText(time.Format("15:04:05.9999"))
}

func (v *view) PushRequestEvent(url string) {
	widgets.AppendText(v.datetimeLabel, "REQUEST "+url+"\n")
}

func (v *view) PushSleepEvent(url string, delay time.Duration) {
	widgets.AppendText(v.datetimeLabel, "SLEEP "+url+"\n")
}

func (v *view) PushFetchEvent(url string) {
	widgets.AppendText(v.datetimeLabel, "FETCH "+url+"\n")
}

func (v *view) PushCalibrateEvent(url string) {
	widgets.AppendText(v.datetimeLabel, "CALIBRATE "+url+"\n")
}

func (v *view) DisableInput() {
	v.submitButton.Disable()
}

func (v *view) EnableInput() {
	v.submitButton.Enable()
}
