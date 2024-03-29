// Code generated by ent, DO NOT EDIT.

package person

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the person type in the database.
	Label = "person"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uuid"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// Table holds the table name of the person in the database.
	Table = "persons"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "rss_authors"
	// AuthorInverseTable is the table name for the RSSAuthor entity.
	// It exists in this package in order to avoid circular dependency with the "rssauthor" package.
	AuthorInverseTable = "rss_authors"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "rss_author_person"
)

// Columns holds all SQL columns for person fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldEmail,
	FieldName,
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
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Person queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAuthorCount orders the results by author count.
func ByAuthorCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorStep(), opts...)
	}
}

// ByAuthor orders the results by author terms.
func ByAuthor(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AuthorTable, AuthorColumn),
	)
}
