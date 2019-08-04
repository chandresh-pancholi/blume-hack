package model

//https://mholt.github.io/json-to-go/ boon library

type DetectText struct {
	TextDetections []TextDetections `json:"TextDetections"`
}

type TextDetections struct {
	Confidence   float64  `json:"Confidence"`
	DetectedText string   `json:"DetectedText"`
	Geometry     Geometry `json:"Geometry"`
	ID           int64    `json:"id"`
	Type         string   `json:"type"`
}

type Geometry struct {
	BoundingBox BoundingBox `json:"BoundingBox"`
	Polygon     []struct {
		X float64 `json:"X"`
		Y float64 `json:"Y"`
	} `json:"Polygon"`
}

type BoundingBox struct {
	Height float64 `json:"Height"`
	Left   float64 `json:"Left"`
	Top    float64 `json:"Top"`
	Width  float64 `json:"Width"`
}
