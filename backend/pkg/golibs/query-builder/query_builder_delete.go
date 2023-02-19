package query_builder

import (
	"strings"
)

// Query Builder struct
// Not Thread safety
type QBD struct {
	returns []string
	table   string
	where   Condition
}

// Add return
func (q *QBD) Return(column ...string) *QBD {
	q.returns = append(q.returns, column...)
	return q
}

// Get arguments
func (q *QBD) GetArguments() []interface{} {
	var arguments []interface{}
	arguments = append(arguments, q.where.GetArguments()...)
	return arguments
}

// Make SQL query
func (q *QBD) String() string {
	var result = make([]string, 0)

	result = append(result, "DELETE FROM "+q.table)

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
func (q *QBD) Where() *Condition {
	return &q.where
}

// New Query Builder
func NewDeleteQB(table string) *QBD {
	return &QBD{
		where: Condition{operator: ConditionOperatorAnd},
		table: table,
	}
}
