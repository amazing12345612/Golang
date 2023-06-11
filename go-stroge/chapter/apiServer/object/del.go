package objects

import (
	"chapter/es"
	"log"
	"net/http"
	"strings"
)

// Url = 127.0.0.1:12345/
func del(w http.ResponseWriter, r *http.Request) {

	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	version, e := es.SearchLatestVersion(name)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	e = es.PutMetadata(name, version.Version+1, 0, "")
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
