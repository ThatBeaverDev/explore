package loader

import (
	"explore/loader/registry"
	"explore/types"
	"fmt"
)

type PackLoader struct {
	Packs []*types.Pack

	Registry *registry.Registry
}

func New() (*PackLoader, error) {
	packs := []*types.Pack{}

	pack, err := loadPack("./content")
	if err != nil {
		return nil, fmt.Errorf("failed to load pack './content': %w", err)
	}
	packs = append(packs, pack)

	loaderRegistry := registry.New(packs)

	return &PackLoader{
		Packs:    packs,
		Registry: loaderRegistry,
	}, nil
}
