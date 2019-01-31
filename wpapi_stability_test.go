package wpapi

import (
	"os"
	"testing"
)

func TestStabilityList(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	stabilityList, err := StabilityList()
	if err != nil {
		t.Error("Stability List (Func) Failed: " + err.Error())
	}

	res := len(stabilityList)

	if res != 338 {
		t.Errorf("Stability List (Func) Failed: Returned '%d', expected '%d'", res, 338)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestStabilityVersion(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	// Test Version
	res, err := StabilityVersion("1.2.1")
	if res != "insecure" {
		t.Error("Stability Version (Func) Failed: Returned '" + res + "', expected 'insecure'")
	}

	if err != nil {
		t.Error("Stability Version (Func) Failed: " + err.Error())
	}

	_, err2 := StabilityVersion("1.0.2")
	if err2 == nil {
		t.Error("Stability Version (Func) Failed: Expected Error to be returned for non existant version")
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")

	// We should now be calling against the live API
	res3, err3 := StabilityVersion("1.0.2")
	if res3 != "insecure" {
		t.Error("Stability Version (Func) Failed: Returned '" + res + "', expected 'insecure' when testing live API")
	}

	if err3 != nil {
		t.Error("Stability Version (Func) Failed: " + err.Error())
	}
}

func TestStabilityLatest(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	// Latest Version
	res1, err1 := StabilityLatest("5.0.3")
	if res1 != true {
		t.Error("Stability Latest (Func) Failed: Returned 'false', expected 'true' for version '5.0.3'")
	}

	if err1 != nil {
		t.Error("Stability Latest (Func) Failed (Returned an Error): " + err1.Error())
	}

	// Not Latest Version
	res2, err2 := StabilityLatest("5.0.2")
	if res2 != false {
		t.Error("Stability Latest (Func) Failed: Returned 'true', expected 'false' for version '5.0.2'")
	}

	if err2 != nil {
		t.Error("Stability Latest (Func) Failed (Returned an Error): " + err2.Error())
	}

	// Missing version number
	_, err3 := StabilityLatest("1.0.2")
	if err3 == nil {
		t.Error("Stability Latest (Func) Failed: Expected an error to be returned for missing version number")
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestStabilityNotLatest(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	// Latest version
	res1, err1 := StabilityNotLatest("5.0.3")

	if res1 != false {
		t.Error("Stability Not Latest (Func) Failed: Returned 'true', expected 'false' for version '5.0.3'")
	}

	if err1 != nil {
		t.Error("Stability Not Latest (Func) Failed (Returned an Error): " + err1.Error())
	}

	// Not Latest Version
	res2, err2 := StabilityNotLatest("5.0.2")

	if res2 != true {
		t.Error("Stability Not Latest (Func) Failed: Returned 'false', expected 'true' for version '5.0.2'")
	}

	if err2 != nil {
		t.Error("Stability Not Latest (Func) Failed (Returned an Error): " + err2.Error())
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestStabilityOutdated(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	// Latest Version
	res1, err1 := StabilityOutdated("5.0.3")
	if res1 != false {
		t.Error("Stability Outdated (Func) Failed: Returned 'true', expected 'false' for version '5.0.3'")
	}

	if err1 != nil {
		t.Error("Stability Outdated (Func) Failed (Returned an Error): " + err1.Error())
	}

	// Outdated Version
	res2, err2 := StabilityOutdated("5.0.2")
	if res2 != true {
		t.Error("Stability Outdated (Func) Failed: Returned 'false', expected 'true' for version '5.0.2'")
	}

	if err2 != nil {
		t.Error("Stability Outdated (Func) Failed (Returned an Error): " + err2.Error())
	}

	// Missing version number
	_, err3 := StabilityOutdated("1.0.2")
	if err3 == nil {
		t.Error("Stability Outdated (Func) Failed: Expected an error to be returned for missing version number")
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestStabilityInsecure(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	// Latest Version
	res1, err1 := StabilityInsecure("5.0.3")
	if res1 != false {
		t.Error("Stability Insecure (Func) Failed: Returned 'true', expected 'false' for version '5.0.3'")
	}

	if err1 != nil {
		t.Error("Stability Insecure (Func) Failed (Returned an Error): " + err1.Error())
	}

	// Insecure Version
	res2, err2 := StabilityInsecure("1.2.1")
	if res2 != true {
		t.Error("Stability Insecure (Func) Failed: Returned 'false', expected 'true' for version '1.2.1'")
	}

	if err2 != nil {
		t.Error("Stability Insecure (Func) Failed (Returned an Error): " + err2.Error())
	}

	// Missing version number
	_, err3 := StabilityInsecure("1.0.2")
	if err3 == nil {
		t.Error("Stability Insecure (Func) Failed: Expected an error to be returned for missing version number")
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}
