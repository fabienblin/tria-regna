package main

func CuttingJet(active *Monster, passive *Monster) {
	active.MagicAttack(passive)
}

func WaveHit(active *Monster, passive *Monster) {
	active.Stun(passive)
}