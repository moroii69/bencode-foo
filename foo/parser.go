package foo

import (
	"errors"
	"fmt"
)

func (d *Decoder) decoder() (any, error) {
	if d.pos >= len(d.data) {
		return nil, errors.New("end of file..")
	}

	// i -> integer, l -> list, d-> dictionary, 0-9(length) -> string
	switch d.data[d.pos] {
	case 'i':
		return d.decodeInt()
	case 'l':
		return d.decodeList()
	case 'd':
		return d.decodeDict()
	default:
		if d.data[d.pos] >= '0' && d.data[d.pos] <= '9' {
			return d.decodeStr()
		}
	}
	return nil, fmt.Errorf("invalid token")
}

// TODO
// implement decoders for int, list, dict, str..
// decodeInt()
// decodeList()
// decodeDict()
// decodeStr()
