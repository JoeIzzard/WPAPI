package wpapi

import (
	"os"
	"strconv"
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

func TestPluginRatingOneStar(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginRatingOneStar("jetpack")
	res := 256
	name := "Plugin Rating One Star"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginRatingTwoStar(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginRatingTwoStar("jetpack")
	res := 62
	name := "Plugin Rating Two Star"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginRatingThreeStar(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginRatingThreeStar("jetpack")
	res := 75
	name := "Plugin Rating Three Star"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginRatingFourStar(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginRatingFourStar("jetpack")
	res := 122
	name := "Plugin Rating Four Star"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginRatingFiveStar(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginRatingFiveStar("jetpack")
	res := 844
	name := "Plugin Rating Five Star"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginRatingNumber(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginRatingNumber("jetpack")
	res := 1359
	name := "Plugin Rating Number"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginSupport(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	plug, err := PluginSupport("jetpack")
	name := "Plugin Support"
	threads := 198
	solved := 177
	open := 21
	percent := 89

	// Total Threads
	if plug.Threads.Total != threads {
		t.Errorf(name+" (Func) Failed [Total]: Returned '%d', expected '%d'", plug.Threads.Total, threads)
	}

	// Solved Threads
	if plug.Threads.Solved != solved {
		t.Errorf(name+" (Func) Failed [Solved]: Returned '%d', expected '%d'", plug.Threads.Solved, solved)
	}

	// Open Threads
	if plug.Threads.Open != open {
		t.Errorf(name+" (Func) Failed [Open]: Returned '%d', expected '%d'", plug.Threads.Open, open)
	}

	// Percentage
	if plug.SolvedPercentage != percent {
		t.Errorf(name+" (Func) Failed [Solved Percentage]: Returned '%d', expected '%d'", plug.SolvedPercentage, percent)
	}

	// Error
	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginSupportThreads(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginSupportThreads("jetpack")
	res := 198
	name := "Plugin Support Threads"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginSupportThreadsOpen(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginSupportThreadsOpen("jetpack")
	res := 21
	name := "Plugin Support Threads Open"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginSupportThreadsClosed(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginSupportThreadsSolved("jetpack")
	res := 177
	name := "Plugin Support Threads Solved"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginSupportThreadsSolvedPercentage(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginSupportThreadsSolvedPercentage("jetpack")
	res := 89
	name := "Plugin Support Threads Solved Percentage"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginDownloads(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginDownloads("jetpack")
	res := 107462381
	name := "Plugin Downloads"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginActiveInstalls(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginActiveInstalls("jetpack")
	res := 5000000
	name := "Plugin Active Installs"

	if test != res {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", test, res)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginLastUpdated(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginLastUpdate("jetpack")
	res := "2019-01-10 2:48pm GMT"
	name := "Plugin Last Update"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginAdded(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginAdded("jetpack")
	res := "2011-01-20"
	name := "Plugin Added"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginHomepage(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginHomepage("jetpack")
	res := "https://jetpack.com"
	name := "Plugin Homepage"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginDownloadLink(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginDownloadLink("jetpack")
	res := "https://downloads.wordpress.org/plugin/jetpack.6.9.zip"
	name := "Plugin Download Link"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginScreenshots(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	// Test with Screenshots
	test, err := PluginScreenshots("jetpack")
	len1 := len(test)
	lenRes := 8
	name := "Plugin Screenshots"

	if len1 != lenRes {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", len1, lenRes)
	}

	if err != nil {
		t.Error(err)
	}

	// Test with no Screenshots
	test2, err2 := PluginScreenshots("hello-dolly")
	len2 := len(test2)
	lenRes2 := 0

	if len2 != lenRes2 {
		t.Errorf(name+" (Func) Failed [No Screenshots Test]: Returned '%d', expected '%d'", len2, lenRes2)
	}

	if err2 != nil {
		t.Error(err2)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginTags(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	// Test with Tags
	test, err := PluginTags("jetpack")
	len1 := len(test)
	lenRes := 5
	name := "Plugin Tags"

	if len1 != lenRes {
		t.Errorf(name+" (Func) Failed: Returned '%d', expected '%d'", len1, lenRes)
	}

	if err != nil {
		t.Error(err)
	}

	// Test with no tags
	test2, err2 := PluginTags("hello-dolly")
	len2 := len(test2)
	lenRes2 := 0

	if len2 != lenRes2 {
		t.Errorf(name+" (Func) Failed [No Screenshots Test]: Returned '%d', expected '%d'", len2, lenRes2)
	}

	if err2 != nil {
		t.Error(err2)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginVersion(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	name := "Plugin Version"

	// Actual Release
	test1, err1 := PluginVersion("jetpack", "6.9")
	res1Version := "6.9"
	res1PreRelease := false
	res1Link := "https://downloads.wordpress.org/plugin/jetpack.6.9.zip"

	if test1.Version != res1Version {
		t.Error(name + " (Func) Failed [Version]: Returned '" + test1.Version + "', expected '" + res1Version + "'")
	}

	if test1.PreRelease != res1PreRelease {
		t.Error(name + " (Func) Failed [PreRelease]: Returned '" + strconv.FormatBool(test1.PreRelease) + "', expected '" + strconv.FormatBool(res1PreRelease) + "'")
	}

	if test1.Link != res1Link {
		t.Error(name + " (Func) Failed [Link]: Returned '" + test1.Link + "', expected '" + res1Link + "'")
	}

	if err1 != nil {
		t.Error(err1)
	}

	// Beta Release
	test2, err2 := PluginVersion("jetpack", "7.0-beta")
	res2Version := "7.0-beta"
	res2PreRelease := true
	res2Link := "https://downloads.wordpress.org/plugin/jetpack.7.0-beta.zip"

	if test2.Version != res2Version {
		t.Error(name + " (Func) Failed [Version]: Returned '" + test2.Version + "', expected '" + res2Version + "'")
	}

	if test2.PreRelease != res2PreRelease {
		t.Error(name + " (Func) Failed [PreRelease]: Returned '" + strconv.FormatBool(test2.PreRelease) + "', expected '" + strconv.FormatBool(res2PreRelease) + "'")
	}

	if test2.Link != res2Link {
		t.Error(name + " (Func) Failed [Link]: Returned '" + test2.Link + "', expected '" + res2Link + "'")
	}

	if err2 != nil {
		t.Error(err2)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginVersions(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	name := "Plugin Versions"
	vers, err := PluginVersions("jetpack")

	// Actual Release
	test1 := vers["6.9"]
	res1Version := "6.9"
	res1PreRelease := false
	res1Link := "https://downloads.wordpress.org/plugin/jetpack.6.9.zip"

	if test1.Version != res1Version {
		t.Error(name + " (Func) Failed [Version]: Returned '" + test1.Version + "', expected '" + res1Version + "'")
	}

	if test1.PreRelease != res1PreRelease {
		t.Error(name + " (Func) Failed [PreRelease]: Returned '" + strconv.FormatBool(test1.PreRelease) + "', expected '" + strconv.FormatBool(res1PreRelease) + "'")
	}

	if test1.Link != res1Link {
		t.Error(name + " (Func) Failed [Link]: Returned '" + test1.Link + "', expected '" + res1Link + "'")
	}

	// Beta Release
	test2 := vers["7.0-beta"]
	res2Version := "7.0-beta"
	res2PreRelease := true
	res2Link := "https://downloads.wordpress.org/plugin/jetpack.7.0-beta.zip"

	if test2.Version != res2Version {
		t.Error(name + " (Func) Failed [Version]: Returned '" + test2.Version + "', expected '" + res2Version + "'")
	}

	if test2.PreRelease != res2PreRelease {
		t.Error(name + " (Func) Failed [PreRelease]: Returned '" + strconv.FormatBool(test2.PreRelease) + "', expected '" + strconv.FormatBool(res2PreRelease) + "'")
	}

	if test2.Link != res2Link {
		t.Error(name + " (Func) Failed [Link]: Returned '" + test2.Link + "', expected '" + res2Link + "'")
	}

	// Count Check
	test3 := len(vers)
	res3 := 76

	if test3 != res3 {
		t.Errorf(name+" (Func) Failed [Count]: Returned '%d', expected '%d'", test3, res3)
	}

	// Error Check
	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}
