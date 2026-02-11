package foo

import (
	"errors"
	"fmt"
	"strconv"
)

func (d *Decoder) decode() (any, error) {
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
	parsedNum, err := strconv.ParseInt(numberString, 10, 64) // base 10, 64 bit signed

	if err != nil {
		return 0, err
	}

	// move cursor past 'e'
	d.pos++

	return parsedNum, nil
}

func (d *Decoder) decodeList() ([]any, error) {
	d.pos++ // skip 'l'
	var list []any

	// currentByte := d.data[d.pos]
	// commented out this cus we have to read currentByte on moving cursor, not just once.

	for d.data[d.pos] != 'e' {
		v, err := d.decode()

		if err != nil {
			return nil, err
		}
		list = append(list, v) // adds decoded val to slice. prev_list + new val.
	} // exits when 'e' is current byte.

	// move cursor past 'e'
	d.pos++
	return list, nil
}

func (d *Decoder) decodeDict() (map[string]any, error) {
	d.pos++                   // skip 'd'
	m := make(map[string]any) // create empty map. we grow it dynamically

	for d.data[d.pos] != 'e' {

		// decode the key
		k, err := d.decodeStr() // dictionary keys are strings
		if err != nil {
			return nil, err
		}

		// decode val
		v, err := d.decode()
		if err != nil {
			return nil, err
		}

		m[string(k)] = v // store the kv pair.. ex: "cow" → "moo"
	}

	// move cursor past 'e'
	d.pos++

	// return the map
	return m, nil
}

func (d *Decoder) decodeStr() ([]byte, error) {
	// here we dont do pos++ cus initial val is the length
	start := d.pos

	for d.data[d.pos] != ':' { // till it reaches :
		d.pos++
	}

	// extract bytes for length
	lengthBytes := d.data[start:d.pos]

	// convert the []byte to string
	lengthStr := string(lengthBytes) // cus strconv expects string

	//convert str to int
	length, err := strconv.Atoi(lengthStr) // same as parseInt used eariler
	if err != nil || length < 0 {
		return nil, fmt.Errorf("invalid string length")
	}

	// skip past :
	d.pos++

	if d.pos+length > len(d.data) {
		return nil, fmt.Errorf("string length exceeds input size")
	}

	// slice out the actual string bytes
	result := d.data[d.pos : d.pos+length]

	// advance cursor past string
	d.pos += length
	return result, nil
}
