// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"

	"github.com/nibbleshift/gorb/ent/bench"
	"github.com/nibbleshift/gorb/ent/benchresult"
	"github.com/nibbleshift/gorb/ent/predicate"
)

// BenchWhereInput represents a where input for filtering Bench queries.
type BenchWhereInput struct {
	Predicates []predicate.Bench  `json:"-"`
	Not        *BenchWhereInput   `json:"not,omitempty"`
	Or         []*BenchWhereInput `json:"or,omitempty"`
	And        []*BenchWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "os" field predicates.
	Os             *string  `json:"os,omitempty"`
	OsNEQ          *string  `json:"osNEQ,omitempty"`
	OsIn           []string `json:"osIn,omitempty"`
	OsNotIn        []string `json:"osNotIn,omitempty"`
	OsGT           *string  `json:"osGT,omitempty"`
	OsGTE          *string  `json:"osGTE,omitempty"`
	OsLT           *string  `json:"osLT,omitempty"`
	OsLTE          *string  `json:"osLTE,omitempty"`
	OsContains     *string  `json:"osContains,omitempty"`
	OsHasPrefix    *string  `json:"osHasPrefix,omitempty"`
	OsHasSuffix    *string  `json:"osHasSuffix,omitempty"`
	OsEqualFold    *string  `json:"osEqualFold,omitempty"`
	OsContainsFold *string  `json:"osContainsFold,omitempty"`

	// "arch" field predicates.
	Arch             *string  `json:"arch,omitempty"`
	ArchNEQ          *string  `json:"archNEQ,omitempty"`
	ArchIn           []string `json:"archIn,omitempty"`
	ArchNotIn        []string `json:"archNotIn,omitempty"`
	ArchGT           *string  `json:"archGT,omitempty"`
	ArchGTE          *string  `json:"archGTE,omitempty"`
	ArchLT           *string  `json:"archLT,omitempty"`
	ArchLTE          *string  `json:"archLTE,omitempty"`
	ArchContains     *string  `json:"archContains,omitempty"`
	ArchHasPrefix    *string  `json:"archHasPrefix,omitempty"`
	ArchHasSuffix    *string  `json:"archHasSuffix,omitempty"`
	ArchEqualFold    *string  `json:"archEqualFold,omitempty"`
	ArchContainsFold *string  `json:"archContainsFold,omitempty"`

	// "cpu" field predicates.
	CPU             *string  `json:"cpu,omitempty"`
	CPUNEQ          *string  `json:"cpuNEQ,omitempty"`
	CPUIn           []string `json:"cpuIn,omitempty"`
	CPUNotIn        []string `json:"cpuNotIn,omitempty"`
	CPUGT           *string  `json:"cpuGT,omitempty"`
	CPUGTE          *string  `json:"cpuGTE,omitempty"`
	CPULT           *string  `json:"cpuLT,omitempty"`
	CPULTE          *string  `json:"cpuLTE,omitempty"`
	CPUContains     *string  `json:"cpuContains,omitempty"`
	CPUHasPrefix    *string  `json:"cpuHasPrefix,omitempty"`
	CPUHasSuffix    *string  `json:"cpuHasSuffix,omitempty"`
	CPUEqualFold    *string  `json:"cpuEqualFold,omitempty"`
	CPUContainsFold *string  `json:"cpuContainsFold,omitempty"`

	// "package" field predicates.
	Package             *string  `json:"package,omitempty"`
	PackageNEQ          *string  `json:"packageNEQ,omitempty"`
	PackageIn           []string `json:"packageIn,omitempty"`
	PackageNotIn        []string `json:"packageNotIn,omitempty"`
	PackageGT           *string  `json:"packageGT,omitempty"`
	PackageGTE          *string  `json:"packageGTE,omitempty"`
	PackageLT           *string  `json:"packageLT,omitempty"`
	PackageLTE          *string  `json:"packageLTE,omitempty"`
	PackageContains     *string  `json:"packageContains,omitempty"`
	PackageHasPrefix    *string  `json:"packageHasPrefix,omitempty"`
	PackageHasSuffix    *string  `json:"packageHasSuffix,omitempty"`
	PackageEqualFold    *string  `json:"packageEqualFold,omitempty"`
	PackageContainsFold *string  `json:"packageContainsFold,omitempty"`

	// "pass" field predicates.
	Pass    *bool `json:"pass,omitempty"`
	PassNEQ *bool `json:"passNEQ,omitempty"`

	// "bench_result" edge predicates.
	HasBenchResult     *bool                    `json:"hasBenchResult,omitempty"`
	HasBenchResultWith []*BenchResultWhereInput `json:"hasBenchResultWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *BenchWhereInput) AddPredicates(predicates ...predicate.Bench) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the BenchWhereInput filter on the BenchQuery builder.
func (i *BenchWhereInput) Filter(q *BenchQuery) (*BenchQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyBenchWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyBenchWhereInput is returned in case the BenchWhereInput is empty.
var ErrEmptyBenchWhereInput = errors.New("ent: empty predicate BenchWhereInput")

// P returns a predicate for filtering benches.
// An error is returned if the input is empty or invalid.
func (i *BenchWhereInput) P() (predicate.Bench, error) {
	var predicates []predicate.Bench
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, bench.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Bench, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, bench.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Bench, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, bench.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, bench.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, bench.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, bench.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, bench.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, bench.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, bench.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, bench.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, bench.IDLTE(*i.IDLTE))
	}
	if i.Os != nil {
		predicates = append(predicates, bench.OsEQ(*i.Os))
	}
	if i.OsNEQ != nil {
		predicates = append(predicates, bench.OsNEQ(*i.OsNEQ))
	}
	if len(i.OsIn) > 0 {
		predicates = append(predicates, bench.OsIn(i.OsIn...))
	}
	if len(i.OsNotIn) > 0 {
		predicates = append(predicates, bench.OsNotIn(i.OsNotIn...))
	}
	if i.OsGT != nil {
		predicates = append(predicates, bench.OsGT(*i.OsGT))
	}
	if i.OsGTE != nil {
		predicates = append(predicates, bench.OsGTE(*i.OsGTE))
	}
	if i.OsLT != nil {
		predicates = append(predicates, bench.OsLT(*i.OsLT))
	}
	if i.OsLTE != nil {
		predicates = append(predicates, bench.OsLTE(*i.OsLTE))
	}
	if i.OsContains != nil {
		predicates = append(predicates, bench.OsContains(*i.OsContains))
	}
	if i.OsHasPrefix != nil {
		predicates = append(predicates, bench.OsHasPrefix(*i.OsHasPrefix))
	}
	if i.OsHasSuffix != nil {
		predicates = append(predicates, bench.OsHasSuffix(*i.OsHasSuffix))
	}
	if i.OsEqualFold != nil {
		predicates = append(predicates, bench.OsEqualFold(*i.OsEqualFold))
	}
	if i.OsContainsFold != nil {
		predicates = append(predicates, bench.OsContainsFold(*i.OsContainsFold))
	}
	if i.Arch != nil {
		predicates = append(predicates, bench.ArchEQ(*i.Arch))
	}
	if i.ArchNEQ != nil {
		predicates = append(predicates, bench.ArchNEQ(*i.ArchNEQ))
	}
	if len(i.ArchIn) > 0 {
		predicates = append(predicates, bench.ArchIn(i.ArchIn...))
	}
	if len(i.ArchNotIn) > 0 {
		predicates = append(predicates, bench.ArchNotIn(i.ArchNotIn...))
	}
	if i.ArchGT != nil {
		predicates = append(predicates, bench.ArchGT(*i.ArchGT))
	}
	if i.ArchGTE != nil {
		predicates = append(predicates, bench.ArchGTE(*i.ArchGTE))
	}
	if i.ArchLT != nil {
		predicates = append(predicates, bench.ArchLT(*i.ArchLT))
	}
	if i.ArchLTE != nil {
		predicates = append(predicates, bench.ArchLTE(*i.ArchLTE))
	}
	if i.ArchContains != nil {
		predicates = append(predicates, bench.ArchContains(*i.ArchContains))
	}
	if i.ArchHasPrefix != nil {
		predicates = append(predicates, bench.ArchHasPrefix(*i.ArchHasPrefix))
	}
	if i.ArchHasSuffix != nil {
		predicates = append(predicates, bench.ArchHasSuffix(*i.ArchHasSuffix))
	}
	if i.ArchEqualFold != nil {
		predicates = append(predicates, bench.ArchEqualFold(*i.ArchEqualFold))
	}
	if i.ArchContainsFold != nil {
		predicates = append(predicates, bench.ArchContainsFold(*i.ArchContainsFold))
	}
	if i.CPU != nil {
		predicates = append(predicates, bench.CPUEQ(*i.CPU))
	}
	if i.CPUNEQ != nil {
		predicates = append(predicates, bench.CPUNEQ(*i.CPUNEQ))
	}
	if len(i.CPUIn) > 0 {
		predicates = append(predicates, bench.CPUIn(i.CPUIn...))
	}
	if len(i.CPUNotIn) > 0 {
		predicates = append(predicates, bench.CPUNotIn(i.CPUNotIn...))
	}
	if i.CPUGT != nil {
		predicates = append(predicates, bench.CPUGT(*i.CPUGT))
	}
	if i.CPUGTE != nil {
		predicates = append(predicates, bench.CPUGTE(*i.CPUGTE))
	}
	if i.CPULT != nil {
		predicates = append(predicates, bench.CPULT(*i.CPULT))
	}
	if i.CPULTE != nil {
		predicates = append(predicates, bench.CPULTE(*i.CPULTE))
	}
	if i.CPUContains != nil {
		predicates = append(predicates, bench.CPUContains(*i.CPUContains))
	}
	if i.CPUHasPrefix != nil {
		predicates = append(predicates, bench.CPUHasPrefix(*i.CPUHasPrefix))
	}
	if i.CPUHasSuffix != nil {
		predicates = append(predicates, bench.CPUHasSuffix(*i.CPUHasSuffix))
	}
	if i.CPUEqualFold != nil {
		predicates = append(predicates, bench.CPUEqualFold(*i.CPUEqualFold))
	}
	if i.CPUContainsFold != nil {
		predicates = append(predicates, bench.CPUContainsFold(*i.CPUContainsFold))
	}
	if i.Package != nil {
		predicates = append(predicates, bench.PackageEQ(*i.Package))
	}
	if i.PackageNEQ != nil {
		predicates = append(predicates, bench.PackageNEQ(*i.PackageNEQ))
	}
	if len(i.PackageIn) > 0 {
		predicates = append(predicates, bench.PackageIn(i.PackageIn...))
	}
	if len(i.PackageNotIn) > 0 {
		predicates = append(predicates, bench.PackageNotIn(i.PackageNotIn...))
	}
	if i.PackageGT != nil {
		predicates = append(predicates, bench.PackageGT(*i.PackageGT))
	}
	if i.PackageGTE != nil {
		predicates = append(predicates, bench.PackageGTE(*i.PackageGTE))
	}
	if i.PackageLT != nil {
		predicates = append(predicates, bench.PackageLT(*i.PackageLT))
	}
	if i.PackageLTE != nil {
		predicates = append(predicates, bench.PackageLTE(*i.PackageLTE))
	}
	if i.PackageContains != nil {
		predicates = append(predicates, bench.PackageContains(*i.PackageContains))
	}
	if i.PackageHasPrefix != nil {
		predicates = append(predicates, bench.PackageHasPrefix(*i.PackageHasPrefix))
	}
	if i.PackageHasSuffix != nil {
		predicates = append(predicates, bench.PackageHasSuffix(*i.PackageHasSuffix))
	}
	if i.PackageEqualFold != nil {
		predicates = append(predicates, bench.PackageEqualFold(*i.PackageEqualFold))
	}
	if i.PackageContainsFold != nil {
		predicates = append(predicates, bench.PackageContainsFold(*i.PackageContainsFold))
	}
	if i.Pass != nil {
		predicates = append(predicates, bench.PassEQ(*i.Pass))
	}
	if i.PassNEQ != nil {
		predicates = append(predicates, bench.PassNEQ(*i.PassNEQ))
	}

	if i.HasBenchResult != nil {
		p := bench.HasBenchResult()
		if !*i.HasBenchResult {
			p = bench.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasBenchResultWith) > 0 {
		with := make([]predicate.BenchResult, 0, len(i.HasBenchResultWith))
		for _, w := range i.HasBenchResultWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasBenchResultWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, bench.HasBenchResultWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyBenchWhereInput
	case 1:
		return predicates[0], nil
	default:
		return bench.And(predicates...), nil
	}
}

// BenchResultWhereInput represents a where input for filtering BenchResult queries.
type BenchResultWhereInput struct {
	Predicates []predicate.BenchResult  `json:"-"`
	Not        *BenchResultWhereInput   `json:"not,omitempty"`
	Or         []*BenchResultWhereInput `json:"or,omitempty"`
	And        []*BenchResultWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "n" field predicates.
	N      *int64  `json:"n,omitempty"`
	NNEQ   *int64  `json:"nNEQ,omitempty"`
	NIn    []int64 `json:"nIn,omitempty"`
	NNotIn []int64 `json:"nNotIn,omitempty"`
	NGT    *int64  `json:"nGT,omitempty"`
	NGTE   *int64  `json:"nGTE,omitempty"`
	NLT    *int64  `json:"nLT,omitempty"`
	NLTE   *int64  `json:"nLTE,omitempty"`

	// "ns_per_op" field predicates.
	NsPerOp      *float64  `json:"nsPerOp,omitempty"`
	NsPerOpNEQ   *float64  `json:"nsPerOpNEQ,omitempty"`
	NsPerOpIn    []float64 `json:"nsPerOpIn,omitempty"`
	NsPerOpNotIn []float64 `json:"nsPerOpNotIn,omitempty"`
	NsPerOpGT    *float64  `json:"nsPerOpGT,omitempty"`
	NsPerOpGTE   *float64  `json:"nsPerOpGTE,omitempty"`
	NsPerOpLT    *float64  `json:"nsPerOpLT,omitempty"`
	NsPerOpLTE   *float64  `json:"nsPerOpLTE,omitempty"`

	// "alloced_bytes_per_op" field predicates.
	AllocedBytesPerOp      *uint64  `json:"allocedBytesPerOp,omitempty"`
	AllocedBytesPerOpNEQ   *uint64  `json:"allocedBytesPerOpNEQ,omitempty"`
	AllocedBytesPerOpIn    []uint64 `json:"allocedBytesPerOpIn,omitempty"`
	AllocedBytesPerOpNotIn []uint64 `json:"allocedBytesPerOpNotIn,omitempty"`
	AllocedBytesPerOpGT    *uint64  `json:"allocedBytesPerOpGT,omitempty"`
	AllocedBytesPerOpGTE   *uint64  `json:"allocedBytesPerOpGTE,omitempty"`
	AllocedBytesPerOpLT    *uint64  `json:"allocedBytesPerOpLT,omitempty"`
	AllocedBytesPerOpLTE   *uint64  `json:"allocedBytesPerOpLTE,omitempty"`

	// "allocs_per_op" field predicates.
	AllocsPerOp      *uint64  `json:"allocsPerOp,omitempty"`
	AllocsPerOpNEQ   *uint64  `json:"allocsPerOpNEQ,omitempty"`
	AllocsPerOpIn    []uint64 `json:"allocsPerOpIn,omitempty"`
	AllocsPerOpNotIn []uint64 `json:"allocsPerOpNotIn,omitempty"`
	AllocsPerOpGT    *uint64  `json:"allocsPerOpGT,omitempty"`
	AllocsPerOpGTE   *uint64  `json:"allocsPerOpGTE,omitempty"`
	AllocsPerOpLT    *uint64  `json:"allocsPerOpLT,omitempty"`
	AllocsPerOpLTE   *uint64  `json:"allocsPerOpLTE,omitempty"`

	// "mb_per_s" field predicates.
	MBPerS      *float64  `json:"mbPerS,omitempty"`
	MBPerSNEQ   *float64  `json:"mbPerSNEQ,omitempty"`
	MBPerSIn    []float64 `json:"mbPerSIn,omitempty"`
	MBPerSNotIn []float64 `json:"mbPerSNotIn,omitempty"`
	MBPerSGT    *float64  `json:"mbPerSGT,omitempty"`
	MBPerSGTE   *float64  `json:"mbPerSGTE,omitempty"`
	MBPerSLT    *float64  `json:"mbPerSLT,omitempty"`
	MBPerSLTE   *float64  `json:"mbPerSLTE,omitempty"`

	// "measured" field predicates.
	Measured      *int  `json:"measured,omitempty"`
	MeasuredNEQ   *int  `json:"measuredNEQ,omitempty"`
	MeasuredIn    []int `json:"measuredIn,omitempty"`
	MeasuredNotIn []int `json:"measuredNotIn,omitempty"`
	MeasuredGT    *int  `json:"measuredGT,omitempty"`
	MeasuredGTE   *int  `json:"measuredGTE,omitempty"`
	MeasuredLT    *int  `json:"measuredLT,omitempty"`
	MeasuredLTE   *int  `json:"measuredLTE,omitempty"`

	// "ord" field predicates.
	Ord      *int  `json:"ord,omitempty"`
	OrdNEQ   *int  `json:"ordNEQ,omitempty"`
	OrdIn    []int `json:"ordIn,omitempty"`
	OrdNotIn []int `json:"ordNotIn,omitempty"`
	OrdGT    *int  `json:"ordGT,omitempty"`
	OrdGTE   *int  `json:"ordGTE,omitempty"`
	OrdLT    *int  `json:"ordLT,omitempty"`
	OrdLTE   *int  `json:"ordLTE,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *BenchResultWhereInput) AddPredicates(predicates ...predicate.BenchResult) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the BenchResultWhereInput filter on the BenchResultQuery builder.
func (i *BenchResultWhereInput) Filter(q *BenchResultQuery) (*BenchResultQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyBenchResultWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyBenchResultWhereInput is returned in case the BenchResultWhereInput is empty.
var ErrEmptyBenchResultWhereInput = errors.New("ent: empty predicate BenchResultWhereInput")

// P returns a predicate for filtering benchresults.
// An error is returned if the input is empty or invalid.
func (i *BenchResultWhereInput) P() (predicate.BenchResult, error) {
	var predicates []predicate.BenchResult
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, benchresult.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.BenchResult, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, benchresult.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.BenchResult, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, benchresult.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, benchresult.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, benchresult.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, benchresult.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, benchresult.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, benchresult.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, benchresult.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, benchresult.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, benchresult.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, benchresult.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, benchresult.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, benchresult.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, benchresult.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, benchresult.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, benchresult.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, benchresult.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, benchresult.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, benchresult.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, benchresult.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, benchresult.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, benchresult.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, benchresult.NameContainsFold(*i.NameContainsFold))
	}
	if i.N != nil {
		predicates = append(predicates, benchresult.NEQ(*i.N))
	}
	if i.NNEQ != nil {
		predicates = append(predicates, benchresult.NNEQ(*i.NNEQ))
	}
	if len(i.NIn) > 0 {
		predicates = append(predicates, benchresult.NIn(i.NIn...))
	}
	if len(i.NNotIn) > 0 {
		predicates = append(predicates, benchresult.NNotIn(i.NNotIn...))
	}
	if i.NGT != nil {
		predicates = append(predicates, benchresult.NGT(*i.NGT))
	}
	if i.NGTE != nil {
		predicates = append(predicates, benchresult.NGTE(*i.NGTE))
	}
	if i.NLT != nil {
		predicates = append(predicates, benchresult.NLT(*i.NLT))
	}
	if i.NLTE != nil {
		predicates = append(predicates, benchresult.NLTE(*i.NLTE))
	}
	if i.NsPerOp != nil {
		predicates = append(predicates, benchresult.NsPerOpEQ(*i.NsPerOp))
	}
	if i.NsPerOpNEQ != nil {
		predicates = append(predicates, benchresult.NsPerOpNEQ(*i.NsPerOpNEQ))
	}
	if len(i.NsPerOpIn) > 0 {
		predicates = append(predicates, benchresult.NsPerOpIn(i.NsPerOpIn...))
	}
	if len(i.NsPerOpNotIn) > 0 {
		predicates = append(predicates, benchresult.NsPerOpNotIn(i.NsPerOpNotIn...))
	}
	if i.NsPerOpGT != nil {
		predicates = append(predicates, benchresult.NsPerOpGT(*i.NsPerOpGT))
	}
	if i.NsPerOpGTE != nil {
		predicates = append(predicates, benchresult.NsPerOpGTE(*i.NsPerOpGTE))
	}
	if i.NsPerOpLT != nil {
		predicates = append(predicates, benchresult.NsPerOpLT(*i.NsPerOpLT))
	}
	if i.NsPerOpLTE != nil {
		predicates = append(predicates, benchresult.NsPerOpLTE(*i.NsPerOpLTE))
	}
	if i.AllocedBytesPerOp != nil {
		predicates = append(predicates, benchresult.AllocedBytesPerOpEQ(*i.AllocedBytesPerOp))
	}
	if i.AllocedBytesPerOpNEQ != nil {
		predicates = append(predicates, benchresult.AllocedBytesPerOpNEQ(*i.AllocedBytesPerOpNEQ))
	}
	if len(i.AllocedBytesPerOpIn) > 0 {
		predicates = append(predicates, benchresult.AllocedBytesPerOpIn(i.AllocedBytesPerOpIn...))
	}
	if len(i.AllocedBytesPerOpNotIn) > 0 {
		predicates = append(predicates, benchresult.AllocedBytesPerOpNotIn(i.AllocedBytesPerOpNotIn...))
	}
	if i.AllocedBytesPerOpGT != nil {
		predicates = append(predicates, benchresult.AllocedBytesPerOpGT(*i.AllocedBytesPerOpGT))
	}
	if i.AllocedBytesPerOpGTE != nil {
		predicates = append(predicates, benchresult.AllocedBytesPerOpGTE(*i.AllocedBytesPerOpGTE))
	}
	if i.AllocedBytesPerOpLT != nil {
		predicates = append(predicates, benchresult.AllocedBytesPerOpLT(*i.AllocedBytesPerOpLT))
	}
	if i.AllocedBytesPerOpLTE != nil {
		predicates = append(predicates, benchresult.AllocedBytesPerOpLTE(*i.AllocedBytesPerOpLTE))
	}
	if i.AllocsPerOp != nil {
		predicates = append(predicates, benchresult.AllocsPerOpEQ(*i.AllocsPerOp))
	}
	if i.AllocsPerOpNEQ != nil {
		predicates = append(predicates, benchresult.AllocsPerOpNEQ(*i.AllocsPerOpNEQ))
	}
	if len(i.AllocsPerOpIn) > 0 {
		predicates = append(predicates, benchresult.AllocsPerOpIn(i.AllocsPerOpIn...))
	}
	if len(i.AllocsPerOpNotIn) > 0 {
		predicates = append(predicates, benchresult.AllocsPerOpNotIn(i.AllocsPerOpNotIn...))
	}
	if i.AllocsPerOpGT != nil {
		predicates = append(predicates, benchresult.AllocsPerOpGT(*i.AllocsPerOpGT))
	}
	if i.AllocsPerOpGTE != nil {
		predicates = append(predicates, benchresult.AllocsPerOpGTE(*i.AllocsPerOpGTE))
	}
	if i.AllocsPerOpLT != nil {
		predicates = append(predicates, benchresult.AllocsPerOpLT(*i.AllocsPerOpLT))
	}
	if i.AllocsPerOpLTE != nil {
		predicates = append(predicates, benchresult.AllocsPerOpLTE(*i.AllocsPerOpLTE))
	}
	if i.MBPerS != nil {
		predicates = append(predicates, benchresult.MBPerSEQ(*i.MBPerS))
	}
	if i.MBPerSNEQ != nil {
		predicates = append(predicates, benchresult.MBPerSNEQ(*i.MBPerSNEQ))
	}
	if len(i.MBPerSIn) > 0 {
		predicates = append(predicates, benchresult.MBPerSIn(i.MBPerSIn...))
	}
	if len(i.MBPerSNotIn) > 0 {
		predicates = append(predicates, benchresult.MBPerSNotIn(i.MBPerSNotIn...))
	}
	if i.MBPerSGT != nil {
		predicates = append(predicates, benchresult.MBPerSGT(*i.MBPerSGT))
	}
	if i.MBPerSGTE != nil {
		predicates = append(predicates, benchresult.MBPerSGTE(*i.MBPerSGTE))
	}
	if i.MBPerSLT != nil {
		predicates = append(predicates, benchresult.MBPerSLT(*i.MBPerSLT))
	}
	if i.MBPerSLTE != nil {
		predicates = append(predicates, benchresult.MBPerSLTE(*i.MBPerSLTE))
	}
	if i.Measured != nil {
		predicates = append(predicates, benchresult.MeasuredEQ(*i.Measured))
	}
	if i.MeasuredNEQ != nil {
		predicates = append(predicates, benchresult.MeasuredNEQ(*i.MeasuredNEQ))
	}
	if len(i.MeasuredIn) > 0 {
		predicates = append(predicates, benchresult.MeasuredIn(i.MeasuredIn...))
	}
	if len(i.MeasuredNotIn) > 0 {
		predicates = append(predicates, benchresult.MeasuredNotIn(i.MeasuredNotIn...))
	}
	if i.MeasuredGT != nil {
		predicates = append(predicates, benchresult.MeasuredGT(*i.MeasuredGT))
	}
	if i.MeasuredGTE != nil {
		predicates = append(predicates, benchresult.MeasuredGTE(*i.MeasuredGTE))
	}
	if i.MeasuredLT != nil {
		predicates = append(predicates, benchresult.MeasuredLT(*i.MeasuredLT))
	}
	if i.MeasuredLTE != nil {
		predicates = append(predicates, benchresult.MeasuredLTE(*i.MeasuredLTE))
	}
	if i.Ord != nil {
		predicates = append(predicates, benchresult.OrdEQ(*i.Ord))
	}
	if i.OrdNEQ != nil {
		predicates = append(predicates, benchresult.OrdNEQ(*i.OrdNEQ))
	}
	if len(i.OrdIn) > 0 {
		predicates = append(predicates, benchresult.OrdIn(i.OrdIn...))
	}
	if len(i.OrdNotIn) > 0 {
		predicates = append(predicates, benchresult.OrdNotIn(i.OrdNotIn...))
	}
	if i.OrdGT != nil {
		predicates = append(predicates, benchresult.OrdGT(*i.OrdGT))
	}
	if i.OrdGTE != nil {
		predicates = append(predicates, benchresult.OrdGTE(*i.OrdGTE))
	}
	if i.OrdLT != nil {
		predicates = append(predicates, benchresult.OrdLT(*i.OrdLT))
	}
	if i.OrdLTE != nil {
		predicates = append(predicates, benchresult.OrdLTE(*i.OrdLTE))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyBenchResultWhereInput
	case 1:
		return predicates[0], nil
	default:
		return benchresult.And(predicates...), nil
	}
}
