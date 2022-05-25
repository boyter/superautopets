package main

import (
	"errors"
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

// BattleState represents the state of a battle, where we can change anything
// without affect
type BattleState struct {
	pet     *Pet
	friends *[]Pet
	foes    *[]Pet
}

// Pet
// https://www.slythergames.com/2021/11/17/super-auto-pets-how-to-level-up/
// Level 1 to 2 is three pets total. This includes the starting one and then two more.
// Level 2 to 3 is six pets total. This includes three at level 2 and then three more on top (assuming these ones are level 1).
// Level 3 is the max so you can no longer gain experience.
//
// When you combine two of the same animal, the game will add +1 to the higher of both the health and attack stat.
// For example, if you combine a 3/1 Pig with a 4/2 Pig, the result will be a 5/3 Pig.
// Experience points combine linearly, meaning you just add the total number of XP between the two pets.
// If the sum of this experience triggers a level-up, you’ll get your +1 Tier pet bonus in the shop!
//
// If both pets have items, you will lose the item from the animal being dragged over.
// Put differently, the animal that receives an upgrade will keep its food.
type Pet struct {
	tier           int
	name           string
	baseAttack     int
	baseHealth     int
	currentLevel   int
	experience     int
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
	return p.baseAttack
}

func (p *Pet) TakeDamage(damage int) {
	// TODO must include modifiers here
	p.baseHealth -= damage
	if p.baseHealth < 0 {
		p.baseHealth = 0
	}
}

func (p *Pet) Fainted() bool {
	return p.baseHealth <= 0
}

type Faint func(state *BattleState) bool
type FaintBuy func(state *BattleState)

func NothingFaint(state *BattleState) bool { return false }

type LevelUp func(state *BattleState)

func NothingLevelUp(state *BattleState) {}

type Sell func(state *BattleState)

func NothingSell(state *BattleState) {}

type Buy func(state *BattleState)

func NothingBuy(state *BattleState) {}

type FriendSummoned func(state *BattleState)

func NothingFriendSummoned(state *BattleState) {}

type BattleStart func(state *BattleState) bool

func NothingBattleStart(state *BattleState) bool { return false }

func RandomPet(level int) (Pet, error) {
	choices := []string{Ant, Fish, Bever, Otter, Cricket, Duck, Horse, Mosquito, Pig}

	rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})

	return CreatePet(choices[0])
}

// ClonePet is used when we start a battle by cloning
// each pet so we can mutate its state in the battle without affecting
// the team, since actions in the battle do not apply after it
func ClonePet(p Pet) Pet {
	return Pet{
		tier:           p.tier,
		name:           p.name,
		baseAttack:     p.baseAttack,
		baseHealth:     p.baseHealth,
		currentLevel:   p.currentLevel,
		experience:     p.experience,
		faint:          p.faint,
		faintBuy:       p.faintBuy,
		levelup:        p.levelup,
		sell:           p.sell,
		buy:            p.buy,
		friendSummoned: p.friendSummoned,
		battleStart:    p.battleStart,
	}
}

func CreatePet(name string) (Pet, error) {
	switch name {
	case Ant:
		return Pet{
			tier:           1,
			name:           Ant,
			baseAttack:     2,
			baseHealth:     1,
			currentLevel:   1,
			experience:     0,
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
			baseHealth:     2,
			currentLevel:   1,
			experience:     0,
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
			baseAttack:     3,
			baseHealth:     2,
			currentLevel:   1,
			experience:     0,
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
			currentLevel:   1,
			experience:     0,
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
			currentLevel:   1,
			experience:     0,
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
			currentLevel:   1,
			experience:     0,
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
			baseAttack:     2,
			baseHealth:     3,
			currentLevel:   1,
			experience:     0,
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
			currentLevel:   1,
			experience:     0,
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
			currentLevel:   1,
			experience:     0,
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
			baseAttack:     4,
			baseHealth:     1,
			currentLevel:   1,
			experience:     0,
			faint:          NothingFaint,
			levelup:        NothingLevelUp,
			sell:           PigSell,
			buy:            NothingBuy,
			friendSummoned: HorseFriendSummoned,
			battleStart:    NothingBattleStart,
		}, nil
	}

	return Pet{}, errors.New("unknown pet")
}

func AntFaint(state *BattleState) bool {
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
		t.baseAttack += 2
		t.baseHealth += 1
	case 2:
		t.baseAttack += 4
		t.baseHealth += 2
	case 3:
		t.baseAttack += 6
		t.baseHealth += 3
	}

	state.pet.faintCalled = true
	return true
}

func CricketFaint(state *BattleState) bool {
	if !state.pet.Fainted() { // have not fainted so return
		return false
	}
	if state.pet.faintCalled { // already fainted
		return false
	}

	// Faint: Summon a 1/1 Cricket

	// update so that this pet becomes a 1/1 zombie cricket
	state.pet.name = ZombieCricket
	state.pet.baseAttack = 1
	state.pet.baseHealth = 1
	state.pet.faint = NothingFaint
	state.pet.faintCalled = true

	return true
}

func FishLevelUp(state *BattleState) {
	// Level-up: Give all friends +1/+1
	buff := 1
	if state.pet.currentLevel > 1 {
		buff = 2
	}

	for i := 0; i < len(*state.friends); i++ {
		(*state.friends)[i].baseHealth += buff
		(*state.friends)[i].baseAttack += buff
	}
}

func BeverSell(state *BattleState) {
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

	t1.baseHealth += buff
	t2.baseHealth += buff
}

func DuckSell(state *BattleState) {}

func PigSell(state *BattleState) {}

func OtterBuy(state *BattleState) {}

func HorseFriendSummoned(state *BattleState) {
}

func MosquitoBattleStart(state *BattleState) bool {
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
	(*state.foes)[choices[0]].faint(&BattleState{
		pet:     &(*state.foes)[choices[0]],
		friends: state.foes,
		foes:    state.friends,
	})

	return true
}
