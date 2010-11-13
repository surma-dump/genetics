package genetics

type Genome []float64

type Genomes []Genome

func (g *Genome) Len() int {
	return len(*g)
}

// Creates a deep copy of the genome
func (g *Genome) Copy() (ng Genome) {
	ng = make(Genome, (*g).Len())
	copy(ng, *g)
	return
}
