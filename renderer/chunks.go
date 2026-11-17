package renderer

import (
	"explore/constants"
	"explore/types"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) DrawChunk(img *ebiten.Image, screenX, screenY float64, chunk *types.Chunk) {
	//vector.FillRect(img, chunkX, chunkY, constants.CHUNK_SIZE, constants.CHUNK_SIZE, color.RGBA{255, 255, 255, 255}, true)

	cx := chunk.Pos[0] * constants.CHUNK_SIZE
	cz := chunk.Pos[1] * constants.CHUNK_SIZE

	textureScale := g.BlockPx / 16.0

	for x, row := range chunk.Blocks {
		for z, col := range row {
			for _, block := range col {
				if block.Block == nil {
					continue
				}

				//blockX := cx + x
				//blockZ := cz - z

				blockScreenX := screenX + (float64(x) * g.BlockPx)
				blockScreenY := screenY + (float64(z) * g.BlockPx)

				op := &ebiten.DrawImageOptions{}
				//op.GeoM.Translate(-8, -8)
				//op.GeoM.Rotate(((4359 * float64(blockX)) + (7543 * float64(blockZ))) * math.Pi / 2)
				//op.GeoM.Translate(8, 8)

				op.GeoM.Scale(textureScale, textureScale)
				op.GeoM.Translate(blockScreenX, blockScreenY)

				if block.Block.Block.Id == "explore.grass" {
					op.ColorScale.SetR(float32(block.Region.GrassTint[0]))
					op.ColorScale.SetG(float32(block.Region.GrassTint[1]))
					op.ColorScale.SetB(float32(block.Region.GrassTint[2]))
				}

				if cx+x == int(math.Floor(g.World.X)) && cz-z == int(math.Floor(g.World.Z)) {

				} else {
					img.DrawImage(block.Block.Top, op)
				}
			}
		}
	}

}
