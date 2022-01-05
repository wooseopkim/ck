package main_page

import (
	"time"

	"fyne.io/fyne/v2/widget"

	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/widgets"
)

type view struct {
	presenter     adapters.Presenter
	datetimeLabel *widget.Label
	submitButton  *widget.Button
	ticker        *time.Ticker
}

func NewView(
	protocolSelect *widget.Select,
	urlEntry *widget.Entry,
	datetimeLabel *widget.Label,
	submitButton *widget.Button,
) adapters.View {
	v := &view{
		datetimeLabel: datetimeLabel,
		submitButton:  submitButton,
	}
	submitButton.OnTapped = func() {
		v.presenter.OnSubmit(protocolSelect.Selected + urlEntry.Text)
	}
	return v
}

func (v *view) Attach(presenter adapters.Presenter) {
	v.presenter = presenter
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
