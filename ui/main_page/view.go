package main_page

import (
	"time"

	"fyne.io/fyne/v2/widget"

	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/widgets"
)

type view struct {
	datetimeLabel *widget.Label
	submitButton  *widget.Button
	ticker        *time.Ticker
}

func NewView(datetimeLabel *widget.Label, submitButton *widget.Button) adapters.View {
	return &view{datetimeLabel: datetimeLabel, submitButton: submitButton}
}

func (v *view) StartTicking(offset time.Duration, interval time.Duration) {
	if v.ticker != nil {
		v.ticker.Stop()
	}
	v.ticker = time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-v.ticker.C:
				v.datetimeLabel.SetText(time.Now().Add(offset).Format("15:04:05.9999"))
			}
		}
	}()
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
