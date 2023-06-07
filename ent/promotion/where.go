// Code generated by ent, DO NOT EDIT.

package promotion

import (
	"storage-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Promotion {
	return predicate.Promotion(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Promotion {
	return predicate.Promotion(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Promotion {
	return predicate.Promotion(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Promotion {
	return predicate.Promotion(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Promotion {
	return predicate.Promotion(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Promotion {
	return predicate.Promotion(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Promotion {
	return predicate.Promotion(sql.FieldLTE(FieldID, id))
}

// Pid applies equality check predicate on the "pid" field. It's identical to PidEQ.
func Pid(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldPid, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldPrice, v))
}

// ExpirationDate applies equality check predicate on the "expiration_date" field. It's identical to ExpirationDateEQ.
func ExpirationDate(v time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldExpirationDate, v))
}

// PidEQ applies the EQ predicate on the "pid" field.
func PidEQ(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldPid, v))
}

// PidNEQ applies the NEQ predicate on the "pid" field.
func PidNEQ(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldNEQ(FieldPid, v))
}

// PidIn applies the In predicate on the "pid" field.
func PidIn(vs ...string) predicate.Promotion {
	return predicate.Promotion(sql.FieldIn(FieldPid, vs...))
}

// PidNotIn applies the NotIn predicate on the "pid" field.
func PidNotIn(vs ...string) predicate.Promotion {
	return predicate.Promotion(sql.FieldNotIn(FieldPid, vs...))
}

// PidGT applies the GT predicate on the "pid" field.
func PidGT(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldGT(FieldPid, v))
}

// PidGTE applies the GTE predicate on the "pid" field.
func PidGTE(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldGTE(FieldPid, v))
}

// PidLT applies the LT predicate on the "pid" field.
func PidLT(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldLT(FieldPid, v))
}

// PidLTE applies the LTE predicate on the "pid" field.
func PidLTE(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldLTE(FieldPid, v))
}

// PidContains applies the Contains predicate on the "pid" field.
func PidContains(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldContains(FieldPid, v))
}

// PidHasPrefix applies the HasPrefix predicate on the "pid" field.
func PidHasPrefix(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldHasPrefix(FieldPid, v))
}

// PidHasSuffix applies the HasSuffix predicate on the "pid" field.
func PidHasSuffix(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldHasSuffix(FieldPid, v))
}

// PidEqualFold applies the EqualFold predicate on the "pid" field.
func PidEqualFold(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldEqualFold(FieldPid, v))
}

// PidContainsFold applies the ContainsFold predicate on the "pid" field.
func PidContainsFold(v string) predicate.Promotion {
	return predicate.Promotion(sql.FieldContainsFold(FieldPid, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldLTE(FieldPrice, v))
}

// ExpirationDateEQ applies the EQ predicate on the "expiration_date" field.
func ExpirationDateEQ(v time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldExpirationDate, v))
}

// ExpirationDateNEQ applies the NEQ predicate on the "expiration_date" field.
func ExpirationDateNEQ(v time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldNEQ(FieldExpirationDate, v))
}

// ExpirationDateIn applies the In predicate on the "expiration_date" field.
func ExpirationDateIn(vs ...time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldIn(FieldExpirationDate, vs...))
}

// ExpirationDateNotIn applies the NotIn predicate on the "expiration_date" field.
func ExpirationDateNotIn(vs ...time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldNotIn(FieldExpirationDate, vs...))
}

// ExpirationDateGT applies the GT predicate on the "expiration_date" field.
func ExpirationDateGT(v time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldGT(FieldExpirationDate, v))
}

// ExpirationDateGTE applies the GTE predicate on the "expiration_date" field.
func ExpirationDateGTE(v time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldGTE(FieldExpirationDate, v))
}

// ExpirationDateLT applies the LT predicate on the "expiration_date" field.
func ExpirationDateLT(v time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldLT(FieldExpirationDate, v))
}

// ExpirationDateLTE applies the LTE predicate on the "expiration_date" field.
func ExpirationDateLTE(v time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldLTE(FieldExpirationDate, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Promotion) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Promotion) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
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
func Not(p predicate.Promotion) predicate.Promotion {
	return predicate.Promotion(func(s *sql.Selector) {
		p(s.Not())
	})
}
