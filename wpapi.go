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
	return apiURLBuilderBaseURL(testing) + "themes/info/1.1/?action=theme_information&request[slug]=" + slug
}

func apiURLBuilderPluginTranslation(slug string, version string, testing bool) string {
	return apiURLBuilderBaseURL(testing) + "translations/plugins/1.0/?slug=" + slug + "&version=" + version
}

func apiURLBuilderThemeTranslation(slug string, version string, testing bool) string {
	return apiURLBuilderBaseURL(testing) + "translations/themes/1.0/?slug=" + slug + "&version=" + version
}
