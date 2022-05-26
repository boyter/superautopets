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

	// if we were doing shop we would want to clone
	// the pets here

	Battle(BattleState{
		friends: &leftPets,
		foes:    &rightPets,
	})

	for i := 0; i < 10; i++ {
		fmt.Println("")
		leftPets = randomTeam()
		rightPets = randomTeam()
		Battle(BattleState{
			friends: &leftPets,
			foes:    &rightPets,
		})
	}

	fmt.Println(CreateShop(1))

	// ok given a shop what states do we have that the network needs...

	// pet for sale 3 * 5 (15 inputs)
	//	type
	//	attack
	//	health
	//
	// item for sale (3 inputs)
	// 	type
	//
	// pets in team 3 * 5 (15 inputs)
	//	type
	//	attack
	//	health
	//
	// gold (1 input)

	// actions we can take

	// decision (what are we doing? 1 output) sell, buy item, buy pet, roll
	//
	// sell pet in team (1 output divided by 5 to determine who should we sell)
	//   1-5
	// buy item and apply to pet in team (or all of them) 1 output to determine who we should sell
	//   1-5
	// buy pet from shop and put on another pet (if same) or in space on team (if space) (2 outputs, who to buy and where to put)
	//   1-5
	// roll

	fmt.Println("////////////////////////////////////////////////////////")

	shop := CreateShop(1)

	rBot := RandomBot{}
	for i := 0; i < 100; i++ {
		fmt.Println(rBot.Decision(shop))
		fmt.Println(shop.BuyPet(0))
	}

	for i := 0; i < 100; i++ {
		shop.BuyPet(0)
	}

}
