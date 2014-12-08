package pkg

import (
	"fmt"
	"runtime"
)

func pkg2() {
	fmt.Println("pkg2")
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc, file, line, ok)
	pc, file, line, ok = runtime.Caller(1)
	fmt.Println(pc, file, line, ok)
	pc, file, line, ok = runtime.Caller(2)
	fmt.Println(pc, file, line, ok)
	fmt.Println("return pkg2")
}
