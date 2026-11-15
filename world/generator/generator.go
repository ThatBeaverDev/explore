package world_generator

import (
	"explore/loader/registry"

	"github.com/aquilax/go-perlin"
)

type Generator struct {
	Registry *registry.Registry

	temperatureNoise perlin.Perlin
	humidityNoise    perlin.Perlin
	altitudeNoise    perlin.Perlin
}

func New(registry *registry.Registry) *Generator {
	return &Generator{
		Registry: registry,

		temperatureNoise: *perlin.NewPerlin(2, 2, 3, 0),
		humidityNoise:    *perlin.NewPerlin(2, 2, 3, 1),
		altitudeNoise:    *perlin.NewPerlin(2, 2, 3, 2),
	}
}
