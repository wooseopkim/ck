package main_page

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/data/binding"
	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/entities"
	"github.com/wooseopkim/ck/v2/usecases"
	"github.com/wooseopkim/ck/v2/widgets"
	"github.com/wooseopkim/goclock/event"
)

const clearPanel = "clearPanel"

type presenter struct {
	inferRemoteTime *usecases.InferRemoteTime

	initialized                 bool
	ticker                      *time.Ticker
	tickerSubscriptionCanceller chan interface{}

	event        binding.Untyped // event.Event
	inputEnabled binding.Bool
	now          binding.Untyped // time.Time
}

func NewViewModel(
	inferRemoteTime *usecases.InferRemoteTime,
) adapters.ViewModel {
	p := &presenter{
		inferRemoteTime:             inferRemoteTime,
		tickerSubscriptionCanceller: make(chan interface{}),
		event:                       binding.NewUntyped(),
		inputEnabled:                binding.NewBool(),
		now:                         binding.NewUntyped(),
	}
	p.inputEnabled.Set(true)
	p.initialize()
	return p
}

func (p *presenter) initialize() {
	if p.initialized {
		return
	}
	p.initialized = true
	go func() {
		for {
			e := <-p.inferRemoteTime.EventChannel()
			p.event.Set(e)
		}
	}()
}

func (p *presenter) OnSubmit(url string) {
	go p.handleSubmit(url)
}

func (p *presenter) handleSubmit(url string) {
	p.inputEnabled.Set(false)

	if p.ticker != nil {
		go func() {
			p.inferRemoteTime.Cancel()
			p.tickerSubscriptionCanceller <- nil
		}()
		p.ticker.Stop()
		p.now.Set(clearPanel)
	}

	offset, err := p.inferRemoteTime.Run(entities.URL(url))
	p.ticker = time.NewTicker(time.Millisecond)

	p.inputEnabled.Set(true)

	if err != nil {
		return
	}

	go func() {
		for {
			select {
			case <-p.tickerSubscriptionCanceller:
				return
			case now := <-p.ticker.C:
				p.now.Set(now.Add(offset))
			}
		}
	}()
}

func (p *presenter) Panel() binding.String {
	panel := binding.NewString()
	target := binding.NewString()
	widgets.OnUntypedChange(p.event, func(value interface{}) {
		switch value.(type) {
		case event.Request:
			url := value.(event.Request).Url
			target.Set(url)
			panel.Set(fmt.Sprintf("%s REQUSTED", url))
		case event.Sleep:
			url := value.(event.Sleep).Url
			delay := value.(event.Sleep).Delay
			ms := delay / time.Millisecond
			target.Set(url)
			panel.Set(fmt.Sprintf("%s SLEEPING FOR %dms", url, ms))
		case event.Fetch:
			url := value.(event.Fetch).Url
			target.Set(url)
			panel.Set(fmt.Sprintf("%s FETCHING", url))
		case event.Calibrate:
			url := value.(event.Calibrate).Url
			target.Set(url)
			panel.Set(fmt.Sprintf("%s CALIBRATING", url))
		}
	})
	widgets.OnUntypedChange(p.now, func(value interface{}) {
		if value == nil || value == clearPanel {
			panel.Set("Welcome to CK!")
			return
		}
		url, _ := target.Get()
		time := value.(time.Time).Format("15:04:05.99")
		panel.Set(fmt.Sprintf("%s %s", url, time))
	})
	return panel
}

func (p *presenter) InputEnabled() binding.Bool {
	return p.inputEnabled
}
