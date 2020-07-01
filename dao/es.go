package dao

import (
	"context"
	"errors"
	"github.com/bitly/go-simplejson"
	"github.com/elastic/go-elasticsearch/esapi"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v8"
	"github.com/shao1f/user.tag.logic/model"
	"log"
	"strconv"
	"strings"
)

type Elastic struct {
	es *elasticsearch7.Client
}

func NewElastic() *Elastic {
	esConfig := elasticsearch7.Config{
		Addresses: []string{"http://49.232.192.87:9200/"},
	}
	esCli, err := elasticsearch7.NewClient(esConfig)
	if err != nil {
		log.Fatal("es new client error,err:", err)
	}
	res, err := esCli.Info()
	if err != nil {
		log.Fatal("get es info err,err:", err)
	}
	if res.IsError() {
		log.Fatal(res.String())
	}
	return &Elastic{
		es: esCli,
	}
}

func (e *Elastic) UploadToES(ctx context.Context, tag *model.UserTag) {
	req := esapi.IndexRequest{
		Index:      "user_tag",
		DocumentID: strconv.Itoa(tag.TagID),
		Body:       strings.NewReader(tag.MustToJSON()),
		Refresh:    "true",
	}
	resp, err := req.Do(context.Background(), e.es)
	if err != nil {
		log.Printf("ESIndexRequestErr: %s", err.Error())
		return
	}

	defer resp.Body.Close()
	if resp.IsError() {
		log.Printf("ESIndexRequestErr: %s", resp.String())
	} else {
		log.Printf("ESIndexRequestOk: %s", resp.String())
	}
}

func (e *Elastic) SearchFromES(ctx context.Context, key string) ([]*model.UserTag, error) {
	// 构建查询语句
	query := model.NormalMap{
		"query": model.NormalMap{
			"match_phrase_prefix": model.NormalMap{
				"tag_name": key,
			},
		},
	}
	jsonBuf := query.MustToJSONBytesBuffer()
	// // 发出查询请求
	// resp, err := e.es.Search(
	// 	e.es.Search.WithContext(context.Background()),
	// 	e.es.Search.WithIndex("user_tag"),
	// 	e.es.Search.WithBody(jsonBuf),
	// )
	req := esapi.SearchRequest{
		Index: []string{"user_tag"},
		Body:  jsonBuf,
	}
	resp, err := req.Do(context.Background(), e.es)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		log.Println(err)
		return nil, errors.New(resp.String())
	}
	js, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	hitsJS := js.GetPath("hits", "hits")
	hits, err := hitsJS.Array()
	if err != nil {
		return nil, err
	}
	hitsLen := len(hits)
	if hitsLen == 0 {
		return []*model.UserTag{}, nil
	}
	tags := make([]*model.UserTag, 0, hitsLen)
	for idx := 0; idx < hitsLen; idx++ {
		sourceJS := hitsJS.GetIndex(idx).Get("_source")

		tagID, err := sourceJS.Get("tag_id").Int()
		if err != nil {
			return nil, err
		}
		tagName, err := sourceJS.Get("tag_name").String()
		if err != nil {
			return nil, err
		}
		tagEntity := &model.UserTag{
			TagID:   tagID,
			TagName: tagName,
		}
		tags = append(tags, tagEntity)
	}
	return tags, nil
}
