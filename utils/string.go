package utils

import (
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

// StructToJsonStr struct转json
func StructToJsonStr(obj interface{}) string {
	info, err := jsoniter.Marshal(obj)
	if err != nil {
		return ""
	}
	return ByteSliceToString(info)
}

// ByteSliceToString byte切片转string
func ByteSliceToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
