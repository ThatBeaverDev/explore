package loader

import (
	"explore/types"
	"explore/util"
	"fmt"
	"path/filepath"

	_ "image/jpeg"
)

func loadPack(dir string) (*types.Pack, error) {
	manifestPath := filepath.Join(dir, "manifest.json")
	manifest, err := util.ReadToType[types.PackManifest](manifestPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest file: %w", err)
	}

	blocks, err := LoadPackBlocks(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to load pack blocks: %w", err)
	}

	regions, err := LoadPackRegions(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to load pack regions: %w", err)
	}

	biomes, err := LoadPackBiomes(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to load pack biomes: %w", err)
	}

	pack := &types.Pack{
		Manifest: manifest,
		Blocks:   blocks,

		Regions: regions,
		Biomes:  biomes,
	}

	return pack, nil
}
