package main_page

import (
	"fmt"
	"time"

	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/entities"
	"github.com/wooseopkim/ck/v2/usecases"
	"github.com/wooseopkim/goclock/event"
)

type presenter struct {
	inferRemoteTime *usecases.InferRemoteTime

	view adapters.View
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
	fmt.Println("OnStart")
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
	offset, err := p.inferRemoteTime.Run(entities.URL(url))
	if err != nil {
		return
	}
	p.view.StartTicking(offset, time.Millisecond)
}
