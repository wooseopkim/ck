package usecases

import (
	"time"

	"github.com/wooseopkim/ck/v2/entities"
	"github.com/wooseopkim/goclock"
	"github.com/wooseopkim/goclock/event"
)

type InferRemoteTime struct {
	eventChannel  chan event.Event
	cancelChannel chan interface{}
}

func NewInferRemoteTime() *InferRemoteTime {
	return &InferRemoteTime{
		make(chan event.Event),
		make(chan interface{}),
	}
}

func (i *InferRemoteTime) Run(url entities.URL) (time.Duration, error) {
	req := goclock.Request{
		URL:           string(url),
		EventChannel:  i.eventChannel,
		CancelChannel: i.cancelChannel,
	}
	gc, err := goclock.Create(req)
	if err != nil {
		return 0, err
	}
	return gc.Offset, nil
}

func (i *InferRemoteTime) Cancel() {
	i.cancelChannel <- nil
}

func (i *InferRemoteTime) EventChannel() chan event.Event {
	return i.eventChannel
}
