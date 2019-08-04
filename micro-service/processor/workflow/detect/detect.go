package detect

import (
	"log"
	"processor/model"
	es "processor/pkg/elasticsearch/config"
	"processor/workflow/item"
	"processor/workflow/store"
	"strings"
)

type DetectWorkflow struct {
	EsClient *es.ESClient
}

func (dw *DetectWorkflow) Trigger(detect model.DetectText) {
	textDetection := detect.TextDetections

	index := 0

	for idx, td := range detect.TextDetections {

		txt := strings.ToLower(td.DetectedText)
		if strings.Contains(txt, "item") || strings.Contains(txt, "qty") {
			log.Printf("triggering store detail workflow")
			store.Trigger(dw.EsClient, textDetection[0:idx-1])

			index = idx
			break;
		}
	}

	log.Printf("triggering item detail workflow")

	td := textDetection[index:]
	item.Trigger(dw.EsClient, td)

}
