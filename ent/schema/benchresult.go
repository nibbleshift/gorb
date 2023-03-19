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
		field.String("name"),                // benchmark name
		field.Int64("n"),                    // number of iterations
		field.Float("ns_per_op"),            // nanoseconds per iteration
		field.Int64("alloced_bytes_per_op"), // bytes allocated per iteration
		field.Int64("allocs_per_op"),        // allocs per iteration
		field.Float("mb_per_s"),             // MB processed per second
		field.Int64("measured"),             // which measurements were recorded
		field.Int64("ord"),                  // ordinal position within a benchmark run
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
