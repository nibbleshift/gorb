package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// BenchResult holds the schema definition for the BenchResult entity.
type BenchResult struct {
	ent.Schema
}

// Fields of the BenchResult.
func (BenchResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),              // benchmark name
		field.Int("N"),                    // number of iterations
		field.Float("NsPerOp"),            // nanoseconds per iteration
		field.Uint64("AllocedBytesPerOp"), // bytes allocated per iteration
		field.Uint64("AllocsPerOp"),       // allocs per iteration
		field.Float("MBPerS"),             // MB processed per second
		field.Int("Measured"),             // which measurements were recorded
		field.Int("Ord"),                  // ordinal position within a benchmark run
	}
}

// Edges of the BenchResult.
func (BenchResult) Edges() []ent.Edge {
	return nil
}

func (BenchResult) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		entgql.RelayConnection(),
	}
}
