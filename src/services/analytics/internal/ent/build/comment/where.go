// Code generated by ent, DO NOT EDIT.

package comment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/varsotech/varsoapi/src/services/analytics/internal/ent/build/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldID, id))
}

// UserUUID applies equality check predicate on the "user_uuid" field. It's identical to UserUUIDEQ.
func UserUUID(v uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserUUID, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldText, v))
}

// WasEdited applies equality check predicate on the "was_edited" field. It's identical to WasEditedEQ.
func WasEdited(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldWasEdited, v))
}

// TotalVotes applies equality check predicate on the "total_votes" field. It's identical to TotalVotesEQ.
func TotalVotes(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldTotalVotes, v))
}

// Upvotes applies equality check predicate on the "upvotes" field. It's identical to UpvotesEQ.
func Upvotes(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpvotes, v))
}

// Downvotes applies equality check predicate on the "downvotes" field. It's identical to DownvotesEQ.
func Downvotes(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldDownvotes, v))
}

// UserUUIDEQ applies the EQ predicate on the "user_uuid" field.
func UserUUIDEQ(v uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserUUID, v))
}

// UserUUIDNEQ applies the NEQ predicate on the "user_uuid" field.
func UserUUIDNEQ(v uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUserUUID, v))
}

// UserUUIDIn applies the In predicate on the "user_uuid" field.
func UserUUIDIn(vs ...uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUserUUID, vs...))
}

// UserUUIDNotIn applies the NotIn predicate on the "user_uuid" field.
func UserUUIDNotIn(vs ...uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUserUUID, vs...))
}

// UserUUIDGT applies the GT predicate on the "user_uuid" field.
func UserUUIDGT(v uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldUserUUID, v))
}

// UserUUIDGTE applies the GTE predicate on the "user_uuid" field.
func UserUUIDGTE(v uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldUserUUID, v))
}

// UserUUIDLT applies the LT predicate on the "user_uuid" field.
func UserUUIDLT(v uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldUserUUID, v))
}

// UserUUIDLTE applies the LTE predicate on the "user_uuid" field.
func UserUUIDLTE(v uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldUserUUID, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldText, v))
}

// WasEditedEQ applies the EQ predicate on the "was_edited" field.
func WasEditedEQ(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldWasEdited, v))
}

// WasEditedNEQ applies the NEQ predicate on the "was_edited" field.
func WasEditedNEQ(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldWasEdited, v))
}

// TotalVotesEQ applies the EQ predicate on the "total_votes" field.
func TotalVotesEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldTotalVotes, v))
}

// TotalVotesNEQ applies the NEQ predicate on the "total_votes" field.
func TotalVotesNEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldTotalVotes, v))
}

// TotalVotesIn applies the In predicate on the "total_votes" field.
func TotalVotesIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldTotalVotes, vs...))
}

// TotalVotesNotIn applies the NotIn predicate on the "total_votes" field.
func TotalVotesNotIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldTotalVotes, vs...))
}

// TotalVotesGT applies the GT predicate on the "total_votes" field.
func TotalVotesGT(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldTotalVotes, v))
}

// TotalVotesGTE applies the GTE predicate on the "total_votes" field.
func TotalVotesGTE(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldTotalVotes, v))
}

// TotalVotesLT applies the LT predicate on the "total_votes" field.
func TotalVotesLT(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldTotalVotes, v))
}

// TotalVotesLTE applies the LTE predicate on the "total_votes" field.
func TotalVotesLTE(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldTotalVotes, v))
}

// UpvotesEQ applies the EQ predicate on the "upvotes" field.
func UpvotesEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpvotes, v))
}

// UpvotesNEQ applies the NEQ predicate on the "upvotes" field.
func UpvotesNEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUpvotes, v))
}

// UpvotesIn applies the In predicate on the "upvotes" field.
func UpvotesIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUpvotes, vs...))
}

// UpvotesNotIn applies the NotIn predicate on the "upvotes" field.
func UpvotesNotIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUpvotes, vs...))
}

// UpvotesGT applies the GT predicate on the "upvotes" field.
func UpvotesGT(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldUpvotes, v))
}

// UpvotesGTE applies the GTE predicate on the "upvotes" field.
func UpvotesGTE(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldUpvotes, v))
}

// UpvotesLT applies the LT predicate on the "upvotes" field.
func UpvotesLT(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldUpvotes, v))
}

// UpvotesLTE applies the LTE predicate on the "upvotes" field.
func UpvotesLTE(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldUpvotes, v))
}

// DownvotesEQ applies the EQ predicate on the "downvotes" field.
func DownvotesEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldDownvotes, v))
}

// DownvotesNEQ applies the NEQ predicate on the "downvotes" field.
func DownvotesNEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldDownvotes, v))
}

// DownvotesIn applies the In predicate on the "downvotes" field.
func DownvotesIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldDownvotes, vs...))
}

// DownvotesNotIn applies the NotIn predicate on the "downvotes" field.
func DownvotesNotIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldDownvotes, vs...))
}

// DownvotesGT applies the GT predicate on the "downvotes" field.
func DownvotesGT(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldDownvotes, v))
}

// DownvotesGTE applies the GTE predicate on the "downvotes" field.
func DownvotesGTE(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldDownvotes, v))
}

// DownvotesLT applies the LT predicate on the "downvotes" field.
func DownvotesLT(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldDownvotes, v))
}

// DownvotesLTE applies the LTE predicate on the "downvotes" field.
func DownvotesLTE(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldDownvotes, v))
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.NotPredicates(p))
}