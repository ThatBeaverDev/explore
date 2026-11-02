package loader

import (
	"explore/types"
	"explore/util"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadPackBiomes(packDir string) ([]*types.PackBiome, error) {
	biomesDir := filepath.Join(packDir, "biomes")
	biomesFiles, err := os.ReadDir(biomesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to list biomes directory, does it exist?: %w", err)
	}

	var biomes []*types.PackBiome
	for _, entry := range biomesFiles {
		if !entry.IsDir() {
			// file

			info, err := entry.Info()
			if err != nil {
				return nil, fmt.Errorf("failed to get info of biome file entry: %w", err)
			}

			name := info.Name()
			_, isJson := strings.CutSuffix(name, ".json")

			if isJson {
				// load it
				path := filepath.Join(biomesDir, name)

				loadedBiome, err := util.ReadToType[types.PackBiome](path)
				if err != nil {
					return nil, fmt.Errorf("failed to load biome: %w", err)
				}

				biomes = append(biomes, loadedBiome)
			}
		}
	}

	return biomes, nil
}
