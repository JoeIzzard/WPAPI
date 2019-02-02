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
