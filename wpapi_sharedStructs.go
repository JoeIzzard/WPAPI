package wpapi

/*
 ----- Raw Data -----
 All JSON responses are unmarshalled into this RawData Struct, which is then formatted into the correct strcut/object for that type
*/
type rawData struct {
	Raw map[string]interface{}
}

/*
----- Stars -----
This struct contains the individual numbers for the 5 star rating
*/
type stars struct {
	One   int
	Two   int
	Three int
	Four  int
	Five  int
}

/*
----- Rating -----
This struct contains all of the Ratings data, including a reference to the Stars struct
*/
type rating struct {
	Overall int
	Stars   stars
	Number  int
}

/*
----- Versions -----
This strutc is used to manage the data for each version of an item
*/
type version struct {
	Version    string
	PreRelease bool
	Link       string
}
