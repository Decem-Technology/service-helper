package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/olivere/elastic/v7"
)

var elasticClient *elastic.Client

// Elastic elasticsearch management olive
type Elastic struct {
}

// CreateElasticConnection init elasticsearch connection
func CreateElasticConnection() {
	addrs := strings.Split(os.Getenv("ELASTICSEARCH_HOST"), ";")
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(addrs...),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(1 * time.Minute),
		elastic.SetRetrier(newRetrier()),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetBasicAuth(os.Getenv("ELASTICSEARCH_USERNAME"), os.Getenv("ELASTICSEARCH_PASSWORD")),
	}
	if os.Getenv("ES_DEBUG") == "true" {
		options = append(options, elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	}
	client, err := elastic.NewClient(
		options...,
	)
	if err != nil {
		panic(fmt.Sprintf("[Elasticsearch] Error creating the client: %v", err))
	}
	elasticClient = client
	fmt.Println("[ElasticSearch] connected")
}

// Client get Elasticsearch client
func (ctl *Elastic) Client() *elastic.Client {
	return elasticClient
}

type retrier struct {
	backoff elastic.Backoff
}

func newRetrier() *retrier {
	return &retrier{
		backoff: elastic.NewExponentialBackoff(10*time.Millisecond, 8*time.Second),
	}
}

func (r *retrier) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {
	// Fail hard on a specific error
	if err == syscall.ECONNREFUSED {
		return 0, false, errors.New("Elasticsearch or network down")
	}

	// Stop after 5 retries
	if retry >= 5 {
		return 0, false, nil
	}

	// Let the backoff strategy decide how long to wait and whether to stop
	wait, stop := r.backoff.Next(retry)
	return wait, stop, nil
}
