package genetics

import (
	"rand"
	"sort"
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
	population.NormalizeFitness()
	sort.Sort(&population)
	for i := 0; i < num; i++ {
		val := rand.Float64()
		p[i] = getSubjectByProbability(val, population)
	}
	return
}

func getSubjectByProbability(val float64, population Population) Subject {
	for i := 1; i < len(population); i++ {
		if Fitness(val) < population[i].Fitness {
			return population[i-1]
		}
	}
	return population[len(population)-1]
}
