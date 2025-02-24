package internal

import (
	"embed"
	"encoding/json"
)

//go:embed config/mergeableRanks.json
var merageRanksFile embed.FS

//go:embed config/specialTokens.json
var specialTokensFile embed.FS

func LoadMergeableRankFile() (map[string]int, error) {
	data, err := merageRanksFile.ReadFile("config/mergeableRanks.json")
	if err != nil {
		return nil, err
	}
	var result map[string]int
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func LoadSpecialTokensFile() (map[string]int, error) {
	data, err := specialTokensFile.ReadFile("config/specialTokens.json")
	if err != nil {
		return nil, err
	}
	var result map[string]int
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
