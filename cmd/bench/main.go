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
	fmt.Printf("| Go | `%sms` | `%sms` | Go (native) |\n",
		formatNum(execCommand("gofib")),
		formatNum(execCommand("gofibtc")))

	// Tengo
	fmt.Printf("| [**Tengo**](https://github.com/d5/tengo) | `%sms` | `%sms` | VM on Go |\n",
		formatNum(execCommand("tengo", dir+"/fib.tengo")),
		formatNum(execCommand("tengo", dir+"/fibtc.tengo")))

	// Lua
	fmt.Printf("| Lua | `%sms` | `%sms` | Lua (native) |\n",
		formatNum(execCommand("lua", dir+"/fib.lua")),
		formatNum(execCommand("lua", dir+"/fibtc.lua")))

	// go-lua
	fmt.Printf("| [go-lua](https://github.com/Shopify/go-lua) | `%sms` | `%sms` | Lua VM on Go |\n",
		formatNum(execCommand("go-lua", dir+"/fib.lua")),
		formatNum(execCommand("go-lua", dir+"/fibtc.lua")))

	// GopherLua (glua)
	fmt.Printf("| [GopherLua](https://github.com/yuin/gopher-lua) | `%sms` | `%sms` | Lua VM on Go |\n",
		formatNum(execCommand("glua", dir+"/fib.lua")),
		formatNum(execCommand("glua", dir+"/fibtc.lua")))

	// Python
	fmt.Printf("| Python | `%sms` | `%sms` | Python (native) |\n",
		formatNum(execCommand("python", dir+"/fib.py")),
		formatNum(execCommand("python", dir+"/fibtc.py")))

	// Gpython
	fmt.Printf("| [gpython](https://github.com/go-python/gpython) | `%sms` | `%sms` | Python Interpreter on Go |\n",
		formatNum(execCommand("gpython", dir+"/fib.py")),
		formatNum(execCommand("gpython", dir+"/fibtc.py")))

	// Starlark
	fmt.Printf("| [starlark-go](https://github.com/google/starlark-go) | `%sms` | `%sms` | Python-like Interpreter on Go |\n",
		formatNum(execCommand("starlark", "-recursion", dir+"/fib.star")),
		formatNum(execCommand("starlark", "-recursion", dir+"/fibtc.star")))

	// otto
	fmt.Printf("| [otto](https://github.com/robertkrimen/otto) | `%sms` | `%sms` | JS Interpreter on Go |\n",
		formatNum(execCommand("otto", dir+"/fib.otto")),
		formatNum(execCommand("otto", dir+"/fibtc.otto")))

	// Anko
	fmt.Printf("| [Anko](https://github.com/mattn/anko) | `%sms` | `%sms` | Interpreter on Go |\n",
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
	// change ns to ms (with rounding)
	n = (n + 500000) / 1000000

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
