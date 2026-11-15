package world_generator

import (
	"explore/types"
	"math"
)

func (g *Generator) BlockRegionType(chunkPos types.ChunkPos) *types.PackRegion {
	x := float64(chunkPos[0])
	y := float64(chunkPos[1])

	temperature := g.temperatureNoise.Noise2D(x, y)
	humidity := g.humidityNoise.Noise2D(x, y)
	altitude := g.altitudeNoise.Noise2D(x, y)

	closestAverage := 1e60
	var bestRegion *types.PackRegion

	for _, region := range g.Registry.Regions {
		temperatureDiff := math.Abs(region.Temperature - temperature)
		humidityDiff := math.Abs(region.Humidity - humidity)
		altitudeDiff := math.Abs(region.Altitude - altitude)

		sum := temperatureDiff + humidityDiff + altitudeDiff
		average := sum / 3

		if average < closestAverage {
			closestAverage = average
			bestRegion = region
		}
	}

	return bestRegion
}
