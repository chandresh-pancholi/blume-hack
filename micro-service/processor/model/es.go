package model

//Hit is
type Hit struct {
	ID     string                 `json:"_id"`
	Source map[string]interface{} `json:"_source"`
}

//HitsObject is a part of ES query response
type HitsObject struct {
	Total Total `json:"total"`
	Hits  []Hit `json:"hits"`
}

//ESResponse is ES search query response
type ESResponse struct {
	Took float32    `json:"took"`
	Hits HitsObject `json:"hits"`
}

//Total is
type Total struct {
	value int `json:"value"`
}

type ESSuggestResponse struct {
	Took       float32          `json:"took"`
	Hits       HitsObject       `json:"hits"`
	Suggestion SuggestionObject `json:"suggestion"`
}

type SuggestionObject struct {
	MySuggestion []MySuggestionObject `json:"my_suggestion"`
}

type MySuggestionObject struct {
	Text    string          `json:"text"`
	Offest  int64           `json:"offest"`
	Length  int64           `json:"length"`
	Options []OptionsObject `json:"options"`
}

type OptionsObject struct {
	Text  string  `json:"text"`
	Score float64 `json:"score"`
	Freq  int64   `json:"freq"`
}
