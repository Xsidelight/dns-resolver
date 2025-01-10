package model

import (
	"bytes"
	"encoding/binary"
)

type Message struct {
	Header    Header
	Questions []Question
}

func (m *Message) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, m.Header.ID); err != nil {
		return nil, err
	}

	var flags uint16
	if m.Header.GetQR() {
		flags |= 1 << QR_BIT
	}
	flags |= 1 << RD_BIT
	if err := binary.Write(buf, binary.BigEndian, flags); err != nil {
		return nil, err
	}

	qdcount := uint16(len(m.Questions))
	ancount := uint16(0)
	nscount := uint16(0)
	arcount := uint16(0)
	if err := binary.Write(buf, binary.BigEndian, qdcount); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, ancount); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, nscount); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, arcount); err != nil {
		return nil, err
	}

	for _, q := range m.Questions {
		encodedName := q.EncodeName(q.Name)
		if _, err := buf.Write(encodedName); err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.BigEndian, q.Type); err != nil {
			return nil, err
		}
		if err := binary.Write(buf, binary.BigEndian, q.Class); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
