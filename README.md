[![CircleCI](https://circleci.com/gh/HnH/onion/tree/master.svg?style=svg&circle-token=cd6ef5c602e0f89a80488349a1e4fbe034b8d717)](https://circleci.com/gh/HnH/onion/tree/master)
[![codecov](https://codecov.io/gh/HnH/onion/branch/master/graph/badge.svg)](https://codecov.io/gh/HnH/onion)
[![Go Report Card](https://goreportcard.com/badge/github.com/HnH/onion)](https://goreportcard.com/report/github.com/HnH/onion)
[![GoDoc](https://godoc.org/github.com/HnH/onion?status.svg)](https://godoc.org/github.com/HnH/onion)

# App

```go
var app = onion.New(
    context.Background(),
    onion.WithLogger(log),
    onion.WithContainer(di.NewContainer()),
    onion.WithProviders(cfg),
    onion.WithLayers(infrastructure.New(), domain.New(), application.New(), www.New()),
)

if err = app.Listen(); err != nil {
    log.Error().Err(err).Send()
    os.Exit(1)
}
```