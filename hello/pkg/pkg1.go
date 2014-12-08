package pkg

import (
	"fmt"
	"runtime"
)

func Pkg1() {
	fmt.Println("pkg1")
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc, file, line, ok)
	pc, file, line, ok = runtime.Caller(1)
	fmt.Println(pc, file, line, ok)
	pkg2()
	fmt.Println("return pkg1")
}
