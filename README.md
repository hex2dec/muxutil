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

## Helpers

### Response helper

Response with JSON format.

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
	// ...

	// response error with JSON format
	err := errors.New("bar")
	helper.ResponseErrorJSON(w, http.StatusBadRequest, err)

	// response data with JSON format
	d := map[string]interface{}{
		"foo": "bar",
	}
	helper.ResponseJSON(w, http.StatusOK, d)
}
```

Response with text format.

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
	// ...

	// response error with text format
	err := errors.New("bar")
	helper.ResponseErrorText(w, http.StatusBadRequest, err)

	// response data with text format
	d := "bar"
	helper.ResponseText(w, http.StatusOK, d)
}
```

### Parse request & response body

Parse request body with JSON format.

```go
type Foo struct {
	Name string
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	// ...

	foo := &Foo{}
	if err := helper.ParseRequestBodyJSON(w, r, foo); err != nil {
		helper.ResponseErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println(foo.Name)

	// ...
}
```

Parse response body with JSON format.

```go
type Foo struct {
	Name string
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	// ...

	resp, _ := http.DefaultClient.Do(req)

	foo := &Foo{}
	if err := helper.ParseResponseBodyJSON(resp, foo); err != nil {
		// ...
	}

	fmt.Println(foo.Name)

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

## Server

```go
func main() {
	r := mux.NewRouter()

	s := server.NewServer("127.0.0.1", "8080")
	_ = s.Run(r)
}
```

## License

Licensed under the [Apache 2.0](./LICENSE). The TFCloud Go Team.
