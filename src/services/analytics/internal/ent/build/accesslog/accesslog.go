// Code generated by ent, DO NOT EDIT.

package accesslog

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the accesslog type in the database.
	Label = "access_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldURI holds the string denoting the uri field in the database.
	FieldURI = "uri"
	// FieldForwardedFor holds the string denoting the forwarded_for field in the database.
	FieldForwardedFor = "forwarded_for"
	// FieldForwardedProto holds the string denoting the forwarded_proto field in the database.
	FieldForwardedProto = "forwarded_proto"
	// FieldForwardedHost holds the string denoting the forwarded_host field in the database.
	FieldForwardedHost = "forwarded_host"
	// FieldForwardedPort holds the string denoting the forwarded_port field in the database.
	FieldForwardedPort = "forwarded_port"
	// FieldForwardedServer holds the string denoting the forwarded_server field in the database.
	FieldForwardedServer = "forwarded_server"
	// FieldRequestID holds the string denoting the request_id field in the database.
	FieldRequestID = "request_id"
	// FieldUserAgent holds the string denoting the user_agent field in the database.
	FieldUserAgent = "user_agent"
	// Table holds the table name of the accesslog in the database.
	Table = "access_logs"
)

// Columns holds all SQL columns for accesslog fields.
var Columns = []string{
	FieldID,
	FieldIP,
	FieldURI,
	FieldForwardedFor,
	FieldForwardedProto,
	FieldForwardedHost,
	FieldForwardedPort,
	FieldForwardedServer,
	FieldRequestID,
	FieldUserAgent,
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

// OrderOption defines the ordering options for the AccessLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIP orders the results by the ip field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByURI orders the results by the uri field.
func ByURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURI, opts...).ToFunc()
}

// ByForwardedFor orders the results by the forwarded_for field.
func ByForwardedFor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForwardedFor, opts...).ToFunc()
}

// ByForwardedProto orders the results by the forwarded_proto field.
func ByForwardedProto(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForwardedProto, opts...).ToFunc()
}

// ByForwardedHost orders the results by the forwarded_host field.
func ByForwardedHost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForwardedHost, opts...).ToFunc()
}

// ByForwardedPort orders the results by the forwarded_port field.
func ByForwardedPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForwardedPort, opts...).ToFunc()
}

// ByForwardedServer orders the results by the forwarded_server field.
func ByForwardedServer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldForwardedServer, opts...).ToFunc()
}

// ByRequestID orders the results by the request_id field.
func ByRequestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestID, opts...).ToFunc()
}

// ByUserAgent orders the results by the user_agent field.
func ByUserAgent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserAgent, opts...).ToFunc()
}
