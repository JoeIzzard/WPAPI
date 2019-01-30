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
	// Live URL
	if res := apiURLBuilderBaseURL(false); res != "https://api.wordpress.org/" {
		t.Error("API URL Builder Base URL (Func) Failed: Returned '" + res + "' with option false. Expected 'https://api.wordpress.org/'")
	}

	// Mock Server URL
	if res := apiURLBuilderBaseURL(true); res != "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/" {
		t.Error("API URL Builder Base URL (Func) Failed: Returned '" + res + "' with option true. Expected 'https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/'")
	}
}

func TestApiURLBuilderSecret(t *testing.T) {
	// Live URL
	if res := apiURLBuilderSecret(false); res != "https://api.wordpress.org/secret-key/1.1/salt" {
		t.Error("API URL Builder Secret (Func) Failed: Returned '" + res + "' with the testing option false. Expected 'https://api.wordpress.org/secret-key/1.1/salt'")
	}

	// Mock Server URL
	if res := apiURLBuilderSecret(true); res != "https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/secret-key/1.1/salt" {
		t.Error("API URL Builder Secret (Func) Failed: Returned '" + res + "' with the testing option true. Expected 'https://93167486-1f77-4f71-a2e2-9d3098460682.mock.pstmn.io/secret-key/1.1/salt'")
	}
}
