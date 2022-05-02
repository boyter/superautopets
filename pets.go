package main

import (
	"math/rand"
)

const (
	Ant      string = "ant"
	Fish            = "fish"
	Bever           = "bever"
	Otter           = "otter"
	Sloth           = "sloth"
	Cricket         = "cricket"
	Duck            = "duck"
	Horse           = "horse"
	Mosquito        = "mosquito"
	Pig             = "pig"
)

type Pet struct {
	tier           int
	name           string
	baseAttack     int
	baseHealth     int
	currentAttack  int
	currentHealth  int
	currentLevel   int
	faint          Faint
	levelup        LevelUp
	sell           Sell
	buy            Buy
	friendSummoned FriendSummoned
	battleStart    BattleStart
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

type Buy func(*Pet, []*Pet)

func NothingBuy(pet *Pet, pets []*Pet) {}

type FriendSummoned func(*Pet, []*Pet)

func NothingFriendSummoned(pet *Pet, pets []*Pet) {}

type BattleStart func(pet *Pet, friends []*Pet, foes []*Pet)

func NothingBattleStart(pet *Pet, friends []*Pet, foes []*Pet) {}

func CreatePet(name string) (*Pet, error) {
	switch name {
	case Ant:
		return &Pet{
			tier:           1,
			name:           Ant,
			baseAttack:     2,
			baseHealth:     1,
			currentAttack:  2,
			currentHealth:  1,
			currentLevel:   1,
			faint:          AntFaint,
			levelup:        NothingLevelUp,
			sell:           NothingSell,
			buy:            NothingBuy,
			friendSummoned: NothingFriendSummoned,
			battleStart:    NothingBattleStart,
		}, nil
	case Fish:
		return &Pet{
			tier:           1,
			name:           Fish,
			baseAttack:     2,
			baseHealth:     3,
			currentAttack:  2,
			currentHealth:  3,
			currentLevel:   1,
			faint:          NothingFaint,
			levelup:        FishLevelUp,
			sell:           NothingSell,
			buy:            NothingBuy,
			friendSummoned: NothingFriendSummoned,
			battleStart:    NothingBattleStart,
		}, nil
	case Bever:
		return &Pet{
			tier:           1,
			name:           Bever,
			baseAttack:     2,
			baseHealth:     2,
			currentAttack:  2,
			currentHealth:  2,
			currentLevel:   1,
			faint:          NothingFaint,
			levelup:        NothingLevelUp,
			sell:           BeverSell,
			buy:            NothingBuy,
			friendSummoned: NothingFriendSummoned,
			battleStart:    NothingBattleStart,
		}, nil
	case Otter:
		return &Pet{
			tier:           1,
			name:           Otter,
			baseAttack:     1,
			baseHealth:     2,
			currentAttack:  1,
			currentHealth:  2,
			currentLevel:   1,
			faint:          NothingFaint,
			levelup:        NothingLevelUp,
			sell:           NothingSell,
			buy:            OtterBuy,
			friendSummoned: NothingFriendSummoned,
			battleStart:    NothingBattleStart,
		}, nil
	case Sloth:
		return &Pet{
			tier:           1,
			name:           Sloth,
			baseAttack:     1,
			baseHealth:     1,
			currentAttack:  1,
			currentHealth:  1,
			currentLevel:   1,
			faint:          NothingFaint,
			levelup:        NothingLevelUp,
			sell:           NothingSell,
			buy:            NothingBuy,
			friendSummoned: NothingFriendSummoned,
			battleStart:    NothingBattleStart,
		}, nil
	case Cricket:
		return &Pet{
			tier:           1,
			name:           Cricket,
			baseAttack:     1,
			baseHealth:     2,
			currentAttack:  1,
			currentHealth:  2,
			currentLevel:   1,
			faint:          CricketFaint,
			levelup:        NothingLevelUp,
			sell:           NothingSell,
			buy:            NothingBuy,
			friendSummoned: NothingFriendSummoned,
			battleStart:    NothingBattleStart,
		}, nil
	case Duck:
		return &Pet{
			tier:           1,
			name:           Duck,
			baseAttack:     1,
			baseHealth:     3,
			currentAttack:  1,
			currentHealth:  3,
			currentLevel:   1,
			faint:          NothingFaint,
			levelup:        NothingLevelUp,
			sell:           DuckSell,
			buy:            NothingBuy,
			friendSummoned: NothingFriendSummoned,
			battleStart:    NothingBattleStart,
		}, nil
	case Horse:
		return &Pet{
			tier:           1,
			name:           Horse,
			baseAttack:     2,
			baseHealth:     1,
			currentAttack:  2,
			currentHealth:  1,
			currentLevel:   1,
			faint:          NothingFaint,
			levelup:        NothingLevelUp,
			sell:           NothingSell,
			buy:            NothingBuy,
			friendSummoned: HorseFriendSummoned,
			battleStart:    NothingBattleStart,
		}, nil
	case Mosquito:
		return &Pet{
			tier:           1,
			name:           Mosquito,
			baseAttack:     2,
			baseHealth:     2,
			currentAttack:  2,
			currentHealth:  2,
			currentLevel:   2,
			faint:          NothingFaint,
			levelup:        NothingLevelUp,
			sell:           NothingSell,
			buy:            NothingBuy,
			friendSummoned: HorseFriendSummoned,
			battleStart:    MosquitoBattleStart,
		}, nil
	case Pig:
		return &Pet{
			tier:           1,
			name:           Pig,
			baseAttack:     3,
			baseHealth:     1,
			currentAttack:  3,
			currentHealth:  1,
			currentLevel:   2,
			faint:          NothingFaint,
			levelup:        NothingLevelUp,
			sell:           PigSell,
			buy:            NothingBuy,
			friendSummoned: HorseFriendSummoned,
			battleStart:    MosquitoBattleStart,
		}, nil
	}

	return &Pet{}, nil
}

func AntFaint(pet *Pet, pets []*Pet) {
	if !pet.Fainted() { // have not fainted so return
		return
	}

	// Faint: Give a random friend +2/+1

	// find which ones have not fainted, if any
	choices := nonFaintedIndex(pets)

	// if no choices return
	if len(choices) == 0 {
		return
	}

	// pick a random choice to buff
	c := rand.Intn(len(choices))
	t := pets[choices[c]]

	// buff the choice
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

func CricketFaint(pet *Pet, pets []*Pet) {}

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

	// TODO ensure its not the same friend
	t1 := pets[rand.Intn(len(pets))]
	t2 := pets[rand.Intn(len(pets))]

	t1.currentHealth += buff
	t2.currentHealth += buff
}

func DuckSell(pet *Pet, pets []*Pet) {
}

func PigSell(pet *Pet, pets []*Pet) {
}

func OtterBuy(pet *Pet, pets []*Pet) {
}

func HorseFriendSummoned(pet *Pet, pets []*Pet) {
}

func MosquitoBattleStart(pet *Pet, friends []*Pet, foes []*Pet) {
	// Start of battle: Deal 1 damage to a random enemy
	choices := nonFaintedIndex(foes)

	if len(choices) == 0 {
		return
	}

	// random to attack
	c := rand.Intn(len(choices))
	foes[c].TakeDamage(1)

	// now call its fainted function just in case
	foes[c].faint(foes[c], foes)
}
