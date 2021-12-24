package CurseAPIWrapper

import (
	"encoding/json"
	"errors"
)

var (
	// TODO change to coreURL
	v1ApiURL = "https://api.curseforge.com/v1"
	apiKey   = ""
)

func Initialize(key string) {
	if key == "" {
		panic("CurseAPIWrapper: No API key provided")
	}
	apiKey = key
}

func GetMod(id string) (GetModResults, error) {
	resp, err := makeRequest("GET", v1ApiURL+"/mods/"+id, nil)
	if err != nil {
		return GetModResults{}, errors.New("[CurseAPIWrapper] makeRequest error: " + err.Error())
	}
	if resp.StatusCode != 200 {
		return GetModResults{}, errors.New("[CurseAPIWrapper] makeRequest error: " + resp.Status)
	}
	var results GetModResults
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return GetModResults{}, errors.New("[CurseAPIWrapper] json.NewDecoder error: " + err.Error())
	}
	return results, nil
}

func GetMods(ids []int) (GetModsResults, error) {
	body, _ := json.Marshal(RequestMods{ModIds: ids})
	resp, err := makeRequest("POST", v1ApiURL+"/mods", body)
	if err != nil {
		return GetModsResults{}, errors.New("[CurseAPIWrapper] makeRequest error: " + err.Error())
	}
	if resp.StatusCode != 200 {
		return GetModsResults{}, errors.New("[CurseAPIWrapper] makeRequest error: " + resp.Status)
	}
	var results GetModsResults
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return GetModsResults{}, errors.New("[CurseAPIWrapper] json.NewDecoder error: " + err.Error())
	}
	return results, nil
}
