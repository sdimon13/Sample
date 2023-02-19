package query_builder

import (
	"strings"

	"git.sample.ru/sample/pkg/golibs/helpy/v2"
)

const defaultPage = 1
const defaultPerPage = 10

func FiltersToQuery(q *QB, col string, f *helpy.FilterItem) {
	s := f.GetValue()
	s = strings.Trim(s, " ")
	switch f.GetOp() {
	case helpy.Operation_eq:
		q.Where().AddExpression(col+" = ?", s)
	case helpy.Operation_neq:
		q.Where().AddExpression(col+" != ?", s)
	case helpy.Operation_gt:
		q.Where().AddExpression(col+" > ?", s)
	case helpy.Operation_gte:
		q.Where().AddExpression(col+" >= ?", s)
	case helpy.Operation_lt:
		q.Where().AddExpression(col+" < ?", s)
	case helpy.Operation_lte:
		q.Where().AddExpression(col+" <= ?", s)
	case helpy.Operation_begins:
		q.Where().AddExpression(col+" ILIKE ?", s+"%")
	case helpy.Operation_contains:
		q.Where().AddExpression(col+" ILIKE ?", "%"+s+"%")
	case helpy.Operation_equal:
		q.Where().AddExpression(col+" = ?", s)
	}
}

func FiltersMultiToQuery(q *QB, col string, f *helpy.FilterItemMulti) {
	s := f.GetValues()

	switch f.GetOp() {
	case helpy.OperationMulti_in:
		q.Where().AddExpression(col+" = ANY(?)", s)
	case helpy.OperationMulti_nin:
		q.Where().AddExpression(col+" != ALL (?)", s)
	}
}

func Sort(q *QB, s *helpy.FilterSort) {
	if s != nil && len(s.GetField()) != 0 {
		dir := " ASC"
		if s.GetDir() == helpy.SortDirection_desc {
			dir = " DESC"
		}
		q.AddOrder(s.GetField() + dir)
	}
}

// Page In case you want to get a complete list of items without pagination, you can pass a value of -1 to the perPage field
func Page(q *QB, p *helpy.FilterPage) {
	page := p.GetPage()
	perPage := p.GetPerPage()

	if perPage == -1 {
		// If we get -1 in the perPage field, do not apply the default pagination
		return
	}

	if page <= 0 {
		page = defaultPage
	}
	if perPage <= 0 {
		perPage = defaultPerPage
	}

	q.SetPagination(int(perPage), int((page-1)*perPage))
}
