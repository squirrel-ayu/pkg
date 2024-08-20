package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"runtime"
	"strings"
)

// GetStack 获取堆栈信息
func GetStack() string {
	return fmt.Sprintf("%s", stack())
}

// stack 获取堆栈信息
func stack() []byte {
	buf := make([]byte, 9128)
	n := runtime.Stack(buf, false)
	return buf[:n]
}

// GetFileName 获取当前文件名
func GetFileName() string {
	_, fileName, _, ok := runtime.Caller(1)
	if ok {
		return fileName
	} else {
		return ""
	}
}

// GetFuncName 获取当前函数名
func GetFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		return runtime.FuncForPC(pc).Name()
	} else {
		return ""
	}
}

// GetCallerFuncName 获取调用者的函数名
func GetCallerFuncName() string {
	pc, _, _, ok := runtime.Caller(2)
	if ok {
		return runtime.FuncForPC(pc).Name()
	} else {
		return ""
	}
}

// ReceiveStruct 格式化打印结构体
func ReceiveStruct(t interface{}) {
	b, err := jsoniter.Marshal(t)
	if err != nil {
		log.Fatalln(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	_, _ = out.WriteTo(os.Stdout)
}

// GetGoroutineId 获取协程id
func GetGoroutineId() string {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	infos := strings.Split(ByteSliceToString(buf[:n]), ":")
	if len(infos) > 0 {
		return infos[0]
	}
	return ""
}
