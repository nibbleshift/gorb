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
		field.String("OS"),
		field.String("Arch"),
		field.String("CPU"),
		field.String("Package"),
		field.Bool("Pass"),
	}
}

// Edges of the Bench.
func (Bench) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("results", BenchResult.Type),
	}
}

func (Bench) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
