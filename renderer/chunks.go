package renderer

import (
	"explore/types"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) DrawChunk(img *ebiten.Image, screenX, screenY float64, chunk *types.Chunk) {
	//vector.FillRect(img, chunkX, chunkY, constants.CHUNK_SIZE, constants.CHUNK_SIZE, color.RGBA{255, 255, 255, 255}, true)

	for x, row := range chunk.Blocks {
		for z, col := range row {
			for _, block := range col {
				if block.Block == nil {
					continue
				}

				posX := screenX + (float64(x) * g.Scale)
				posY := screenY + (float64(z) * g.Scale)

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(posX, posY)

				img.DrawImage(block.Block.Top, op)
			}
		}
	}

}
