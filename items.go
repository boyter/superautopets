package main

import (
	"errors"
	"math/rand"
)

const (
	Apple string = "Apple"
	Honey        = "Honey"
)

// Item
// https://www.slythergames.com/2021/11/25/super-auto-pets-all-items-and-foods/
type Item struct {
	name string
}

func RandomItem(level int) (Item, error) {
	choices := []string{Apple, Honey}

	rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})

	return CreateItem(choices[0])
}

func CreateItem(name string) (Item, error) {
	switch name {
	case Apple:
		return Item{
			name: Apple,
		}, nil
	case Honey:
		return Item{
			name: Honey,
		}, nil
	}

	return Item{}, errors.New("missing")
}
