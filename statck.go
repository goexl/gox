package gox

import (
	"fmt"
	"runtime"
	"strings"
)

var _ = Stack

// Stack 返回当前执行堆栈
func Stack() string {
	stack := 10
	skip := 4
	callers := make([]uintptr, stack+1)
	count := runtime.Callers(skip+2, callers)
	frames := runtime.CallersFrames(callers[:count])

	stacks := make([]string, 0, stack)
	for {
		frame, more := frames.Next()
		stacks = append(stacks, fmt.Sprintf("%s()\n\t%s:%d", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}

	return strings.Join(stacks, "\n")
}
