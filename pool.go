package dom

import (
	"bytes"
	"sync"
)

var bufPool = &sync.Pool{New: func() any { return new(bytes.Buffer) }}

func getBuffer() *bytes.Buffer {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	return b
}

func putBuffer(b *bytes.Buffer) {
	bufPool.Put(b)
}
