package loader

import (
	"explore/types"
	"explore/util"
	"fmt"
	"path/filepath"
	"sync"

	_ "image/jpeg"
)

func loadPack(dir string) (*types.Pack, error) {
	manifestPath := filepath.Join(dir, "manifest.json")
	manifest, err := util.ReadToType[types.PackManifest](manifestPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest file: %w", err)
	}

	var wg sync.WaitGroup

	var blocks []*types.LoadedBlock
	var regions []*types.PackRegion

	wg.Go(func() {
		packBlocks, ierr := LoadPackBlocks(manifest.Id, dir)
		blocks = packBlocks
		if ierr != nil && err == nil {
			err = fmt.Errorf("failed to load pack blocks: %w", ierr)
		}
	})
	wg.Go(func() {
		packRegions, ierr := LoadPackRegions(manifest.Id, dir)
		regions = packRegions
		if ierr != nil && err == nil {
			err = fmt.Errorf("failed to load pack regions: %w", ierr)
		}
	})

	wg.Wait()
	if err != nil {
		return nil, fmt.Errorf("failed to load pack blocks: %w", err)
	}

	pack := &types.Pack{
		Manifest: manifest,
		Blocks:   blocks,

		Regions: regions,
	}

	return pack, nil
}
