package main

func RockThrow(active *Monster, passive *Monster) {
	active.MagicAttack(passive)
	active.PhysicalAttack(passive)
}

func BarkShield(active *Monster, passive *Monster) {
	active.Shield(100)
}