package genetics

import (
	"rand"
)

// A mutator has the chance of mutating the
// subjects of a new generation
type Mutator interface {
	Mutate(s Subject) Subject
}

// The NopMutator does nothing
type NopMutator struct{}

func NewNopMutator() Mutator {
	return new(NopMutator)
}

func (m *NopMutator) Mutate(s Subject) Subject {
	return s
}


// Resets each value in the genome
// with a fixed probability
type RandomMutator struct {
	prob float64
}

func NewRandomMutator(prob float64) Mutator {
	m := new(RandomMutator)
	m.prob = prob
	return m
}

func (m *RandomMutator) Mutate(s Subject) Subject {
	for i := range s.Genome {
		v := rand.Float64()
		if v < m.prob {
			s.Genome[i] = rand.Float64()
		}
	}
	return s
}
