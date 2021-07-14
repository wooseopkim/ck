package main_page

import (
	"time"

	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/entities"
	"github.com/wooseopkim/ck/v2/usecases"
)

type presenter struct {
	inferRemoteTime *usecases.InferRemoteTime

	view adapters.View
}

func NewPresenter(view adapters.View, inferRemoteTime *usecases.InferRemoteTime) adapters.Presenter {
	return &presenter{
		view:            view,
		inferRemoteTime: inferRemoteTime,
	}
}

func (p *presenter) OnSubmit(url string) {
	offset, err := p.inferRemoteTime.Run(entities.URL(url))
	if err != nil {
		return
	}
	p.view.StartTicking(offset, time.Millisecond)
}
