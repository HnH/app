package mocks

import (
	"context"

	"github.com/HnH/di"
)

type Provider struct {
	ID string
}

func (p *Provider) Provide(container di.Container) error {
	return nil
}

type Layer struct {
	ID          string
	Initialized bool
}

func (l *Layer) Initialize(context.Context) error {
	l.Initialized = true
	return nil
}

func (l *Layer) Shutdown() <-chan error {
	var out = make(chan error)
	close(out)
	return out
}
