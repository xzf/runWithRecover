package errMock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
)

func recoverPanic() (result string) {
	panicInfo := recover()
	if panicInfo == nil {
		return
	}
	byteSlice, err := json.MarshalIndent(panicInfo, "", "\t")
	if err != nil {
		fmt.Println("[ygbu6nt38t] recoverPanic json.MarshalIndent err:", err)
		result = fmt.Sprintf("%s", panicInfo)
		return
	}
	result = string(byteSlice)
	return result
}

func readStack() string {
	content := make([]byte, 102400)
	index := runtime.Stack(content, true)
	if index > 0 {
		end := bytes.Index(content[1:], []byte(`
goroutine `))
		if end > 0 {
			content = content[:end]
		}
	}
	return string(content)
}
