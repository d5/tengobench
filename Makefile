setup:
	go install github.com/d5/tengo/v2/cmd/tengo
	go install github.com/d5/tengobench/cmd/gofib
	go install github.com/d5/tengobench/cmd/gofibtc
	go install github.com/yuin/gopher-lua/cmd/glua
	go install github.com/robertkrimen/otto/otto
	go install github.com/d5/tengobench/cmd/go-lua
	go install github.com/mattn/anko
	go install go.starlark.net/cmd/starlark
	go install github.com/go-python/gpython
	go install github.com/dop251/goja/goja
	go install github.com/containous/yaegi/cmd/yaegi

start:
	@go run ./cmd/bench ./code

fmt:
	@go fmt ./...
