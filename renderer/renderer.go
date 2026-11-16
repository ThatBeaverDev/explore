package renderer

import (
	"explore/constants"
	"explore/types"
	worlds "explore/world"
	world_generator "explore/world/generator"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	width   int
	height  int
	BlockPx float64

	World     *types.World
	Generator *world_generator.Generator
}

func New(world *types.World, generator *world_generator.Generator) *Game {
	return &Game{
		BlockPx: 16,

		World:     world,
		Generator: generator,
	}
}

func (g *Game) Update() error {
	dt := 1.0 / 60.0

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.World.Z += constants.PLAYER_SPEED * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.World.Z -= constants.PLAYER_SPEED * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.World.X -= constants.PLAYER_SPEED * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.World.X += constants.PLAYER_SPEED * dt
	}

	_, dy := ebiten.Wheel()
	if dy > 0 {
		g.BlockPx /= 1.2
	} else if dy < 0 {
		g.BlockPx *= 1.2
	}

	worlds.InsureChunkPresence(g.World, g.Generator, [2]int{int(g.World.X), int(g.World.Z)})

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var playerX float64 = g.World.X
	var playerY float64 = g.World.Z

	cx := float64(g.width) / 2
	cy := float64(g.height) / 2

	chunks := worlds.ActiveChunks(g.World, [2]int{int(g.World.X), int(g.World.Z)})

	for _, chunkPos := range chunks {
		chunk, ok := g.World.LoadedChunks[chunkPos]
		if !ok || chunk.Loaded == false {
			continue // not loaded. that's fine.
		}

		chunkX := float64(chunkPos[0] * constants.CHUNK_SIZE)
		chunkY := float64(chunkPos[1] * constants.CHUNK_SIZE)

		relativeX := playerX - chunkX
		relativeY := playerY - chunkY

		screenX := cx - (relativeX * g.BlockPx)
		screenY := cy + (relativeY * g.BlockPx)

		g.DrawChunk(screen, screenX, screenY, chunk)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %f, Y: %f", g.World.X, g.World.Z))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.width = outsideWidth
	g.height = outsideHeight

	return outsideWidth, outsideHeight
}

func (g *Game) StartGame() {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Explore")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
