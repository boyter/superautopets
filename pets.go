package main

import (
	"fmt"
	"math/rand"
)

type Pet struct {
	name          string
	baseAttack    int
	baseHealth    int
	currentAttack int
	currentHealth int
	currentLevel  int
	faint         Faint
	levelup       LevelUp
	sell          Sell
}

func (p *Pet) CurrentAttack() int {
	// TODO must include modifiers here
	return p.currentAttack
}

func (p *Pet) TakeDamage(damage int) {
	// TODO must include modifiers here
	p.currentHealth -= damage
	if p.currentHealth < 0 {
		p.currentHealth = 0
	}
}

func (p *Pet) Fainted() bool {
	return p.currentHealth <= 0
}

type Faint func(*Pet, []*Pet)

func NothingFaint(pet *Pet, pets []*Pet) {}

type LevelUp func(*Pet, []*Pet)

func NothingLevelUp(pet *Pet, pets []*Pet) {}

type Sell func(*Pet, []*Pet)

func NothingSell(pet *Pet, pets []*Pet) {}

func CreatePet(name string) (*Pet, error) {
	switch name {
	case "ant":
		return &Pet{
			name:          name,
			baseAttack:    2,
			baseHealth:    1,
			currentAttack: 2,
			currentHealth: 1,
			currentLevel:  1,
			faint:         AntFaint,
			levelup:       NothingLevelUp,
			sell:          NothingSell,
		}, nil
	case "fish":
		return &Pet{
			name:          name,
			baseAttack:    2,
			baseHealth:    3,
			currentAttack: 2,
			currentHealth: 3,
			currentLevel:  1,
			faint:         NothingFaint,
			levelup:       FishLevelUp,
			sell:          NothingSell,
		}, nil
	case "bever":
		return &Pet{
			name:          name,
			baseAttack:    2,
			baseHealth:    2,
			currentAttack: 2,
			currentHealth: 2,
			currentLevel:  1,
			faint:         NothingFaint,
			levelup:       NothingLevelUp,
			sell:          BeverSell,
		}, nil
	}

	return &Pet{}, nil
}

func AntFaint(pet *Pet, pets []*Pet) {
	// Faint: Give a random friend +2/+1
	// TODO ignore this pet, or any other that has fainted
	t := pets[rand.Intn(len(pets))]

	switch pet.currentLevel {
	case 1:
		t.currentAttack += 2
		t.currentHealth += 1
	case 2:
		t.currentAttack += 4
		t.currentHealth += 2
	case 3:
		t.currentAttack += 6
		t.currentHealth += 3
	}
}

func FishLevelUp(pet *Pet, pets []*Pet) {
	// Level-up: Give all friends +1/+1
	buff := 1
	if pet.currentLevel > 1 {
		buff = 2
	}

	for i := 0; i < len(pets); i++ {
		pets[i].currentHealth += buff
		pets[i].currentAttack += buff
	}
}

func BeverSell(pet *Pet, pets []*Pet) {
	// Sell: Give two random friends +1 health
	buff := 1
	switch pet.currentLevel {
	case 2:
		buff = 2
	case 3:
		buff = 3
	}

	fmt.Println(buff)
}
