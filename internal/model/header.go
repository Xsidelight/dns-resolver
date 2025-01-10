package model

type Header struct {
	ID    uint16
	Flags uint16
}

const (
	QR_BIT       = 15
	OPCODE_SHIFT = 11
	AA_BIT       = 10
	TC_BIT       = 9
	RD_BIT       = 8
	RA_BIT       = 7
	Z_SHIFT      = 4
	RCODE_MASK   = 0x000F
)

func (h *Header) SetQR(qr bool) {
	if qr {
		h.Flags |= (1 << QR_BIT)
	} else {
		h.Flags &^= (1 << QR_BIT)
	}
}

func (h *Header) SetRD(rd bool) {
	if rd {
		h.Flags |= (1 << RD_BIT)
	} else {
		h.Flags &^= (1 << RD_BIT)
	}
}

func (h *Header) GetQR() bool {
	return h.Flags&(1<<QR_BIT) != 0
}
