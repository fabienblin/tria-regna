package main

import (
	"image/color"
	"log"

	"github.com/spf13/viper"
)

type Global struct {
	ScreenWidth		int
	ScreenHeight	int
	Zoom			float64
	Scale			float64	// Échelle du bruit Perlin
	Octaves			float64		// Nombre d'octaves pour le bruit Perlin
	Persistence		float64	// Persistance pour le bruit Perlin
	Lacunarity		int32		// Lacunarité pour le bruit Perlin
	Threshold		float64	// Seuil pour déterminer si un pixel est terre ou eau
	Seed			int64		// Graine pour la génération aléatoire

	// map displacement
	Speed			float64

	// altitude limits
	SeaAltitude		float64
	CoastAltitude	float64
	GroundAltitude	float64
	SnowAltitude	float64

	// rivers
	WaterSourceAltitude float64

	// continents
	ContinentThreshold	float64

	// colors
	Blue	color.RGBA
	Red		color.RGBA
	Green	color.RGBA
	White	color.RGBA
	Black	color.RGBA
}

func initGlobals() *Global {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatalf("fatal error config file: %s", err.Error())
		panic(nil)
	}
	a := viper.GetFloat64("Scale")
	_ = a
	return &Global{
		ScreenWidth:			viper.GetInt("ScreenWidth"),
		ScreenHeight:			viper.GetInt("ScreenHeight"),
		Zoom:					viper.GetFloat64("Zoom"),
		Scale:					viper.GetFloat64("Scale"),
		Octaves:				viper.GetFloat64("Octaves"),
		Persistence:			viper.GetFloat64("Persistence"),
		Lacunarity:				viper.GetInt32("Lacunarity"),
		Threshold:				viper.GetFloat64("Threshold"),
		Seed:					viper.GetInt64("Seed"),
		Speed:					viper.GetFloat64("Speed"),
		SeaAltitude:			viper.GetFloat64("SeaAltitude"),
		CoastAltitude:			viper.GetFloat64("CoastAltitude"),
		GroundAltitude:			viper.GetFloat64("GroundAltitude"),
		SnowAltitude:			viper.GetFloat64("SnowAltitude"),
		WaterSourceAltitude:	viper.GetFloat64("WaterSourceAltitude"),
		ContinentThreshold:		viper.GetFloat64("ContinentThreshold"),
		Blue:					color.RGBA{0, 0, 255, 255},
		Red:					color.RGBA{255, 0, 0, 255},
		Green:					color.RGBA{0, 255, 0, 255},
		White:					color.RGBA{255, 255, 255, 255},
		Black:					color.RGBA{0, 0, 0, 255},
	}
}

type Position struct {
	x, y int
}

func XOFF(x int) float64 {
	return (float64(x)-game.X) * globals.Scale
}

func YOFF(y int) float64 {
	return (float64(y)-game.Y) * globals.Scale
}
