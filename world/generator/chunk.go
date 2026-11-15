package world_generator

import (
	"explore/constants"
	"explore/types"
	"fmt"
)

func (g *Generator) GenerateChunk(pos types.ChunkPos) types.Chunk {
	chunk := types.Chunk{
		Pos:    pos,
		Blocks: [constants.CHUNK_SIZE][constants.CHUNK_SIZE][constants.WORLD_HEIGHT_LIMIT]types.ChunkBlock{},
	}

	for x := range constants.CHUNK_SIZE {
		for z := range constants.CHUNK_SIZE {
			region := g.BlockRegionType([2]int{x, z})
			surfaceBlock, ok := g.Registry.Blocks[region.SurfaceBlock]
			if !ok {
				fmt.Println(g.Registry)
				panic("Region " + region.Id + " requests block ID " + region.SurfaceBlock + " for surface but that block does not exist.")
			}

			chunk.Blocks[x][z] = [constants.WORLD_HEIGHT_LIMIT]types.ChunkBlock{
				{Block: surfaceBlock},
			}
		}
	}

	return chunk
}
