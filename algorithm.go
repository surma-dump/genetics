// The genetics package is a framework for implementing
// genetic algorithms.
// Genome are modeled by slices of float64s, each of which
// has a value in the range of [0.0;1.0)
// The interpretation of these genomes and everything else
// can be modified by implementing the interfaces manually.
package genetics


// All the single parts stacked together
// result in a runnable genetic algorithm
type Algorithm struct {
	PopulationSize int
	GenomeLength   int
	population     Population
	Initializer    Initializer
	Evaluator      Evaluator
	Selector       Selector
	Breeder        Breeder
	Mutator        Mutator
}

// Starts a new population of size PopulationSize
// and inits every subject using the Initializer
func (a *Algorithm) Init() {
	a.population = make(Population, a.PopulationSize)
	for i := range a.population {
		a.population[i] = Subject{Genome: a.Initializer.NewGenome(a.GenomeLength)}
	}
	a.evaluateAll()
}

// Creates the next generation
// At this point, the populations has to be evaluated already
func (a *Algorithm) Next() {
	newpop := a.Breeder.Breed(a.population, a.Selector)
	for i := range newpop {
		newpop[i] = a.Mutator.Mutate(newpop[i])
	}
	a.population = newpop
}

// Returns the current population
func (a *Algorithm) Population() Population {
	return a.population
}

func (a *Algorithm) evaluateAll() {
	for i := range a.population {
		a.Evaluator.Evaluate(a.population[i])
	}
}
