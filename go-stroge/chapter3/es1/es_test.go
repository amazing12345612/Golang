package es1

import (
	"chapter2/config"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestT(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://10.203.209.28:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println(err)

	p1 := Metadata{Name: "1", Version: 2, Size: 1, Hash: "1"}
	put1, err := client.Index().
		Index("metadata").
		Id("1_2").
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
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://10.203.209.28:9200"))
	if err != nil {
		panic(err)
	}
	get, err := client.Get().Index("metadata").Id("1_1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	m := Metadata{}
	json.Unmarshal(get.Source, &m)
	fmt.Println(m.Name)
	if get.Found {
		fmt.Printf("got document %s in version %d from index %s,type %s \n", get.Id, get.Version, get.Index, get.Type)
	}

}

func TestLatestversion(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://10.203.209.28:9200"))
	if err != nil {
		panic(err)
	}
	q := elastic.NewTermQuery("name", "1")
	searchResult, err := client.Search("metadata").Query(q).Sort("version", false).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 Metadata
		r := searchResult.Each(reflect.TypeOf(b1))[0]
		fmt.Println(r)
	}

}
func TestSearchallversion(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.Eshost))
	if err != nil {
		panic(err)
	}

	q := elastic.NewTermQuery("name", "1")
	searchResult, err := client.Search("metadata").Query(q).Sort("name", true).Do(context.Background())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 Metadata
		r := searchResult.Each(reflect.TypeOf(b1))
		for _, item := range r {
			fmt.Println(item.(Metadata))
		}

	}
}
