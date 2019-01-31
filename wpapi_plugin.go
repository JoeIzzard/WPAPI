package wpapi

import (
	"math"
	"os"
	"reflect"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mvdan/xurls"
)

// ---------- Plugin Structs ---------- //
type author struct {
	Name    string
	Profile string
	Link    string
	Slug    string
}

type phpVersion struct {
	Provided bool
	Version  string
}

type threads struct {
	Total  int
	Solved int
	Open   int
}

type support struct {
	Threads          threads
	SolvedPercentage int
}

type screenshot struct {
	Caption string
	SRC     string
}

type plugin struct {
	Name         string
	Slug         string
	Version      string
	Author       author
	Tested       string
	RequiredWP   string
	RequiresPHP  phpVersion
	Rating       rating
	Support      support
	Downloads    int
	LastUpdate   string
	Added        string
	Homepage     string
	DownloadLink string
	Screenshots  []screenshot
	Tags         []string
	Versions     []version
	DonateLink   string
}

func getPlugin(slug string) (raw rawData, err error) {
	// Determine if we should use the mock server
	debugEnv := os.Getenv("WPAPIDEBUG")

	var debug = false

	if debugEnv == "1" {
		debug = true
	}

	// Build the URL
	url := apiURLBuilderPlugin(slug, debug)

	raw, err = fetchData(url)

	return raw, err
}

func generatePlugin(slug string) (plug plugin, err error) {
	raw, rawErr := getPlugin(slug)

	if rawErr != nil {
		return plug, rawErr
	}

	// Name
	plug.Name = raw.Raw["name"].(string)
	plug.Slug = raw.Raw["slug"].(string)

	// Version
	plug.Version = raw.Raw["version"].(string)

	// Author
	plug.Author.Profile = raw.Raw["author_profile"].(string)
	plug.Author.Name = strip.StripTags(raw.Raw["author"].(string))
	plug.Author.Slug = strings.TrimLeft(raw.Raw["author_profile"].(string), "https://profiles.wordpress.org/")
	plug.Author.Link = xurls.Strict().FindString(raw.Raw["author"].(string))

	// Tested
	plug.Tested = raw.Raw["tested"].(string)

	// Required WordPress Version
	plug.RequiredWP = raw.Raw["requires"].(string)

	// // Required PHP Version
	switch raw.Raw["requires_php"].(type) {
	case bool:
		plug.RequiresPHP.Provided = false
		plug.RequiresPHP.Version = ""
	case string:
		plug.RequiresPHP.Provided = true
		plug.RequiresPHP.Version = raw.Raw["requires_php"].(string)
	default:
		plug.RequiresPHP.Provided = false
		plug.RequiresPHP.Version = ""
	}

	// ** Ratings! **
	plug.Rating.Overall = int(raw.Raw["rating"].(float64))

	ratings := raw.Raw["ratings"].(map[string]interface{})
	plug.Rating.Stars.One = int(ratings["1"].(float64))
	plug.Rating.Stars.Two = int(ratings["2"].(float64))
	plug.Rating.Stars.Three = int(ratings["3"].(float64))
	plug.Rating.Stars.Four = int(ratings["4"].(float64))
	plug.Rating.Stars.Five = int(ratings["5"].(float64))

	plug.Rating.Number = int(raw.Raw["num_ratings"].(float64))

	// Support!
	plug.Support.Threads.Total = int(raw.Raw["support_threads"].(float64))
	plug.Support.Threads.Solved = int(raw.Raw["support_threads_resolved"].(float64))
	plug.Support.Threads.Open = plug.Support.Threads.Total - plug.Support.Threads.Solved

	// Solved Percent
	plug.Support.SolvedPercentage = int(math.Round((float64(plug.Support.Threads.Solved) / float64(plug.Support.Threads.Total)) * 100))

	// Downloads
	plug.Downloads = int(raw.Raw["downloaded"].(float64))

	// Last Update
	plug.LastUpdate = raw.Raw["last_updated"].(string)

	// Added
	plug.Added = raw.Raw["added"].(string)

	// Homepage
	plug.Homepage = raw.Raw["homepage"].(string)

	// Download Link
	plug.DownloadLink = raw.Raw["download_link"].(string)

	// Reflect Set Up
	var t map[string]interface{}
	ans := reflect.TypeOf(t).Kind()

	// *** Screenshots ***
	if reflect.TypeOf(raw.Raw["screenshots"]).Kind() == ans {
		screenshotsMap := raw.Raw["screenshots"].(map[string]interface{})
		for _, value := range screenshotsMap {
			newS := screenshot{}
			current := value.(map[string]interface{})
			newS.Caption = current["caption"].(string)
			newS.SRC = current["src"].(string)
			plug.Screenshots = append(plug.Screenshots, newS)
		}
	}

	// *** Tags ***
	if reflect.TypeOf(raw.Raw["tags"]).Kind() == ans {
		tagsMap := raw.Raw["tags"].(map[string]interface{})
		for _, value := range tagsMap {
			plug.Tags = append(plug.Tags, value.(string))
		}
	}

	// *** Versions ***
	if reflect.TypeOf(raw.Raw["versions"]).Kind() == ans {
		versionMap := raw.Raw["versions"].(map[string]interface{})
		for key, value := range versionMap {
			newVersion := version{}

			if strings.Contains(key, "beta") || strings.Contains(key, "alpha") {
				newVersion.PreRelease = true
			} else {
				newVersion.PreRelease = false
			}

			newVersion.Version = key
			newVersion.Link = value.(string)

			plug.Versions = append(plug.Versions, newVersion)
		}
	}

	// Donate Link
	plug.DonateLink = raw.Raw["donate_link"].(string)

	return plug, err
}

// GetPlugin returns the plugin struct for a given plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func GetPlugin(slug string) (plug plugin, err error) {
	plug, err = generatePlugin(slug)
	return plug, err
}

// PluginName returns the name of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginName(slug string) (name string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Name, err
}

// PluginVersion returns the version string of the plugin as a string. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginVersion(slug string) (version string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Version, err
}

// PluginAuthorName returns the name of the plugins author stripped from the author field in the results. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAuthorName(slug string) (authorName string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Author.Name, err
}

// PluginAuthorProfile returns the profile URL of the plugin author of WordPress.org. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAuthorProfile(slug string) (authorProfile string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Author.Profile, err
}

// PluginAuthorLink returns the link included in the author field returned. This is sometimes different to the link to the official profile on Wordpress.org. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAuthorLink(slug string) (authorLink string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Author.Link, err
}

// PluginAuthorSlug returns the slug of the official profile. This is the result of stripping the generic URL from the Profile field. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAuthorSlug(slug string) (authorSlug string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Author.Slug, err
}

// PluginTested returns the version number (in string form) for the highest version of WordPress that is supported. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginTested(slug string) (version string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Tested, err
}

// PluginRequiredWPVersion returns the version number (in string form) for the lowest version of WordPress that is supported. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRequiredWPVersion(slug string) (version string, err error) {
	plug, err := generatePlugin(slug)
	return plug.RequiredWP, err
}

// PluginRequiredPHP returns the phpVersion struct for the plugin. This contains a bool indicating if it is specified and the version number (in string format) if it is specified. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRequiredPHP(slug string) (version phpVersion, err error) {
	plug, err := generatePlugin(slug)
	return plug.RequiresPHP, err
}

// PluginRatingOverall returns the overall rating of the plugin expressed as an int out of 100. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingOverall(slug string) (overall int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Rating.Overall, err
}

// PluginRatingOneStar returns the number (as an int) of one star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingOneStar(slug string) (star int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Rating.Stars.One, err
}

// PluginRatingTwoStar returns the number (as an int) of two star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingTwoStar(slug string) (star int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Rating.Stars.Two, err
}

// PluginRatingThreeStar returns the number (as an int) of three star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingThreeStar(slug string) (star int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Rating.Stars.Three, err
}

// PluginRatingFourStar returns the number (as an int) of four star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingFourStar(slug string) (star int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Rating.Stars.Four, err
}

// PluginRatingFiveStar returns the number (as an int) of five star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingFiveStar(slug string) (star int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Rating.Stars.Five, err
}

// PluginRatingNumber returns the total count of ratings (as an int) for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingNumber(slug string) (number int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Rating.Number, err
}

// PluginSupportThreads returns the total number (as an int) of support threads for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginSupportThreads(slug string) (threads int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Support.Threads.Total, err
}

// PluginSupportThreadsSolved returns the total number (as an int) of solved support threads for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginSupportThreadsSolved(slug string) (threadsSolved int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Support.Threads.Solved, err
}

// PluginSupportThreadsOpen returns the total number (as an int) of open support threads for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginSupportThreadsOpen(slug string) (threadsOpen int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Support.Threads.Open, err
}

// PluginSupportThreadsSolvedPercentage returns the calculated closure percentage of support threads (expressed as an int) for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginSupportThreadsSolvedPercentage(slug string) (solvedPercent int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Support.SolvedPercentage, err
}

// PluginDownloads returns the total download count (as an int) for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginDownloads(slug string) (downloads int, err error) {
	plug, err := generatePlugin(slug)
	return plug.Downloads, err
}

// PluginLastUpdate returns the date and time of the last update to the plugin as a string. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginLastUpdate(slug string) (updated string, err error) {
	plug, err := generatePlugin(slug)
	return plug.LastUpdate, err
}

// PluginAdded returns the date the plugin was originally created, in string format. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAdded(slug string) (added string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Added, err
}

// PluginHomepage returns the homepage URL for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginHomepage(slug string) (homepage string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Homepage, err
}

// PluginDownloadLink returns the current download link for the plugin. In general this isn't the trunk download link, but the most recent versioned download. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginDownloadLink(slug string) (download string, err error) {
	plug, err := generatePlugin(slug)
	return plug.DownloadLink, err
}

// PluginScreenshots returns an array of screenshot structs containing the information for all of the plugins screenshots. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginScreenshots(slug string) (screenshots []screenshot, err error) {
	plug, err := generatePlugin(slug)
	return plug.Screenshots, err
}

// PluginTags returns an array of strings containing the Tags for the plugin. These are not the sanatized versions. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginTags(slug string) (tags []string, err error) {
	plug, err := generatePlugin(slug)
	return plug.Tags, err
}

// PluginVersions returns an array of version structs containin the information for all of the plugins available versions. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginVersions(slug string) (versions []version, err error) {
	plug, err := generatePlugin(slug)
	return plug.Versions, err
}

// PluginDonateLink returns the the URL to the plugins donation page. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginDonateLink(slug string) (download string, err error) {
	plug, err := generatePlugin(slug)
	return plug.DonateLink, err
}
