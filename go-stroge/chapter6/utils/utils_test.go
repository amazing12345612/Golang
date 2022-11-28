package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
	file, _ := os.Open("/Users/cjm/1.go")
	fmt.Println(CalculateHash(file))
}

type Test struct {
	Name string `json:"name"`
}

func TestMain(t *testing.T) {
	e := Test{Name: "1234"}
	b, _ := json.Marshal(e)
	c := base64.StdEncoding.EncodeToString(b)
	a, _ := RsaEncrypt([]byte(c))
	fmt.Println(a)
	f, _ := RsaDecrypt(a)
	d, _ := base64.StdEncoding.DecodeString(string(f))
	y := Test{}
	json.Unmarshal(d, &y)
	fmt.Println(y)

}
func TestHash(t *testing.T) {
	file, _ := os.Open("/Users/cjm/Desktop/区块链与隐私保护综述.docx")
	d := CalculateHash(file)
	fmt.Println(d)
}
