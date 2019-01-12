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

	fmt.Println("| Lang | fib(35) | fibt(35) |  Type  |")
	fmt.Println("| :--- |    ---: |     ---: |  :---: |")

	// Go
	fmt.Print("| Go |")
	start := time.Now()
	fib(35)
	fmt.Printf(" `%s` |", formatNum(time.Since(start).Nanoseconds()))
	start = time.Now()
	fibTC(35, 0, 1)
	fmt.Printf(" `%s` |", formatNum(time.Since(start).Nanoseconds()))
	fmt.Println(" Go (native) |")

	// Tengo
	fmt.Printf("| **Tengo** | `%s` | `%s` | Go-VM |\n",
		formatNum(execCommand("tengo", dir+"/fib.tengo")),
		formatNum(execCommand("tengo", dir+"/fibtc.tengo")))

	// Lua
	fmt.Printf("| Lua | `%s` | `%s` | Lua (native) |\n",
		formatNum(execCommand("lua", dir+"/fib.lua")),
		formatNum(execCommand("lua", dir+"/fibtc.lua")))

	// go-lua
	fmt.Printf("| go-lua | `%s` | `%s` | Go-Lua-VM |\n",
		formatNum(execCommand("go-lua", dir+"/fib.lua")),
		formatNum(execCommand("go-lua", dir+"/fibtc.lua")))

	// GopherLua (glua)
	fmt.Printf("| GopherLua | `%s` | `%s` | Go-Lua-VM |\n",
		formatNum(execCommand("glua", dir+"/fib.lua")),
		formatNum(execCommand("glua", dir+"/fibtc.lua")))

	// Python
	fmt.Printf("| Python | `%s` | `%s` | Python (native) |\n",
		formatNum(execCommand("python", dir+"/fib.py")),
		formatNum(execCommand("python", dir+"/fibtc.py")))

	// otto
	fmt.Printf("| otto | `%s` | `%s` | Go-JS-Interpreter |\n",
		formatNum(execCommand("otto", dir+"/fib.otto")),
		formatNum(execCommand("otto", dir+"/fibtc.otto")))

	// Anko
	fmt.Printf("| Anko | `%s` | `%s` | Go-Interpreter |\n",
		formatNum(execCommand("anko", dir+"/fib.ank")),
		formatNum(execCommand("anko", dir+"/fibtc.ank")))
}

func execCommand(name string, args ...string) int64 {
	cmd := exec.Command(name, args...)

	start := time.Now()

	_ = cmd.Run()

	return time.Since(start).Nanoseconds()
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func fibTC(n, a, b int) int {
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else {
		return fibTC(n-1, b, a+b)
	}
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
