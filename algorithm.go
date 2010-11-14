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
	Population   *Population
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
	a.Population = NewPopulation(a.PopulationSize())
	for i := range a.Population.Subjects {
		a.Population.Subjects[i] = &Subject{Genome: a.Initializer.NewGenome(a.GenomeLength)}
		a.Evaluator.Evaluate(a.Population.Subjects[i])
		a.Population.FitnessSum += a.Population.Subjects[i].Fitness
	}
	return
}

// Creates the next generation
func (a *Algorithm) CreateNextGenerationParallel() *Population {
	newpop := a.Breeder.Breed(a.Population, a.Selector)
	c := make(chan int)
	defer close(c)
	for i := range newpop.Subjects {
		go func() {
			a.Mutator.Mutate(newpop.Subjects[i])
			a.Evaluator.Evaluate(newpop.Subjects[i])
			newpop.FitnessSum += newpop.Subjects[i].Fitness
			c <- 1
		}()
	}
	for count := 0; count < a.PopulationSize(); {
		count += <-c
	}
	return newpop
}

// Creates the next generation
func (a *Algorithm) CreateNextGeneration() *Population {
	newpop := a.Breeder.Breed(a.Population, a.Selector)
	for i := range newpop.Subjects {
		a.Mutator.Mutate(newpop.Subjects[i])
		a.Evaluator.Evaluate(newpop.Subjects[i])
		newpop.FitnessSum += newpop.Subjects[i].Fitness
	}
	return newpop
}

// Replaces the old population with the new one
func (a *Algorithm) Evolve() {
	a.Population = a.CreateNextGeneration()
}

// Replaces the old population with the new one
func (a *Algorithm) EvolveParallel() {
	a.Population = a.CreateNextGenerationParallel()
}
