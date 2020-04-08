package main

import (
	"math/rand"
	"time"
)

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
