// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PromptRequestsColumns holds the columns for the "prompt_requests" table.
	PromptRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "identifier", Type: field.TypeUUID, Unique: true},
		{Name: "prompt", Type: field.TypeString},
		{Name: "is_queued", Type: field.TypeBool, Default: false},
		{Name: "is_answered", Type: field.TypeBool, Default: false},
		{Name: "create_date", Type: field.TypeTime},
		{Name: "user_prompt_requests", Type: field.TypeInt, Nullable: true},
	}
	// PromptRequestsTable holds the schema information for the "prompt_requests" table.
	PromptRequestsTable = &schema.Table{
		Name:       "prompt_requests",
		Columns:    PromptRequestsColumns,
		PrimaryKey: []*schema.Column{PromptRequestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prompt_requests_users_prompt_requests",
				Columns:    []*schema.Column{PromptRequestsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PromptResponsesColumns holds the columns for the "prompt_responses" table.
	PromptResponsesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "response", Type: field.TypeString},
		{Name: "create_date", Type: field.TypeTime},
		{Name: "prompt_request_prompt_response", Type: field.TypeInt, Unique: true},
	}
	// PromptResponsesTable holds the schema information for the "prompt_responses" table.
	PromptResponsesTable = &schema.Table{
		Name:       "prompt_responses",
		Columns:    PromptResponsesColumns,
		PrimaryKey: []*schema.Column{PromptResponsesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prompt_responses_prompt_requests_prompt_response",
				Columns:    []*schema.Column{PromptResponsesColumns[3]},
				RefColumns: []*schema.Column{PromptRequestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeBytes},
		{Name: "salt", Type: field.TypeBytes},
		{Name: "create_date", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PromptRequestsTable,
		PromptResponsesTable,
		UsersTable,
	}
)

func init() {
	PromptRequestsTable.ForeignKeys[0].RefTable = UsersTable
	PromptResponsesTable.ForeignKeys[0].RefTable = PromptRequestsTable
}
