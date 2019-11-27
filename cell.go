package karta

import (
	"image/color"

	"github.com/pzsz/voronoi"
)

// Cell is the smalles unit on the map
type Cell struct {
	Index          int            `json:"index"`
	CenterDistance float64        `json:"center_distance"`
	NoiseLevel     float64        `json:"noise_level"`
	Elevation      float64        `json:"elevation"`
	Precipitation  float64        `json:"precipitation"`
	Land           bool           `json:"land"`
	Biome          Biome          `json:"biome"`
	Site           voronoi.Vertex `json:"site"`
	FillColor      color.RGBA     `json:"fill_color"`
	StrokeColor    color.RGBA     `json:"stroke_color"`
}

func (c *Cell) adjustElevation(u float64) {
	d := c.CenterDistance
	if d < u*3.3 {
		c.Elevation += 0.3
	}

	if d < u*2.3 {
		c.Elevation += 0.6
	}

	if d < u*1.3 {
		c.Elevation += 0.9
	}

	if c.Elevation < 0 {
		c.Land = false
	}
}

func (c *Cell) addLake() {
	if c.NoiseLevel < -.7 {
		c.Biome = SaltwaterLake
	} else if c.NoiseLevel < -.3 {
		c.Biome = FreshwaterLake
	}
}

func (c *Cell) calculateBiome() Biome {
	if c.Land {

		if c.NoiseLevel < -0.75 {
			c.addLake()
			return c.Biome
		}

		if c.Elevation < 1 {
			return Beach
		}

		if c.Elevation > 8 {
			if c.Precipitation < .2 {
				return BareTundra
			}
			if c.Precipitation < .5 {
				return Tundra
			}
			return Snow
		}

		if c.Elevation > 6 {
			if c.Precipitation < 1.0/3.0 {
				return TemperateDesert
			}
			if c.Precipitation < 2.0/3.0 {
				return Shrubland
			}
			return Tiaga
		}

		if c.Elevation > 3 {
			if c.Precipitation < 1.0/6.0 {
				return TemperateDesert
			}
			if c.Precipitation < 0.5 {
				return Grassland
			}
			if c.Precipitation < 5.0/6.0 {
				return DeciduousForest
			}
			return TemperateRainforest
		}

		if c.Precipitation < 1.0/6.0 {
			return SubtropicalDesert
		}

		if c.Precipitation < 1.0/3.0 {
			return Grassland
		}

		if c.Precipitation < 2.0/3.0 {
			return TropicalSeasonalForest
		}

		return TropicalRainforest

	}

	if c.Elevation < -1.5 {
		return DeepOcean
	}
	if c.Elevation < -1 {
		return ShallowOcean
	}
	return CoastalOcean

}

// Biome are defined by moisture and elevation
type Biome string

const (
	DeepOcean              Biome = "deep_ocean"
	ShallowOcean           Biome = "shallow_ocean"
	CoastalOcean           Biome = "coastal_ocean"
	SaltwaterLake          Biome = "saltwater_lake"
	FreshwaterLake         Biome = "freshwater_lake"
	Beach                  Biome = "beach"
	BareTundra             Biome = "bare_tundra"
	Tundra                 Biome = "tundra"
	Snow                   Biome = "snow"
	TemperateDesert        Biome = "temperate_desert"
	Shrubland              Biome = "shrubland"
	Tiaga                  Biome = "tiaga"
	Grassland              Biome = "grassland"
	DeciduousForest        Biome = "deciduous_forest"
	TemperateRainforest    Biome = "temperate_rainforest"
	SubtropicalDesert      Biome = "subtropical_desert"
	TropicalSeasonalForest Biome = "tropical_seasonal_forest"
	TropicalRainforest     Biome = "tropical_rainforest"
)
