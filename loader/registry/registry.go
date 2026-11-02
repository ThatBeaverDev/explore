package registry

import (
	"explore/types"
)

type BlocksMap map[string]*types.LoadedBlock
type RegionsMap map[string]*types.PackRegion
type BiomesMap map[string]*types.PackBiome

type Registry struct {
	Blocks  BlocksMap
	Regions RegionsMap
	Biomes  BiomesMap
}

func New(packs []*types.Pack) *Registry {
	blocks := make(BlocksMap)
	regions := make(RegionsMap)
	biomes := make(BiomesMap)

	for _, pack := range packs {
		for _, block := range pack.Blocks {
			blocks[pack.Manifest.Id+"."+block.Block.Id] = block
		}

		for _, region := range pack.Regions {
			regions[pack.Manifest.Id+"."+region.Id] = region
		}

		for _, biome := range pack.Biomes {
			biomes[pack.Manifest.Id+"."+biome.Id] = biome
		}
	}

	return &Registry{
		Blocks: blocks,

		Regions: regions,
		Biomes:  biomes,
	}
}
