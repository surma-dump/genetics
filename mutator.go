package genetics

// A mutator has the chance of mutating the
// subjects of a new generation
type Mutator interface {
	Mutate(s Subject) Subject
}

// The NopMutator does nothing
type NopMutator struct{}

func (n *NopMutator) Mutate(s Subject) Subject {
	return s
}
