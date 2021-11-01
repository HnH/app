package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/HnH/di"
	"github.com/rs/zerolog"
)

// Application interface
type Application interface {
	Listen() error
}

// Config interface
type Config interface {
	Register(di.Container) error
}

// New creates new application
func New(ctx context.Context, options ...Option) Application {
	var app = &application{
		context: ctx,
	}

	for _, opt := range options {
		opt(app)
	}

	return app
}

type application struct {
	context   context.Context
	container di.Container
	logger    zerolog.Logger
	config    Config
	layers    layers
}

// SetContainer sets container for Application
func (self *application) SetContainer(container di.Container) {
	self.container = container
}

// SetLogger sets logger for Application
func (self *application) SetLogger(log zerolog.Logger) {
	self.logger = log
}

// SetConfig sets Config for Application
func (self *application) SetConfig(cfg Config) {
	self.config = cfg
}

// SetLayers sets layers for Application
func (self *application) SetLayers(layers ...Layer) {
	self.layers = layers
}

func (self *application) init() (err error) {
	if self.container == nil {
		self.container = di.NewContainer()
	}

	self.context = di.CtxWith(self.context, self.container).Raw()

	if self.config != nil {
		if err = self.config.Register(self.container); err != nil {
			return
		}
	}

	if err = self.container.Implementation(self.logger); err != nil {
		return
	}

	return self.layers.initialize(self.context)
}

// Listen for system signals
func (self *application) Listen() (err error) {
	if err = self.init(); err != nil {
		return
	}

	self.logger.Info().Msg("application instance started")

	var chanSignal = make(chan os.Signal, 1)
	signal.Notify(chanSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for sig := range chanSignal {
		self.logger.Info().Str("SIGNAL", sig.String()).Msg("process termination signal received")

		if err := self.layers.shutdown(self.context); err != nil {
			self.logger.Error().Err(err)
		}

		signal.Stop(chanSignal)
		close(chanSignal)
	}

	self.logger.Info().Msg("application instance terminated")

	return
}
