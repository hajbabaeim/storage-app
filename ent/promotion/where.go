// Code generated by ent, DO NOT EDIT.

package promotion

import (
	"storage-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Promotion {
	return predicate.Promotion(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Promotion {
	return predicate.Promotion(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Promotion {
	return predicate.Promotion(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Promotion {
	return predicate.Promotion(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Promotion {
	return predicate.Promotion(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Promotion {
	return predicate.Promotion(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Promotion {
	return predicate.Promotion(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Promotion {
	return predicate.Promotion(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Promotion {
	return predicate.Promotion(sql.FieldContainsFold(FieldID, id))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldPrice, v))
}

// ExpirationDate applies equality check predicate on the "expiration_date" field. It's identical to ExpirationDateEQ.
func ExpirationDate(v time.Time) predicate.Promotion {
	return predicate.Promotion(sql.FieldEQ(FieldExpirationDate, v))
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