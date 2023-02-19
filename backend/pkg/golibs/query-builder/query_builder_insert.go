package query_builder

import (
	"fmt"
	"strings"
)

// Query Builder struct
// Not Thread safety
type QBI struct {
	returns []string
	table   string
	values  []*InsertValues
}

// Add return
func (q *QBI) Return(column ...string) *QBI {
	q.returns = append(q.returns, column...)
	return q
}

func (q *QBI) AddValues(values *InsertValues) error {
	if len(q.values) > 0 {
		if len(q.values[0].fields) != len(values.fields) {
			return fmt.Errorf("incorrect number of fields for insert \n0 element:(%s) \n %d element:(%s)",
				strings.Join(q.values[0].fields, ", "),
				len(q.values),
				strings.Join(values.fields, ", "),
			)
		}
		for i, v := range values.fields {
			if v != q.values[0].fields[i] {
				return fmt.Errorf("incorrect column order for insert \n0 element:(%s) \n %d element:(%s)",
					strings.Join(q.values[0].fields, ", "),
					len(q.values),
					strings.Join(values.fields, ", "),
				)
			}
		}
	}
	q.values = append(q.values, values)
	return nil
}

// Get arguments
func (q *QBI) GetArguments() []interface{} {
	var arguments []interface{}
	for _, v := range q.values {
		arguments = append(arguments, v.arguments...)
	}
	return arguments
}

func (q *QBI) getInsertForSql() string {
	var column []string
	var fullSet []string

	for _, value := range q.values[0].fields {
		column = append(column, value)
	}
	for _, v := range q.values {
		fullSet = append(fullSet, "("+strings.Join(v.set, ", ")+")")
	}
	return " (" + strings.Join(column, ", ") + ") VALUES " +
		strings.Join(fullSet, ", ")
}

// Make SQL query
func (q *QBI) String() string {
	var result = make([]string, 0)

	// Insert table
	if len(q.values) > 0 {
		result = append(result, "INSERT INTO "+q.table+q.getInsertForSql())
	}

	// Select columns
	if len(q.returns) > 0 {
		result = append(result, "RETURNING "+strings.Join(q.returns, ", "))
	}

	return preparePositionalArgsQuery(strings.Join(result, " "))
}

type InsertValues struct {
	fields    []string
	set       []string
	arguments []interface{}
}

func (iv *InsertValues) Set(field string, value interface{}) *InsertValues {
	iv.fields = append(iv.fields, field)
	iv.set = append(iv.set, "?")
	iv.arguments = append(iv.arguments, value)
	return iv
}

func (iv *InsertValues) SetSql(field string, value interface{}) *InsertValues {
	iv.fields = append(iv.fields, field)
	iv.set = append(iv.set, value.(string))
	return iv
}

// Get arguments
func (iv *InsertValues) GetArguments() []interface{} {
	return iv.arguments
}

// New list values
func NewInsertValues() *InsertValues {
	return &InsertValues{}
}

// New Query Builder
func NewInsertQB(table string) *QBI {
	return &QBI{
		table: table,
	}
}
