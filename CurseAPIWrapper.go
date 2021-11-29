package CurseAPIWrapper

import "encoding/json"

var (
	// TODO change to coreURL
	v1ApiURL = "https://api.curseforge.com/v1"
	apiKey = ""
)

func Initialize(key string)  {
	if key == "" {
        panic("CurseAPIWrapper: No API key provided")
    }
	apiKey = key
}

func GetMod(id int) GetModResults {
    return GetModResults{}
}

func GetMods(ids []int) (GetModsResults, error) {
	body, _ := json.Marshal(RequestMods{ModIds: ids})
	resp, err := makeRequest("POST", v1ApiURL + "/mods", body)
	if err != nil {
		return GetModsResults{}, err
	}
	var results GetModsResults
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return GetModsResults{}, err
	}
    return results, nil
}