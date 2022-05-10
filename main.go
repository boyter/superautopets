// https://superauto.pet
// https://twoaveragegamers.com/ultimate-guide-to-super-auto-pets-game-mechanics/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var log = true

func main() {
	rand.Seed(time.Now().UnixNano())

	var leftPets []Pet
	var rightPets []Pet

	pet, _ := CreatePet(Mosquito)
	leftPets = append(leftPets, pet)
	pet, _ = CreatePet(Mosquito)
	leftPets = append(leftPets, pet)

	pet, _ = CreatePet(Cricket)
	rightPets = append(rightPets, pet)

	Battle(BattleMutableState{
		friends: &leftPets,
		foes:    &rightPets,
	})

	for i := 0; i < 1000; i++ {
		fmt.Println("")
		leftPets = randomTeam()
		rightPets = randomTeam()
		Battle(BattleMutableState{
			friends: &leftPets,
			foes:    &rightPets,
		})
	}
}
