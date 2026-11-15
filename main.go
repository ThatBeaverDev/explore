package main

import (
	"explore/loader"
	"explore/renderer"
	worlds "explore/world"
	world_generator "explore/world/generator"
	"fmt"
)

func main() {
	packloader, err := loader.New()
	if err != nil {
		panic(err)
	}

	for _, pack := range packloader.Packs {
		for _, block := range pack.Blocks {
			fmt.Println(block.Block)
			fmt.Println(block.Side)
			fmt.Println(block.Top)
		}

		for _, region := range pack.Regions {
			fmt.Println(region)
		}
	}

	world := worlds.New()
	generator := world_generator.New(packloader.Registry)

	gui := renderer.New(world, generator)
	gui.StartGame()
}
