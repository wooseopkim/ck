package usecases

import (
	"net/http"
	"time"

	"github.com/wooseopkim/ck/v2/entities"
	"github.com/wooseopkim/goclock/event"
)

type TestRemoteTime struct {
	eventChannel  chan event.Event
	cancelChannel chan interface{}
}

func NewTestRemoteTime() *TestRemoteTime {
	return &TestRemoteTime{
		make(chan event.Event),
		make(chan interface{}),
	}
}

func (t *TestRemoteTime) Run(url entities.URL) (time.Time, error) {
	resp, err := http.Get(string(url))
	if err != nil {
		return time.Time{}, err
	}
	date, err := time.Parse(time.RFC1123, resp.Header.Get("Date"))
	if err != nil {
		return time.Time{}, err
	}
	return date.Local(), nil
}

func (t *TestRemoteTime) Cancel() {
	t.cancelChannel <- nil
}

func (t *TestRemoteTime) EventChannel() chan event.Event {
	return t.eventChannel
}
