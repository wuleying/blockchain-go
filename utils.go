package main

import (
	"bytes"
	"encoding/binary"
	"github.com/go-clog/clog"
)

// IntToHex 整数转十六进制
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		clog.Fatal(ClogSkipDisplayInfo, err.Error())
	}

	return buff.Bytes()
}
