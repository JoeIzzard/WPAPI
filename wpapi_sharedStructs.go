package wpapi

/*
 ----- Raw Data -----
 All JSON responses are unmarshalled into this RawData Struct, which is then formatted into the correct strcut/object for that type
*/
type rawData struct {
	Raw map[string]interface{}
}
