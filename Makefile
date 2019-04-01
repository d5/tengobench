setup:
	go install github.com/d5/tengo/cmd/tengo
	go install github.com/d5/tengobench/cmd/gofib
	go install github.com/d5/tengobench/cmd/gofibtc
	go get github.com/yuin/gopher-lua/cmd/glua
	go get github.com/robertkrimen/otto/otto
	go get github.com/d5/tengobench/cmd/go-lua
	go get github.com/mattn/anko
	go get go.starlark.net/cmd/starlark
	go get github.com/go-python/gpython
	go get github.com/dop251/goja/goja

vet:
	@go vet ./...

start: vet
	@go run ./cmd/bench ./code

fmt:
	@go fmt ./...
