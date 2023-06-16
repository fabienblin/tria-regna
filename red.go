package main

func FireBall(active *Monster, passive *Monster) {
	active.MagicAttack(passive)
}

func LightSpeed(active *Monster, passive *Monster) {
	active.Speed *= 2
}