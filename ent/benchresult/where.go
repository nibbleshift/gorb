// Code generated by ent, DO NOT EDIT.

package benchresult

import (
	"entgo.io/ent/dialect/sql"
	"github.com/nibbleshift/gorb/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldName, v))
}

// N applies equality check predicate on the "n" field. It's identical to NEQ.
func N(v int64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldN, v))
}

// NsPerOp applies equality check predicate on the "ns_per_op" field. It's identical to NsPerOpEQ.
func NsPerOp(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldNsPerOp, v))
}

// AllocedBytesPerOp applies equality check predicate on the "alloced_bytes_per_op" field. It's identical to AllocedBytesPerOpEQ.
func AllocedBytesPerOp(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldAllocedBytesPerOp, v))
}

// AllocsPerOp applies equality check predicate on the "allocs_per_op" field. It's identical to AllocsPerOpEQ.
func AllocsPerOp(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldAllocsPerOp, v))
}

// MBPerS applies equality check predicate on the "mb_per_s" field. It's identical to MBPerSEQ.
func MBPerS(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldMBPerS, v))
}

// Measured applies equality check predicate on the "measured" field. It's identical to MeasuredEQ.
func Measured(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldMeasured, v))
}

// Ord applies equality check predicate on the "ord" field. It's identical to OrdEQ.
func Ord(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldOrd, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldContainsFold(FieldName, v))
}

// NEQ applies the EQ predicate on the "n" field.
func NEQ(v int64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldN, v))
}

// NNEQ applies the NEQ predicate on the "n" field.
func NNEQ(v int64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNEQ(FieldN, v))
}

// NIn applies the In predicate on the "n" field.
func NIn(vs ...int64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldIn(FieldN, vs...))
}

// NNotIn applies the NotIn predicate on the "n" field.
func NNotIn(vs ...int64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNotIn(FieldN, vs...))
}

// NGT applies the GT predicate on the "n" field.
func NGT(v int64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGT(FieldN, v))
}

// NGTE applies the GTE predicate on the "n" field.
func NGTE(v int64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGTE(FieldN, v))
}

// NLT applies the LT predicate on the "n" field.
func NLT(v int64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLT(FieldN, v))
}

// NLTE applies the LTE predicate on the "n" field.
func NLTE(v int64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLTE(FieldN, v))
}

// NsPerOpEQ applies the EQ predicate on the "ns_per_op" field.
func NsPerOpEQ(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldNsPerOp, v))
}

// NsPerOpNEQ applies the NEQ predicate on the "ns_per_op" field.
func NsPerOpNEQ(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNEQ(FieldNsPerOp, v))
}

// NsPerOpIn applies the In predicate on the "ns_per_op" field.
func NsPerOpIn(vs ...float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldIn(FieldNsPerOp, vs...))
}

// NsPerOpNotIn applies the NotIn predicate on the "ns_per_op" field.
func NsPerOpNotIn(vs ...float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNotIn(FieldNsPerOp, vs...))
}

// NsPerOpGT applies the GT predicate on the "ns_per_op" field.
func NsPerOpGT(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGT(FieldNsPerOp, v))
}

// NsPerOpGTE applies the GTE predicate on the "ns_per_op" field.
func NsPerOpGTE(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGTE(FieldNsPerOp, v))
}

// NsPerOpLT applies the LT predicate on the "ns_per_op" field.
func NsPerOpLT(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLT(FieldNsPerOp, v))
}

// NsPerOpLTE applies the LTE predicate on the "ns_per_op" field.
func NsPerOpLTE(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLTE(FieldNsPerOp, v))
}

// AllocedBytesPerOpEQ applies the EQ predicate on the "alloced_bytes_per_op" field.
func AllocedBytesPerOpEQ(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldAllocedBytesPerOp, v))
}

// AllocedBytesPerOpNEQ applies the NEQ predicate on the "alloced_bytes_per_op" field.
func AllocedBytesPerOpNEQ(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNEQ(FieldAllocedBytesPerOp, v))
}

// AllocedBytesPerOpIn applies the In predicate on the "alloced_bytes_per_op" field.
func AllocedBytesPerOpIn(vs ...uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldIn(FieldAllocedBytesPerOp, vs...))
}

// AllocedBytesPerOpNotIn applies the NotIn predicate on the "alloced_bytes_per_op" field.
func AllocedBytesPerOpNotIn(vs ...uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNotIn(FieldAllocedBytesPerOp, vs...))
}

// AllocedBytesPerOpGT applies the GT predicate on the "alloced_bytes_per_op" field.
func AllocedBytesPerOpGT(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGT(FieldAllocedBytesPerOp, v))
}

// AllocedBytesPerOpGTE applies the GTE predicate on the "alloced_bytes_per_op" field.
func AllocedBytesPerOpGTE(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGTE(FieldAllocedBytesPerOp, v))
}

// AllocedBytesPerOpLT applies the LT predicate on the "alloced_bytes_per_op" field.
func AllocedBytesPerOpLT(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLT(FieldAllocedBytesPerOp, v))
}

// AllocedBytesPerOpLTE applies the LTE predicate on the "alloced_bytes_per_op" field.
func AllocedBytesPerOpLTE(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLTE(FieldAllocedBytesPerOp, v))
}

// AllocsPerOpEQ applies the EQ predicate on the "allocs_per_op" field.
func AllocsPerOpEQ(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldAllocsPerOp, v))
}

// AllocsPerOpNEQ applies the NEQ predicate on the "allocs_per_op" field.
func AllocsPerOpNEQ(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNEQ(FieldAllocsPerOp, v))
}

// AllocsPerOpIn applies the In predicate on the "allocs_per_op" field.
func AllocsPerOpIn(vs ...uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldIn(FieldAllocsPerOp, vs...))
}

// AllocsPerOpNotIn applies the NotIn predicate on the "allocs_per_op" field.
func AllocsPerOpNotIn(vs ...uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNotIn(FieldAllocsPerOp, vs...))
}

// AllocsPerOpGT applies the GT predicate on the "allocs_per_op" field.
func AllocsPerOpGT(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGT(FieldAllocsPerOp, v))
}

// AllocsPerOpGTE applies the GTE predicate on the "allocs_per_op" field.
func AllocsPerOpGTE(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGTE(FieldAllocsPerOp, v))
}

// AllocsPerOpLT applies the LT predicate on the "allocs_per_op" field.
func AllocsPerOpLT(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLT(FieldAllocsPerOp, v))
}

// AllocsPerOpLTE applies the LTE predicate on the "allocs_per_op" field.
func AllocsPerOpLTE(v uint64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLTE(FieldAllocsPerOp, v))
}

// MBPerSEQ applies the EQ predicate on the "mb_per_s" field.
func MBPerSEQ(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldMBPerS, v))
}

// MBPerSNEQ applies the NEQ predicate on the "mb_per_s" field.
func MBPerSNEQ(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNEQ(FieldMBPerS, v))
}

// MBPerSIn applies the In predicate on the "mb_per_s" field.
func MBPerSIn(vs ...float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldIn(FieldMBPerS, vs...))
}

// MBPerSNotIn applies the NotIn predicate on the "mb_per_s" field.
func MBPerSNotIn(vs ...float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNotIn(FieldMBPerS, vs...))
}

// MBPerSGT applies the GT predicate on the "mb_per_s" field.
func MBPerSGT(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGT(FieldMBPerS, v))
}

// MBPerSGTE applies the GTE predicate on the "mb_per_s" field.
func MBPerSGTE(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGTE(FieldMBPerS, v))
}

// MBPerSLT applies the LT predicate on the "mb_per_s" field.
func MBPerSLT(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLT(FieldMBPerS, v))
}

// MBPerSLTE applies the LTE predicate on the "mb_per_s" field.
func MBPerSLTE(v float64) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLTE(FieldMBPerS, v))
}

// MeasuredEQ applies the EQ predicate on the "measured" field.
func MeasuredEQ(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldMeasured, v))
}

// MeasuredNEQ applies the NEQ predicate on the "measured" field.
func MeasuredNEQ(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNEQ(FieldMeasured, v))
}

// MeasuredIn applies the In predicate on the "measured" field.
func MeasuredIn(vs ...int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldIn(FieldMeasured, vs...))
}

// MeasuredNotIn applies the NotIn predicate on the "measured" field.
func MeasuredNotIn(vs ...int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNotIn(FieldMeasured, vs...))
}

// MeasuredGT applies the GT predicate on the "measured" field.
func MeasuredGT(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGT(FieldMeasured, v))
}

// MeasuredGTE applies the GTE predicate on the "measured" field.
func MeasuredGTE(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGTE(FieldMeasured, v))
}

// MeasuredLT applies the LT predicate on the "measured" field.
func MeasuredLT(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLT(FieldMeasured, v))
}

// MeasuredLTE applies the LTE predicate on the "measured" field.
func MeasuredLTE(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLTE(FieldMeasured, v))
}

// OrdEQ applies the EQ predicate on the "ord" field.
func OrdEQ(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldEQ(FieldOrd, v))
}

// OrdNEQ applies the NEQ predicate on the "ord" field.
func OrdNEQ(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNEQ(FieldOrd, v))
}

// OrdIn applies the In predicate on the "ord" field.
func OrdIn(vs ...int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldIn(FieldOrd, vs...))
}

// OrdNotIn applies the NotIn predicate on the "ord" field.
func OrdNotIn(vs ...int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldNotIn(FieldOrd, vs...))
}

// OrdGT applies the GT predicate on the "ord" field.
func OrdGT(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGT(FieldOrd, v))
}

// OrdGTE applies the GTE predicate on the "ord" field.
func OrdGTE(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldGTE(FieldOrd, v))
}

// OrdLT applies the LT predicate on the "ord" field.
func OrdLT(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLT(FieldOrd, v))
}

// OrdLTE applies the LTE predicate on the "ord" field.
func OrdLTE(v int) predicate.BenchResult {
	return predicate.BenchResult(sql.FieldLTE(FieldOrd, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BenchResult) predicate.BenchResult {
	return predicate.BenchResult(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BenchResult) predicate.BenchResult {
	return predicate.BenchResult(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BenchResult) predicate.BenchResult {
	return predicate.BenchResult(func(s *sql.Selector) {
		p(s.Not())
	})
}
