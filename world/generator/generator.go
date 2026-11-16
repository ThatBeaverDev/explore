package world_generator

import (
	"explore/loader/registry"
	"explore/types"
	"sync"

	fastnoise "github.com/setanarut/fastnoise"
)

type Generator struct {
	Registry        *registry.Registry
	ChunkQueue      []*types.Chunk
	ChunkQueueMutex sync.RWMutex

	seed int

	temperatureNoise *fastnoise.NoiseState[float64]
	humidityNoise    *fastnoise.NoiseState[float64]
	altitudeNoise    *fastnoise.NoiseState[float64]
	blockNoise       *fastnoise.NoiseState[float64]
}

func New(registry *registry.Registry, seed int) (*Generator, error) {
	blockNoise := fastnoise.NewNoiseState[float64]()
	blockNoise.SetNoiseType(fastnoise.Perlin)
	blockNoise.Seed = seed

	temperatureNoise := fastnoise.NewNoiseState[float64]()
	temperatureNoise.SetNoiseType(fastnoise.Perlin)
	temperatureNoise.Seed = seed + 1

	humidityNoise := fastnoise.NewNoiseState[float64]()
	humidityNoise.SetNoiseType(fastnoise.Perlin)
	humidityNoise.Seed = seed + 2

	altitudeNoise := fastnoise.NewNoiseState[float64]()
	altitudeNoise.SetNoiseType(fastnoise.Perlin)
	altitudeNoise.Seed = seed + 3

	return &Generator{
		Registry:        registry,
		ChunkQueue:      []*types.Chunk{},
		ChunkQueueMutex: sync.RWMutex{},

		temperatureNoise: temperatureNoise,
		humidityNoise:    humidityNoise,
		altitudeNoise:    altitudeNoise,

		seed:       seed,
		blockNoise: blockNoise,
	}, nil
}
