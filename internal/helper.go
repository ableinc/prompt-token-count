package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func LoadJSONFile(filename string) (map[string]int, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filepath.Join(dir, "config", filename))
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
