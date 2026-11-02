package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadToType[T any](path string) (*T, error) {
	var Data *T

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	defer file.Close()

	// Decode the JSON configuration into the config struct.
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to open decode JSON config: %w", err)
	}

	return Data, nil
}
