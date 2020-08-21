package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	latestReleaseURL = "https://factorio.com/api/latest-releases"
)

// LatestRelease fetches the latest experimental and stable Factorio versions.
//
// Supported build arguments include "alpha", "demo", and "headless".
func LatestRelease(build string) (experimental, stable Version, err error) {
	resp, err := http.Get(latestReleaseURL)
	if err != nil {
		return NilVersion, NilVersion, fmt.Errorf("get: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return NilVersion, NilVersion, fmt.Errorf("response status: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	m := struct {
		Experimental map[string]Version `json:"experimental"`
		Stable       map[string]Version `json:"stable"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return NilVersion, NilVersion, fmt.Errorf("decode body: %v", err)
	}
	experimental, _ = m.Experimental[build]
	stable, _ = m.Stable[build]
	return experimental, stable, nil
}
