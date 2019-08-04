package item

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/jdkato/prose.v2"
	"log"
	"processor/model"
	"processor/pkg/elasticsearch"
	es "processor/pkg/elasticsearch/config"
	"processor/workflow/store"
	"strings"
)

func Trigger(esClient *es.ESClient, textDetection []model.TextDetections) {
	item := new(model.Item)

	m := make(map[string]string)

	columnToken := tokenize(textDetection[0].DetectedText)
	columnSize := len(columnToken)

	for idx, ct := range columnToken  {
		key := fmt.Sprintf("%s:%s", idx, ct)
		m[key] = ""
	}
	for _, tdx := range textDetection {
		esSuggestResponse, err := elasticsearch.SuggestSearch(esClient, "item-index", tdx.DetectedText, "title")
		if err != nil {
			log.Fatalf("elasticsearch suggest query failed. text %s, error %v", tdx.DetectedText, err)
		}

		newText := store.NewText(*esSuggestResponse, strings.ToLower(tdx.DetectedText))

		token := tokenize(newText)
		for i := 0; i <= columnSize - 2 ; i++{
			key := fmt.Sprintf("%s:%s", columnSize-i-1, columnToken[columnSize-i-1])
			m[key] = token[len(token)- i - 1]
		}


		if isItemComplete(item) {
			esResponse, err := elasticsearch.Search(esClient, "item-index", "name", item.Name)
			if err != nil {
				log.Fatalf("es search query failed. title %s, err %v ", item.Name, err)
			}
			if len(esResponse.Hits.Hits) == 0 {
				//ingest into elasticsearch item index
				fmt.Printf("item %v", item)
				elasticsearch.IndexDocument(esClient, "item-index", "item-type", uuid.NewV4().String(), item)
			}
		}
	}
}

func findOrder(td string) map[int]string {
	token := tokenize(td)
	m := make(map[int]string)

	for idx, txt := range token  {
		m[idx] = txt
	}
	return m
}

func tokenize(td string) []string  {
	var token []string
	doc, err := prose.NewDocument(td)
	if err != nil {
		log.Fatalf("prose err. err %v ", err)
	}

	for _, tok := range doc.Tokens() {
		token = append(token, tok.Text)
	}
	return token
}

func isItemComplete(item *model.Item)  bool {
	if len(item.Name) == 0  || len(item.Price) == 0 || len(item.Quantity) == 0{
		return false
	}
	return true
}

func reverse(input string) string  {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}