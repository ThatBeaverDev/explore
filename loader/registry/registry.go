package registry

import (
	"explore/types"
)

type BlocksMap map[string]*types.LoadedBlock
type RegionsMap map[string]*types.PackRegion

type Registry struct {
	Blocks  BlocksMap
	Regions RegionsMap
}

func New(packs []*types.Pack) *Registry {
	blocks := make(BlocksMap)
	regions := make(RegionsMap)

	for _, pack := range packs {
		for _, block := range pack.Blocks {
			blocks[block.Block.Id] = block
		}

		for _, region := range pack.Regions {
			regions[region.Id] = region
		}
	}

	return &Registry{
		Blocks: blocks,

		Regions: regions,
	}
}
