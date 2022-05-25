package main

import (
	"fmt"
	"github.com/yaricom/goNEAT/v2/experiment"
	"github.com/yaricom/goNEAT/v2/neat"
	"github.com/yaricom/goNEAT/v2/neat/genetics"
	"os"
)

type SuperAutoPetsGenerationEvaluator struct{}

func (ex SuperAutoPetsGenerationEvaluator) GenerationEvaluate(
	pop *genetics.Population,
	epoch *experiment.Generation,
	context *neat.Options,
) (err error) {
	// Calculate the fitness of all organisms in the population
	for _, org := range pop.Organisms {
		res, err := ex.orgEvaluate(org)
		if err != nil {
			return err
		}

		if res && (epoch.Best == nil || org.Fitness > epoch.Best.Fitness) {
			epoch.Solved = true
			epoch.WinnerNodes = len(org.Genotype.Nodes)
			epoch.WinnerGenes = org.Genotype.Extrons()
			epoch.WinnerEvals = context.PopSize*epoch.Id + org.Genotype.Id
			epoch.Best = org
			if epoch.WinnerNodes == 5 {
				neat.InfoLog(fmt.Sprintf("Dumped optimal genome\n"))
			}
		}
	}

	epoch.FillPopulationStatistics(pop)

	// if we have a best candidate now save it
	if epoch.Best != nil {
		//bestOrgPath := fmt.Sprintf("best_%v_%04d", epoch.TrialId, epoch.Id)
		bestOrgPath := "best"
		file, err := os.Create(bestOrgPath)
		if err != nil {
			neat.ErrorLog(fmt.Sprintf("Failed to dump population, reason: %s\n", err))
		} else {
			org := epoch.Best
			_, _ = fmt.Fprintf(file, "/* Organism #%d Fitness: %.3f Error: %.3f */\n",
				org.Genotype.Id, org.Fitness, org.Error)
			_ = org.Genotype.Write(file)
		}
	}

	return nil
}

func (e *SuperAutoPetsGenerationEvaluator) orgEvaluate(organism *genetics.Organism) (bool, error) {
	//game := CreateGame()
	//b := CooperateBot{}
	//
	//netDepth, _ := organism.Phenotype.MaxActivationDepthFast(0) // The max depth of the network to be activated
	//
	//for !game.GameOver() {
	//	// get the game state
	//	state := game.State()
	//
	//	// set up our input
	//	err := organism.Phenotype.LoadSensors([]float64{
	//		float64(state.aPrevious),
	//		float64(state.bPrevious),
	//	})
	//	if err != nil {
	//		return false, err
	//	}
	//
	//	// run the network
	//	_, err = organism.Phenotype.ForwardSteps(netDepth)
	//	if err != nil {
	//		return false, err
	//	}
	//
	//	// based on what the network says play!
	//	decision := Cooperate
	//	if organism.Phenotype.Outputs[0].Activation > 0.5 {
	//		decision = Defect
	//	}
	//
	//	game.Play(gameDecision{
	//		aChoice: decision,
	//		bChoice: b.Decision(state),
	//	})
	//}

	//organism.Fitness = float64(game.AScore)
	//organism.Error = 0.0
	//organism.IsWinner = game.AScore > 20

	return organism.IsWinner, nil
}
