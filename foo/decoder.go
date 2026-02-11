package foo

func NewDecoder(b []byte) *Decoder {
	return &Decoder{data: b} // this creates the struct, pos defaults to 0
}

// main api
func Decode(b []byte) (any, error) {
	d := NewDecoder(b)
	return d.decode() // private method in parser
}	