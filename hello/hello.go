package main

import (
	"./pkg"
	"fmt"
	"runtime"
)

func myfunc() {
	fmt.Println("in myfunc")
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc, file, line, ok)
	pc, file, line, ok = runtime.Caller(1)
	fmt.Println(pc, file, line, ok)
	fmt.Println("return myfunc")
}

func main() {
	fmt.Println("Hello, Wold")
	fmt.Println("in main")
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc, file, line, ok)
	myfunc()
	pkg.Pkg1()
	fmt.Println("return main")
}
