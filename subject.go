package genetics

// Represents a single subject of a
// population
type Subject struct {
	Genome   Genome
	Fitness  Fitness
}

func (s *Subject) GenomeLength() int {
	return len(s.Genome)
}

type Population []Subject

// Adjusts all fitness values so their sum is 1.0
func (p *Population) NormalizeFitness() {
	fitnessSum := Fitness(0.0)
	for i := range *p {
		fitnessSum += (*p)[i].Fitness
	}

	for i := range *p {
		(*p)[i].Fitness /= fitnessSum
	}
}

// The Less() function of sort.Interface
func (p *Population) Less(i, j int) bool {
	return (*p)[i].Fitness < (*p)[j].Fitness
}

// The Len() function of sort.Interface
func (p *Population) Len() int {
	return len(*p)
}

// The Swap() function of sort.Interface
func (p *Population) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}


