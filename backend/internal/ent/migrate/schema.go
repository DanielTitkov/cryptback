// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChallengesColumns holds the columns for the "challenges" table.
	ChallengesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "challenge", Type: field.TypeString, Unique: true},
		{Name: "text", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "author", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"first", "name", "frequency"}},
	}
	// ChallengesTable holds the schema information for the "challenges" table.
	ChallengesTable = &schema.Table{
		Name:       "challenges",
		Columns:    ChallengesColumns,
		PrimaryKey: []*schema.Column{ChallengesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "ip_address", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChallengesTable,
		UsersTable,
	}
)

func init() {
}
