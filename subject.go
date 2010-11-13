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

