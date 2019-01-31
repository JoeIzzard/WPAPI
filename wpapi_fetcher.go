package wpapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchData(url string) (raw rawData, err error) {
	// Empty Data, returned in the event of an error or used to store the Raw Data
	var data = rawData{}

	wpClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return data, err
	}

	req.Header.Set("User-Agent", "WPAPIServ-Go")

	res, getErr := wpClient.Do(req)
	if getErr != nil {
		return data, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return data, readErr
	}

	jsonErr := json.Unmarshal([]byte(body), &data.Raw)
	if jsonErr != nil {
		return data, jsonErr
	}

	return data, nil
}
