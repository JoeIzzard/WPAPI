package wpapi

import (
	"strings"
)

func apiURLBuilder(slug string, rtype string) string {
	var url string

	// Decide which core url to use
	if rtype == "plugin" {
		url = "https://api.wordpress.org/plugins/info/1.0/"
	}

	// Compile and return the string
	var b strings.Builder
	b.WriteString(url)
	b.WriteString(slug)
	b.WriteString(".json")
	return b.String()
}
