setup:
	go install github.com/d5/tengo/cmd/tengo
	go install github.com/d5/tengobench/cmd/gofib
	go install github.com/d5/tengobench/cmd/gofibtc
	go get -u github.com/yuin/gopher-lua/cmd/glua
	go get -u github.com/robertkrimen/otto/otto
	go get -u github.com/d5/tengobench/cmd/go-lua
	go get -u github.com/mattn/anko
	go get -u go.starlark.net/cmd/starlark
	go get -u github.com/go-python/gpython
	go get -u github.com/dop251/goja/goja
	go get -u github.com/containous/yaegi/cmd/yaegi

vet:
	@go vet ./...

start: vet
	@go run ./cmd/bench ./code

fmt:
	@go fmt ./...
