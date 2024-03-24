// Code generated by ent, DO NOT EDIT.

package build

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/newsitem"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/person"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/predicate"
)

// PersonUpdate is the builder for updating Person entities.
type PersonUpdate struct {
	config
	hooks    []Hook
	mutation *PersonMutation
}

// Where appends a list predicates to the PersonUpdate builder.
func (pu *PersonUpdate) Where(ps ...predicate.Person) *PersonUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *PersonUpdate) SetUpdateTime(t time.Time) *PersonUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// SetEmail sets the "email" field.
func (pu *PersonUpdate) SetEmail(s string) *PersonUpdate {
	pu.mutation.SetEmail(s)
	return pu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableEmail(s *string) *PersonUpdate {
	if s != nil {
		pu.SetEmail(*s)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *PersonUpdate) SetName(s string) *PersonUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableName(s *string) *PersonUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// AddItemIDs adds the "item" edge to the NewsItem entity by IDs.
func (pu *PersonUpdate) AddItemIDs(ids ...uuid.UUID) *PersonUpdate {
	pu.mutation.AddItemIDs(ids...)
	return pu
}

// AddItem adds the "item" edges to the NewsItem entity.
func (pu *PersonUpdate) AddItem(n ...*NewsItem) *PersonUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return pu.AddItemIDs(ids...)
}

// Mutation returns the PersonMutation object of the builder.
func (pu *PersonUpdate) Mutation() *PersonMutation {
	return pu.mutation
}

// ClearItem clears all "item" edges to the NewsItem entity.
func (pu *PersonUpdate) ClearItem() *PersonUpdate {
	pu.mutation.ClearItem()
	return pu
}

// RemoveItemIDs removes the "item" edge to NewsItem entities by IDs.
func (pu *PersonUpdate) RemoveItemIDs(ids ...uuid.UUID) *PersonUpdate {
	pu.mutation.RemoveItemIDs(ids...)
	return pu
}

// RemoveItem removes "item" edges to NewsItem entities.
func (pu *PersonUpdate) RemoveItem(n ...*NewsItem) *PersonUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return pu.RemoveItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PersonUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PersonUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PersonUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PersonUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PersonUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		v := person.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

func (pu *PersonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.SetField(person.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
	}
	if pu.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   person.ItemTable,
			Columns: person.ItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedItemIDs(); len(nodes) > 0 && !pu.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   person.ItemTable,
			Columns: person.ItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   person.ItemTable,
			Columns: person.ItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PersonUpdateOne is the builder for updating a single Person entity.
type PersonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PersonMutation
}

// SetUpdateTime sets the "update_time" field.
func (puo *PersonUpdateOne) SetUpdateTime(t time.Time) *PersonUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// SetEmail sets the "email" field.
func (puo *PersonUpdateOne) SetEmail(s string) *PersonUpdateOne {
	puo.mutation.SetEmail(s)
	return puo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableEmail(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetEmail(*s)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *PersonUpdateOne) SetName(s string) *PersonUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableName(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// AddItemIDs adds the "item" edge to the NewsItem entity by IDs.
func (puo *PersonUpdateOne) AddItemIDs(ids ...uuid.UUID) *PersonUpdateOne {
	puo.mutation.AddItemIDs(ids...)
	return puo
}

// AddItem adds the "item" edges to the NewsItem entity.
func (puo *PersonUpdateOne) AddItem(n ...*NewsItem) *PersonUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return puo.AddItemIDs(ids...)
}

// Mutation returns the PersonMutation object of the builder.
func (puo *PersonUpdateOne) Mutation() *PersonMutation {
	return puo.mutation
}

// ClearItem clears all "item" edges to the NewsItem entity.
func (puo *PersonUpdateOne) ClearItem() *PersonUpdateOne {
	puo.mutation.ClearItem()
	return puo
}

// RemoveItemIDs removes the "item" edge to NewsItem entities by IDs.
func (puo *PersonUpdateOne) RemoveItemIDs(ids ...uuid.UUID) *PersonUpdateOne {
	puo.mutation.RemoveItemIDs(ids...)
	return puo
}

// RemoveItem removes "item" edges to NewsItem entities.
func (puo *PersonUpdateOne) RemoveItem(n ...*NewsItem) *PersonUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return puo.RemoveItemIDs(ids...)
}

// Where appends a list predicates to the PersonUpdate builder.
func (puo *PersonUpdateOne) Where(ps ...predicate.Person) *PersonUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PersonUpdateOne) Select(field string, fields ...string) *PersonUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Person entity.
func (puo *PersonUpdateOne) Save(ctx context.Context) (*Person, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PersonUpdateOne) SaveX(ctx context.Context) *Person {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PersonUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PersonUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PersonUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		v := person.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

func (puo *PersonUpdateOne) sqlSave(ctx context.Context) (_node *Person, err error) {
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`build: missing "Person.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, person.FieldID)
		for _, f := range fields {
			if !person.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("build: invalid field %q for query", f)}
			}
			if f != person.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.SetField(person.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
	}
	if puo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   person.ItemTable,
			Columns: person.ItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedItemIDs(); len(nodes) > 0 && !puo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   person.ItemTable,
			Columns: person.ItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   person.ItemTable,
			Columns: person.ItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Person{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
