package storage

type Pair struct {
	First  string
	Second int64
}

var Chache = make(map[string]Pair)
