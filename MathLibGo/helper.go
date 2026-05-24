package helper

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

type Hashable interface {
	int | string
}

func Hash(s any) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	return b.Bytes()
}

func computeHash_Many[R Hashable](v1 R, v2 ...R) int {
	h := fnv.New64a()
	h.Write(Hash(v1))

	for _, val := range v2 {
		h.Write(Hash(val))
	}

	return int(h.Sum64())
}

func CacheWrapper_ManyToSingle[R Hashable](f func(p1 R, p2 ...R) R) func(p1 R, p2 ...R) any {
	m := make(map[int]*R)

	return func(p1 R, p2 ...R) any {
		inputHash := computeHash_Many(p1, p2...)

		// TODO - Make the key for the map a constant value for any list of input
		if cv, ok := m[inputHash]; ok == true {
			// fmt.Printf("Cached %d! - Hash: %d\n", *cv, inputHash)
			return *cv
		}
		cv := f(p1, p2...)
		m[inputHash] = &cv

		return cv
	}
}

func CacheWrapper_SingleToSingle[R, S any](f func(p R) S) func(p R) S {
	// l, _ := lru.New[any, *S](16384)
	m := make(map[any]*S)

	return func(p R) S {
		// if cv, ok := l.Get(p); ok == true {
		if cv, ok := m[p]; ok {
			// fmt.Printf("Cached %d!\n", *cv)
			return *cv
		}
		cv := f(p)
		// l.Add(p, &cv)
		m[p] = &cv

		return cv
	}
}
