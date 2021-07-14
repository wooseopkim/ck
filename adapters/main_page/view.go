package main_page

import "time"

type View interface {
	StartTicking(offset time.Duration, interval time.Duration)
}
