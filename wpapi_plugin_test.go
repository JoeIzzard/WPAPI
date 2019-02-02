package wpapi

import (
	"os"
	"testing"
)

func TestPluginName(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginName("jetpack")
	res := "Jetpack by WordPress.com"
	name := "Plugin Name"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginCurrentVersion(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginCurrentVersion("jetpack")
	res := "6.9"
	name := "Plugin Current Version"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginAuthorName(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginAuthorName("jetpack")
	res := "Automattic"
	name := "Plugin Author Name"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginAuthorProfile(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginAuthorProfile("jetpack")
	res := "https://profiles.wordpress.org/automattic"
	name := "Plugin Author Profile"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginAuthorLink(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginAuthorLink("jetpack")
	res := "https://jetpack.com"
	name := "Plugin Author Link"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginAuthorSlug(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginAuthorSlug("jetpack")
	res := "automattic"
	name := "Plugin Author Slug"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginTested(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginTested("jetpack")
	res := "5.0.3"
	name := "Plugin Tested"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginRequiredWPVersion(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginRequiredWPVersion("jetpack")
	res := "4.8"
	name := "Plugin Required WP Version"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginRequiredPHP(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	// False Test
	test, err := PluginRequiredPHP("jetpack")
	name := "Plugin Required PHP"

	if test.Provided != false {
		t.Error(name + " (Func) Failed: Didn't return 'false' with Jetpack")
	}

	if err != nil {
		t.Error(err)
	}

	// Specified Test
	test2, err2 := PluginRequiredPHP("wordpress-seo")

	if test2.Provided != true {
		t.Error(name + " (Func) Failed: Didn't return 'true' with WordPress SEO")
	}

	if test2.Version != "5.2.4" {
		t.Error(name + " (Func) Failed: Returned '" + test2.Version + "', expected '5.2.4'")
	}

	if err2 != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginRating(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginRating("jetpack")
	name := "Plugin Rating"
	overall := 78
	oneStar := 256
	twoStar := 62
	threeStar := 75
	fourStar := 122
	fiveStar := 844
	totalRatings := 1359

	// Overall
	if test.Overall != overall {
		t.Errorf(name+" (Func) Failed: Overall rating returned incorrect, Returned '%d', expected '%d'", test.Overall, overall)
	}

	// One Star
	if test.Stars.One != oneStar {
		t.Errorf(name+" (Func) Failed: One Star rating returned incorrect, Returned '%d', expected '%d'", test.Stars.One, oneStar)
	}

	// Two Star
	if test.Stars.Two != twoStar {
		t.Errorf(name+" (Func) Failed: Two Star rating returned incorrect, Returned '%d', expected '%d'", test.Stars.Two, twoStar)
	}

	// Three Star
	if test.Stars.Three != threeStar {
		t.Errorf(name+" (Func) Failed: Three Star rating returned incorrect, Returned '%d', expected '%d'", test.Stars.Three, threeStar)
	}

	// Four Star
	if test.Stars.Four != fourStar {
		t.Errorf(name+" (Func) Failed: Four Star rating returned incorrect, Returned '%d', expected '%d'", test.Stars.Four, fourStar)
	}

	// Five Star
	if test.Stars.Five != fiveStar {
		t.Errorf(name+" (Func) Failed: Five Star rating returned incorrect, Returned '%d', expected '%d'", test.Stars.Five, fiveStar)
	}

	// Number
	if test.Number != totalRatings {
		t.Errorf(name+" (Func) Failed: Total Number of Ratings returned incorrect, Returned '%d', expected '%d'", test.Number, totalRatings)
	}

	// Error Check
	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginRatingOverall(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginRatingOverall("jetpack")
	res := 78
	name := "Plugin Rating Overall"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}
