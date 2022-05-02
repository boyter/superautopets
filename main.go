// https://superauto.pet

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var log = true

func main() {
	rand.Seed(time.Now().UnixNano())

	var leftPets []*Pet
	var rightPets []*Pet

	pet, _ := CreatePet(Mosquito)
	leftPets = append(leftPets, pet)

	pet, _ = CreatePet(Sloth)
	rightPets = append(rightPets, pet)

	Battle(leftPets, rightPets)

	fmt.Println("")
	leftPets = randomTeam()
	rightPets = randomTeam()
	Battle(leftPets, rightPets)
}

func Battle(left []*Pet, right []*Pet) {
	if log {
		printTeam(left)
		printTeam(right)
	}

	// start by calling the pre abilities of each pet
	for i := 0; i < len(left); i++ {
		left[i].battleStart(left[i], left, right)
	}
	for i := 0; i < len(right); i++ {
		right[i].battleStart(right[i], right, left)
	}

	for {
		if log {
			printTeam(left)
			printTeam(right)
		}

		// get the first non fainted pet of each, then fight them
		l := firstNonFainted(left)
		r := firstNonFainted(right)

		if l == nil || r == nil {
			// TODO handle draws and who actually won
			if l == nil && r != nil {
				if log {
					fmt.Println("left lost")
				}
				return
			}
			if l != nil && r == nil {
				if log {
					fmt.Println("right lost")
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
			l.faint(l, left)
			if log {
				fmt.Println(l.name, "fainted - left")
			}
		}

		if r.Fainted() {
			r.faint(l, right)
			if log {
				fmt.Println(r.name, "fainted - right")
			}
		}
	}
}

// returns first non fainted pet, if the return is nil that means
// all pets have fainted and possibly have lost the game
func firstNonFainted(pets []*Pet) *Pet {
	for i := 0; i < len(pets); i++ {
		if pets[i].currentHealth > 0 {
			return pets[i]
		}
	}

	return nil
}

func printTeam(pets []*Pet) {
	for i, p := range pets {
		fmt.Println(fmt.Sprintf("pos:%v n:%v a:%v h:%v l:%v", i, p.name, p.currentAttack, p.currentHealth, p.currentLevel))
	}
}

func nonFaintedIndex(pets []*Pet) []int {
	choices := []int{}
	for i := 0; i < len(pets); i++ {
		if !pets[i].Fainted() {
			choices = append(choices, i)
		}
	}
	return choices
}

// creates a random team of pets
func randomTeam() []*Pet {
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

	var pets []*Pet
	for i := 0; i < petCount; i++ {
		pet, _ := CreatePet(choices[rand.Intn(len(choices))])
		pets = append(pets, pet)
	}

	return pets
}
