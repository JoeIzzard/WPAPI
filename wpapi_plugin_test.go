package wpapi

import (
	"os"
	"strconv"
	"testing"
)

func TestGetPlugin(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	Plugin, err := GetPlugin("jetpack")
	name := "Get Plugin"

	res := &plugin{
		Name:    "Jetpack by WordPress.com",
		Slug:    "jetpack",
		Version: "6.9",
		Author: author{
			Name:    "Automattic",
			Profile: "https://profiles.wordpress.org/automattic",
			Slug:    "automattic",
			Link:    "https://jetpack.com",
		},
		Tested:     "5.0.3",
		RequiredWP: "4.8",
		RequiresPHP: phpVersion{
			Provided: false,
			Version:  "",
		},
		Rating: rating{
			Overall: 78,
			Stars: stars{
				One:   256,
				Two:   62,
				Three: 75,
				Four:  122,
				Five:  844,
			},
			Number: 1359,
		},
		Support: support{
			Threads: threads{
				Total:  198,
				Solved: 177,
				Open:   21,
			},
			SolvedPercentage: 89,
		},
		Downloads:      107462381,
		ActiveInstalls: 5000000,
		LastUpdate:     "2019-01-10 2:48pm GMT",
		Added:          "2011-01-20",
		Homepage:       "https://jetpack.com",
		DownloadLink:   "https://downloads.wordpress.org/plugin/jetpack.6.9.zip",
	}
	lenScreenshots := 8
	lenTags := 5
	lenVersions := 76

	// Test the Name
	if Plugin.Name != res.Name {
		t.Error(name + " (Func) Failed [Name]: Returned '" + Plugin.Name + "', expected '" + res.Name + "'")
	}

	// Slug
	if Plugin.Slug != res.Slug {
		t.Error(name + " (Func) Failed [Slug]: Returned '" + Plugin.Slug + "', expected '" + res.Slug + "'")
	}

	// ---- Author ----
	// Name
	if Plugin.Author.Name != res.Author.Name {
		t.Error(name + " (Func) Failed [Author Name]: Returned '" + Plugin.Author.Name + "', expected '" + res.Author.Name + "'")
	}

	// Profile
	if Plugin.Author.Profile != res.Author.Profile {
		t.Error(name + " (Func) Failed [Author Profile]: Returned '" + Plugin.Author.Profile + "', expected '" + res.Author.Profile + "'")
	}

	// Link
	if Plugin.Author.Link != res.Author.Link {
		t.Error(name + " (Func) Failed [Author Link]: Returned '" + Plugin.Author.Link + "', expected '" + res.Author.Link + "'")
	}

	// Slug
	if Plugin.Author.Slug != res.Author.Slug {
		t.Error(name + " (Func) Failed [Author Slug]: Returned '" + Plugin.Author.Slug + "', expected '" + res.Author.Slug + "'")
	}

	// Tested
	if Plugin.Tested != res.Tested {
		t.Error(name + " (Func) Failed [Tested]: Returned '" + Plugin.Tested + "', expected '" + res.Tested + "'")
	}

	// Required WP
	if Plugin.RequiredWP != res.RequiredWP {
		t.Error(name + " (Func) Failed [Required WP]: Returned '" + Plugin.RequiredWP + "', expected '" + res.RequiredWP + "'")
	}

	// Required PHP
	if Plugin.RequiresPHP.Provided != false {
		t.Error(name + " (Func) Failed [Required PHP Provided]: Returned '" + strconv.FormatBool(Plugin.RequiresPHP.Provided) + "', expected '" + strconv.FormatBool(res.RequiresPHP.Provided) + "'")
	}

	if Plugin.RequiresPHP.Version != res.RequiresPHP.Version {
		t.Error(name + " (Func) Failed [Required PHP Version]: Returned '" + Plugin.RequiresPHP.Version + "', expected '" + res.RequiresPHP.Version + "'")
	}

	// ----- Rating -----
	// Overall
	if Plugin.Rating.Overall != res.Rating.Overall {
		t.Errorf(name+" (Func) Failed [Rating - Overall]: Returned '%d', expected '%d'", Plugin.Rating.Overall, res.Rating.Overall)
	}

	// Stars 1
	if Plugin.Rating.Stars.One != res.Rating.Stars.One {
		t.Errorf(name+" (Func) Failed [Rating - One Stars]: Returned '%d', expected '%d'", Plugin.Rating.Stars.One, res.Rating.Stars.One)
	}

	// Stars 12
	if Plugin.Rating.Stars.Two != res.Rating.Stars.Two {
		t.Errorf(name+" (Func) Failed [Rating - Two Stars]: Returned '%d', expected '%d'", Plugin.Rating.Stars.Two, res.Rating.Stars.Two)
	}

	// Stars 3
	if Plugin.Rating.Stars.Three != res.Rating.Stars.Three {
		t.Errorf(name+" (Func) Failed [Rating - Three Stars]: Returned '%d', expected '%d'", Plugin.Rating.Stars.Three, res.Rating.Stars.Three)
	}

	// Stars 4
	if Plugin.Rating.Stars.Four != res.Rating.Stars.Four {
		t.Errorf(name+" (Func) Failed [Rating - Four Stars]: Returned '%d', expected '%d'", Plugin.Rating.Stars.Four, res.Rating.Stars.Four)
	}

	// Stars 5
	if Plugin.Rating.Stars.Five != res.Rating.Stars.Five {
		t.Errorf(name+" (Func) Failed [Rating - Five Stars]: Returned '%d', expected '%d'", Plugin.Rating.Stars.Five, res.Rating.Stars.Five)
	}

	// Number
	if Plugin.Rating.Number != res.Rating.Number {
		t.Errorf(name+" (Func) Failed [Rating - Number]: Returned '%d', expected '%d'", Plugin.Rating.Number, res.Rating.Number)
	}

	// ----- Support -----
	// Support Threads
	if Plugin.Support.Threads.Total != res.Support.Threads.Total {
		t.Errorf(name+" (Func) Failed [Support - Total Threads]: Returned '%d', expected '%d'", Plugin.Support.Threads.Total, res.Support.Threads.Total)
	}

	// Support Threads Solved
	if Plugin.Support.Threads.Solved != res.Support.Threads.Solved {
		t.Errorf(name+" (Func) Failed [Support - Solved Threads]: Returned '%d', expected '%d'", Plugin.Support.Threads.Solved, res.Support.Threads.Solved)
	}

	// Support Threads Open
	if Plugin.Support.Threads.Open != res.Support.Threads.Open {
		t.Errorf(name+" (Func) Failed [Support - Open Threads]: Returned '%d', expected '%d'", Plugin.Support.Threads.Open, res.Support.Threads.Open)
	}

	// Solved Percentage
	if Plugin.Support.SolvedPercentage != res.Support.SolvedPercentage {
		t.Errorf(name+" (Func) Failed [Support - Solved Percentage]: Returned '%d', expected '%d'", Plugin.Support.SolvedPercentage, res.Support.SolvedPercentage)
	}

	// Downloads
	if Plugin.Downloads != res.Downloads {
		t.Errorf(name+" (Func) Failed [Downloads]: Returned '%d', expected '%d'", Plugin.Downloads, res.Downloads)
	}

	// Last Update
	if Plugin.LastUpdate != res.LastUpdate {
		t.Error(name + " (Func) Failed [Last Update]: Returned '" + Plugin.LastUpdate + "', expected '" + res.LastUpdate + "'")
	}

	// Added
	if Plugin.Added != res.Added {
		t.Error(name + " (Func) Failed [Added]: Returned '" + Plugin.Added + "', expected '" + res.Added + "'")
	}

	// Homepage
	if Plugin.Homepage != res.Homepage {
		t.Error(name + " (Func) Failed [Downloads]: Returned '" + Plugin.Homepage + "', expected '" + res.Homepage + "'")
	}

	// Download Link
	if Plugin.DownloadLink != res.DownloadLink {
		t.Error(name + " (Func) Failed [Download Link]: Returned '" + Plugin.DownloadLink + "', expected '" + res.DownloadLink + "'")
	}

	// Screenshots
	if len(Plugin.Screenshots) != lenScreenshots {
		t.Errorf(name+" (Func) Failed [Screenshots - Length]: Returned '%d', expected '%d'", len(Plugin.Screenshots), lenScreenshots)
	}

	// Tags
	if len(Plugin.Tags) != lenTags {
		t.Errorf(name+" (Func) Failed [Tags - Length]: Returned '%d', expected '%d'", len(Plugin.Tags), lenTags)
	}

	// Versions
	if len(Plugin.Versions) != lenVersions {
		t.Errorf(name+" (Func) Failed [Versions - Length]: Returned '%d', expected '%d'", len(Plugin.Versions), lenVersions)
	}

	// Error Check
	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

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

	// Non existent Release
	_, err := PluginVersion("jetpack", "FakeyMcFakeFake")

	if err == nil {
		t.Error(name + " (Func) Failed: Didn't return an error for non existent version")
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

func TestPluginContributor(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginContributor("jetpack", "automattic")
	name := "Plugin Contributor"
	resSlug := "automattic"
	resLink := "https://profiles.wordpress.org/automattic"

	if test.Slug != resSlug {
		t.Error(name + " (Func) Failed [Slug]: Returned '" + test.Slug + "', expected '" + resSlug + "'")
	}

	if test.URL != resLink {
		t.Error(name + " (Func) Failed [URL]: Returned '" + test.URL + "', expected '" + resLink + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Non Existent Contributor
	_, err2 := PluginContributor("jetpack", "FakeyMcFakeFake")

	if err2 == nil {
		t.Error(name + " (Func) Failed: Didn't return an error for fake contributor")
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginContributors(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginContributors("jetpack")
	name := "Plugin Contributors"
	resSlug := "automattic"
	resLink := "https://profiles.wordpress.org/automattic"
	testCount := len(test)
	resCount := 95

	if contrib, ok := test[resSlug]; ok {
		if contrib.Slug != resSlug {
			t.Error(name + " (Func) Failed [Slug]: Returned '" + contrib.Slug + "', expected '" + resSlug + "'")
		}

		if contrib.URL != resLink {
			t.Error(name + " (Func) Failed [URL]: Returned '" + contrib.URL + "', expected '" + resLink + "'")
		}
	} else {
		t.Error(name + " (Func) Failed: Missing a test slug")
	}

	if testCount != resCount {
		t.Errorf(name+" (Func) Failed [Count]: Returned '%d', expected '%d'", testCount, resCount)
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}

func TestPluginDonateLink(t *testing.T) {
	// Enable Debug Mode
	os.Setenv("WPAPIDEBUG", "1")

	test, err := PluginDonateLink("wordpress-seo")
	name := "Plugin Donate Link"
	res := "https://yoa.st/1up"

	if test != res {
		t.Error(name + " (Func) Failed: Returned '" + test + "', expected '" + res + "'")
	}

	if err != nil {
		t.Error(err)
	}

	// Disable Debug Mode
	os.Setenv("WPAPIDEBUG", "0")
}