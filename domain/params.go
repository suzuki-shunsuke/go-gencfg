package domain

type (
	// Params is array of Param which can be sorted by param's name.
	Params []Param
)

// Len is the implementation of the sort.Interface interface.
func (p Params) Len() int {
	return len(p)
}

// Swap is the implementation of the sort.Interface interface.
func (p Params) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less is the implementation of the sort.Interface interface.
func (p Params) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}
