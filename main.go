package main

import (
	"fmt"

	"github.com/moroii69/bencode-foo/foo"
)

func main() {

	tests := []string{
		"4:spam",                   // string
		"i42e",                     // int
		"l4:spam4:eggsi7ee",        // list
		"d3:cow3:moo4:spam4:eggse", // dict
	}

	for _, t := range tests {
		fmt.Println("input:", t)

		v, err := foo.Decode([]byte(t))
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		fmt.Printf("output: %#v\n", v)
		fmt.Println("note: 0x.. values above are hex bytes (raw binary data)")

		// if it's a string ([]byte), also print as text
		if b, ok := v.([]byte); ok {
			fmt.Println("as string:", string(b))
			fmt.Println("reason: bencode strings are binary-safe, so decoder returns []byte")
		}

		fmt.Println("-----")
	}
}
