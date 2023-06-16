package main

import (
	"math/rand"
	"time"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"
)

var globals *Global
var game *GameInstance

func init() {
	rand.Seed(time.Now().UnixNano())

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatalf("fatal error config file: %s", err.Error())
		panic(nil)
	}

	globals = initGlobals()
	game = initGame()
}

func main() {
	// Initialise the game


	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
