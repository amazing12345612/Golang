package es

import (
	"chapter3/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"

	// "strconv"

	"github.com/olivere/elastic/v7"
)

type Metadata struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
	Size    int64  `json:"size"`
	Hash    string `json:"hash"`
}

func getMetadata(name string, versionId int) (meta Metadata, e error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.Eshost))
	if err != nil {
		log.Fatal("连接错误")
	}
	id := name + "_" + strconv.Itoa(versionId)
	get, err := client.Get().Index("metadata").Id(id).Do(context.Background())
	if err != nil {
		log.Fatal("获取错误")
	}
	json.Unmarshal(get.Source, &meta)
	if get.Found {
		fmt.Printf("got document %s in from index %s,type %s \n", get.Id, get.Index, get.Type)
	}
	return
}

func SearchLatestVersion(name string) (meta Metadata, e error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.Eshost))
	if err != nil {
		log.Fatal("连接失败")
	}
	q := elastic.NewTermQuery("name", name)
	searchResult, e := client.Search("metadata").Query(q).Sort("version", false).Do(context.Background())
	if e != nil {
		log.Fatal("获取失败")
	}
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 Metadata
		meta = searchResult.Each(reflect.TypeOf(b1))[0].(Metadata)
	}
	return
}

func GetMetadata(name string, version int) (Metadata, error) {
	if version == 0 {
		return SearchLatestVersion(name)
	}
	return getMetadata(name, version)
}

//发送
func PutMetadata(name string, version int, size int64, hash string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.Eshost))
	if err != nil {
		log.Fatal("连接失败")
	}

	p1 := Metadata{Name: name, Version: version, Size: size, Hash: hash}
	put1, err := client.Index().
		Index("metadata").
		Id(name + "_" + strconv.Itoa(version)).
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if put1.Result == "updated" {
		return PutMetadata(name, version+1, size, hash)
	}

	return nil
}

func AddVersion(name, hash string, size int64) error {
	version, e := SearchLatestVersion(name)
	if e != nil {
		return e
	}
	return PutMetadata(name, version.Version+1, size, hash)
}
func SearchAllVersions(name string, from, size int) (m []Metadata, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.Eshost))
	if err != nil {
		log.Fatal("连接失败")
	}

	q := elastic.NewTermQuery("name", name)
	searchResult, err := client.Search("metadata").From(from).Size(size).Query(q).Sort("version", true).Do(context.Background())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 Metadata
		r := searchResult.Each(reflect.TypeOf(b1))
		for _, item := range r {
			fmt.Println(item.(Metadata))
			m = append(m, item.(Metadata))
		}

	}
	return
}
func DelMetadata(name string, version int) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.Eshost))
	if err != nil {
		log.Fatal("连接失败")
	}
	res, err := client.Delete().Index("metadata").Id(name + "_" + strconv.Itoa(version)).Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("delete result %s", res.Result)
}
func DelTaggetMetadata(name string) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.Eshost))
	if err != nil {
		log.Fatal("连接失败")
	}
	q := elastic.NewTermQuery("name", name)
	res, err := client.DeleteByQuery().Index("metadata").Query(q).Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
