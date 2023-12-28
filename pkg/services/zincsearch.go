package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/disto-stack/j-mails-indexer/pkg/types"
)

type ZincsearchService struct {
	configService *Config
}

type IndexPropertiesValues struct {
	Type          string `json:"type"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store"`
	Highlightable bool   `json:"highlightable"`
}

type IndexProperties struct {
	MessageId IndexPropertiesValues `json:"message_id"`
	Date      IndexPropertiesValues `json:"date"`
	From      IndexPropertiesValues `json:"from"`
	To        IndexPropertiesValues `json:"to"`
	Subject   IndexPropertiesValues `json:"subject"`
	Content   IndexPropertiesValues `json:"content"`
}

type Mappings struct {
	Properties IndexProperties `json:"properties"`
}

type Index struct {
	Name        string   `json:"name"`
	StorageType string   `json:"storage_type"`
	ShardNum    int8     `json:"shard_num"`
	Mappings    Mappings `json:"mappings"`
}

func (z *ZincsearchService) SetDependencies(config *Config) {
	z.configService = config
}

func (z *ZincsearchService) UploadIndex() error {
	index := createIndex()
	indexUrl := z.configService.GetZincsearchUrl() + "/api/" + "index"

	json, err := json.Marshal(index)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", indexUrl, strings.NewReader(string(json)))
	if err != nil {
		return err
	}

	req.SetBasicAuth("admin", "12345678")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	res, err := http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		fmt.Println("Error in petition", string(body))
	}

	defer req.Body.Close()

	fmt.Println("Index created!")
	return nil
}

func (z *ZincsearchService) UploadBulkData(data []types.EmailData) error {
	for _, emailData := range data {
		json, err := json.Marshal(emailData)
		if err != nil {
			log.Fatal(err)
		}

		bulkDataUrl := z.configService.zincsearchUrl + "/api/" + "_bulkv2"
		req, err := http.NewRequest("POST", bulkDataUrl, bytes.NewReader(json))
		if err != nil {
			return err
		}

		req.SetBasicAuth("admin", "12345678")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

		client := &http.Client{}
		res, err := client.Do(req)

		if res.StatusCode != http.StatusOK {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}

			fmt.Println("Error in petition", string(body))
		}

		defer req.Body.Close()
	}

	fmt.Println("Data uploaded!")
	return nil
}

func createIndex() Index {
	index := Index{
		Name:        "email",
		StorageType: "disk",
		ShardNum:    1,
		Mappings: Mappings{
			Properties: IndexProperties{
				MessageId: IndexPropertiesValues{
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: false,
				},
				From: IndexPropertiesValues{
					Type:          "text",
					Index:         false,
					Store:         false,
					Highlightable: true,
				},
				To: IndexPropertiesValues{
					Type:          "text",
					Index:         false,
					Store:         false,
					Highlightable: true,
				},
				Subject: IndexPropertiesValues{
					Type:          "text",
					Index:         false,
					Store:         false,
					Highlightable: true,
				},
				Date: IndexPropertiesValues{
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
				Content: IndexPropertiesValues{
					Type:          "text",
					Index:         true,
					Store:         false,
					Highlightable: true,
				},
			},
		},
	}

	return index
}
