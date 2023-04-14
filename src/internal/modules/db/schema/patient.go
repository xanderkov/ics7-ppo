package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return nil
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doctor", Doctor.Type).
			Ref("treats"),
	}
}
