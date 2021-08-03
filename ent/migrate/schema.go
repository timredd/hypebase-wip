// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ServiceTwitchesColumns holds the columns for the "service_twitches" table.
	ServiceTwitchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "access_token", Type: field.TypeString},
		{Name: "refresh_token", Type: field.TypeString},
		{Name: "scopes", Type: field.TypeJSON},
	}
	// ServiceTwitchesTable holds the schema information for the "service_twitches" table.
	ServiceTwitchesTable = &schema.Table{
		Name:        "service_twitches",
		Columns:     ServiceTwitchesColumns,
		PrimaryKey:  []*schema.Column{ServiceTwitchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UserSessionsColumns holds the columns for the "user_sessions" table.
	UserSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "session_id", Type: field.TypeString, Unique: true},
		{Name: "login_time", Type: field.TypeTime},
		{Name: "last_seen_time", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// UserSessionsTable holds the schema information for the "user_sessions" table.
	UserSessionsTable = &schema.Table{
		Name:       "user_sessions",
		Columns:    UserSessionsColumns,
		PrimaryKey: []*schema.Column{UserSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_sessions_users_user",
				Columns:    []*schema.Column{UserSessionsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ServiceTwitchesTable,
		UsersTable,
		UserSessionsTable,
	}
)

func init() {
	UserSessionsTable.ForeignKeys[0].RefTable = UsersTable
}
