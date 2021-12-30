package main_page

import "time"

type View interface {
	StartTicking(offset time.Duration, interval time.Duration)
	PushRequestEvent(url string)
	PushSleepEvent(url string, delay time.Duration)
	PushFetchEvent(url string)
	PushCalibrateEvent(url string)
	DisableInput()
	EnableInput()
}
