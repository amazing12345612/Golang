package test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	u, _ := url.Parse("https://127.0.0.1/objects/text10/version=1")
	q := u.Query()
	name := strings.Split(u.EscapedPath(), "/")[2]
	versionId := u.Query()["version"]
	fmt.Println(q.Get("version"))
	fmt.Println(name, versionId)
}
