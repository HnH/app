# onion

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