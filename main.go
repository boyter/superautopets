// https://superauto.pet

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var leftPets []*Pet
	var rightPets []*Pet

	pet, _ := CreatePet(Ant)
	leftPets = append(leftPets, pet)

	pet, _ = CreatePet(Duck)
	rightPets = append(rightPets, pet)

	Battle(leftPets, rightPets)

}

func Battle(left []*Pet, right []*Pet) {
	for {
		// get the first non fainted pet of each, then fight them
		l := firstNonFainted(left)
		r := firstNonFainted(right)

		if l == nil || r == nil {
			// TODO handle draws and who actually won
			if l == nil && r != nil {
				fmt.Println("left lost")
				return
			}
			if l != nil && r == nil {
				fmt.Println("right lost")
				return
			}

			fmt.Println("draw")
			return
		}

		printTeam(left)
		printTeam(right)

		// now we fight! both at the same time more or less
		// but keep in mind we need to apply modifiers at some point
		for !l.Fainted() && !r.Fainted() {
			l.TakeDamage(r.CurrentAttack())
			r.TakeDamage(l.CurrentAttack())
		}

		if l.Fainted() {
			l.faint(l, left)
			fmt.Println(l.name, "fainted - left")
		}

		if r.Fainted() {
			r.faint(l, right)
			fmt.Println(r.name, "fainted - right")
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
