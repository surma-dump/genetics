// The genetics package is a framework for implementing
// genetic algorithms.
// Genome are modeled by slices of float64s, each of which
// has a value in the range of [0.0;1.0)
// The interpretation of these genomes and everything else
// can be modified by implementing the interfaces manually.
package genetics


// All the single parts stacked together
// result in a runnable genetic algorithm
// The population size is soley defined by
// the size of the population slice
type Algorithm struct {
	GenomeLength   int
	Population     Population
	Initializer    Initializer
	Evaluator      Evaluator
	Selector       Selector
	Breeder        Breeder
	Mutator        Mutator
}

// Returns the size of the population array
func (a *Algorithm) PopulationSize() int {
	return len(a.Population)
}

// Starts a new population of size PopulationSize
// and inits every subject using the Initializer
func (a *Algorithm) Init() {
	a.Population = make(Population, a.PopulationSize())
	for i := range a.Population {
		a.Population[i] = Subject{Genome: a.Initializer.NewGenome(a.GenomeLength)}
	}
	a.evaluateAll()
}

// Creates the next generation
// At this point, the populations has to be evaluated already
func (a *Algorithm) Next() {
	newpop := a.Breeder.Breed(a.Population, a.Selector)
	for i := range newpop {
		newpop[i] = a.Mutator.Mutate(newpop[i])
	}
	a.Population = newpop
}

func (a *Algorithm) evaluateAll() {
	for i := range a.Population {
		a.Evaluator.Evaluate(a.Population[i])
	}
}
