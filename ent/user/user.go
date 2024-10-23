// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldAPIKey holds the string denoting the api_key field in the database.
	FieldAPIKey = "api_key"
	// FieldCreateDate holds the string denoting the create_date field in the database.
	FieldCreateDate = "create_date"
	// EdgePromptRequests holds the string denoting the prompt_requests edge name in mutations.
	EdgePromptRequests = "prompt_requests"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PromptRequestsTable is the table that holds the prompt_requests relation/edge.
	PromptRequestsTable = "prompt_requests"
	// PromptRequestsInverseTable is the table name for the PromptRequest entity.
	// It exists in this package in order to avoid circular dependency with the "promptrequest" package.
	PromptRequestsInverseTable = "prompt_requests"
	// PromptRequestsColumn is the table column denoting the prompt_requests relation/edge.
	PromptRequestsColumn = "user_prompt_requests"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldAPIKey,
	FieldCreateDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateDate holds the default value on creation for the "create_date" field.
	DefaultCreateDate func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByAPIKey orders the results by the api_key field.
func ByAPIKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAPIKey, opts...).ToFunc()
}

// ByCreateDate orders the results by the create_date field.
func ByCreateDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateDate, opts...).ToFunc()
}

// ByPromptRequestsCount orders the results by prompt_requests count.
func ByPromptRequestsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPromptRequestsStep(), opts...)
	}
}

// ByPromptRequests orders the results by prompt_requests terms.
func ByPromptRequests(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPromptRequestsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPromptRequestsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PromptRequestsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PromptRequestsTable, PromptRequestsColumn),
	)
}