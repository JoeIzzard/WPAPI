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
