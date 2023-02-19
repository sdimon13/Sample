/**
 * BSD 2-Clause License
 *
 * Copyright (c) 2022, Dmitry Soloma
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this
 *    list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package query_builder

import (
	"fmt"
	"strings"
)

// With SQL
type sqlWith struct {
	keys    map[int]string
	queries []*QB
}

// SQL Pagination limit offset
type sqlPagination struct {
	Limit  int
	Offset int
}

// Len of with queries
func (w *sqlWith) Len() int {
	return len(w.keys)
}

// Query Builder struct
// Not Thread safety
type QB struct {
	with       sqlWith
	columns    []string
	from       []string
	join       []string
	where      Condition
	orders     []string
	group      []string
	union      []*QB
	except     []*QB
	intersect  []*QB
	having     Condition
	pagination sqlPagination
	SubQuery   bool
}

// Add With
func (f *QB) With(name string, qb *QB) *QB {
	if name != "" {
		f.with.queries = append(f.with.queries, qb)
		f.with.keys[len(f.with.queries)-1] = name
	}
	return f
}

// Reset With
func (f *QB) ResetWith() *QB {
	f.with.queries = make([]*QB, 0)
	f.with.keys = make(map[int]string, 0)
	return f
}

// Union
func (f *QB) Union(qb *QB) *QB {
	if qb != nil {
		f.union = append(f.union, qb)
	}
	return f
}

// Except
func (f *QB) Except(qb *QB) *QB {
	if qb != nil {
		f.except = append(f.except, qb)
	}
	return f
}

// Intersect
func (f *QB) Intersect(qb *QB) *QB {
	if qb != nil {
		f.intersect = append(f.intersect, qb)
	}
	return f
}

// Reset Intersect
func (f *QB) ResetIntersect() *QB {
	f.intersect = make([]*QB, 0)
	return f
}

// Reset Union
func (f *QB) ResetUnion() *QB {
	f.union = make([]*QB, 0)
	return f
}

// Reset Except
func (f *QB) ResetExcept() *QB {
	f.except = make([]*QB, 0)
	return f
}

// Get With
func (f *QB) GetWith(name string) *QB {
	for i, key := range f.with.keys {
		if key == name {
			return f.with.queries[i]
		}
	}
	return nil
}

// Add column
func (f *QB) Columns(column ...string) *QB {
	f.columns = append(f.columns, column...)
	return f
}

// Reset column
func (f *QB) ResetColumns() *QB {
	f.columns = []string{}
	return f
}

// Add from
func (f *QB) From(table ...string) *QB {
	f.from = append(f.from, table...)
	return f
}

// Reset column
func (f *QB) ResetFrom() *QB {
	f.from = []string{}
	return f
}

// Add join
func (f *QB) Relate(relation ...string) *QB {
	f.join = append(f.join, relation...)
	return f
}

// Reset join
func (f *QB) ResetRelations() *QB {
	f.join = []string{}
	return f
}

// Where conditions
func (f *QB) Where() *Condition {
	return &f.where
}

// Where conditions
func (f *QB) Having() *Condition {
	return &f.having
}

// Add Order
func (f *QB) AddOrder(expression ...string) *QB {
	f.orders = append(f.orders, expression...)
	return f
}

// Reset Order
func (f *QB) ResetOrder() *QB {
	f.orders = []string{}
	return f
}

// Add Group
func (f *QB) GroupBy(fields ...string) *QB {
	f.group = append(f.group, fields...)
	return f
}

// Reset Group
func (f *QB) ResetGroupBy() *QB {
	f.group = []string{}
	return f
}

// Set pagination
func (f *QB) SetPagination(limit int, offset int) *QB {
	f.pagination = sqlPagination{Limit: limit, Offset: offset}
	return f
}

// Get arguments
func (f *QB) GetArguments() []interface{} {
	arguments := make([]interface{}, 0)
	if f.with.Len() > 0 {
		for _, w := range f.with.queries {
			arguments = append(arguments, w.GetArguments()...)
		}
	}

	arguments = append(arguments, append(f.where.GetArguments(), f.having.GetArguments()...)...)

	if len(f.union) > 0 {
		for _, u := range f.union {
			arguments = append(arguments, u.GetArguments()...)
		}
	}

	if len(f.except) > 0 {
		for _, u := range f.except {
			arguments = append(arguments, u.GetArguments()...)
		}
	}

	if len(f.intersect) > 0 {
		for _, i := range f.intersect {
			arguments = append(arguments, i.GetArguments()...)
		}
	}
	return arguments
}

// Make SQL query
func (f *QB) String() string {
	var result = make([]string, 0)
	var with = make([]string, 0)
	var union = make([]string, 0)
	var except = make([]string, 0)
	var intersect = make([]string, 0)

	// With render
	if f.with.Len() > 0 {
		for index, w := range f.with.queries {
			with = append(with, f.with.keys[index]+" AS ("+w.String()+")")
		}
		result = append(result, "WITH "+strings.Join(with, ", "))
	}

	// Select columns
	if len(f.columns) > 0 {
		result = append(result, "SELECT "+strings.Join(f.columns, ", "))
	}

	// From table
	if len(f.from) > 0 {
		result = append(result, "FROM "+strings.Join(f.from, ", "))
	}

	// From table
	if len(f.join) > 0 {
		result = append(result, strings.Join(f.join, " "))
	}

	// Where conditions
	if len(f.where.expression) > 0 || f.where.merge != nil {
		result = append(result, "WHERE "+f.where.String())
	}

	// Prepare groups
	if len(f.group) > 0 {
		result = append(result, "GROUP BY "+strings.Join(f.group, ", "))
	}

	// Prepare having expression
	if len(f.having.expression) > 0 || f.having.merge != nil {
		result = append(result, "HAVING "+f.having.String())
	}

	// Prepare orders
	if len(f.orders) > 0 {
		result = append(result, "ORDER BY "+strings.Join(f.orders, ", "))
	}

	// Prepare pagination
	if f.pagination.Limit > 0 {
		result = append(result, fmt.Sprintf("LIMIT %v OFFSET %v", f.pagination.Limit, f.pagination.Offset))
	}

	// Union render
	if len(f.union) > 0 {
		for _, u := range f.union {
			union = append(union, u.String())
		}
		result = append(result, "UNION "+strings.Join(union, " UNION "))
	}

	// Except render
	if len(f.except) > 0 {
		for _, u := range f.except {
			except = append(except, u.String())
		}
		result = append(result, "EXCEPT "+strings.Join(except, " EXCEPT "))
	}

	// Intersect render
	if len(f.intersect) > 0 {
		for _, i := range f.intersect {
			intersect = append(intersect, i.String())
		}
		result = append(result, "INTERSECT "+strings.Join(intersect, " INTERSECT "))
	}

	// Check if the query is for sub query
	if f.SubQuery {
		return "(" + strings.Join(result, " ") + ")"
	}

	return preparePositionalArgsQuery(strings.Join(result, " "))
}

// New Query Builder
func NewQB() *QB {
	return &QB{
		with: sqlWith{
			keys:    make(map[int]string),
			queries: make([]*QB, 0),
		},
		where:  Condition{operator: ConditionOperatorAnd},
		having: Condition{operator: ConditionOperatorAnd},
	}
}
