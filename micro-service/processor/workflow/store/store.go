package store

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"processor/model"
	"processor/pkg/elasticsearch"
	es "processor/pkg/elasticsearch/config"
	"strings"
)

func Trigger(esClient *es.ESClient, textDetection []model.TextDetections) *model.Store {
	st := new(model.Store)
	address := ""
	for _, tdx := range textDetection {
		esResponse, err := elasticsearch.SuggestSearch(esClient, "store-index", tdx.DetectedText, "title")
		if err != nil {
			log.Fatalf("elasticsearch suggest query failed. text %s, error %v", tdx.DetectedText, err)
		}

		//fmt.Println(esResponse)

		newText := NewText(*esResponse, strings.ToLower(tdx.DetectedText))
		if len(st.Title) == 0 {
			st.Title = SetTitle(newText)
		} else if len(st.Address) == 0 {
			if !strings.Contains(newText, "ph") && !strings.Contains(newText, "phone") && !strings.Contains(newText, "tin") {
				address = fmt.Sprintf("%s %s",address, newText)
			} else {
				st.Address = address
			}
		} else if len(st.PhoneNo) == 0 {
			st.PhoneNo = setPhoneNo(newText)
		} else if len(st.Tin) == 0 {
			st.Tin = setTinNo(newText)
		}
	}

	if isComplete(*st) {
		esResponse, err := elasticsearch.Search(esClient, "store-index", "title", st.Title)
		if err != nil {
			log.Fatalf("es search query failed. title %s, err %v ", st.Title, err)
		}
		if len(esResponse.Hits.Hits) == 0 {
			fmt.Printf("store %v ", st)
			//ingest into elasticsearch store index
			elasticsearch.IndexDocument(esClient, "store-index", "store-type", uuid.NewV4().String(), st)
		}
	}

	return st
}

func NewText(esSuggestResponse model.ESSuggestResponse, originalText string) string {
	mySugestion := esSuggestResponse.Suggestion.MySuggestion

	for _, suggestion := range mySugestion {
		word := suggestion.Text

		for _, opt := range suggestion.Options {
			strings.Replace(originalText, word, opt.Text, 1)
		}
	}

	return originalText
}

func isComplete(store model.Store) bool {
	fmt.Println(store.Title)
	if len(store.Title) == 0 { /* store.Address == "" || store.PhoneNo == "" || store.Tin == "" */
		return false
	}
	return true
}

func setPhoneNo(text string) string {
	if strings.Contains(text, "ph no.") {
		return strings.Replace(text, "ph no.", "", 1)
	} else if strings.Contains(text, "ph") {
		return strings.Replace(text, "ph", "", 1)
	} else if strings.Contains(text, "phone") {
		return strings.Replace(text, "phone", "", 1)
	} else {
		return ""
	}
}

func setTinNo(text string) string {
	if strings.Contains(text, "tin") {
		return strings.Replace(text, "tin", "", 1)
	} else if strings.Contains(text, "tin no") {
		return strings.Replace(text, "tine no", "", 1)
	} else {
		return ""
	}
}

func SetTitle(text string) string {
	return text;
}
