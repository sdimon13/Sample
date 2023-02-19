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

import "strings"

const (
	ConditionOperatorAnd = "AND"
	ConditionOperatorOr  = "OR"
	ConditionOperatorXor = "XOR"
)

// Merge with Condition
type merge struct {
	operator  string
	condition []*Condition
}

// Condition type
type Condition struct {
	operator   string
	expression []string
	argument   []interface{}
	merge      *merge
}

// Init Condition
func NewSqlCondition(operator string) *Condition {
	return &Condition{operator: operator}
}

// Get string of conditions
func (c *Condition) String() string {
	if c.merge != nil {
		var slaves []string
		for i := range (*c.merge).condition {
			slaves = append(slaves, (*c.merge).condition[i].String())
		}
		if c.expression != nil {
			slaves = append(slaves, "("+strings.Join(c.expression, " "+c.operator+" ")+")")
		}
		return "(" + strings.Join(slaves, " "+c.merge.operator+" ") + ")"
	} else {
		if c.expression != nil {
			return "(" + strings.Join(c.expression, " "+c.operator+" ") + ")"
		}
		return ""
	}
}

// Check if condition is empty
func (c *Condition) IsEmpty() bool {
	return c == nil || (len(c.expression) == 0 && c.merge == nil)
}

// Get arguments
func (c *Condition) GetArguments() []interface{} {
	var arguments = make([]interface{}, 0)
	if c.merge != nil {
		for i := range (*c.merge).condition {
			arguments = append(arguments, (*c.merge).condition[i].GetArguments()...)
		}
	}
	return append(arguments, c.argument...)
}

// Add expression
func (c *Condition) AddExpression(expression string, values ...interface{}) *Condition {
	c.expression = append(c.expression, expression)
	c.argument = append(c.argument, values...)
	return c
}

// Add argument
func (c *Condition) AddArgument(values ...interface{}) *Condition {
	c.argument = append(c.argument, values...)
	return c
}

// Merge with conditions
func (c *Condition) Merge(operator string, conditions ...*Condition) *Condition {
	if len(conditions) > 0 {
		for i := range conditions {
			if conditions[i] == nil {
				continue
			}
			if c.merge == nil {
				c.merge = &merge{operator: operator, condition: []*Condition{}}
			}
			c.merge.condition = append(c.merge.condition, conditions[i])
		}
	}
	return c
}
