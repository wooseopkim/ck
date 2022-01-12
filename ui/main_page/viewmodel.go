package main_page

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2/data/binding"
	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/entities"
	"github.com/wooseopkim/ck/v2/usecases"
	"github.com/wooseopkim/ck/v2/widgets"
	"github.com/wooseopkim/goclock/event"
)

const (
	clearPanel      = "clearPanel"
	timeTemplate    = "15:04:05.99"
	timeTemplateLen = len(timeTemplate)
)

type viewModel struct {
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
	v := &viewModel{
		inferRemoteTime:             inferRemoteTime,
		tickerSubscriptionCanceller: make(chan interface{}),
		event:                       binding.NewUntyped(),
		inputEnabled:                binding.NewBool(),
		now:                         binding.NewUntyped(),
	}
	v.inputEnabled.Set(true)
	v.initialize()
	return v
}

func (v *viewModel) initialize() {
	if v.initialized {
		return
	}
	v.initialized = true
	go func() {
		for {
			e := <-v.inferRemoteTime.EventChannel()
			v.event.Set(e)
		}
	}()
}

func (v *viewModel) OnSubmit(url string) {
	go v.handleSubmit(url)
}

func (v *viewModel) handleSubmit(url string) {
	v.inputEnabled.Set(false)

	if v.ticker != nil {
		go func() {
			v.inferRemoteTime.Cancel()
			v.tickerSubscriptionCanceller <- nil
		}()
		v.ticker.Stop()
		v.now.Set(clearPanel)
	}

	offset, err := v.inferRemoteTime.Run(entities.URL(url))
	v.ticker = time.NewTicker(time.Millisecond)

	v.inputEnabled.Set(true)

	if err != nil {
		return
	}

	go func() {
		for {
			select {
			case <-v.tickerSubscriptionCanceller:
				return
			case now := <-v.ticker.C:
				v.now.Set(now.Add(offset))
			}
		}
	}()
}

func (v *viewModel) Panel() binding.String {
	panel := binding.NewString()
	widgets.OnUntypedChange(v.event, func(value interface{}) {
		switch value.(type) {
		case event.Request:
			panel.Set("REQUSTED")
		case event.Sleep:
			delay := value.(event.Sleep).Delay
			ms := delay / time.Millisecond
			panel.Set(fmt.Sprintf("SLEEPING FOR %dms", ms))
		case event.Fetch:
			panel.Set("FETCHING")
		case event.Calibrate:
			panel.Set("CALIBRATING")
		}
	})
	widgets.OnUntypedChange(v.now, func(value interface{}) {
		if value == nil || value == clearPanel {
			panel.Set("Type URL above and click the button")
			return
		}
		time := value.(time.Time).Format(timeTemplate)
		timeLen := len(time)
		if timeLen < timeTemplateLen {
			time = time + strings.Repeat("0", timeTemplateLen-timeLen)
		}
		panel.Set(time)
	})
	return panel
}

func (v *viewModel) Target() binding.String {
	target := binding.NewString()
	widgets.OnUntypedChange(v.event, func(value interface{}) {
		switch value.(type) {
		case event.Request:
			url := value.(event.Request).Url
			target.Set(url)
		case event.Sleep:
			url := value.(event.Sleep).Url
			target.Set(url)
		case event.Fetch:
			url := value.(event.Fetch).Url
			target.Set(url)
		case event.Calibrate:
			url := value.(event.Calibrate).Url
			target.Set(url)
		}
	})
	target.Set("Welcome to CK!")
	return target
}

func (v *viewModel) InputEnabled() binding.Bool {
	return v.inputEnabled
}
