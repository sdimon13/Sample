# query-builder

## SQL request builder

### Select

Example: Simple query
```
	q := NewQB()
	q.Columns("id", "value")
	q.From("schema.table")

	db.Query(q.String(), q.GetArguments()...)
```

Example: Query with JOIN and filter
```
	q := NewQB()
	q.Columns("ft.id", "ft.value")
	q.From("schema.table ft")
	q.Relate("JOIN other.table st ON ft.some_id = st.id")
	q.Where().AddExpression("ft.id = ?", 10)
	q.Where().AddExpression("ft.deleted_at IS NULL")

	db.Query(q.String(), q.GetArguments()...)
```

Exmple: Query with block "WITH"
```
	testQ := NewQB()
	testQ.Columns("id", "value")
	testQ.From("some.table")
	testQ.Where().AddExpression("deleted_at IS NULL")

	q := NewQB()
	q.With("test_q", testQ)
	q.Columns("ft.id", "ft.value", "tq.value")
	q.From("schema.table ft")
	q.Relate("JOIN test_q tq ON ft.some_id = tq.id")

	db.Query(q.String(), q.GetArguments()...)
```

### Update
```
	q := NewUpdateQB("some.table")
	q.Where().AddExpression("id = ?", 10)
	q.Where().AddExpression("deleted_at IS NULL")
	q.Return("value", "created_by")

	q.Set("value", "home")
	q.SetSql("updated_at", "now()")

	db.Query(q.String(), q.GetArguments()...)
```

### Insert
```
	q := NewInsertQB("some.table")
	q.Return("id")

	q.AddValues(
		NewInsertValues().
			Set("value", "hello").
			SetSql("created_at", "now()").
			SetSql("updated_at", "now()"),
	)
	// You can check the error if you are not sure.
	// An error will be returned if the number of columns in AddValues()
	// is different from the number of columns in the first query,
	// or if the columns have a different order in the first query.
	err := q.AddValues(
		NewInsertValues().
			Set("value", "bye").
			SetSql("created_at", "now()").
			SetSql("updated_at", "now()"),
	)

	db.Query(q.String(), q.GetArguments()...)
```

### Raw
```
	q := NewRawQB()
	q.Append("SELECT")
	q.AddList((&RawList{}).
		Add("id").
		Add("value"),
	)
	q.Append("FROM some.table")
	q.Append("WHERE")
	q.AddAndList((&RawList{}).
		Add("id = ?", 10).
		Add("deleted_at IS NULL"),
	)

	db.Query(ctx, q.String(), q.GetArguments()...)
```

Or without RawList:
```
	q := NewRawQB()
	q.Append("SELECT id, value")
	q.Append("FROM some.table")
	q.Append("WHERE")
	q.Append("id = ?", 10)
	q.Append("AND deleted_at IS NULL")

	db.Query(ctx, q.String(), q.GetArguments()...)
```
