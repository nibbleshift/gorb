// Code generated by ent, DO NOT EDIT.

package ent

// CreateBenchInput represents a mutation input for creating benches.
type CreateBenchInput struct {
	Os             string
	Arch           string
	CPU            string
	Package        string
	Pass           bool
	BenchResultIDs []int
}

// Mutate applies the CreateBenchInput on the BenchMutation builder.
func (i *CreateBenchInput) Mutate(m *BenchMutation) {
	m.SetOs(i.Os)
	m.SetArch(i.Arch)
	m.SetCPU(i.CPU)
	m.SetPackage(i.Package)
	m.SetPass(i.Pass)
	if v := i.BenchResultIDs; len(v) > 0 {
		m.AddBenchResultIDs(v...)
	}
}

// SetInput applies the change-set in the CreateBenchInput on the BenchCreate builder.
func (c *BenchCreate) SetInput(i CreateBenchInput) *BenchCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateBenchResultInput represents a mutation input for creating benchresults.
type CreateBenchResultInput struct {
	Name              string
	N                 int64
	NsPerOp           float64
	AllocedBytesPerOp int64
	AllocsPerOp       int64
	MBPerS            float64
	Measured          int64
	Ord               int64
}

// Mutate applies the CreateBenchResultInput on the BenchResultMutation builder.
func (i *CreateBenchResultInput) Mutate(m *BenchResultMutation) {
	m.SetName(i.Name)
	m.SetN(i.N)
	m.SetNsPerOp(i.NsPerOp)
	m.SetAllocedBytesPerOp(i.AllocedBytesPerOp)
	m.SetAllocsPerOp(i.AllocsPerOp)
	m.SetMBPerS(i.MBPerS)
	m.SetMeasured(i.Measured)
	m.SetOrd(i.Ord)
}

// SetInput applies the change-set in the CreateBenchResultInput on the BenchResultCreate builder.
func (c *BenchResultCreate) SetInput(i CreateBenchResultInput) *BenchResultCreate {
	i.Mutate(c.Mutation())
	return c
}
