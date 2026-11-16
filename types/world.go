package types

import (
	"explore/constants"
)

type ChunkBlock struct {
	Block *LoadedBlock
}

type ChunkBlocks [constants.CHUNK_SIZE][constants.CHUNK_SIZE][constants.WORLD_HEIGHT_LIMIT]ChunkBlock

type ChunkPos [2]int
type Chunk struct {
	Pos    ChunkPos    `json:"pos"`
	Blocks ChunkBlocks `json:"blocks"`
	Loaded bool        `json:"-"`
}

type ChunkMap map[ChunkPos]*Chunk

type World struct {
	Version      int
	LoadedChunks ChunkMap

	X float64
	Z float64
}
