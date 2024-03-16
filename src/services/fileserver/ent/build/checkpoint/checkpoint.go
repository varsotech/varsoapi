// Code generated by ent, DO NOT EDIT.

package checkpoint

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the checkpoint type in the database.
	Label = "checkpoint"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUserUUID holds the string denoting the user_uuid field in the database.
	FieldUserUUID = "user_uuid"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// Table holds the table name of the checkpoint in the database.
	Table = "checkpoints"
)

// Columns holds all SQL columns for checkpoint fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUserUUID,
	FieldKey,
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
)

// OrderOption defines the ordering options for the Checkpoint queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUserUUID orders the results by the user_uuid field.
func ByUserUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserUUID, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}