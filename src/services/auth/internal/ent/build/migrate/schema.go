// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "permissions", Type: field.TypeJSON},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_uuid",
				Unique:  true,
				Columns: []*schema.Column{RolesColumns[1]},
			},
			{
				Name:    "role_key",
				Unique:  true,
				Columns: []*schema.Column{RolesColumns[2]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "ban_time", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_uuid",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[0]},
			},
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[3]},
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeInt},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RolesTable,
		UsersTable,
		UserRolesTable,
	}
)

func init() {
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[1].RefTable = RolesTable
}
