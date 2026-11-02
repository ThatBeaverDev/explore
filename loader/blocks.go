package loader

import (
	"explore/types"
	"explore/util"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadPackBlocks(packDir string) ([]*types.LoadedBlock, error) {
	blocksDir := filepath.Join(packDir, "blocks")
	blocksFiles, err := os.ReadDir(blocksDir)
	if err != nil {
		return nil, fmt.Errorf("failed to list blocks directory, does it exist?: %w", err)
	}

	var blocks []*types.LoadedBlock
	for _, entry := range blocksFiles {
		if !entry.IsDir() {
			// file

			info, err := entry.Info()
			if err != nil {
				return nil, fmt.Errorf("failed to get info of block file entry: %w", err)
			}

			name := info.Name()
			_, isJson := strings.CutSuffix(name, ".json")

			if isJson {
				// load it
				path := filepath.Join(blocksDir, name)

				loadedBlock, err := loadBlock(path)
				if err != nil {
					return nil, fmt.Errorf("failed to load block: %w", err)
				}

				blocks = append(blocks, loadedBlock)
			}
		}
	}

	return blocks, nil
}

func loadBlock(dir string) (*types.LoadedBlock, error) {
	block, err := util.ReadToType[types.PackBlock](dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read block json entry "+dir+": %w", err)
	}

	topImagePath := filepath.Join(filepath.Dir(dir), block.Textures.Top)
	topImage, _, err := ebitenutil.NewImageFromFile(topImagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read top texture file for block "+filepath.Base(dir)+": %w", err)
	}

	sideImagePath := filepath.Join(filepath.Dir(dir), block.Textures.Side)
	sideImage, _, err := ebitenutil.NewImageFromFile(sideImagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read side texture file for block "+filepath.Base(dir)+": %w", err)
	}

	return &types.LoadedBlock{
		Block: block,

		Top:  topImage,
		Side: sideImage,
	}, nil
}
