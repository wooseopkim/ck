package main_page

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2/data/binding"
	adapters "github.com/wooseopkim/ck/v2/adapters/main_page"
	"github.com/wooseopkim/ck/v2/entities"
	"github.com/wooseopkim/ck/v2/extensions/bindings"
	"github.com/wooseopkim/ck/v2/usecases"
	"github.com/wooseopkim/goclock/event"
)

const (
	clearPanel = "clearPanel"
)

type viewModel struct {
	inferRemoteTime *usecases.InferRemoteTime
	testRemoteTime  *usecases.TestRemoteTime

	initialized                 bool
	ticker                      *time.Ticker
	tickerSubscriptionCanceller chan interface{}

	inferRemoteTimeEvent binding.Untyped // event.Event
	inputEnabled         binding.Bool
	target               binding.String
	now                  binding.Untyped // time.Time
	testerReady          binding.Bool
	testResult           binding.Untyped // adapters.TestResult
}

func NewViewModel(
	inferRemoteTime *usecases.InferRemoteTime,
	testRemoteTime *usecases.TestRemoteTime,
) adapters.ViewModel {
	v := &viewModel{
		inferRemoteTime:             inferRemoteTime,
		testRemoteTime:              testRemoteTime,
		tickerSubscriptionCanceller: make(chan interface{}),
		inferRemoteTimeEvent:        binding.NewUntyped(),
		inputEnabled:                binding.NewBool(),
		target:                      binding.NewString(),
		now:                         binding.NewUntyped(),
		testerReady:                 binding.NewBool(),
		testResult:                  binding.NewUntyped(),
	}
	v.target.Set("Welcome to CK!")
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
			v.inferRemoteTimeEvent.Set(e)
		}
	}()
}

func (v *viewModel) OnSubmit(url string) {
	go v.handleSubmit(url)
}

func (v *viewModel) handleSubmit(url string) {
	v.inputEnabled.Set(false)
	v.testerReady.Set(false)

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
	v.testerReady.Set(true)

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

func (v *viewModel) OnTest() {
	url, _ := v.target.Get()
	go v.handleTest(url)
}

func (v *viewModel) handleTest(url string) {
	v.inputEnabled.Set(false)

	clientTime := time.Now()
	clockTime, _ := v.now.Get()
	remoteTime, err := v.testRemoteTime.Run(entities.URL(url))

	v.inputEnabled.Set(true)

	if err != nil {
		return
	}

	v.testResult.Set(adapters.TestResult{
		RemoteTime: remoteTime,
		ClientTime: clientTime,
		ClockTime:  clockTime.(time.Time),
	})
}

func (v *viewModel) Panel() binding.String {
	panel := binding.NewString()
	bindings.OnUntypedChange(v.inferRemoteTimeEvent, func(value interface{}) {
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
	bindings.OnUntypedChange(v.now, func(value interface{}) {
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
	bindings.OnUntypedChange(v.inferRemoteTimeEvent, func(value interface{}) {
		switch value.(type) {
		case event.Request:
			url := value.(event.Request).Url
			v.target.Set(url)
		case event.Sleep:
			url := value.(event.Sleep).Url
			v.target.Set(url)
		case event.Fetch:
			url := value.(event.Fetch).Url
			v.target.Set(url)
		case event.Calibrate:
			url := value.(event.Calibrate).Url
			v.target.Set(url)
		}
	})
	return v.target
}

func (v *viewModel) TesterReady() binding.Bool {
	return v.testerReady
}

func (v *viewModel) TestResult() binding.Untyped {
	return v.testResult
}

func (v *viewModel) InputEnabled() binding.Bool {
	return v.inputEnabled
}
