package main_page

import (
	"time"

	"fyne.io/fyne/v2/widget"

	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
)

type view struct {
	datetimeLabel *widget.Label
	ticker        *time.Ticker
}

func NewView(datetimeLabel *widget.Label) adapters.View {
	return &view{datetimeLabel: datetimeLabel}
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
