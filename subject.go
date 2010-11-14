package genetics

// Represents a single subject of a
// population
type Subject struct {
	Genome  Genome
	Fitness Fitness
	// Stores abitrary data for others to access
	Data interface{}
}

func (s *Subject) GenomeLength() int {
	return len(s.Genome)
}

type Population struct {
	Subjects   []*Subject
	FitnessSum Fitness
}

func NewPopulation(size int) (p *Population) {
	p = new(Population)
	p.Subjects = make([]*Subject, size)
	p.FitnessSum = Fitness(-1.0)
	return
}

// Adjusts all fitness values so their sum is 1.0
func (p *Population) NormalizeFitness() {
	fitnessSum := Fitness(0.0)
	for i := range (*p).Subjects {
		fitnessSum += (*p).Subjects[i].Fitness
	}

	for i := range (*p).Subjects {
		(*p).Subjects[i].Fitness /= fitnessSum
	}
}

// The Less() function of sort.Interface
func (p *Population) Less(i, j int) bool {
	return (*p).Subjects[i].Fitness < (*p).Subjects[j].Fitness
}

// The Len() function of sort.Interface
func (p *Population) Len() int {
	return len((*p).Subjects)
}

// The Swap() function of sort.Interface
func (p *Population) Swap(i, j int) {
	(*p).Subjects[i], (*p).Subjects[j] = (*p).Subjects[j], (*p).Subjects[i]
}

func (p *Population) Size() int {
	return len((*p).Subjects)
}
