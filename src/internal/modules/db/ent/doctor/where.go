// Code generated by ent, DO NOT EDIT.

package doctor

import (
	"hospital/src/internal/modules/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Doctor {
	return predicate.Doctor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Doctor {
	return predicate.Doctor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Doctor {
	return predicate.Doctor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Doctor {
	return predicate.Doctor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Doctor {
	return predicate.Doctor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Doctor {
	return predicate.Doctor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Doctor {
	return predicate.Doctor(sql.FieldLTE(FieldID, id))
}

// TokenId applies equality check predicate on the "tokenId" field. It's identical to TokenIdEQ.
func TokenId(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldTokenId, v))
}

// Surname applies equality check predicate on the "surname" field. It's identical to SurnameEQ.
func Surname(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldSurname, v))
}

// Speciality applies equality check predicate on the "speciality" field. It's identical to SpecialityEQ.
func Speciality(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldSpeciality, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldRole, v))
}

// TokenIdEQ applies the EQ predicate on the "tokenId" field.
func TokenIdEQ(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldTokenId, v))
}

// TokenIdNEQ applies the NEQ predicate on the "tokenId" field.
func TokenIdNEQ(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldNEQ(FieldTokenId, v))
}

// TokenIdIn applies the In predicate on the "tokenId" field.
func TokenIdIn(vs ...string) predicate.Doctor {
	return predicate.Doctor(sql.FieldIn(FieldTokenId, vs...))
}

// TokenIdNotIn applies the NotIn predicate on the "tokenId" field.
func TokenIdNotIn(vs ...string) predicate.Doctor {
	return predicate.Doctor(sql.FieldNotIn(FieldTokenId, vs...))
}

// TokenIdGT applies the GT predicate on the "tokenId" field.
func TokenIdGT(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldGT(FieldTokenId, v))
}

// TokenIdGTE applies the GTE predicate on the "tokenId" field.
func TokenIdGTE(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldGTE(FieldTokenId, v))
}

// TokenIdLT applies the LT predicate on the "tokenId" field.
func TokenIdLT(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldLT(FieldTokenId, v))
}

// TokenIdLTE applies the LTE predicate on the "tokenId" field.
func TokenIdLTE(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldLTE(FieldTokenId, v))
}

// TokenIdContains applies the Contains predicate on the "tokenId" field.
func TokenIdContains(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldContains(FieldTokenId, v))
}

// TokenIdHasPrefix applies the HasPrefix predicate on the "tokenId" field.
func TokenIdHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldHasPrefix(FieldTokenId, v))
}

// TokenIdHasSuffix applies the HasSuffix predicate on the "tokenId" field.
func TokenIdHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldHasSuffix(FieldTokenId, v))
}

// TokenIdEqualFold applies the EqualFold predicate on the "tokenId" field.
func TokenIdEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEqualFold(FieldTokenId, v))
}

// TokenIdContainsFold applies the ContainsFold predicate on the "tokenId" field.
func TokenIdContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldContainsFold(FieldTokenId, v))
}

// SurnameEQ applies the EQ predicate on the "surname" field.
func SurnameEQ(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldSurname, v))
}

// SurnameNEQ applies the NEQ predicate on the "surname" field.
func SurnameNEQ(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldNEQ(FieldSurname, v))
}

// SurnameIn applies the In predicate on the "surname" field.
func SurnameIn(vs ...string) predicate.Doctor {
	return predicate.Doctor(sql.FieldIn(FieldSurname, vs...))
}

// SurnameNotIn applies the NotIn predicate on the "surname" field.
func SurnameNotIn(vs ...string) predicate.Doctor {
	return predicate.Doctor(sql.FieldNotIn(FieldSurname, vs...))
}

// SurnameGT applies the GT predicate on the "surname" field.
func SurnameGT(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldGT(FieldSurname, v))
}

// SurnameGTE applies the GTE predicate on the "surname" field.
func SurnameGTE(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldGTE(FieldSurname, v))
}

// SurnameLT applies the LT predicate on the "surname" field.
func SurnameLT(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldLT(FieldSurname, v))
}

// SurnameLTE applies the LTE predicate on the "surname" field.
func SurnameLTE(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldLTE(FieldSurname, v))
}

// SurnameContains applies the Contains predicate on the "surname" field.
func SurnameContains(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldContains(FieldSurname, v))
}

// SurnameHasPrefix applies the HasPrefix predicate on the "surname" field.
func SurnameHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldHasPrefix(FieldSurname, v))
}

// SurnameHasSuffix applies the HasSuffix predicate on the "surname" field.
func SurnameHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldHasSuffix(FieldSurname, v))
}

// SurnameEqualFold applies the EqualFold predicate on the "surname" field.
func SurnameEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEqualFold(FieldSurname, v))
}

// SurnameContainsFold applies the ContainsFold predicate on the "surname" field.
func SurnameContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldContainsFold(FieldSurname, v))
}

// SpecialityEQ applies the EQ predicate on the "speciality" field.
func SpecialityEQ(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldSpeciality, v))
}

// SpecialityNEQ applies the NEQ predicate on the "speciality" field.
func SpecialityNEQ(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldNEQ(FieldSpeciality, v))
}

// SpecialityIn applies the In predicate on the "speciality" field.
func SpecialityIn(vs ...string) predicate.Doctor {
	return predicate.Doctor(sql.FieldIn(FieldSpeciality, vs...))
}

// SpecialityNotIn applies the NotIn predicate on the "speciality" field.
func SpecialityNotIn(vs ...string) predicate.Doctor {
	return predicate.Doctor(sql.FieldNotIn(FieldSpeciality, vs...))
}

// SpecialityGT applies the GT predicate on the "speciality" field.
func SpecialityGT(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldGT(FieldSpeciality, v))
}

// SpecialityGTE applies the GTE predicate on the "speciality" field.
func SpecialityGTE(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldGTE(FieldSpeciality, v))
}

// SpecialityLT applies the LT predicate on the "speciality" field.
func SpecialityLT(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldLT(FieldSpeciality, v))
}

// SpecialityLTE applies the LTE predicate on the "speciality" field.
func SpecialityLTE(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldLTE(FieldSpeciality, v))
}

// SpecialityContains applies the Contains predicate on the "speciality" field.
func SpecialityContains(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldContains(FieldSpeciality, v))
}

// SpecialityHasPrefix applies the HasPrefix predicate on the "speciality" field.
func SpecialityHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldHasPrefix(FieldSpeciality, v))
}

// SpecialityHasSuffix applies the HasSuffix predicate on the "speciality" field.
func SpecialityHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldHasSuffix(FieldSpeciality, v))
}

// SpecialityEqualFold applies the EqualFold predicate on the "speciality" field.
func SpecialityEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEqualFold(FieldSpeciality, v))
}

// SpecialityContainsFold applies the ContainsFold predicate on the "speciality" field.
func SpecialityContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldContainsFold(FieldSpeciality, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.Doctor {
	return predicate.Doctor(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.Doctor {
	return predicate.Doctor(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldLTE(FieldRole, v))
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldContains(FieldRole, v))
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldHasPrefix(FieldRole, v))
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldHasSuffix(FieldRole, v))
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldEqualFold(FieldRole, v))
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(sql.FieldContainsFold(FieldRole, v))
}

// HasTreats applies the HasEdge predicate on the "treats" edge.
func HasTreats() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TreatsTable, TreatsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTreatsWith applies the HasEdge predicate on the "treats" edge with a given conditions (other predicates).
func HasTreatsWith(preds ...predicate.Patient) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := newTreatsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
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
func Not(p predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		p(s.Not())
	})
}
