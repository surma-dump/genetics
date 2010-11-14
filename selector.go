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
	p.Subjects = make([]Subject, num)
	for i := 0; i < num; i++ {
		val := rand.Float64() * float64(population.FitnessSum)
		p.Subjects[i] = getSubjectByProbability(val, population)
	}
	return
}

func getSubjectByProbability(val float64, population Population) Subject {
	sum := Fitness(0)
	for _, subject := range population.Subjects {
		sum += subject.Fitness
		if Fitness(val) < sum {
			return subject
		}
	}
	return population.Subjects[population.Size()-1]
}
