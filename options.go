package onion

import (
	"github.com/HnH/di"
	"github.com/rs/zerolog"
)

// Option represents single option type
type Option func(Options)

// Options represents a target for applying an Option
type Options interface {
	SetContainer(di.Container)
	SetLogger(zerolog.Logger)
	SetProviders(...di.Provider)
	SetLayers(...Layer)
}

// WithContainer creates a container Option
func WithContainer(container di.Container) Option {
	return func(o Options) {
		o.SetContainer(container)
	}
}

// WithLogger creates  a logger Option
func WithLogger(log zerolog.Logger) Option {
	return func(o Options) {
		o.SetLogger(log)
	}
}

// WithProviders creates a providers Option
func WithProviders(providers ...di.Provider) Option {
	return func(o Options) {
		o.SetProviders(providers...)
	}
}

// WithLayers creates a layers Option
func WithLayers(layers ...Layer) Option {
	return func(o Options) {
		o.SetLayers(layers...)
	}
}
