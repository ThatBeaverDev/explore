package world_generator

import (
	"explore/constants"
	"explore/types"
	"fmt"
	"math/rand/v2"
	"time"

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
			undergroundBlock, ok := g.Registry.Blocks[region.UndergroundBlock]
			if !ok {
				fmt.Println(g.Registry)
				panic("Region " + region.Id + " requests block ID " + region.UndergroundBlock + " for underground but that block does not exist.")
			}

			random := fastnoise.Value2D(chunkX+x, chunkZ-z, g.blockNoise)

			block := surfaceBlock
			if random > 0 {
				block = undergroundBlock
			}

			chunk.Blocks[x][z] = [constants.WORLD_HEIGHT_LIMIT]types.ChunkBlock{
				{Block: block},
			}
		}
	}

	random := rand.Float64() * 2
	dur, err := time.ParseDuration(fmt.Sprintf("%fs", random))
	if err != nil {
		panic(err)
	}
	time.Sleep(dur)

	g.ChunkQueueMutex.Lock()
	g.ChunkQueue = append(g.ChunkQueue, &chunk)
	g.ChunkQueueMutex.Unlock()
}
