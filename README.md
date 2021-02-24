# muxutil

A collections of helpers for gorilla/mux and Go's net/http.

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
