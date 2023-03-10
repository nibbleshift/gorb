package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bench holds the schema definition for the Bench entity.
type Bench struct {
	ent.Schema
}

// Fields of the Bench.
func (Bench) Fields() []ent.Field {
	return []ent.Field{
		field.String("os"),
		field.String("arch"),
		field.String("cpu"),
		field.String("package"),
		field.Bool("pass"),
	}
}

// Edges of the Bench.
func (Bench) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bench_result", BenchResult.Type),
	}
}

func (Bench) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		entgql.RelayConnection(),
	}
}
