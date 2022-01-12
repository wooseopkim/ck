package main_page

import (
	"time"

	"fyne.io/fyne/v2/data/binding"
)

type ViewModel interface {
	OnSubmit(url string)
	OnTest()
	Panel() binding.String
	Target() binding.String
	TesterReady() binding.Bool
	TestResult() binding.Untyped // TestResult
	InputEnabled() binding.Bool
}

type TestResult struct {
	RemoteTime time.Time
	ClientTime time.Time
	ClockTime  time.Time
}
