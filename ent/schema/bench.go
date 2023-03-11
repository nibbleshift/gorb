package schema

import "entgo.io/ent"

// Bench holds the schema definition for the Bench entity.
type Bench struct {
	ent.Schema
}

// Fields of the Bench.
func (Bench) Fields() []ent.Field {
	return nil
}

// Edges of the Bench.
func (Bench) Edges() []ent.Edge {
	return nil
}
