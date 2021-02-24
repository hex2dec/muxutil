# muxutil

A collections of helpers for gorilla/mux and Go's net/http.

## Renders

```go
func jsonRenderHandler(w http.ResponseWriter, r *http.Request) {
	d := map[string]interface{} {
		"name": "Gorilla/Mux",
	}
	if err := render.JSON(w, http.StatusOK, d); err != nil {
		// do something
	}
}
```
