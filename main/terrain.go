package main

import (
	"image"
	"image/color"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
)

func GenerateRivers(localMinimas []*Position) {
	for _, lm := range localMinimas {
		altitude := game.WaterSources.Noise2D(XOFF(lm.x), YOFF(lm.y))
		isWaterSource := altitude > globals.WaterSourceAltitude
		if isWaterSource {
			// game.Image.Set(lm.x, lm.y, globals.Blue)
			fillWaterSource(game.Image, lm, altitude)
		}
	}
}

// start precedAltitude with -1.0
func fillWaterSource(img *ebiten.Image, currentPosition *Position, currentAltitude float64) {
	if currentAltitude <= globals.SeaAltitude {
		// flowWater(img, currentPosition, precedAltitude)
		return
	}

	// find lowest point just above current point
	// lowestPosition := *currentPosition
	lowestAltitude := 1.0
	spreadLimit := 20
	for spread := 1; spread < spreadLimit; spread++ {
		for x := -1 * spread; x <= 1*spread; x++ {
			for y := -1 * spread; y <= 1*spread; y++ {
				nextAlt := game.Altitudes.Noise2D(XOFF(x+currentPosition.x), YOFF(y+currentPosition.y))
				// nextPos := Position{currentPosition.x + x, currentPosition.y + y}
				if lowestAltitude > nextAlt {
					// lowestPosition = nextPos
					lowestAltitude = nextAlt
					// img.Set(lowestPosition.x + x, lowestPosition.y + y, globals.Blue)
				}
			}
		}

	}

	// fillWaterSource(img, &lowestPosition, lowestAltitude)
	// img.Set(currentPosition.x, currentPosition.y, globals.Blue)

	return
}

func GenerateTerrain() {
	for y := 0; y < globals.ScreenHeight; y++ {
		for x := 0; x < globals.ScreenWidth; x++ {
			continent := game.Continents.Noise2D(XOFF(x)*10, YOFF(y)*10)
			altitude := (game.Altitudes.Noise2D(XOFF(x), YOFF(y)) + 1)
			altitude *= altitude
			if continent > globals.ContinentThreshold {
				grey := uint8(255)
				game.Image.Set(x, y, color.RGBA{grey, grey, grey, 255})

			}


			// sea := altitude < globals.SeaAltitude
			// coast := altitude >= globals.SeaAltitude && altitude < globals.CoastAltitude
			// ground := altitude >= globals.CoastAltitude && altitude < globals.GroundAltitude
			// snow := altitude >= globals.GroundAltitude

			// if ground {
			// 	g := uint8(127.5 * (math.Cos(altitude*math.Pi) + 1))
			// 	r := uint8(127.5 * (-math.Cos(altitude*math.Pi) + 1))
			// 	b := uint8(20 * (math.Cos(altitude*math.Pi) + 1))
			// 	game.Image.Set(x, y, color.RGBA{r, g, b, 255})
			// } else if coast {
			// 	game.Image.Set(x, y, color.RGBA{250, 248, 190, 255})
			// } else if snow {
			// 	grey := uint8(255 - (255*(math.Cos(altitude*math.Pi)+1))/5)
			// 	game.Image.Set(x, y, color.RGBA{grey, grey, grey, 255})
			// } else if sea {
			// 	game.Image.Set(x, y, color.RGBA{0, 0, 255, 255})
			// }
		}
	}
}

func findLocalMinima(img *ebiten.Image) []*Position {
	localMinimas := []*Position{}
	// Parcourt les pixels de l'image
	for y := 0; y < globals.ScreenHeight; y++ {
		for x := 0; x < globals.ScreenWidth; x++ {
			// Vérifie si le pixel est un maximum local
			isLocalMinima := true

			// Obtient la valeur de bruit Perlin normalisée dans la plage [-1, 1]
			currentHeight := game.Altitudes.Noise2D(XOFF(x), YOFF(y))
			// Parcourt les pixels voisins
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					// Vérifie les limites de l'image
					nx := x + dx
					ny := y + dy
					if nx >= 0 && nx < globals.ScreenWidth && ny >= 0 && ny < globals.ScreenHeight {
						// Vérifie si le pixel voisin est plus élevé
						if game.Altitudes.Noise2D(XOFF(nx), YOFF(ny)) < currentHeight {
							isLocalMinima = false
							break
						}
					}
				}
				if !isLocalMinima {
					break
				}
			}
			// Si c'est un maximum local, rempli le pixel en bleu
			if isLocalMinima {
				localMinimas = append(localMinimas, &Position{x, y})
			}

		}
	}
	return localMinimas
}

// TODO: remove params game and altitude
func findLocalMaxima(game *GameInstance, img *image.RGBA, altitudes *perlin.Perlin) {

	// Parcourt les pixels de l'image
	for y := int(game.Y); y < globals.ScreenHeight; y++ {
		for x := int(game.X); x < globals.ScreenWidth; x++ {
			// Vérifie si le pixel est un maximum local
			isLocalMaxima := true
			xoff := float64(x) * globals.Scale
			yoff := float64(y) * globals.Scale

			// Obtient la valeur de bruit Perlin normalisée dans la plage [-1, 1]
			currentHeight := altitudes.Noise2D(xoff, yoff)
			if currentHeight > 0.0 {
				// Parcourt les pixels voisins
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						// Vérifie les limites de l'image
						nx := x + dx
						ny := y + dy
						if nx >= 0 && nx < globals.ScreenWidth && ny >= 0 && ny < globals.ScreenHeight {
							// Vérifie si le pixel voisin est plus élevé
							if altitudes.Noise2D(float64(nx)*globals.Scale, float64(ny)*globals.Scale) > currentHeight {
								isLocalMaxima = false
								break
							}
						}
					}
					if !isLocalMaxima {
						break
					}
				}

				// Si c'est un maximum local, rempli le pixel en bleu
				if isLocalMaxima {
					img.Set(x, y, color.RGBA{255, 0, 0, 255})
				}
			}
		}
	}
}
