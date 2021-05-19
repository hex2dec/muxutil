# muxutil

A collections of helpers for gorilla/mux and Go's net/http.

## Handlers

### Health Check Handler

```go
func main() {
        // ...

        r := mux.NewRouter()

        hc := handler.NewHealthCheck()
        r.HandleFunc("/healthcheck", hc.CheckStatus).Methods(http.MethodGet)

        // ...
}
```

## Middlewares

### Server Header

```go
sh := &middleware.ServerHeader{
	Name: "TFCloud Go Server",
}
r.Use(sh.Middleware)
```

## Renders

### JSON

```go
func jsonRenderHandler(w http.ResponseWriter, r *http.Request) {
	d := map[string]interface{} {
		"team": "tfcloud-go",
	}
	if err := render.JSON(w, http.StatusOK, d); err != nil {
		// do something
	}
}
```

### Text

```go
func textRenderHandler(w http.ResponseWriter, r *http.Request) {
	d := "tfcloud-go"
	if err := render.Text(w, http.StatusOK, d); err != nil {
		// do something
	}
}
```

## License

Licensed under the [Apache 2.0](./LICENSE). The TFCloud Go Team.
