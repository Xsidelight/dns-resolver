package model

import (
	"bytes"
	"strings"
)

type Question struct {
	Name  string
	Type  uint16
	Class uint16
}

func (q *Question) EncodeName(name string) []byte {
	var buf bytes.Buffer
	parts := strings.Split(name, ".")
	for _, part := range parts {
		buf.WriteByte(byte(len(part)))
		buf.WriteString(part)
	}
	buf.WriteByte(0)
	return buf.Bytes()
}
