package main

import (
	"fmt"
	"math/rand"
)

// returns first non fainted pet, if the return is nil that means
// all pets have fainted and possibly have lost the game
func firstNonFainted(pets *[]Pet) *Pet {
	for i := 0; i < len(*pets); i++ {
		if (*pets)[i].baseHealth > 0 {
			return &(*pets)[i]
		}
	}

	return nil
}

func printTeam(pets *[]Pet) {
	for i, p := range *pets {
		fmt.Println(fmt.Sprintf("pos:%v n:%v a:%v h:%v l:%v", i, p.name, p.baseAttack, p.baseHealth, p.currentLevel))
	}
}

func nonFaintedIndex(pets []Pet) []int {
	choices := []int{}
	for i := 0; i < len(pets); i++ {
		p := pets[i].Fainted()
		if !p {
			choices = append(choices, i)
		}
	}
	return choices
}

// creates a random team of pets
func randomTeam() []Pet {
	petCount := rand.Intn(5) + 1

	choices := []string{
		Ant,
		Fish,
		Bever,
		Otter,
		Cricket,
		Duck,
		Horse,
		Mosquito,
		Pig,
	}

	var pets []Pet
	for i := 0; i < petCount; i++ {
		var pet Pet
		// stupidly low chance to get sloth
		if rand.Intn(10_000) == 1 {
			pet, _ = CreatePet(Sloth)
		} else {
			pet, _ = CreatePet(choices[rand.Intn(len(choices))])
		}

		pets = append(pets, pet)
	}

	return pets
}

func Battle(state BattleState) {
	if log {
		printTeam(state.friends)
		printTeam(state.foes)
		fmt.Println()
	}

	hadEffect := false
	// start by calling the pre abilities of each pet
	for i := 0; i < len(*state.friends); i++ {
		t := (*state.friends)[i].battleStart(&BattleState{
			pet:     &(*state.friends)[i],
			friends: state.friends,
			foes:    state.foes,
		})

		if t {
			hadEffect = true
		}
	}
	for i := 0; i < len(*state.foes); i++ {
		t := (*state.foes)[i].battleStart(&BattleState{
			pet:     &(*state.foes)[i],
			friends: state.foes,
			foes:    state.friends,
		})

		if t {
			hadEffect = true
		}
	}

	for hadEffect {
		hadEffect = false
		for i := 0; i < len(*state.friends); i++ {
			t := (*state.friends)[i].faint(&BattleState{
				pet:     &(*state.friends)[i],
				friends: state.friends,
				foes:    state.foes,
			})

			if t {
				hadEffect = true
			}
		}

		for i := 0; i < len(*state.foes); i++ {
			t := (*state.foes)[i].faint(&BattleState{
				pet:     &(*state.foes)[i],
				friends: state.foes,
				foes:    state.friends,
			})

			if t {
				hadEffect = true
			}
		}
	}

	for {
		if log {
			printTeam(state.friends)
			printTeam(state.foes)
			fmt.Println()
		}

		// get the first non fainted pet of each, then fight them
		l := firstNonFainted(state.friends)
		r := firstNonFainted(state.foes)

		if l == nil || r == nil {
			// TODO handle draws and who actually won
			if l == nil && r != nil {
				if log {
					fmt.Println("friends lost")
				}
				return
			}
			if l != nil && r == nil {
				if log {
					fmt.Println("foes lost")
				}
				return
			}

			if log {
				fmt.Println("draw")
			}
			return
		}

		// now we fight! both at the same time more or less
		// but keep in mind we need to apply modifiers at some point
		for !l.Fainted() && !r.Fainted() {
			l.TakeDamage(r.CurrentAttack())
			r.TakeDamage(l.CurrentAttack())
		}

		if l.Fainted() {
			l.faint(&BattleState{
				pet:     l,
				friends: state.friends,
				foes:    state.foes,
			})
		}

		if r.Fainted() {
			r.faint(&BattleState{
				pet:     r,
				friends: state.foes,
				foes:    state.friends,
			})
		}
	}
}
