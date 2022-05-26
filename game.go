package main

type Game struct {
	Friends []Pet
	Foes    []Pet
}

// a normal game consists of X rounds where you lose a life on each...

// the core part of the game is being presented with shop options, which go until no gold left, then
// fight... so lets create a random bot

func CreateGame(friends []Pet, foes []Pet) Game {
	return Game{
		Friends: friends,
		Foes:    foes,
	}
}
