// Code generated by ent, DO NOT EDIT.

package promotion

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the promotion type in the database.
	Label = "promotion"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldExpirationDate holds the string denoting the expiration_date field in the database.
	FieldExpirationDate = "expiration_date"
	// Table holds the table name of the promotion in the database.
	Table = "promotions"
)

// Columns holds all SQL columns for promotion fields.
var Columns = []string{
	FieldID,
	FieldPid,
	FieldPrice,
	FieldExpirationDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Promotion queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPid orders the results by the pid field.
func ByPid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPid, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByExpirationDate orders the results by the expiration_date field.
func ByExpirationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpirationDate, opts...).ToFunc()
}
