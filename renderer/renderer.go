package renderer

import (
	"explore/constants"
	"explore/types"
	worlds "explore/world"
	world_generator "explore/world/generator"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	width  int
	height int
	Scale  float64

	World     *types.World
	Generator *world_generator.Generator
}

func New(world *types.World, generator *world_generator.Generator) *Game {
	return &Game{
		Scale: 16,

		World:     world,
		Generator: generator,
	}
}

func (g Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.World.Z += 0.2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.World.Z -= 0.2
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.World.X += 0.2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.World.X -= 0.2
	}

	worlds.InsureChunkPresence(g.World, g.Generator, [2]int{0, 0})

	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	var x float64 = g.World.X
	var y float64 = g.World.Z

	cx := float64(g.width) / 2
	cy := float64(g.height) / 2

	chunks := worlds.ActiveChunks(g.World, [2]int{int(g.World.X), int(g.World.Z)})

	for _, pos := range chunks {
		chunk, ok := g.World.LoadedChunks[pos]
		if !ok {
			continue // not loaded. that's fine.
		}

		rx := x - float64(pos[0])
		ry := y - float64(pos[1])

		screenX := cx + (rx * g.Scale * constants.CHUNK_SIZE)
		screenY := cy + (ry * g.Scale * constants.CHUNK_SIZE)

		g.DrawChunk(screen, screenX, screenY, chunk)
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.width = outsideWidth
	g.height = outsideHeight

	return outsideWidth, outsideHeight
}

func (g Game) StartGame() {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Explore")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
