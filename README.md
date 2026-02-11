# bencode-foo

a minimal bencode decoder written in go.

this project exists mainly to understand how bencoding works internally
(strings, ints, lists, dicts) using a manual cursor + recursive parsing.

just learning + experimentation (as part of learning working of bittorrent internals)

---

## features

- decodes bencode strings, integers, lists, and dictionaries
- binary-safe (strings are returned as []byte)
- recursive descent parser
- minimal main.go for quick testing

---

## structure
```
bencode-foo/
├── main.go # tiny test runner
└── foo/ # decoder package
    ├── types.go # decoder struct
    ├── decoder.go # public api
    └── parser.go # parsing logic
```

---

## run
```bash
go run .
```

## example
```go
v, err := foo.Decode([]byte("4:spam"))
```

returns:
```
[]byte("spam")
```
