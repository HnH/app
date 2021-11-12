package onion

import (
	"context"
	"reflect"

	"github.com/HnH/di"
	"github.com/rs/zerolog"
)

// Layer is an Application layer abstraction
type Layer interface {
	Initialize(context.Context) error
}

// Shutdowner is a layer that has his state shutdown logic
type Shutdowner interface {
	Shutdown() <-chan error
}

type layers []Layer

func (lrs layers) initialize(ctx context.Context) (err error) {
	var log zerolog.Logger
	if err = di.Ctx(ctx).Resolver().Resolve(&log); err != nil {
		return
	}

	// Initialize application layers
	log.Info().Msgf("starting to initialize %d layer(s)", len(lrs))

	for _, l := range lrs {
		if err = l.Initialize(ctx); err != nil {
			return
		}

		if t, ok := l.(di.Provider); ok {
			if err = t.Provide(di.Ctx(ctx).Container()); err != nil {
				return
			}
		}

		log.Info().Msgf("%s layer initialized", reflect.TypeOf(l).String())
	}

	log.Info().Msgf("all layers successfully initialized")

	return
}

func (lrs layers) shutdown(ctx context.Context) (err error) {
	var log zerolog.Logger
	if err = di.Ctx(ctx).Resolver().Resolve(&log); err != nil {
		return
	}

	log.Info().Msgf("starting to gracefully shut down %d layer(s)", len(lrs))

	// Loop and close layers in reversed order
	for n := len(lrs) - 1; n >= 0; n-- {
		if t, ok := lrs[n].(Shutdowner); ok {
			for err = range t.Shutdown() {
				if err != nil {
					log.Error().Err(err).Send()
				}
			}
		}

		log.Info().Msgf("%s layer gracefully shut down", reflect.TypeOf(lrs[n]).String())
		lrs = lrs[:n] // pop
	}

	log.Info().Msgf("all layers gracefully shut down")

	return
}
