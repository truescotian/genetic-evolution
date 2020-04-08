package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	iterations   int     = 50
	N            int     = 100
	target       string  = "to be or not to be"
	mutationRate float64 = 0.01
)

var targetChar []byte

func randomASCIIByte() byte {
	rand.Seed(time.Now().UnixNano())
	c := 32 + rand.Intn(91)
	return byte(c)
}

func main() {
	population := make([]DNA, N)
	targetChar = []byte(target)
	for i := 0; i < len(population); i++ {
		population[i] = NewDNA()
	}

	found := false
	for z := 0; z < iterations; z++ { // iterations
		// Selection
		for i := 0; i < len(population); i++ {
			population[i].setFitness()
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

}
