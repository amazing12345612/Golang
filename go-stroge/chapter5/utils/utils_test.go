package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
	file, _ := os.Open("/Users/cjm/1.go")
	fmt.Println(CalculateHash(file))
}
