package main

import "errors"

// ShopState contains what is possible to buy, and the state
type ShopState struct {
	round int
	pets  []Pet
	items []Item
	gold  int
}

// BuyPet removes pet from the shop and returns it
func (s *ShopState) BuyPet(i int) (Pet, error) {
	if s.gold < 3 {
		return Pet{}, errors.New("not enough gold")
	}

	p := s.pets[i]
	s.pets = removePet(s.pets, i)
	s.gold -= 3
	return p, nil
}

// BuyItem removes item from the shop and returns it
func (s *ShopState) BuyItem(i int) (Item, error) {
	if s.gold < 3 {
		return Item{}, errors.New("not enough gold")
	}

	p := s.items[i]
	s.items = removeItem(s.items, i)
	s.gold -= 3
	return p, nil
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
func CreateShop(round int) ShopState {
	var pets []Pet   // can be at most 5 except if pets combined to get higher tier offered
	var items []Item // can be at most 2

	pet, _ := RandomPet(1)
	pets = append(pets, pet)
	pet, _ = RandomPet(1)
	pets = append(pets, pet)
	pet, _ = RandomPet(1)
	pets = append(pets, pet)

	// TODO apply modifiers if someone has used canned food to level every pet up
	// TODO apply modifiers such as swan which adds gold

	item, _ := RandomItem(1)
	items = append(items, item)

	return ShopState{
		round: round,
		pets:  pets,
		items: items,
		gold:  10,
	}
}
