package main

type ShopState struct {
	round int
	pets  []Pet
	items []Item
	gold  int
}

func CreateShop() ShopState {
	var pets []Pet

	pet, _ := RandomPet(1)
	pets = append(pets, pet)
	pet, _ = RandomPet(1)
	pets = append(pets, pet)
	pet, _ = RandomPet(1)
	pets = append(pets, pet)

	return ShopState{
		round: 1,
		pets:  pets,
		items: nil,
		gold:  10,
	}
}
