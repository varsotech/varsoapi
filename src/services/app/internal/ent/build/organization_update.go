// Code generated by ent, DO NOT EDIT.

package build

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/organization"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/predicate"
)

// OrganizationUpdate is the builder for updating Organization entities.
type OrganizationUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationMutation
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ou *OrganizationUpdate) Where(ps ...predicate.Organization) *OrganizationUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUniqueName sets the "unique_name" field.
func (ou *OrganizationUpdate) SetUniqueName(s string) *OrganizationUpdate {
	ou.mutation.SetUniqueName(s)
	return ou
}

// SetNillableUniqueName sets the "unique_name" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableUniqueName(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetUniqueName(*s)
	}
	return ou
}

// ClearUniqueName clears the value of the "unique_name" field.
func (ou *OrganizationUpdate) ClearUniqueName() *OrganizationUpdate {
	ou.mutation.ClearUniqueName()
	return ou
}

// SetName sets the "name" field.
func (ou *OrganizationUpdate) SetName(s string) *OrganizationUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableName(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetName(*s)
	}
	return ou
}

// SetWebsiteURL sets the "website_url" field.
func (ou *OrganizationUpdate) SetWebsiteURL(s string) *OrganizationUpdate {
	ou.mutation.SetWebsiteURL(s)
	return ou
}

// SetNillableWebsiteURL sets the "website_url" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableWebsiteURL(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetWebsiteURL(*s)
	}
	return ou
}

// SetRssFeedURL sets the "rss_feed_url" field.
func (ou *OrganizationUpdate) SetRssFeedURL(s string) *OrganizationUpdate {
	ou.mutation.SetRssFeedURL(s)
	return ou
}

// SetNillableRssFeedURL sets the "rss_feed_url" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableRssFeedURL(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetRssFeedURL(*s)
	}
	return ou
}

// Mutation returns the OrganizationMutation object of the builder.
func (ou *OrganizationUpdate) Mutation() *OrganizationMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrganizationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrganizationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrganizationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrganizationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrganizationUpdate) check() error {
	if v, ok := ou.mutation.UniqueName(); ok {
		if err := organization.UniqueNameValidator(v); err != nil {
			return &ValidationError{Name: "unique_name", err: fmt.Errorf(`build: validator failed for field "Organization.unique_name": %w`, err)}
		}
	}
	return nil
}

func (ou *OrganizationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UniqueName(); ok {
		_spec.SetField(organization.FieldUniqueName, field.TypeString, value)
	}
	if ou.mutation.UniqueNameCleared() {
		_spec.ClearField(organization.FieldUniqueName, field.TypeString)
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.WebsiteURL(); ok {
		_spec.SetField(organization.FieldWebsiteURL, field.TypeString, value)
	}
	if value, ok := ou.mutation.RssFeedURL(); ok {
		_spec.SetField(organization.FieldRssFeedURL, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrganizationUpdateOne is the builder for updating a single Organization entity.
type OrganizationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationMutation
}

// SetUniqueName sets the "unique_name" field.
func (ouo *OrganizationUpdateOne) SetUniqueName(s string) *OrganizationUpdateOne {
	ouo.mutation.SetUniqueName(s)
	return ouo
}

// SetNillableUniqueName sets the "unique_name" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableUniqueName(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetUniqueName(*s)
	}
	return ouo
}

// ClearUniqueName clears the value of the "unique_name" field.
func (ouo *OrganizationUpdateOne) ClearUniqueName() *OrganizationUpdateOne {
	ouo.mutation.ClearUniqueName()
	return ouo
}

// SetName sets the "name" field.
func (ouo *OrganizationUpdateOne) SetName(s string) *OrganizationUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableName(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetName(*s)
	}
	return ouo
}

// SetWebsiteURL sets the "website_url" field.
func (ouo *OrganizationUpdateOne) SetWebsiteURL(s string) *OrganizationUpdateOne {
	ouo.mutation.SetWebsiteURL(s)
	return ouo
}

// SetNillableWebsiteURL sets the "website_url" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableWebsiteURL(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetWebsiteURL(*s)
	}
	return ouo
}

// SetRssFeedURL sets the "rss_feed_url" field.
func (ouo *OrganizationUpdateOne) SetRssFeedURL(s string) *OrganizationUpdateOne {
	ouo.mutation.SetRssFeedURL(s)
	return ouo
}

// SetNillableRssFeedURL sets the "rss_feed_url" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableRssFeedURL(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetRssFeedURL(*s)
	}
	return ouo
}

// Mutation returns the OrganizationMutation object of the builder.
func (ouo *OrganizationUpdateOne) Mutation() *OrganizationMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ouo *OrganizationUpdateOne) Where(ps ...predicate.Organization) *OrganizationUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrganizationUpdateOne) Select(field string, fields ...string) *OrganizationUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Organization entity.
func (ouo *OrganizationUpdateOne) Save(ctx context.Context) (*Organization, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) SaveX(ctx context.Context) *Organization {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrganizationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrganizationUpdateOne) check() error {
	if v, ok := ouo.mutation.UniqueName(); ok {
		if err := organization.UniqueNameValidator(v); err != nil {
			return &ValidationError{Name: "unique_name", err: fmt.Errorf(`build: validator failed for field "Organization.unique_name": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrganizationUpdateOne) sqlSave(ctx context.Context) (_node *Organization, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`build: missing "Organization.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organization.FieldID)
		for _, f := range fields {
			if !organization.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("build: invalid field %q for query", f)}
			}
			if f != organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UniqueName(); ok {
		_spec.SetField(organization.FieldUniqueName, field.TypeString, value)
	}
	if ouo.mutation.UniqueNameCleared() {
		_spec.ClearField(organization.FieldUniqueName, field.TypeString)
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.WebsiteURL(); ok {
		_spec.SetField(organization.FieldWebsiteURL, field.TypeString, value)
	}
	if value, ok := ouo.mutation.RssFeedURL(); ok {
		_spec.SetField(organization.FieldRssFeedURL, field.TypeString, value)
	}
	_node = &Organization{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}