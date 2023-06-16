package main

import (

)

type Player struct {
	Name string
	Pos Position
	Monsters []Monster
}

func (player *Player) Move(direction Position) {
	player.Pos.x += direction.x
	player.Pos.y += direction.y
}