// Code generated by ent, DO NOT EDIT.

package build

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/predicate"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/rssauthor"
)

// RSSAuthorDelete is the builder for deleting a RSSAuthor entity.
type RSSAuthorDelete struct {
	config
	hooks    []Hook
	mutation *RSSAuthorMutation
}

// Where appends a list predicates to the RSSAuthorDelete builder.
func (rad *RSSAuthorDelete) Where(ps ...predicate.RSSAuthor) *RSSAuthorDelete {
	rad.mutation.Where(ps...)
	return rad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rad *RSSAuthorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rad.sqlExec, rad.mutation, rad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rad *RSSAuthorDelete) ExecX(ctx context.Context) int {
	n, err := rad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rad *RSSAuthorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rssauthor.Table, sqlgraph.NewFieldSpec(rssauthor.FieldID, field.TypeUUID))
	if ps := rad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rad.mutation.done = true
	return affected, err
}

// RSSAuthorDeleteOne is the builder for deleting a single RSSAuthor entity.
type RSSAuthorDeleteOne struct {
	rad *RSSAuthorDelete
}

// Where appends a list predicates to the RSSAuthorDelete builder.
func (rado *RSSAuthorDeleteOne) Where(ps ...predicate.RSSAuthor) *RSSAuthorDeleteOne {
	rado.rad.mutation.Where(ps...)
	return rado
}

// Exec executes the deletion query.
func (rado *RSSAuthorDeleteOne) Exec(ctx context.Context) error {
	n, err := rado.rad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rssauthor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rado *RSSAuthorDeleteOne) ExecX(ctx context.Context) {
	if err := rado.Exec(ctx); err != nil {
		panic(err)
	}
}
