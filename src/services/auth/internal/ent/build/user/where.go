// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeleteTime, v))
}

// BanTime applies equality check predicate on the "ban_time" field. It's identical to BanTimeEQ.
func BanTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBanTime, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Salt applies equality check predicate on the "salt" field. It's identical to SaltEQ.
func Salt(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSalt, v))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDeleteTime))
}

// BanTimeEQ applies the EQ predicate on the "ban_time" field.
func BanTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBanTime, v))
}

// BanTimeNEQ applies the NEQ predicate on the "ban_time" field.
func BanTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBanTime, v))
}

// BanTimeIn applies the In predicate on the "ban_time" field.
func BanTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldBanTime, vs...))
}

// BanTimeNotIn applies the NotIn predicate on the "ban_time" field.
func BanTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBanTime, vs...))
}

// BanTimeGT applies the GT predicate on the "ban_time" field.
func BanTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldBanTime, v))
}

// BanTimeGTE applies the GTE predicate on the "ban_time" field.
func BanTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBanTime, v))
}

// BanTimeLT applies the LT predicate on the "ban_time" field.
func BanTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldBanTime, v))
}

// BanTimeLTE applies the LTE predicate on the "ban_time" field.
func BanTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBanTime, v))
}

// BanTimeIsNil applies the IsNil predicate on the "ban_time" field.
func BanTimeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldBanTime))
}

// BanTimeNotNil applies the NotNil predicate on the "ban_time" field.
func BanTimeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldBanTime))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// SaltEQ applies the EQ predicate on the "salt" field.
func SaltEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSalt, v))
}

// SaltNEQ applies the NEQ predicate on the "salt" field.
func SaltNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSalt, v))
}

// SaltIn applies the In predicate on the "salt" field.
func SaltIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldSalt, vs...))
}

// SaltNotIn applies the NotIn predicate on the "salt" field.
func SaltNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSalt, vs...))
}

// SaltGT applies the GT predicate on the "salt" field.
func SaltGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldSalt, v))
}

// SaltGTE applies the GTE predicate on the "salt" field.
func SaltGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldSalt, v))
}

// SaltLT applies the LT predicate on the "salt" field.
func SaltLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldSalt, v))
}

// SaltLTE applies the LTE predicate on the "salt" field.
func SaltLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldSalt, v))
}

// SaltContains applies the Contains predicate on the "salt" field.
func SaltContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldSalt, v))
}

// SaltHasPrefix applies the HasPrefix predicate on the "salt" field.
func SaltHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldSalt, v))
}

// SaltHasSuffix applies the HasSuffix predicate on the "salt" field.
func SaltHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldSalt, v))
}

// SaltEqualFold applies the EqualFold predicate on the "salt" field.
func SaltEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldSalt, v))
}

// SaltContainsFold applies the ContainsFold predicate on the "salt" field.
func SaltContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldSalt, v))
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newRolesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
