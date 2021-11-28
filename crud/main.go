package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var es *elasticsearch.Client

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	var err error
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	//ctx := context.Background()
	ping, err := es.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	defer ping.Body.Close()
	log.Printf("%+v", ping)

	// create index
	if err = mapping(); err != nil {
		log.Fatalln(err)
	}

}

func createIndex() error {

	//request :=esapi.IndicesCreateRequest{
	//	Index:               "crud-index-1",
	//	//Body:                strings.NewReader(`{ "title" : "1" }`),
	//
	//}
	request := esapi.IndexRequest{
		Index:      "crud-index-1",
		Body:       strings.NewReader(`{ "title" : "1" }`),
		DocumentID: "1",
	}

	res, err := request.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.New(res.String())
	}

	log.Printf("%+v", res.String())
	return nil
}

type ii struct {
	Query struct {
		Match struct {
			Title string `json:"title"`
		} `json:"match"`
	} `json:"query"`
}

func getItems() error {
	indexMatch := new(ii)
	indexMatch.Query.Match.Title = "1"

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(*indexMatch); err != nil {
		log.Fatalf("Error encoding query: %s", err)
		return err
	}

	search, err := es.Search(
		es.Search.WithIndex("crud-index-1"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return err
	}
	defer search.Body.Close()
	if search.IsError() {
		return errors.New(search.String())
	}

	var r AutoGenerated
	if err = json.NewDecoder(search.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return err
	}

	log.Printf("111 %+v", r)
	return nil

}

type AutoGenerated struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Title string `json:"title"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type test_mapping struct {
	Mappings struct {
		DynamicDateFormats []string `json:"dynamic_date_formats"`
	} `json:"mappings"`
}

func mapping() error {

	buf := strings.NewReader(`{
		"mappings": {
			"dynamic_date_formats": ["MM/dd/yyyy"]
		}
	}`)

	req := esapi.IndicesPutMappingRequest{
		Index: []string{"my-index-000002"},
		Body:  buf,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}
