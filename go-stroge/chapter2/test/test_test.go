package test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(putHandler))
	defer s.Close()
	fmt.Println(s.URL[7:])
	req, _ := newfileUploadRequest(s.URL[7:], "/Users/cjm/1.go")
	client := &http.Client{}
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(resp)
	}
}
func putHandler(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create("/Users/cjm/2.go")
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer f.Close()
	io.Copy(f, r.Body)
}

func newfileUploadRequest(uri string, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	_, err = io.Copy(body, file)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", "11")
	return request, err
}
