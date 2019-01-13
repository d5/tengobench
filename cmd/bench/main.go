package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"
)

func main() {
	dir := os.Args[1]

	fmt.Println("| | fib(35) | fibt(35) |  Type  |")
	fmt.Println("| :--- |    ---: |     ---: |  :---: |")

	// Go
	fmt.Printf("| Go | `%s` | `%s` | Go (native) |\n",
		formatNum(execCommand("gofib")),
		formatNum(execCommand("gofibtc")))

	// Tengo
	fmt.Printf("| [**Tengo**](https://github.com/d5/tengo) | `%s` | `%s` | Go-VM |\n",
		formatNum(execCommand("tengo", dir+"/fib.tengo")),
		formatNum(execCommand("tengo", dir+"/fibtc.tengo")))

	// Lua
	fmt.Printf("| Lua | `%s` | `%s` | Lua (native) |\n",
		formatNum(execCommand("lua", dir+"/fib.lua")),
		formatNum(execCommand("lua", dir+"/fibtc.lua")))

	// go-lua
	fmt.Printf("| [go-lua](https://github.com/Shopify/go-lua) | `%s` | `%s` | Go-Lua-VM |\n",
		formatNum(execCommand("go-lua", dir+"/fib.lua")),
		formatNum(execCommand("go-lua", dir+"/fibtc.lua")))

	// GopherLua (glua)
	fmt.Printf("| [GopherLua](https://github.com/yuin/gopher-lua) | `%s` | `%s` | Go-Lua-VM |\n",
		formatNum(execCommand("glua", dir+"/fib.lua")),
		formatNum(execCommand("glua", dir+"/fibtc.lua")))

	// Python
	fmt.Printf("| Python | `%s` | `%s` | Python (native) |\n",
		formatNum(execCommand("python", dir+"/fib.py")),
		formatNum(execCommand("python", dir+"/fibtc.py")))

	// otto
	fmt.Printf("| [otto](https://github.com/robertkrimen/otto) | `%s` | `%s` | Go-JS-Interpreter |\n",
		formatNum(execCommand("otto", dir+"/fib.otto")),
		formatNum(execCommand("otto", dir+"/fibtc.otto")))

	// Anko
	fmt.Printf("| [Anko](https://github.com/mattn/anko) | `%s` | `%s` | Go-Interpreter |\n",
		formatNum(execCommand("anko", dir+"/fib.ank")),
		formatNum(execCommand("anko", dir+"/fibtc.ank")))
}

func execCommand(name string, args ...string) int64 {
	cmd := exec.Command(name, args...)

	start := time.Now()

	_ = cmd.Run()

	return time.Since(start).Nanoseconds()
}

func formatNum(n int64) string {
	if n == 0 {
		return "0"
	}

	var neg bool
	a := n
	if a < 0 {
		neg = true
		a = -a
	}

	digits := int(math.Log10(float64(a))) + 1
	outlen := digits + (digits-1)/3
	if neg {
		outlen += 1
	}

	out := make([]rune, outlen, outlen)

	var i, j int
	for a > 0 {
		d := a % 10

		out[outlen-j-1] = '0' + rune(d)
		i++
		j++

		if i%3 == 0 && j < outlen {
			out[outlen-j-1] = ','
			j++
		}

		a /= 10
	}

	if neg {
		out[0] = '-'
	}

	return string(out)
}
