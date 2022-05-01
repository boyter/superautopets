package main

import "math/rand"

type Pet interface {
	GetAttack() int
	GetHealth() int
	GetLevel() int

	//Faint(pets []Pet)
	//LevelUp(pets []Pet)
}

type Ant struct {
	baseAttack    int
	baseHealth    int
	currentAttack int
	currentHealth int
	currentLevel  int
}

func CreateAnt() Ant {
	return Ant{
		baseAttack:    2,
		baseHealth:    1,
		currentAttack: 2,
		currentHealth: 1,
	}
}

func (a Ant) GetAttack() int {
	return a.currentAttack
}

func (a Ant) GetHealth() int {
	return a.currentHealth
}

func (a Ant) GetLevel() int {
	return a.currentLevel
}

func (a Ant) Faint(pets []Pet) {
	t := pets[rand.Intn(len(pets))]

	switch a.GetLevel() {
	case 1:
		t.GetAttack()
	case 2:
		t.GetAttack()
	case 3:
		t.GetAttack()
	}
}

type Fish struct {
	baseAttack    int
	baseHealth    int
	currentAttack int
	currentHealth int
	currentLevel  int
}

func CreateFish() Fish {
	return Fish{
		baseAttack:    2,
		baseHealth:    3,
		currentAttack: 2,
		currentHealth: 3,
	}
}

func (a Fish) GetAttack() int {
	return a.currentAttack
}

func (a Fish) GetHealth() int {
	return a.currentHealth
}

func (a Fish) GetLevel() int {
	return a.currentLevel
}
