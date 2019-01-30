package wpapi

func apiURLBuilderBaseURL(testing bool) string {
	var url string
	// Switch the Testing URL
	switch testing {
	case true:
		url = "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/"
	case false:
		url = "https://api.wordpress.org/"
	}

	return url
}

func apiURLBuilderSecret(testing bool) string {
	return apiURLBuilderBaseURL(testing) + "secret-key/1.1/salt"
}

func apiURLBuilderStability(testing bool) string {
	return apiURLBuilderBaseURL(testing) + "core/stable-check/1.0"
}
