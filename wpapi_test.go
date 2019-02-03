package wpapi

import "testing"

// ----- IMPORTANT NOTE -----
/*
To facilitate testing of the endpoint data, we have a second base url using a postman mock server.
If this needs to be changed, it needs to be changed in all of the following locations:
  - apiURLBuilderBaseURL [wpapi.go]
  - TestApiURLBuilderBaseURL [wpapi_test.go]
  - TestApiURLBuilderSecret [wpapi_test.go]
  - TestApiURLBuilderStability [wpapi_test.go]
  - TestApiURLBuilderPluginTranslation [wpapi_test.go]
  - TestApiURLBuilderThemeTranslation [wpapi_test.go]
  - TestApiURLBuilderPlugin [wpapi_test.go]
  - TestApiURLBuilderTheme [wpapi_test.go]

The current mock server address is: https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/

When changing this address, please note that the trailing slash is required.

The current data for this mock server, along with information on creating all the endpoints can be found in the 'Test Data' Directory. This will be updated as the data changes.
*/

func TestApiURLBuilderBaseURL(t *testing.T) {
	live := "https://api.wordpress.org/"
	mock := "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/"
	funcName := "API URL Builder Base URL"

	// Live URL
	if res := apiURLBuilderBaseURL(false); res != live {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with option false. Expected '" + live + "'")
	}

	// Mock Server URL
	if res := apiURLBuilderBaseURL(true); res != mock {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with option true. Expected '" + mock + "'")
	}
}

func TestApiURLBuilderSecret(t *testing.T) {
	endpoint := "secret-key/1.1/salt"
	live := "https://api.wordpress.org/" + endpoint
	mock := "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/" + endpoint
	funcName := "API URL Builder Secret"

	// Live URL
	if res := apiURLBuilderSecret(false); res != live {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option false. Expected '" + live + "'")
	}

	// Mock Server URL
	if res := apiURLBuilderSecret(true); res != mock {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option true. Expected '" + mock + "'")
	}
}

func TestApiURLBuilderStability(t *testing.T) {
	endpoint := "core/stable-check/1.0"
	live := "https://api.wordpress.org/" + endpoint
	mock := "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/" + endpoint
	funcName := "API URL Builder Stability"

	// Live URL
	if res := apiURLBuilderStability(false); res != live {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option false. Expected '" + live + "'")
	}

	// Mock Server URL
	if res := apiURLBuilderStability(true); res != mock {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option true. Expected '" + mock + "'")
	}
}

func TestApiURLBuilderPlugin(t *testing.T) {
	slug := "jetpack"
	allFields := "&request[fields][active_installs]=1&request[fields][contributors]=1"
	endpoint := "plugins/info/1.1/?action=plugin_information&request[slug]=" + slug + allFields
	live := "https://api.wordpress.org/" + endpoint
	mock := "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/" + endpoint
	funcName := "API URL Builder Plugin"

	// Live URL
	if res := apiURLBuilderPlugin(slug, false); res != live {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option false. Expected '" + live + "'")
	}

	// Mock Server URL
	if res := apiURLBuilderPlugin(slug, true); res != mock {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option true. Expected '" + mock + "'")
	}
}

func TestApiURLBuilderTheme(t *testing.T) {
	slug := "twentynineteen"
	allFields := "&request[fields][description]=1&request[fields][sections]=1&request[fields][rating]=1&request[fields][ratings]=1&request[fields][downloaded]=1&request[fields][download_link]=1&request[fields][last_updated]=1&request[fields][homepage]=1&request[fields][tags]=1&request[fields][template]=1&request[fields][parent]=1&request[fields][versions]=1&request[fields][screenshot_url]=1&request[fields][active_installs]=1"
	endpoint := "themes/info/1.1/?action=theme_information&request[slug]=" + slug + allFields
	live := "https://api.wordpress.org/" + endpoint
	mock := "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/" + endpoint
	funcName := "API URL Builder Theme"

	// Live URL
	if res := apiURLBuilderTheme(slug, false); res != live {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option false. Expected '" + live + "'")
	}

	// Mock Server URL
	if res := apiURLBuilderTheme(slug, true); res != mock {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option true. Expected '" + mock + "'")
	}
}

func TestApiURLBuilderPluginTranslation(t *testing.T) {
	slug := "jetpack"
	version := "6.8"
	endpoint := "translations/plugins/1.0/?slug=" + slug + "&version=" + version
	live := "https://api.wordpress.org/" + endpoint
	mock := "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/" + endpoint
	funcName := "API URL Builder Plugin Translation"

	// Live URL
	if res := apiURLBuilderPluginTranslation(slug, version, false); res != live {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option false. Expected '" + live + "'")
	}

	// Mock Server URL
	if res := apiURLBuilderPluginTranslation(slug, version, true); res != mock {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option true. Expected '" + mock + "'")
	}
}

func TestApiURLBuilderThemeTranslation(t *testing.T) {
	slug := "twentynineteen"
	version := "1.1"
	endpoint := "translations/themes/1.0/?slug=" + slug + "&version=" + version
	live := "https://api.wordpress.org/" + endpoint
	mock := "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/" + endpoint
	funcName := "API URL Builder Theme Translation"

	// Live URL
	if res := apiURLBuilderThemeTranslation(slug, version, false); res != live {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option false. Expected '" + live + "'")
	}

	// Mock Server URL
	if res := apiURLBuilderThemeTranslation(slug, version, true); res != mock {
		t.Error(funcName + " (Func) Failed: Returned '" + res + "' with the testing option true. Expected '" + mock + "'")
	}
}
