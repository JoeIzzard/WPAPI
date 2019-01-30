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
	endpoint := "plugin/info/1.0/" + slug + ".json"
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
	endpoint := "themes/info/1.1/?action=theme_information&request[slug]=" + slug
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
