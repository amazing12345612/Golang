package es

import (
	"context"
	"fmt"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestT(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println(err)

	p1 := Metadata{Name: "2", Version: 2, Size: 2, Hash: "2"}
	put1, err := client.Index().
		Index("metadata").
		Id("2_2").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

type Employee struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func TestGet(t *testing.T) {
	getMetadata("1", 1)
	getMetadata("1", 2)
}
func TestSearchLatest(t *testing.T) {
	fmt.Println(SearchLatestVersion("1"))
}

func TestGetMedate(t *testing.T) {
	fmt.Println(GetMetadata("1", 1))
	fmt.Println(GetMetadata("1", 0))
}
func TestPutMedate(t *testing.T) {
	PutMetadata("1", 1, 1, "1")
}

func TestAddversion(t *testing.T) {
	AddVersion("1", "1", 1)
}
func TestSearchAllversion(t *testing.T) {
	SearchAllVersions("1", 0, 100)
}
func TestDeledate(t *testing.T) {
	fmt.Println(GetMetadata("1", 2))
	DelMetadata("1", 2)
}
func TestDeleTargetdate(t *testing.T) {
	DelTaggetMetadata("1")
}
