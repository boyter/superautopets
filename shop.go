package main

// ShopState contains what is possible to buy, and the state
type ShopState struct {
	round int
	pets  []Pet
	items []Item
	gold  int
}

// CreateShop
//
// Turn	Pet Tier
// 1	1
// 3	2
// 5	3
// 7	4
// 9	5
// 11	6
func CreateShop(round int, gold int) ShopState {
	var pets []Pet   // can be at most 5 except if pets combined to get higher tier offered
	var items []Item // can be at most 2

	pet, _ := RandomPet(1)
	pets = append(pets, pet)
	pet, _ = RandomPet(1)
	pets = append(pets, pet)
	pet, _ = RandomPet(1)
	pets = append(pets, pet)

	// TODO apply modifiers if someone has used canned food to level every pet up

	item, _ := RandomItem(1)
	items = append(items, item)

	return ShopState{
		round: round,
		pets:  pets,
		items: items,
		gold:  gold,
	}
}
