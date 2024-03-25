// Code generated by ent, DO NOT EDIT.

package build

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/organization"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/person"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/rssauthor"
)

// RSSAuthor is the model entity for the RSSAuthor schema.
type RSSAuthor struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RSSAuthorQuery when eager-loading is set.
	Edges                   RSSAuthorEdges `json:"edges"`
	organization_author     *uuid.UUID
	rss_author_person       *uuid.UUID
	rss_author_organization *uuid.UUID
	selectValues            sql.SelectValues
}

// RSSAuthorEdges holds the relations/edges for other nodes in the graph.
type RSSAuthorEdges struct {
	// Person holds the value of the person edge.
	Person *Person `json:"person,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// Newsitem holds the value of the newsitem edge.
	Newsitem []*NewsItem `json:"newsitem,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PersonOrErr returns the Person value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RSSAuthorEdges) PersonOrErr() (*Person, error) {
	if e.Person != nil {
		return e.Person, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: person.Label}
	}
	return nil, &NotLoadedError{edge: "person"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RSSAuthorEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// NewsitemOrErr returns the Newsitem value or an error if the edge
// was not loaded in eager-loading.
func (e RSSAuthorEdges) NewsitemOrErr() ([]*NewsItem, error) {
	if e.loadedTypes[2] {
		return e.Newsitem, nil
	}
	return nil, &NotLoadedError{edge: "newsitem"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RSSAuthor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rssauthor.FieldEmail, rssauthor.FieldName:
			values[i] = new(sql.NullString)
		case rssauthor.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case rssauthor.FieldID:
			values[i] = new(uuid.UUID)
		case rssauthor.ForeignKeys[0]: // organization_author
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case rssauthor.ForeignKeys[1]: // rss_author_person
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case rssauthor.ForeignKeys[2]: // rss_author_organization
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RSSAuthor fields.
func (ra *RSSAuthor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rssauthor.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ra.ID = *value
			}
		case rssauthor.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ra.CreateTime = value.Time
			}
		case rssauthor.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				ra.Email = value.String
			}
		case rssauthor.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ra.Name = value.String
			}
		case rssauthor.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_author", values[i])
			} else if value.Valid {
				ra.organization_author = new(uuid.UUID)
				*ra.organization_author = *value.S.(*uuid.UUID)
			}
		case rssauthor.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field rss_author_person", values[i])
			} else if value.Valid {
				ra.rss_author_person = new(uuid.UUID)
				*ra.rss_author_person = *value.S.(*uuid.UUID)
			}
		case rssauthor.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field rss_author_organization", values[i])
			} else if value.Valid {
				ra.rss_author_organization = new(uuid.UUID)
				*ra.rss_author_organization = *value.S.(*uuid.UUID)
			}
		default:
			ra.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RSSAuthor.
// This includes values selected through modifiers, order, etc.
func (ra *RSSAuthor) Value(name string) (ent.Value, error) {
	return ra.selectValues.Get(name)
}

// QueryPerson queries the "person" edge of the RSSAuthor entity.
func (ra *RSSAuthor) QueryPerson() *PersonQuery {
	return NewRSSAuthorClient(ra.config).QueryPerson(ra)
}

// QueryOrganization queries the "organization" edge of the RSSAuthor entity.
func (ra *RSSAuthor) QueryOrganization() *OrganizationQuery {
	return NewRSSAuthorClient(ra.config).QueryOrganization(ra)
}

// QueryNewsitem queries the "newsitem" edge of the RSSAuthor entity.
func (ra *RSSAuthor) QueryNewsitem() *NewsItemQuery {
	return NewRSSAuthorClient(ra.config).QueryNewsitem(ra)
}

// Update returns a builder for updating this RSSAuthor.
// Note that you need to call RSSAuthor.Unwrap() before calling this method if this RSSAuthor
// was returned from a transaction, and the transaction was committed or rolled back.
func (ra *RSSAuthor) Update() *RSSAuthorUpdateOne {
	return NewRSSAuthorClient(ra.config).UpdateOne(ra)
}

// Unwrap unwraps the RSSAuthor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ra *RSSAuthor) Unwrap() *RSSAuthor {
	_tx, ok := ra.config.driver.(*txDriver)
	if !ok {
		panic("build: RSSAuthor is not a transactional entity")
	}
	ra.config.driver = _tx.drv
	return ra
}

// String implements the fmt.Stringer.
func (ra *RSSAuthor) String() string {
	var builder strings.Builder
	builder.WriteString("RSSAuthor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ra.ID))
	builder.WriteString("create_time=")
	builder.WriteString(ra.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(ra.Email)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ra.Name)
	builder.WriteByte(')')
	return builder.String()
}

// RSSAuthors is a parsable slice of RSSAuthor.
type RSSAuthors []*RSSAuthor
