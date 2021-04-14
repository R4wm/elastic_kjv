package query

import (
	"fmt"
	"strings"
)

// returns query_string: "text: just AND text: faith AND text: alone"
func GetQueryString(keywords []string) string {
	var b strings.Builder

	if len(keywords) == 0 {
		return ""
	}

	lastKeywordNum := (len(keywords) - 1)
	for i, v := range keywords {
		b.WriteString(fmt.Sprintf("text: %s", v))
		if i != lastKeywordNum {
			b.WriteString(" AND ")
		}
	}

	return b.String()
}
