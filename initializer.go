package genetics

import (
	"time"
	"rand"
)

// A initializer creates genomes from scratch
type Initializer interface {
	NewGenome(len int) Genome
}

// RandomInitializer returns genomes which
// contain purely random genes.
type RandomInitializer struct{}

func NewRandomInitializer() Initializer {
	return new(RandomInitializer)
}

func (init *RandomInitializer) NewGenome(len int) (g Genome) {
	rand.Seed(time.Nanoseconds())
	g = make(Genome, len)
	for i := range g {
		g[i] = rand.Float64()
	}
	return
}
