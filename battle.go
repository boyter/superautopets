package main

import (
	"fmt"
	"math/rand"
)

// returns first non fainted pet, if the return is nil that means
// all pets have fainted and possibly have lost the game
func firstNonFainted(pets *[]Pet) *Pet {
	for i := 0; i < len(*pets); i++ {
		if (*pets)[i].currentHealth > 0 {
			return &(*pets)[i]
		}
	}

	return nil
}

func printTeam(pets *[]Pet) {
	for i, p := range *pets {
		fmt.Println(fmt.Sprintf("pos:%v n:%v a:%v h:%v l:%v", i, p.name, p.currentAttack, p.currentHealth, p.currentLevel))
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
		Sloth,
		Cricket,
		Duck,
		Horse,
		Mosquito,
		Pig,
	}

	var pets []Pet
	for i := 0; i < petCount; i++ {
		pet, _ := CreatePet(choices[rand.Intn(len(choices))])
		pets = append(pets, pet)
	}

	return pets
}

func Battle(state BattleMutableState) {
	if log {
		printTeam(state.friends)
		printTeam(state.foes)
		fmt.Println()
	}

	hadEffect := false
	// start by calling the pre abilities of each pet
	for i := 0; i < len(*state.friends); i++ {
		t := (*state.friends)[i].battleStart(&BattleMutableState{
			pet:     &(*state.friends)[i],
			friends: state.friends,
			foes:    state.foes,
		})

		if t {
			hadEffect = true
		}
	}
	for i := 0; i < len(*state.foes); i++ {
		t := (*state.foes)[i].battleStart(&BattleMutableState{
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
			t := (*state.friends)[i].faint(&BattleMutableState{
				pet:     &(*state.friends)[i],
				friends: state.friends,
				foes:    state.foes,
			})

			if t {
				hadEffect = true
			}
		}

		for i := 0; i < len(*state.foes); i++ {
			t := (*state.foes)[i].faint(&BattleMutableState{
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
			l.faint(&BattleMutableState{
				pet:     l,
				friends: state.friends,
				foes:    state.foes,
			})
		}

		if r.Fainted() {
			r.faint(&BattleMutableState{
				pet:     r,
				friends: state.foes,
				foes:    state.friends,
			})
		}
	}
}
