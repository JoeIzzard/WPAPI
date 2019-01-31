package wpapi

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
