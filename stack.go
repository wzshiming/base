package base

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

func DebugStack() {
	if !DEBUG_ {
		return
	}
	ret := ""
	for i := 4; i < 10; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			ret += fmt.Sprintf("frame %v: < func: %v, file: %v, line: %v >\n", i, runtime.FuncForPC(funcName).Name(), file, line)
		}
	}
	DEBUG(ret)
}

func PrintPanicStack() {
	if x := recover(); x != nil {
		PANIC(x)
		DebugStack()
	}
}

func PanicErr(err *error, msg ...string) {
	if x := recover(); x != nil {
		PANIC(x)
		if len(msg) == 0 {
			*err = errors.New("panic")
		} else {
			*err = errors.New(strings.Join(msg, " "))
		}
	}
}
