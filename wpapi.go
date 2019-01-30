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

func apiURLBuilderPlugin(slug string, testing bool) string {
	return apiURLBuilderBaseURL(testing) + "plugin/info/1.0/" + slug + ".json"
}

func apiURLBuilderTheme(slug string, testing bool) string {
	// --- Note ---
	/*
		The Wordpress.org API emmits certains fields by default. The allFields variables tells the API to return all of the available fields in it's response
	*/
	allFields := "&request[fields][description]=1&request[fields][sections]=1&request[fields][rating]=1&request[fields][ratings]=1&request[fields][downloaded]=1&request[fields][download_link]=1&request[fields][last_updated]=1&request[fields][homepage]=1&request[fields][tags]=1&request[fields][template]=1&request[fields][parent]=1&request[fields][versions]=1&request[fields][screenshot_url]=1&request[fields][active_installs]=1"
	return apiURLBuilderBaseURL(testing) + "themes/info/1.1/?action=theme_information&request[slug]=" + slug + allFields
}

func apiURLBuilderPluginTranslation(slug string, version string, testing bool) string {
	return apiURLBuilderBaseURL(testing) + "translations/plugins/1.0/?slug=" + slug + "&version=" + version
}

func apiURLBuilderThemeTranslation(slug string, version string, testing bool) string {
	return apiURLBuilderBaseURL(testing) + "translations/themes/1.0/?slug=" + slug + "&version=" + version
}
