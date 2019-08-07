package utils

import (
	"bytes"
	"runtime"
)

var bytesAsk = []byte{'?'}

func GetGoId() []byte {
	var buf [32]byte
	sb := buf[:]
	runtime.Stack(sb, false)
	i := bytes.IndexByte(sb, ' ')
	if i > 0 {
		j := bytes.IndexByte(buf[i+1:], ' ')
		if j > 0 {
			return buf[i+1 : i+j+1]
		}
	}
	return bytesAsk
}
