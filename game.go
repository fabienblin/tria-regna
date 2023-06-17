package main

import (
	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameInstance struct {
	Image        *ebiten.Image
	X, Y         float64
	Altitudes    *perlin.Perlin
	WaterSources *perlin.Perlin
	Continents   *perlin.Perlin
}

func initGame() *GameInstance {
	g := &GameInstance{
		Image:        ebiten.NewImage(globals.ScreenWidth, globals.ScreenHeight),
		X:            float64(globals.ScreenWidth) / 2,
		Y:            float64(globals.ScreenHeight) / 2,
		Altitudes:    perlin.NewPerlin(globals.Octaves, globals.Persistence, globals.Lacunarity, globals.Seed),
		WaterSources: perlin.NewPerlin(globals.Octaves, globals.Persistence, globals.Lacunarity, globals.Seed),
		Continents:   perlin.NewPerlin(2, 3, 2, globals.Seed+1),
	}

	// Set the game window properties
	ebiten.SetWindowSize(globals.ScreenWidth, globals.ScreenHeight)
	ebiten.SetFullscreen(true)

	ebiten.SetWindowTitle("Tria Regna")
	ebiten.SetMaxTPS(25)

	return g
}

func (g *GameInstance) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Y += globals.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Y -= globals.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.X += globals.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.X -= globals.Speed
	}

	// zoom
	perlinCenterX := XOFF(globals.ScreenWidth / 2)
	perlinCenterY := YOFF(globals.ScreenHeight / 2)

	if ebiten.IsKeyPressed(ebiten.KeyKPAdd) {
		if globals.Scale*globals.Zoom < 1 {
			globals.Scale *= globals.Zoom // Increase the scale factor

			newPerlinCenterX := XOFF(globals.ScreenWidth / 2)
			newPerlinCenterY := YOFF(globals.ScreenHeight / 2)

			offsetX := (newPerlinCenterX - perlinCenterX) / globals.Scale
			offsetY := (newPerlinCenterY - perlinCenterY) / globals.Scale

			g.X += (offsetX)
			g.Y += (offsetY)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyKPSubtract) {
		if globals.Scale/globals.Zoom > 0 {
			globals.Scale /= globals.Zoom // Decrease the scale factor

			newPerlinCenterX := XOFF(globals.ScreenWidth / 2)
			newPerlinCenterY := YOFF(globals.ScreenHeight / 2)

			offsetX := (newPerlinCenterX - perlinCenterX) / globals.Scale
			offsetY := (newPerlinCenterY - perlinCenterY) / globals.Scale

			g.X += (offsetX)
			g.Y += (offsetY)
		}
	}

	return nil
}

func (g *GameInstance) Draw(screen *ebiten.Image) {
	// op := &ebiten.DrawImageOptions{}
	GenerateTerrain()
	// GenerateRivers(findLocalMinima(g.Image))

	screen.DrawImage(g.Image, nil)
}

func (g *GameInstance) Layout(outsideWidth, outsideHeight int) (int, int) {
	return globals.ScreenWidth, globals.ScreenHeight
}
