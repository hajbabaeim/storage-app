// ent/schema/promotion.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Promotion holds the schema definition for the Promotion entity.
type Promotion struct {
	ent.Schema
}

// Fields of the Promotion.
func (Promotion) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Float("price"),
		field.Time("expiration_date"),
	}
}

// Edges of the Promotion.
func (Promotion) Edges() []ent.Edge {
	return []ent.Edge{}
}
