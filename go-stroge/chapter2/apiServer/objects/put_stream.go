package objects

import (
	"chapter3/apiServer/heartbeat"
	"chapter3/objectstream"
	"errors"
)

func putStream(hash string, size int64) (*objectstream.TempPutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, errors.New("cannot find any dataServer")
	}

	return objectstream.NewTempPutStream(server, hash, size)
}
