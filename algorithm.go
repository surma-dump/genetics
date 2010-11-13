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
	population     []Subject
	Initializer    Initializer
	Evaluator      Evaluator
	Selector       Selector
	Breeder        Breeder
	Mutator        Mutator
}
