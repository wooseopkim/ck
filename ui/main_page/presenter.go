package main_page

import (
	"time"

	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/entities"
	"github.com/wooseopkim/ck/v2/usecases"
	"github.com/wooseopkim/goclock/event"
)

type presenter struct {
	inferRemoteTime *usecases.InferRemoteTime

	view   adapters.View
	ticker *time.Ticker
}

func NewPresenter(
	view adapters.View,
	inferRemoteTime *usecases.InferRemoteTime,
) adapters.Presenter {
	p := &presenter{
		view:            view,
		inferRemoteTime: inferRemoteTime,
	}
	p.OnStart()
	return p
}

func (p *presenter) OnStart() {
	go func() {
		for {
			e := <-p.inferRemoteTime.EventChannel()
			switch e.(type) {
			case event.Request:
				p.view.PushRequestEvent(e.(event.Request).Url)
			case event.Sleep:
				typed := e.(event.Sleep)
				p.view.PushSleepEvent(typed.Url, typed.Delay)
			case event.Fetch:
				p.view.PushFetchEvent(e.(event.Fetch).Url)
			case event.Calibrate:
				p.view.PushCalibrateEvent(e.(event.Calibrate).Url)
			}
		}
	}()
}

func (p *presenter) OnSubmit(url string) {
	if p.ticker != nil {
		p.ticker.Stop()
		p.ticker = nil
	}
	p.view.DisableInput()
	offset, err := p.inferRemoteTime.Run(entities.URL(url))
	p.view.EnableInput()
	if err != nil {
		return
	}
	p.ticker = time.NewTicker(time.Millisecond)
	go func(ticker *time.Ticker) {
		for {
			if p.ticker == nil {
				return
			}
			select {
			case <-ticker.C:
				p.view.DisplayTime(time.Now().Add(offset))
			}
		}
	}(p.ticker)
}
