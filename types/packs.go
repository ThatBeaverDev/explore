package types

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type PackManifest struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Description string `json:"description,omitempty"`
}

type PackBlockTextures struct {
	Top  string `json:"top"`
	Side string `json:"side"`
}

type PackBlock struct {
	Name string `json:"name"`
	Id   string `json:"id"`

	Type     string            `json:"type"` // block
	Textures PackBlockTextures `json:"textures"`
}

type LoadedBlock struct {
	Block *PackBlock

	Top  *ebiten.Image
	Side *ebiten.Image
}

type PackRegion struct {
	Name string `json:"name"`
	Id   string `json:"id"`

	GrassTint [3]float32 `json:"grassTint"`

	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Altitude    float64 `json:"altitude"`

	SurfaceBlock           string  `json:"surfaceBlock"`
	SurfacePerlinAltBlock  string  `json:"surfacePerlinBlock"`
	SurfacePerlinThreshold float64 `json:"surfacePerlinThreshold"`
	UndergroundBlock       string  `json:"undergroundBlock"`
}

type Pack struct {
	Manifest *PackManifest

	Blocks []*LoadedBlock

	Regions []*PackRegion
}
