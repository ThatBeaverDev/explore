package world_generator

import (
	"explore/types"
	"math"

	"github.com/setanarut/fastnoise"
)

func (g *Generator) BlockRegionType(chunkPos types.ChunkPos) *types.PackRegion {
	x := float64(chunkPos[0])
	y := float64(chunkPos[1])

	temperature := fastnoise.Value2D(x, y, g.temperatureNoise)
	humidity := fastnoise.Value2D(x, y, g.humidityNoise)
	altitude := fastnoise.Value2D(x, y, g.altitudeNoise)

	closestAverage := 1e60
	var bestRegion *types.PackRegion

	for _, region := range g.Registry.Regions {
		temperatureDiff := math.Abs(region.Temperature - float64(temperature))
		humidityDiff := math.Abs(region.Humidity - float64(humidity))
		altitudeDiff := math.Abs(region.Altitude - float64(altitude))

		sum := temperatureDiff + humidityDiff + altitudeDiff
		average := sum / 3

		if average < closestAverage {
			closestAverage = average
			bestRegion = region
		}
	}

	return bestRegion
}
