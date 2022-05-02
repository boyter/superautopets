package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkBattle(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	log = false

	leftPets := randomTeam()
	rightPets := randomTeam()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Battle(leftPets, rightPets)
	}
}
