package genetics

import (
	"rand"
)

// A breeder uses the selector to choose
// parents and make them have dirty animal sex
// to create a whole new generation
type Breeder interface {
	Breed(oldpop *Population, sel Selector) *Population
}

// The TwoParentBreeder uses 2 subjects as parents
// for a new child. It choses a random crossover point
// in the genome and swaps the parts after that point.
type TwoParentBreeder struct{}

func NewTwoParentBreeder() Breeder {
	return new(TwoParentBreeder)
}

func (c *TwoParentBreeder) Breed(oldpop *Population, sel Selector) (newpop *Population) {
	newpop = NewPopulation(oldpop.Size())
	for i := range newpop.Subjects {
		parents := sel.Select(2, oldpop)
		crosspoint := int(rand.Float64() * float64(parents.Subjects[0].GenomeLength()))
		g1, g2 := crossover(parents.Subjects[0].Genome, parents.Subjects[1].Genome, crosspoint)
		newpop.Subjects[i] = &Subject{Genome: selectAChild(g1, g2)}
	}
	return
}

func crossover(g1, g2 Genome, point int) (ng1, ng2 Genome) {
	ng1 = g1.Copy()
	ng2 = g2.Copy()
	copy(ng1[point:], g2[point:])
	copy(ng2[point:], g1[point:])
	return
}

func selectAChild(g1, g2 Genome) Genome {
	if rand.Float64() >= 0.5 {
		return g1
	}
	return g2
}
