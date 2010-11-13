package genetics

import (
	"rand"
)

// Selects some subjects of the population
type Selector interface {
	// Selects num individuals
	Select(num int, population Population) Population
}

// Selects subject by stochastical criteria based on the
// fitness of the subject.
type StochasticSelector struct{}

func NewStochasticSelector() Selector {
	return new(StochasticSelector)
}

func (sel *StochasticSelector) Select(num int, population Population) (p Population) {
	p = make(Population, num)
	p.NormalizeFitness()
	for i := 0; i < num; i++ {
		val := rand.Float64()
		p[i] = getSubjectByPropability(val, population)
	}
	return
}

func getSubjectByPropability(val float64, population Population) Subject {
	for i := 0; i < len(population); i++ {
		if Fitness(val) < population[i].Fitness {
			return population[i-1]
		}
	}
	return population[len(population)-1]
}
