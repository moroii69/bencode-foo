package foo

type Decoder struct {
	data []byte 
	pos  int // current read pointer.. instead of slicing repeatedly, we move an index. (much faster)
}
