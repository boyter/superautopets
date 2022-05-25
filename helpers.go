package main

func removePet(slice []Pet, s int) []Pet {
	return append(slice[:s], slice[s+1:]...)
}

func removeItem(slice []Item, s int) []Item {
	return append(slice[:s], slice[s+1:]...)
}
