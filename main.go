package main

import (
	"explore/loader"
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

		for _, biome := range pack.Biomes {
			fmt.Println(biome)
		}
	}
}
