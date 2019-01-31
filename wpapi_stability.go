package wpapi

import (
	"errors"
	"os"
	"reflect"
)

func getStabilityList() (data map[string]string, err error) {
	var url string
	// Check the WPAPI Environment Variable to see if we are in debug mode
	debug := os.Getenv("WPAPIDEBUG")
	if debug == "1" {
		url = apiURLBuilderStability(true)
	} else {
		url = apiURLBuilderStability(false)
	}

	fetchData, fetchErr := fetchData(url)
	if fetchErr != nil {
		return data, fetchErr
	}

	var t map[string]interface{}
	ans := reflect.TypeOf(t).Kind()
	data = make(map[string]string)

	if reflect.TypeOf(fetchData.Raw).Kind() == ans {
		stabilityMap := fetchData.Raw
		for version, stability := range stabilityMap {
			data[version] = stability.(string)
		}
	}

	return data, nil
}

func StabilityList() (data map[string]string, err error) {
	return getStabilityList()
}

func StabilityVersion(version string) (stability string, err error) {
	stabilityList, err := getStabilityList()
	if stability, ok := stabilityList[version]; ok {
		return stability, err
	} else {
		return "", errors.New("Version not present in stability list")
	}
}

func StabilityLatest(version string) (latest bool, err error) {
	stability, err := StabilityVersion(version)
	if err != nil {
		return false, err
	}

	if stability == "latest" {
		return true, err
	} else {
		return false, err
	}
}

func StabilityNotLatest(version string) (latest bool, err error) {
	latest, err = StabilityLatest(version)
	return !latest, err
}

func StabilityOutdated(version string) (outdated bool, err error) {
	stability, err := StabilityVersion(version)
	if err != nil {
		return false, err
	}

	if stability == "outdated" {
		return true, err
	} else {
		return false, err
	}
}

func StabilityInsecure(version string) (insecure bool, err error) {
	stability, err := StabilityVersion(version)
	if err != nil {
		return false, err
	}

	if stability == "insecure" {
		return true, err
	} else {
		return false, err
	}
}
