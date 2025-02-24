package internal

import (
	"embed"
	"encoding/json"
)

//go:embed mergeableRanks.json
var merageRanksFile embed.FS

//go:embed specialTokens.json
var specialTokensFile embed.FS

func LoadMergeableRankFile() (map[string]int, error) {
	data, err := merageRanksFile.ReadFile("mergeableRanks.json")
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
	data, err := specialTokensFile.ReadFile("specialTokens.json")
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
