package usecases

import (
	"time"

	"github.com/wooseopkim/ck/v2/entities"
)

type InferRemoteTime struct {
	run func(string) (time.Duration, error)
}

func NewInferRemoteTime(run func(string) (time.Duration, error)) *InferRemoteTime {
	return &InferRemoteTime{run: run}
}

func (i *InferRemoteTime) Run(url entities.URL) (time.Duration, error) {
	offset, err := i.run(string(url))
	return offset, err
}
