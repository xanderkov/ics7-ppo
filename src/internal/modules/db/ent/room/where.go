// Code generated by ent, DO NOT EDIT.

package room

import (
	"hospital/internal/modules/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldID, id))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int32) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldNumber, v))
}

// Floor applies equality check predicate on the "floor" field. It's identical to FloorEQ.
func Floor(v int32) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldFloor, v))
}

// NumberBeds applies equality check predicate on the "numberBeds" field. It's identical to NumberBedsEQ.
func NumberBeds(v int32) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldNumberBeds, v))
}

// NumberPatients applies equality check predicate on the "numberPatients" field. It's identical to NumberPatientsEQ.
func NumberPatients(v int32) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldNumberPatients, v))
}

// TypeRoom applies equality check predicate on the "typeRoom" field. It's identical to TypeRoomEQ.
func TypeRoom(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldTypeRoom, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int32) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int32) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int32) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int32) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int32) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int32) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int32) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int32) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldNumber, v))
}

// FloorEQ applies the EQ predicate on the "floor" field.
func FloorEQ(v int32) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldFloor, v))
}

// FloorNEQ applies the NEQ predicate on the "floor" field.
func FloorNEQ(v int32) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldFloor, v))
}

// FloorIn applies the In predicate on the "floor" field.
func FloorIn(vs ...int32) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldFloor, vs...))
}

// FloorNotIn applies the NotIn predicate on the "floor" field.
func FloorNotIn(vs ...int32) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldFloor, vs...))
}

// FloorGT applies the GT predicate on the "floor" field.
func FloorGT(v int32) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldFloor, v))
}

// FloorGTE applies the GTE predicate on the "floor" field.
func FloorGTE(v int32) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldFloor, v))
}

// FloorLT applies the LT predicate on the "floor" field.
func FloorLT(v int32) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldFloor, v))
}

// FloorLTE applies the LTE predicate on the "floor" field.
func FloorLTE(v int32) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldFloor, v))
}

// NumberBedsEQ applies the EQ predicate on the "numberBeds" field.
func NumberBedsEQ(v int32) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldNumberBeds, v))
}

// NumberBedsNEQ applies the NEQ predicate on the "numberBeds" field.
func NumberBedsNEQ(v int32) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldNumberBeds, v))
}

// NumberBedsIn applies the In predicate on the "numberBeds" field.
func NumberBedsIn(vs ...int32) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldNumberBeds, vs...))
}

// NumberBedsNotIn applies the NotIn predicate on the "numberBeds" field.
func NumberBedsNotIn(vs ...int32) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldNumberBeds, vs...))
}

// NumberBedsGT applies the GT predicate on the "numberBeds" field.
func NumberBedsGT(v int32) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldNumberBeds, v))
}

// NumberBedsGTE applies the GTE predicate on the "numberBeds" field.
func NumberBedsGTE(v int32) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldNumberBeds, v))
}

// NumberBedsLT applies the LT predicate on the "numberBeds" field.
func NumberBedsLT(v int32) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldNumberBeds, v))
}

// NumberBedsLTE applies the LTE predicate on the "numberBeds" field.
func NumberBedsLTE(v int32) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldNumberBeds, v))
}

// NumberPatientsEQ applies the EQ predicate on the "numberPatients" field.
func NumberPatientsEQ(v int32) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldNumberPatients, v))
}

// NumberPatientsNEQ applies the NEQ predicate on the "numberPatients" field.
func NumberPatientsNEQ(v int32) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldNumberPatients, v))
}

// NumberPatientsIn applies the In predicate on the "numberPatients" field.
func NumberPatientsIn(vs ...int32) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldNumberPatients, vs...))
}

// NumberPatientsNotIn applies the NotIn predicate on the "numberPatients" field.
func NumberPatientsNotIn(vs ...int32) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldNumberPatients, vs...))
}

// NumberPatientsGT applies the GT predicate on the "numberPatients" field.
func NumberPatientsGT(v int32) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldNumberPatients, v))
}

// NumberPatientsGTE applies the GTE predicate on the "numberPatients" field.
func NumberPatientsGTE(v int32) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldNumberPatients, v))
}

// NumberPatientsLT applies the LT predicate on the "numberPatients" field.
func NumberPatientsLT(v int32) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldNumberPatients, v))
}

// NumberPatientsLTE applies the LTE predicate on the "numberPatients" field.
func NumberPatientsLTE(v int32) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldNumberPatients, v))
}

// TypeRoomEQ applies the EQ predicate on the "typeRoom" field.
func TypeRoomEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldEQ(FieldTypeRoom, v))
}

// TypeRoomNEQ applies the NEQ predicate on the "typeRoom" field.
func TypeRoomNEQ(v string) predicate.Room {
	return predicate.Room(sql.FieldNEQ(FieldTypeRoom, v))
}

// TypeRoomIn applies the In predicate on the "typeRoom" field.
func TypeRoomIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldIn(FieldTypeRoom, vs...))
}

// TypeRoomNotIn applies the NotIn predicate on the "typeRoom" field.
func TypeRoomNotIn(vs ...string) predicate.Room {
	return predicate.Room(sql.FieldNotIn(FieldTypeRoom, vs...))
}

// TypeRoomGT applies the GT predicate on the "typeRoom" field.
func TypeRoomGT(v string) predicate.Room {
	return predicate.Room(sql.FieldGT(FieldTypeRoom, v))
}

// TypeRoomGTE applies the GTE predicate on the "typeRoom" field.
func TypeRoomGTE(v string) predicate.Room {
	return predicate.Room(sql.FieldGTE(FieldTypeRoom, v))
}

// TypeRoomLT applies the LT predicate on the "typeRoom" field.
func TypeRoomLT(v string) predicate.Room {
	return predicate.Room(sql.FieldLT(FieldTypeRoom, v))
}

// TypeRoomLTE applies the LTE predicate on the "typeRoom" field.
func TypeRoomLTE(v string) predicate.Room {
	return predicate.Room(sql.FieldLTE(FieldTypeRoom, v))
}

// TypeRoomContains applies the Contains predicate on the "typeRoom" field.
func TypeRoomContains(v string) predicate.Room {
	return predicate.Room(sql.FieldContains(FieldTypeRoom, v))
}

// TypeRoomHasPrefix applies the HasPrefix predicate on the "typeRoom" field.
func TypeRoomHasPrefix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasPrefix(FieldTypeRoom, v))
}

// TypeRoomHasSuffix applies the HasSuffix predicate on the "typeRoom" field.
func TypeRoomHasSuffix(v string) predicate.Room {
	return predicate.Room(sql.FieldHasSuffix(FieldTypeRoom, v))
}

// TypeRoomEqualFold applies the EqualFold predicate on the "typeRoom" field.
func TypeRoomEqualFold(v string) predicate.Room {
	return predicate.Room(sql.FieldEqualFold(FieldTypeRoom, v))
}

// TypeRoomContainsFold applies the ContainsFold predicate on the "typeRoom" field.
func TypeRoomContainsFold(v string) predicate.Room {
	return predicate.Room(sql.FieldContainsFold(FieldTypeRoom, v))
}

// HasContains applies the HasEdge predicate on the "contains" edge.
func HasContains() predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ContainsTable, ContainsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContainsWith applies the HasEdge predicate on the "contains" edge with a given conditions (other predicates).
func HasContainsWith(preds ...predicate.Patient) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		step := newContainsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
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
func Not(p predicate.Room) predicate.Room {
	return predicate.Room(func(s *sql.Selector) {
		p(s.Not())
	})
}
