// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"storage-app/ent/promotion"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Promotion is the model entity for the Promotion schema.
type Promotion struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// ExpirationDate holds the value of the "expiration_date" field.
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Promotion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case promotion.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case promotion.FieldID:
			values[i] = new(sql.NullString)
		case promotion.FieldExpirationDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Promotion fields.
func (pr *Promotion) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case promotion.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case promotion.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = value.Float64
			}
		case promotion.FieldExpirationDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_date", values[i])
			} else if value.Valid {
				pr.ExpirationDate = value.Time
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Promotion.
// This includes values selected through modifiers, order, etc.
func (pr *Promotion) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Promotion.
// Note that you need to call Promotion.Unwrap() before calling this method if this Promotion
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Promotion) Update() *PromotionUpdateOne {
	return NewPromotionClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Promotion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Promotion) Unwrap() *Promotion {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Promotion is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Promotion) String() string {
	var builder strings.Builder
	builder.WriteString("Promotion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", ")
	builder.WriteString("expiration_date=")
	builder.WriteString(pr.ExpirationDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Promotions is a parsable slice of Promotion.
type Promotions []*Promotion
