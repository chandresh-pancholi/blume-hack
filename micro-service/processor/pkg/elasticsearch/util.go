package elasticsearch

import (
	"fmt"
	"io"
	"strings"
)

const sugesterQuery = `
	"suggest" : {
		"my-suggestion" : {
			"text" : %q,
		"term" : {
        	"field" : "%q"
      	}
		}
	}`
const searchMatch = `
	"query" : {
		"match" : {
			%q : %q
		}
	}`


func buildSuggestQuery(value ...string) io.Reader {
	var b strings.Builder

	b.WriteString("{\n")

	if len(value) > 0 && value[0] != "" && value[0] != "null" {
		b.WriteString(fmt.Sprintf(sugesterQuery, value[0], value[1]))
	}

	b.WriteString("\n}")

	// fmt.Printf("%s\n", b.String())
	return strings.NewReader(b.String())
}

func buildQuery(query string, value ...string) io.Reader {
	var b strings.Builder

	b.WriteString("{\n")

	if len(value) > 0 && value[0] != "" && value[0] != "null" {
		b.WriteString(fmt.Sprintf(searchMatch, query, value[0]))
	}

	b.WriteString("\n}")

	// fmt.Printf("%s\n", b.String())
	return strings.NewReader(b.String())
}