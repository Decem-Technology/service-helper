package bootstrap

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var esClient *elasticsearch.Client

// Elasticsearch elasticsearch management
type Elasticsearch struct {
}

// CreateElasticsearchConnection init elasticsearch connection
func CreateElasticsearchConnection() {
	addrs := strings.Split(os.Getenv("ELASTICSEARCH_HOST"), ";")
	cfg := elasticsearch.Config{
		Addresses: addrs,
		Username:  os.Getenv("ELASTICSEARCH_USERNAME"),
		Password:  os.Getenv("ELASTICSEARCH_PASSWORD"),
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(fmt.Sprintf("[Elasticsearch] Error creating the client: %v", err))
	}
	esClient = es

	res, err := es.Info()
	if err != nil {
		panic(fmt.Sprintf("[Elasticsearch] Error getting response: %s", err))
	}
	defer res.Body.Close()
	fmt.Println("[ElasticSearch] connected")
}

// Client get Elasticsearch client
func (ctl *Elasticsearch) Client() *elasticsearch.Client {
	return esClient
}

// SearchBuiderData build search data
func (ctl *Elasticsearch) SearchBuiderData(ctx context.Context, read *strings.Reader, index string) (*esapi.Response, error) {
	res, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(read),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SearchData search data
func (ctl *Elasticsearch) SearchData(ctx context.Context, read *strings.Reader, index string, response interface{}) (*esapi.Response, error) {
	res, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(read),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return res, err
	}
	return res, nil
}

// DecodeResponse decode response model or map
func (ctl *Elasticsearch) DecodeResponse(body io.ReadCloser, response interface{}) error {
	if err := json.NewDecoder(body).Decode(response); err != nil {
		return err
	}
	return nil
}
