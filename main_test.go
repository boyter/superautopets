package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkBattle(b *testing.B) {
	//_ = os.Remove("profile.pprof") // remove so we always get a fresh one
	//f, _ := os.Create("profile.pprof")
	//_ = pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	rand.Seed(time.Now().UnixNano())
	log = false

	leftPets := randomTeam()
	rightPets := randomTeam()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Battle(BattleMutableState{
			friends: &leftPets,
			foes:    &rightPets,
		})
	}
}
