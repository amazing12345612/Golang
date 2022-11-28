package rs

import (
	"bytes"
	"fmt"
	"io"

	"testing"
)

type Cache struct {
	a io.Writer
	b []byte
}

func TestGet(t *testing.T) {
	// file, _ := os.Open("/Users/cjm/1.go")
	c := bytes.NewReader([]byte{1, 2, 3, 4, 5})
	// c := strings.NewReader("1234456")
	a := []byte{}
	io.ReadFull(c, a)
	fmt.Println(a)
}
