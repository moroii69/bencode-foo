package foo

import (
	"errors"
	"fmt"
	"strconv"
)

func (d *Decoder) decoder() (any, error) {
	if d.pos >= len(d.data) {
		return nil, errors.New("end of file..")
	}

	currentByte := d.data[d.pos]
	// i -> integer, l -> list, d-> dictionary, 0-9(length) -> string
	switch currentByte {
	case 'i':
		return d.decodeInt()
	case 'l':
		return d.decodeList()
	case 'd':
		return d.decodeDict()
	default:
		if currentByte >= '0' && currentByte <= '9' {
			return d.decodeStr()
		}
	}
	return nil, fmt.Errorf("invalid token")
}

func (d *Decoder) decodeInt() (int64, error) {
	d.pos++        // skip 'i'
	start := d.pos // mark where number begins. will slice from here...

	for d.data[d.pos] != 'e' { // assuming we have vald input
		d.pos++ // move forward.. byte-by-byte
	}

	numberBytes := d.data[start:d.pos]                       // extract bytes.
	numberString := string(numberBytes)                      // need to convert to string for strconv to parse
	parsedNum, err := strconv.ParseInt(numberString, 64, 10) // base 10, 64 bit signed

	if err != nil {
		return 0, err
	}

	// move cursor past 'e'
	d.pos++

	return parsedNum, nil
}

func (d *Decoder) decodeList() ([]any, error){
	d.pos++ // skip 'l'
	var list []any
	
	// currentByte := d.data[d.pos]
	// commented out this cus we have to read currentByte on moving cursor, not just once.
	
	for d.data[d.pos] != 'e' {
		v, err := d.decode()
		
		if err != nil {
			return nil, err
		}
		list = append(list,v) // adds decoded val to slice. prev_list + new val. 		
	} // exits when 'e' is current byte.
	
	// move cursor past 'e'
	d.pos++
	return list, nil	
}