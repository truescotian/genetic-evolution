package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomASCIIByte() byte {
	rand.Seed(time.Now().UnixNano())
	c := 32 + rand.Intn(91)
	return byte(c)
}

type gene []byte

func NewGenes() (genes gene) {
	genes = make([]byte, 18)
	return genes
}

type DNA struct {
	genes   []byte // 18 bytes
	fitness float64
}

func NewDNA() (dna DNA) {

	dna = DNA{
		genes: NewGenes(),
	}

	for i := 0; i < len(dna.genes); i++ {
		dna.genes[i] = randomASCIIByte()
	}

	return dna
}

// calculates the fitness of the DNA's genes
func (dna *DNA) setFitness() {
	score := 0
	for idx, b := range dna.genes {
		if b == targetChar[idx] { // compare ASCII values
			score++
		}
	}
	dna.fitness = float64(score) / float64(len(target))
}

// crossover Creates a child using random midpoint method of crossover
// which is the first section of genes taken from parent A, and the second
// section from parent B.
func (dna *DNA) crossover(partner DNA) (child DNA) {
	child = NewDNA()

	// get a random index from genes as a midpoint
	midPoint := rand.Intn(len(dna.genes))

	for i := 0; i < len(dna.genes); i++ {
		// before midpoint copy genes from A, after midpoint copy from B
		if i > midPoint {
			child.genes[i] = dna.genes[i]
		} else {
			child.genes[i] = partner.genes[i]
		}
	}

	return child
}

func (dna *DNA) mutate() {
	for i := 0; i < len(dna.genes); i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Float64() < mutationRate {
			dna.genes[i] = randomASCIIByte()
		}
	}
}

// Convert to string -- PHENOTYPE
func (dna *DNA) getPhrase() (phrase string) {
	return string(dna.genes)
}

const (
	N            int     = 100
	target       string  = "to be or not to be"
	mutationRate float64 = 0.01
)

var targetChar []byte

func main() {
	population := make([]DNA, N)
	targetChar = []byte(target)
	maxFitness := 0.00
	var bestFitness DNA
	for i := 0; i < len(population); i++ {
		population[i] = NewDNA()
	}

	found := false
	for z := 0; z < 5000; z++ {
		// Selection
		for i := 0; i < len(population); i++ {
			population[i].setFitness()
			if population[i].fitness > maxFitness {
				maxFitness = population[i].fitness
				bestFitness = population[i]
				fmt.Println("new max fitness: ", population[i].fitness)
				fmt.Println("new best fit phrase: ", population[i].getPhrase())
			}
			fmt.Println("Generation: ", z, "Fitness: ", population[i].fitness, "Phrase: ", population[i].getPhrase())
			if population[i].getPhrase() == target {
				found = true
				break
			}
		}

		if found {
			break
		}

		// Mating pool
		// we want to make it so there is a percent score
		// is proportional to the amount they would get picked
		// from a basket of potential parents.
		// So if 'a' has a 5% of being picked, we want
		// 'a' to be in the pool 5% of the time, which is N=fitness*100
		matingPool := make([]DNA, 0)
		for i := 0; i < len(population); i++ {
			n := int(population[i].fitness * 100)
			for j := 0; j < n; j++ {
				matingPool = append(matingPool, population[i])
			}
		}

		// Reproduction
		max := len(matingPool)
		for i := 0; i < len(population); i++ {
			a := rand.Intn(max)
			b := rand.Intn(max)

			parentA := matingPool[a]
			parentB := matingPool[b]
			child := parentA.crossover(parentB)
			child.mutate()

			// overwrite the population with new child
			population[i] = child
		}

	}

	fmt.Println("best phrase: ", bestFitness.getPhrase())
	fmt.Println("best fitness: ", maxFitness)

}
