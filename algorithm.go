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
	GenomeLength int
	Population   Population
	Initializer  Initializer
	Evaluator    Evaluator
	Selector     Selector
	Breeder      Breeder
	Mutator      Mutator
}

// Returns the size of the population array
func (a *Algorithm) PopulationSize() int {
	return a.Population.Size()
}

// Starts a new population of size PopulationSize
// and inits every subject using the Initializer
func (a *Algorithm) Init() {
	a.Population.Subjects = make([]Subject, a.PopulationSize())
	a.Population.FitnessSum = Fitness(0.0)
	for i := range a.Population.Subjects {
		a.Population.Subjects[i] = Subject{Genome: a.Initializer.NewGenome(a.GenomeLength)}
		a.Population.Subjects[i].Fitness = a.Evaluator.Evaluate(a.Population.Subjects[i])
		a.Population.FitnessSum += a.Population.Subjects[i].Fitness
	}
	return
}

// Creates the next generation
func (a *Algorithm) CreateNextGeneration() Population {
	newpop := a.Breeder.Breed(a.Population, a.Selector)
	newpop.FitnessSum = Fitness(0.0)
	for i := range newpop.Subjects {
		newpop.Subjects[i] = a.Mutator.Mutate(newpop.Subjects[i])
		newpop.Subjects[i].Fitness = a.Evaluator.Evaluate(newpop.Subjects[i])
		newpop.FitnessSum += newpop.Subjects[i].Fitness
	}
	return newpop
}

// Replaces the old population with the new one
func (a *Algorithm) Evolve() {
	a.Population = a.CreateNextGeneration()
}
