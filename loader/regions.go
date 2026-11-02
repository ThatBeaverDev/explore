package loader

import (
	"explore/types"
	"explore/util"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadPackRegions(packDir string) ([]*types.PackRegion, error) {
	regionsDir := filepath.Join(packDir, "regions")
	regionsFiles, err := os.ReadDir(regionsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to list regions directory, does it exist?: %w", err)
	}

	var regions []*types.PackRegion
	for _, entry := range regionsFiles {
		if !entry.IsDir() {
			// file

			info, err := entry.Info()
			if err != nil {
				return nil, fmt.Errorf("failed to get info of region file entry: %w", err)
			}

			name := info.Name()
			_, isJson := strings.CutSuffix(name, ".json")

			if isJson {
				// load it
				path := filepath.Join(regionsDir, name)

				loadedRegion, err := util.ReadToType[types.PackRegion](path)
				if err != nil {
					return nil, fmt.Errorf("failed to load region: %w", err)
				}

				regions = append(regions, loadedRegion)
			}
		}
	}

	return regions, nil
}
