package main

import "fmt"

type Character struct {
	Name        string
	HP          int 
	AttackPower int
}

func (me Character) Attack(enemy *Character) {
	fmt.Printf("%sの攻撃! %sに%dダメージ\n", me.Name, enemy.Name, me.AttackPower)
	enemy.HP -= me.AttackPower
}

func main() {
	a := Character {
		Name: "A",
		HP:   100,
		AttackPower: 30,
	}

	b := Character {
		Name: "B",
		HP:   90,
		AttackPower: 40,
	}

	fmt.Println(b.HP)
	a.Attack(&b)
	fmt.Println(b.HP)
}