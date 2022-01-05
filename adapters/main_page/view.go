package main_page

import "time"

type View interface {
	PushRequestEvent(url string)
	PushSleepEvent(url string, delay time.Duration)
	PushFetchEvent(url string)
	PushCalibrateEvent(url string)
	DisableInput()
	EnableInput()
	DisplayTime(time time.Time)
}
