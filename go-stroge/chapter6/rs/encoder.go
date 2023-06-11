package rs

import (
	"github.com/klauspost/reedsolomon"
	"io"
)

type encoder struct {
	writers []io.Writer
	enc     reedsolomon.Encoder
	cache   []byte
}

func NewEncoder(writers []io.Writer) *encoder {
	enc, _ := reedsolomon.New(DATA_SHARDS, PARITY_SHARDS)
	return &encoder{writers, enc, nil}
}

// 若是缓存满了，则写入
func (e *encoder) Write(p []byte) (n int, err error) {
	//输入byte长度 假设为25000
	length := len(p)

	current := 0

	for length != 0 {
		//cache 0 next == 16000
		next := BLOCK_SIZE - len(e.cache)

		//
		if next > length {
			next = length
		}
		//添加满
		e.cache = append(e.cache, p[current:current+next]...)
		if len(e.cache) == BLOCK_SIZE {
			e.Flush()
		}
		//current == 16000
		current += next
		//length == 9000
		length -= next
	}
	return len(p), nil
}

// 调用patch发送
func (e *encoder) Flush() {
	if len(e.cache) == 0 {
		return
	}
	shards, _ := e.enc.Split(e.cache)
	e.enc.Encode(shards)
	for i := range shards {
		e.writers[i].Write(shards[i])
	}
	e.cache = []byte{}
}
