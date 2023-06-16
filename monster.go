package main

import (
)

type Monster struct {
	Name string
	MagicType int // RED | GREEN | BLUE
	Hp int
	ManaPool int
	Armor int
	Resist int
	Atk int
	MagicPow int
	Pos Position
	IsStunned bool
	Speed int
}

func (monster *Monster) Move(direction Position) {
	monster.Pos.x += direction.x
	monster.Pos.y += direction.y
}

func (monster *Monster) PhysicalAttack(enemy *Monster) {
	enemy.physicalDefend(monster.Atk)
}

func (monster *Monster) physicalDefend(damage int) {
	monster.Hp -= damage - monster.Armor
}

func (monster *Monster) MagicAttack(enemy *Monster) {
	enemy.MagicDefend(monster.MagicPow, monster.MagicType)
}

func (monster *Monster) MagicDefend(damage int, magicType int) {
	monster.Hp -= (int(float64(damage) * weaknesses(magicType, monster.MagicType))) - monster.Resist
}

func (monster *Monster) Shield(shield int) {
	monster.Resist += shield
	monster.Armor += shield
}

func (monster *Monster) Stun(enemy *Monster) {
	enemy.IsStunned = true
}
