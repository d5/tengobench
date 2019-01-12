# Tengo Language Benchmark

To run the benchmark, you will need the following tools installed in your PATH.

- Python
- Lua 
- [Tengo](https://github.com/d5/tengo): `go install github.com/d5/tengo/cmd/tengo`
- [GopherLua](https://github.com/yuin/gopher-lua): `go get github.com/yuin/gopher-lua/cmd/glua`
- [otto](https://github.com/robertkrimen/otto): `go get -v github.com/robertkrimen/otto/otto`
- [go-lua](https://github.com/Shopify/go-lua): `go install github.com/d5/tengobench/cmd/go-lua`
- [Anko](https://github.com/mattn/anko): `go install github.com/mattn/anko`

To start the benchmark:

```bash
make start
```

You can see all the source codes used in this benchmark in [code](https://github.com/d5/tengobench/tree/master/code) directory.
