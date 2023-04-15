package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("number").Unique(),
		field.Int32("floor"),
		field.Int32("numberBeds"),
		field.Int32("numberPatients"),
		field.String("typeRoom"),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contains", Patient.Type),
	}
}
