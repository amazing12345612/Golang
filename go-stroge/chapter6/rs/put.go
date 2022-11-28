package rs

import (
	"chapter6/objectstream"
	"fmt"
	"io"
)

type RSPutStream struct {
	*encoder
}

func NewRSPutStream(dataServers []string, hash string, size int64) (*RSPutStream, error) {
	if len(dataServers) != ALL_SHARDS {
		return nil, fmt.Errorf("dataServers number mismatch")
	}
	//每片多少字节
	perShard := (size + DATA_SHARDS - 1) / DATA_SHARDS
	//假设是100000字节，那么perShard就是25000
	writers := make([]io.Writer, ALL_SHARDS)
	var e error
	for i := range writers {
		//向temp节点发送信息
		writers[i], e = objectstream.NewTempPutStream(dataServers[i],
			fmt.Sprintf("%s.%d", hash, i), perShard)
		if e != nil {
			return nil, e
		}
	}
	enc := NewEncoder(writers)

	return &RSPutStream{enc}, nil
}

func (s *RSPutStream) Commit(success bool) {
	s.Flush()
	for i := range s.writers {
		s.writers[i].(*objectstream.TempPutStream).Commit(success)
	}
}
