package world_generator

import (
	"explore/constants"
	"explore/types"
	"fmt"

	"github.com/setanarut/fastnoise"
)

func (g *Generator) GenerateChunk(pos types.ChunkPos) {
	chunk := types.Chunk{
		Pos:    pos,
		Blocks: [constants.CHUNK_SIZE][constants.CHUNK_SIZE][constants.WORLD_HEIGHT_LIMIT]types.ChunkBlock{},
		Loaded: true,
	}

	chunkX := chunk.Pos[0] * constants.CHUNK_SIZE
	chunkZ := chunk.Pos[1] * constants.CHUNK_SIZE

	for x := range constants.CHUNK_SIZE {
		for z := range constants.CHUNK_SIZE {
			region := g.BlockRegionType([2]int{x, z})
			surfaceBlock, ok := g.Registry.Blocks[region.SurfaceBlock]
			if !ok {
				fmt.Println(g.Registry)
				panic("Region " + region.Id + " requests block ID " + region.SurfaceBlock + " for surface but that block does not exist.")
			}

			block := surfaceBlock
			if region.SurfacePerlinAltBlock != "" {
				perlinBlock, ok := g.Registry.Blocks[region.SurfacePerlinAltBlock]
				if !ok {
					fmt.Println(g.Registry)
					panic("Region " + region.Id + " requests block ID " + region.UndergroundBlock + " for underground but that block does not exist.")
				}

				random := fastnoise.Value2D(10*(chunkX+x), 10*(chunkZ-z), g.blockNoise)
				if random > region.SurfacePerlinThreshold {
					block = perlinBlock
				}
			}

			chunk.Blocks[x][z] = [constants.WORLD_HEIGHT_LIMIT]types.ChunkBlock{
				{
					Block:  block,
					Region: region,
				},
			}
		}
	}

	g.ChunkQueueMutex.Lock()
	g.ChunkQueue = append(g.ChunkQueue, &chunk)
	g.ChunkQueueMutex.Unlock()
}
