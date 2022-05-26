package main

import (
	"github.com/yaricom/goNEAT/v2/neat/genetics"
	"github.com/yaricom/goNEAT/v2/neat/network"
	"math/rand"
	"strings"
)

const (
	DecisionBuyPet = iota
	DecisionSellPet
	DecisionBuyItem
	DecisionRoll
)

type Decision int

type BotDecision struct {
	Decision Decision
	Value    int // if decision is buy or sell, which one should we do that for?
}

// Bot is an interface that must be implemented to perform actions against the shop
type Bot interface {
	Decision(state ShopState) BotDecision
}

// RandomBot is a bot that just does random actions as something to test things with
type RandomBot struct{}

func (r RandomBot) Decision(state ShopState) BotDecision {
	var d Decision
	switch rand.Intn(4) {
	case 0:
		d = DecisionBuyPet
	case 1:
		d = DecisionSellPet
	case 2:
		d = DecisionBuyItem
	case 3:
		d = DecisionRoll
	default:
		panic("ARGH!")
	}

	return BotDecision{
		Decision: d,
		Value:    rand.Intn(len(state.pets)),
	}
}

type NeuralNetworkBot struct {
	net *network.Network
}

func (r NeuralNetworkBot) Decision(state ShopState) BotDecision {
	//_ = r.net.LoadSensors([]float64{
	//	float64(state.aPrevious),
	//	float64(state.bPrevious),
	//})
	//
	//_, _ = r.net.Activate()
	//outputs := r.net.ReadOutputs()
	//
	//// based on what the network says play!
	//decision := Cooperate
	//if outputs[0] > 0.5 {
	//	decision = Defect
	//}
	//
	//return decision

	return BotDecision{
		Decision: 0,
		Value:    0,
	}
}

func getGenome(genomeStr string) *network.Network {
	genome, _ := genetics.ReadGenome(strings.NewReader(genomeStr), 1)

	net, _ := genome.Genesis(1)

	return net
}
