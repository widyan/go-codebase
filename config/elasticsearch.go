package config

import (
	"codebase/go-codebase/helper"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

func MetadataElasticSearch(logger *helper.CustomLogger) (es *elasticsearch.Client) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELK_METADATA"),
		}, Transport: &http.Transport{
			MaxIdleConnsPerHost:   100,
			ResponseHeaderTimeout: time.Second * 3,
			// TLSClientConfig: &tls.Config{
			// 	MinVersion: tls.VersionTLS11,
			// },
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Panic(err)
	}

	return
}
