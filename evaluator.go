package genetics

// A fitness value has to be non-negative.
type Fitness float64

// A evaluator calcualtes the fitness of one
// or many subjects. 
type Evaluator interface {
	Evaluate(g Subject) Fitness
}
