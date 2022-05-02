// https://superauto.pet

package main

import "fmt"

func main() {
	var leftPets []*Pet
	var rightPets []*Pet

	pet, _ := CreatePet("ant")
	leftPets = append(leftPets, pet)
	pet, _ = CreatePet("ant")
	leftPets = append(leftPets, pet)

	pet, _ = CreatePet("fish")
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

		// now we fight! both at the same time more or less
		// but keep in mind we need to apply modifiers at some point
		for !l.Fainted() && !r.Fainted() {
			l.TakeDamage(r.CurrentAttack())
			r.TakeDamage(l.CurrentAttack())
		}

		if l.Fainted() {
			fmt.Println(l.name, "fainted")
		}
		if r.Fainted() {
			fmt.Println(r.name, "fainted")
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
