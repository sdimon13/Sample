package query_builder

import "strings"

type RawList struct {
	list      []string
	arguments []interface{}
}

func (l *RawList) Add(value string, arguments ...interface{}) *RawList {
	l.list = append(l.list, value)
	l.arguments = append(l.arguments, arguments...)
	return l
}

// Query Builder struct
// Not Thread safety
type RawQB struct {
	query    []string
	argument []interface{}
}

func (f *RawQB) Append(query string, arguments ...interface{}) *RawQB {
	f.query = append(f.query, query)
	f.argument = append(f.argument, arguments...)
	return f
}

func (f *RawQB) AddList(list *RawList) *RawQB {
	f.query = append(f.query, strings.Join(list.list, ", "))
	f.argument = append(f.argument, list.arguments...)
	return f
}
func (f *RawQB) AddAndList(list *RawList) *RawQB {
	f.query = append(f.query, strings.Join(list.list, " AND "))
	f.argument = append(f.argument, list.arguments...)
	return f
}
func (f *RawQB) AddOrList(list *RawList) *RawQB {
	f.query = append(f.query, strings.Join(list.list, " OR "))
	f.argument = append(f.argument, list.arguments...)
	return f
}

func (f *RawQB) GetArguments() []interface{} {
	return f.argument
}
func (f *RawQB) String() string {
	return preparePositionalArgsQuery(strings.Join(f.query, " "))
}

func NewRawQB() *RawQB {
	return &RawQB{}
}
