// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeAudio holds the string denoting the audio edge name in mutations.
	EdgeAudio = "audio"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AudioTable is the table that holds the audio relation/edge.
	AudioTable = "audios"
	// AudioInverseTable is the table name for the Audio entity.
	// It exists in this package in order to avoid circular dependency with the "audio" package.
	AudioInverseTable = "audios"
	// AudioColumn is the table column denoting the audio relation/edge.
	AudioColumn = "user_audio"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldLogin,
	FieldPassword,
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
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// LoginValidator is a validator for the "login" field. It is called by the builders before save.
	LoginValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByLogin orders the results by the login field.
func ByLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogin, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByAudioCount orders the results by audio count.
func ByAudioCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAudioStep(), opts...)
	}
}

// ByAudio orders the results by audio terms.
func ByAudio(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAudioStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAudioStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AudioInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AudioTable, AudioColumn),
	)
}