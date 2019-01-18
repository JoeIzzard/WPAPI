package wpapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mvdan/xurls"
)

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

type stars struct {
	One   int
	Two   int
	Three int
	Four  int
	Five  int
}

type rating struct {
	Overall int
	Stars   stars
	Number  int
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

type version struct {
	Version    string
	PreRelease bool
	Link       string
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

type pluginRaw struct {
	Raw map[string]interface{}
}

func generatePlugin(input pluginRaw, slug string) plugin {
	output := plugin{}

	// Ensure there wasn't an error
	if val, ok := input.Raw["error"]; ok {
		var err strings.Builder
		err.WriteString(val.(string))
		err.WriteString(" - ")
		err.WriteString(slug)
		log.Print(err.String())
		return output
	}

	// Name
	output.Name = input.Raw["name"].(string)
	output.Slug = input.Raw["slug"].(string)

	// Version
	output.Version = input.Raw["version"].(string)

	// // Author
	output.Author.Profile = input.Raw["author_profile"].(string)
	output.Author.Name = strip.StripTags(input.Raw["author"].(string))
	output.Author.Slug = strings.TrimLeft(input.Raw["author_profile"].(string), "https://profiles.wordpress.org/")
	output.Author.Link = xurls.Strict().FindString(input.Raw["author"].(string))

	// // Tested
	output.Tested = input.Raw["tested"].(string)

	// Required WordPress Version
	output.RequiredWP = input.Raw["requires"].(string)

	// // Required PHP Version
	switch input.Raw["requires_php"].(type) {
	case bool:
		output.RequiresPHP.Provided = false
		output.RequiresPHP.Version = ""
	case string:
		output.RequiresPHP.Provided = true
		output.RequiresPHP.Version = input.Raw["requires_php"].(string)
	default:
		output.RequiresPHP.Provided = false
		output.RequiresPHP.Version = ""
	}

	// // ** Ratings! **
	output.Rating.Overall = int(input.Raw["rating"].(float64))

	ratings := input.Raw["ratings"].(map[string]interface{})
	output.Rating.Stars.One = int(ratings["1"].(float64))
	output.Rating.Stars.Two = int(ratings["2"].(float64))
	output.Rating.Stars.Three = int(ratings["3"].(float64))
	output.Rating.Stars.Four = int(ratings["4"].(float64))
	output.Rating.Stars.Five = int(ratings["5"].(float64))

	output.Rating.Number = int(input.Raw["num_ratings"].(float64))

	// Support!
	output.Support.Threads.Total = int(input.Raw["support_threads"].(float64))
	output.Support.Threads.Solved = int(input.Raw["support_threads_resolved"].(float64))
	output.Support.Threads.Open = output.Support.Threads.Total - output.Support.Threads.Solved

	// Solved Percent
	output.Support.SolvedPercentage = int(math.Round((float64(output.Support.Threads.Solved) / float64(output.Support.Threads.Total)) * 100))

	// Downloads
	output.Downloads = int(input.Raw["downloaded"].(float64))

	// Last Update
	output.LastUpdate = input.Raw["last_updated"].(string)

	// Added
	output.Added = input.Raw["added"].(string)

	// Homepage
	output.Homepage = input.Raw["homepage"].(string)

	// Download Link
	output.DownloadLink = input.Raw["download_link"].(string)

	// *** Screenshots ***
	screenshotsMap := input.Raw["screenshots"].(map[string]interface{})
	for _, value := range screenshotsMap {
		newS := screenshot{}
		current := value.(map[string]interface{})
		newS.Caption = current["caption"].(string)
		newS.SRC = current["src"].(string)
		output.Screenshots = append(output.Screenshots, newS)
	}

	// *** Tags ***
	tagsMap := input.Raw["tags"].(map[string]interface{})
	for _, value := range tagsMap {
		output.Tags = append(output.Tags, value.(string))
	}

	// *** Versions ***
	versionMap := input.Raw["versions"].(map[string]interface{})
	for key, value := range versionMap {
		newVersion := version{}

		if strings.Contains(key, "beta") || strings.Contains(key, "alpha") {
			newVersion.PreRelease = true
		} else {
			newVersion.PreRelease = false
		}

		newVersion.Version = key
		newVersion.Link = value.(string)

		output.Versions = append(output.Versions, newVersion)
	}

	// Donate Link
	output.DonateLink = input.Raw["donate_link"].(string)

	// Send it Back!
	return output
}

func fetchPlugin(slug string) plugin {
	url := apiURLBuilder(slug, "plugin")

	wpClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "WPAPIServ-Go")

	res, getErr := wpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	plugin1 := pluginRaw{}
	jsonErr := json.Unmarshal([]byte(body), &plugin1.Raw)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	plugin2 := generatePlugin(plugin1, slug)

	return plugin2
}

// GetPlugin returns the plugin struct for a given plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func GetPlugin(slug string) plugin {
	return fetchPlugin(slug)
}

// PluginName returns the name of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginName(slug string) string {
	return fetchPlugin(slug).Name
}

// PluginVersion returns the version string of the plugin as a string. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginVersion(slug string) string {
	return fetchPlugin(slug).Version
}

// PluginAuthorName returns the name of the plugins author stripped from the author field in the results. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAuthorName(slug string) string {
	return fetchPlugin(slug).Author.Name
}

// PluginAuthorProfile returns the profile URL of the plugin author of WordPress.org. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAuthorProfile(slug string) string {
	return fetchPlugin(slug).Author.Profile
}

// PluginAuthorLink returns the link included in the author field returned. This is sometimes different to the link to the official profile on Wordpress.org. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAuthorLink(slug string) string {
	return fetchPlugin(slug).Author.Link
}

// PluginAuthorSlug returns the slug of the official profile. This is the result of stripping the generic URL from the Profile field. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAuthorSlug(slug string) string {
	return fetchPlugin(slug).Author.Slug
}

// PluginTested returns the version number (in string form) for the highest version of WordPress that is supported. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginTested(slug string) string {
	return fetchPlugin(slug).Tested
}

// PluginRequiredWPVersion returns the version number (in string form) for the lowest version of WordPress that is supported. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRequiredWPVersion(slug string) string {
	return fetchPlugin(slug).RequiredWP
}

// PluginRequiredPHP returns the phpVersion struct for the plugin. This contains a bool indicating if it is specified and the version number (in string format) if it is specified. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRequiredPHP(slug string) phpVersion {
	return fetchPlugin(slug).RequiresPHP
}

// PluginRatingOverall returns the overall rating of the plugin expressed as an int out of 100. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingOverall(slug string) int {
	return fetchPlugin(slug).Rating.Overall
}

// PluginRatingOneStar returns the number (as an int) of one star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingOneStar(slug string) int {
	return fetchPlugin(slug).Rating.Stars.One
}

// PluginRatingTwoStar returns the number (as an int) of two star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingTwoStar(slug string) int {
	return fetchPlugin(slug).Rating.Stars.Two
}

// PluginRatingThreeStar returns the number (as an int) of three star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingThreeStar(slug string) int {
	return fetchPlugin(slug).Rating.Stars.Three
}

// PluginRatingFourStar returns the number (as an int) of four star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingFourStar(slug string) int {
	return fetchPlugin(slug).Rating.Stars.Four
}

// PluginRatingFiveStar returns the number (as an int) of five star ratings of the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingFiveStar(slug string) int {
	return fetchPlugin(slug).Rating.Stars.Five
}

// PluginRatingNumber returns the total count of ratings (as an int) for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginRatingNumber(slug string) int {
	return fetchPlugin(slug).Rating.Number
}

// PluginSupportThreads returns the total number (as an int) of support threads for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginSupportThreads(slug string) int {
	return fetchPlugin(slug).Support.Threads.Total
}

// PluginSupportThreadsSolved returns the total number (as an int) of solved support threads for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginSupportThreadsSolved(slug string) int {
	return fetchPlugin(slug).Support.Threads.Solved
}

// PluginSupportThreadsOpen returns the total number (as an int) of open support threads for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginSupportThreadsOpen(slug string) int {
	return fetchPlugin(slug).Support.Threads.Open
}

// PluginSupportThreadsSolvedPercentage returns the calculated closure percentage of support threads (expressed as an int) for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginSupportThreadsSolvedPercentage(slug string) int {
	return fetchPlugin(slug).Support.SolvedPercentage
}

// PluginDownloads returns the total download count (as an int) for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginDownloads(slug string) int {
	return fetchPlugin(slug).Downloads
}

// PluginLastUpdate returns the date and time of the last update to the plugin as a string. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginLastUpdate(slug string) string {
	return fetchPlugin(slug).LastUpdate
}

// PluginAdded returns the date the plugin was originally created, in string format. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginAdded(slug string) string {
	return fetchPlugin(slug).Added
}

// PluginHomepage returns the homepage URL for the plugin. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginHomepage(slug string) string {
	return fetchPlugin(slug).Homepage
}

// PluginDownloadLink returns the current download link for the plugin. In general this isn't the trunk download link, but the most recent versioned download. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginDownloadLink(slug string) string {
	return fetchPlugin(slug).DownloadLink
}

// PluginScreenshots returns an array of screenshot structs containing the information for all of the plugins screenshots. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginScreenshots(slug string) []screenshot {
	return fetchPlugin(slug).Screenshots
}

// PluginTags returns an array of strings containing the Tags for the plugin. These are not the sanatized versions. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginTags(slug string) []string {
	return fetchPlugin(slug).Tags
}

// PluginVersions returns an array of version structs containin the information for all of the plugins available versions. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginVersions(slug string) []version {
	return fetchPlugin(slug).Versions
}

// PluginDonateLink returns the the URL to the plugins donation page. Uses the slug argument as the identifier provided to the WordPress.org API
func PluginDonateLink(slug string) string {
	return fetchPlugin(slug).DonateLink
}
