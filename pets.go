package main

import (
	"math/rand"
)

const (
	Ant           string = "ant"
	Fish                 = "fish"
	Bever                = "bever"
	Otter                = "otter"
	Sloth                = "sloth"
	Cricket              = "cricket"
	ZombieCricket        = "zombiecricket"
	Duck                 = "duck"
	Horse                = "horse"
	Mosquito             = "mosquito"
	Pig                  = "pig"
)

type GameState struct {
	round int
	pets  []Pet
	gold  int
}

type MutableState struct {
	pet     *Pet
	friends *[]Pet
	foes    *[]Pet
}

type Pet struct {
	tier           int
	name           string
	baseAttack     int
	baseHealth     int
	currentAttack  int
	currentHealth  int
	currentLevel   int
	faint          Faint
	faintCalled    bool
	faintBuy       FaintBuy
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

type Faint func(state *MutableState) bool
type FaintBuy func(state *MutableState)

func NothingFaint(state *MutableState) bool { return false }

type LevelUp func(state *MutableState)

func NothingLevelUp(state *MutableState) {}

type Sell func(state *MutableState)

func NothingSell(state *MutableState) {}

type Buy func(state *MutableState)

func NothingBuy(state *MutableState) {}

type FriendSummoned func(state *MutableState)

func NothingFriendSummoned(state *MutableState) {}

type BattleStart func(state *MutableState) bool

func NothingBattleStart(state *MutableState) bool { return false }

func CreatePet(name string) (Pet, error) {
	switch name {
	case Ant:
		return Pet{
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
		return Pet{
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
		return Pet{
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
		return Pet{
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
		return Pet{
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
		return Pet{
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
		return Pet{
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
		return Pet{
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
		return Pet{
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
		return Pet{
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
			battleStart:    NothingBattleStart,
		}, nil
	}

	return Pet{}, nil
}

func AntFaint(state *MutableState) bool {
	if !state.pet.Fainted() { // have not fainted so return
		return false
	}
	if state.pet.faintCalled { // already fainted
		return false
	}

	// Faint: Give a random friend +2/+1

	// find which ones have not fainted, if any
	choices := nonFaintedIndex(*state.friends)

	// if no choices return
	if len(choices) == 0 {
		return false
	}

	// pick a random choice to buff
	c := rand.Intn(len(choices))
	p := *state.friends
	t := p[choices[c]]

	// buff the choice
	switch state.pet.currentLevel {
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

	state.pet.faintCalled = true
	return true
}

func CricketFaint(state *MutableState) bool {
	if !state.pet.Fainted() { // have not fainted so return
		return false
	}
	if state.pet.faintCalled { // already fainted
		return false
	}

	// Faint: Summon a 1/1 Cricket

	// update so that this pet becomes a 1/1 zombie cricket
	state.pet.name = ZombieCricket
	state.pet.currentAttack = 1
	state.pet.currentHealth = 1
	state.pet.faint = NothingFaint
	state.pet.faintCalled = true

	return true
}

func FishLevelUp(state *MutableState) {
	// Level-up: Give all friends +1/+1
	buff := 1
	if state.pet.currentLevel > 1 {
		buff = 2
	}

	for i := 0; i < len(*state.friends); i++ {
		(*state.friends)[i].currentHealth += buff
		(*state.friends)[i].currentAttack += buff
	}
}

func BeverSell(state *MutableState) {
	// Sell: Give two random friends +1 health
	buff := 1
	switch state.pet.currentLevel {
	case 2:
		buff = 2
	case 3:
		buff = 3
	}

	// TODO ensure its not the same friend
	t1 := (*state.friends)[rand.Intn(len(*state.friends))]
	t2 := (*state.friends)[rand.Intn(len(*state.friends))]

	t1.currentHealth += buff
	t2.currentHealth += buff
}

func DuckSell(state *MutableState) {}

func PigSell(state *MutableState) {}

func OtterBuy(state *MutableState) {}

func HorseFriendSummoned(state *MutableState) {
}

func MosquitoBattleStart(state *MutableState) bool {
	// Start of battle: Deal 1 damage to a random enemy
	choices := nonFaintedIndex(*state.foes)

	if len(choices) == 0 {
		return false
	}

	rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})

	// random to attack
	(*state.foes)[choices[0]].TakeDamage(1)
	// now call its fainted function just in case
	// TODO would it be cleaner to call this after?
	(*state.foes)[choices[0]].faint(&MutableState{
		pet:     &(*state.foes)[choices[0]],
		friends: state.foes,
		foes:    state.friends,
	})

	return true
}
