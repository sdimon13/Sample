package query_builder

import (
	"strings"
)

// Query Builder struct
// Not Thread safety
type QBU struct {
	returns   []string
	table     string
	arguments []interface{}
	fields    []string
	set       []string
	where     Condition
}

// Add return
func (q *QBU) Return(column ...string) *QBU {
	q.returns = append(q.returns, column...)
	return q
}

// Get arguments
func (q *QBU) GetArguments() []interface{} {
	var arguments []interface{}
	arguments = append(arguments, q.arguments...)
	arguments = append(arguments, q.where.GetArguments()...)
	return arguments
}

func (q *QBU) getUpdateForSql() string {
	var set []string

	for i, value := range q.fields {
		set = append(set, value+" = "+q.set[i])
	}
	return strings.Join(set, ", ")
}

// Make SQL query
func (q *QBU) String() string {
	var result = make([]string, 0)

	// Update table
	if len(q.fields) > 0 {
		result = append(result, "UPDATE "+q.table+" SET "+q.getUpdateForSql())
	}

	// Where conditions
	if len(q.where.expression) > 0 || q.where.merge != nil {
		result = append(result, "WHERE "+q.where.String())
	}

	// Return columns
	if len(q.returns) > 0 {
		result = append(result, "RETURNING "+strings.Join(q.returns, ", "))
	}

	return preparePositionalArgsQuery(strings.Join(result, " "))
}

// Where conditions
func (q *QBU) Where() *Condition {
	return &q.where
}

// New Query Builder
func NewUpdateQB(table string) *QBU {
	return &QBU{
		where: Condition{operator: ConditionOperatorAnd},
		table: table,
	}
}

func (q *QBU) Set(field string, value interface{}) *QBU {
	q.fields = append(q.fields, field)
	q.set = append(q.set, "?")
	q.arguments = append(q.arguments, value)
	return q
}

func (q *QBU) SetSql(field string, value interface{}) *QBU {
	q.fields = append(q.fields, field)
	q.set = append(q.set, value.(string))
	return q
}
