package utils

import (
	"encoding/json"
	"os"
)

func LoadManifest(path string) (map[string]interface{}, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var manifest map[string]interface{}
	if err := json.Unmarshal(file, &manifest); err != nil {
		return nil, err
	}

	return manifest, nil
}
