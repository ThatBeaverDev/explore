package worlds

import (
	"explore/constants"
	"explore/types"
	world_generator "explore/world/generator"
	"math"
)

func New() *types.World {
	return &types.World{
		Version:      constants.CURRENT_WORLD_VERSION,
		LoadedChunks: make(types.ChunkMap),
	}
}

func ActiveChunks(world *types.World, pos types.ChunkPos) []types.ChunkPos {
	var chunks []types.ChunkPos

	LOADED_WIDTH_OFFSET := int(math.Ceil(constants.LOADED_WIDTH / 2))

	for rx := range constants.LOADED_WIDTH {
		x := rx - LOADED_WIDTH_OFFSET
		for ry := range constants.LOADED_WIDTH {
			y := ry - LOADED_WIDTH_OFFSET

			pos := [2]int{x, y}

			chunks = append(chunks, pos)
		}
	}

	return chunks
}

func InsureChunkPresence(world *types.World, generator *world_generator.Generator, pos types.ChunkPos) {
	chunks := ActiveChunks(world, pos)

	for _, pos := range chunks {
		_, ok := world.LoadedChunks[pos]
		if !ok {
			// not loaded
			chunk := generator.GenerateChunk(pos)

			world.LoadedChunks[pos] = &chunk
		}

	}
}
