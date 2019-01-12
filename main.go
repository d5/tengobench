package main

import (
	"fmt"
	"math"
	"os/exec"
	"time"

	"github.com/Shopify/go-lua"
)

func main() {
	// Go
	start := time.Now()
	fib(35)
	fmt.Printf("Go     fib(35)  :  %15s\n", formatNum(time.Since(start).Nanoseconds()))
	start = time.Now()
	fibTC(35, 0, 1)
	fmt.Printf("Go     fibtc(35):  %15s\n", formatNum(time.Since(start).Nanoseconds()))

	// Python
	fmt.Printf("Python fib(35)  :  %15s\n", formatNum(execCommand("python", "code/fib.py")))
	fmt.Printf("Python fibtc(35):  %15s\n", formatNum(execCommand("python", "code/fibtc.py")))

	// Anko
	fmt.Printf("Anko   fib(35)  :  %15s\n", formatNum(execCommand("anko", "code/fib.ank")))
	fmt.Printf("Anko   fibtc(35):  %15s\n", formatNum(execCommand("anko", "code/fibtc.ank")))

	// otto
	fmt.Printf("otto   fib(35)  :  %15s\n", formatNum(execCommand("otto", "code/fib.otto")))
	fmt.Printf("otto   fibtc(35):  %15s\n", formatNum(execCommand("otto", "code/fibtc.otto")))

	// Lua
	fmt.Printf("Lua    fib(35)  :  %15s\n", formatNum(execCommand("lua", "code/fib.lua")))
	fmt.Printf("Lua    fibtc(35):  %15s\n", formatNum(execCommand("lua", "code/fibtc.lua")))

	// GopherLua (glua)
	fmt.Printf("glua   fib(35)  :  %15s\n", formatNum(execCommand("glua", "code/fib.lua")))
	fmt.Printf("glua   fibtc(35):  %15s\n", formatNum(execCommand("glua", "code/fibtc.lua")))

	// go-lua
	fmt.Printf("go-lua fib(35)  :  %15s\n", formatNum(runGoLua("code/fib.lua")))
	fmt.Printf("go-lua fibtc(35):  %15s\n", formatNum(runGoLua("code/fibtc.lua")))

	// Tengo
	fmt.Printf("Tengo  fib(35)  :  %15s\n", formatNum(execCommand("tengo", "code/fib.tengo")))
	fmt.Printf("Tengo  fibtc(35):  %15s\n", formatNum(execCommand("tengo", "code/fibtc.tengo")))
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

func runGoLua(f string) int64 {
	l := lua.NewState()
	//lua.OpenLibraries(l)

	start := time.Now()

	_ = lua.DoFile(l, f)

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
