package repository

import (
	"context"
	"git.sample.ru/sample/internal/db"
	"git.sample.ru/sample/internal/entity"
	"git.sample.ru/sample/internal/logger"
	qb "git.sample.ru/sample/pkg/golibs/query-builder"
	"github.com/georgysavva/scany/pgxscan"
	"sync"
)

type Client struct {
	db *db.DB
}

type ClientInterface interface {
	Get(ctx context.Context, id int64) (*entity.Client, error)
	Find(ctx context.Context, phone string, name string) (*entity.Client, error)
	List(ctx context.Context) (*[]entity.Client, error)
	Add(ctx context.Context, e *entity.Client) (*entity.Client, error)
	Update(ctx context.Context, e *entity.Client) (*entity.Client, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

func NewClient(d *db.DB) *Client {
	return &Client{db: d}
}

func (r *Client) Get(ctx context.Context, id int64) (*entity.Client, error) {
	eu := entity.Client{}

	q := qb.NewQB().
		Columns("id", "phone", "name").
		From(eu.GetTable())
	q.Where().AddExpression("deleted_at IS NULL").
		AddExpression("id = ?", id)

	err := pgxscan.Get(ctx, r.db.Client, &eu, q.String(), q.GetArguments()...)
	if err != nil {
		return nil, err
	}

	return &eu, nil
}

func (r *Client) Find(ctx context.Context, phone string, name string) (*entity.Client, error) {
	eu := entity.Client{}

	q := qb.NewQB().
		Columns("id", "phone", "name").
		From(eu.GetTable())
	q.Where().AddExpression("deleted_at IS NULL").
		AddExpression("phone = ?", phone).
		AddExpression("name = ?", name)

	err := pgxscan.Get(ctx, r.db.Client, &eu, q.String(), q.GetArguments()...)
	if err != nil {
		return nil, err
	}

	return &eu, nil
}

func (r *Client) List(ctx context.Context) ([]*entity.Client, error) {
	var el []*entity.Client
	t := entity.Client{}.GetTable()

	q := qb.NewQB().
		Columns("id", "phone", "name").
		From(t)
	q.Where().AddExpression("deleted_at IS NULL")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := pgxscan.Select(ctx, r.db.Client, &el, q.String(), q.GetArguments()...)
		if err != nil {
			logger.Error.Println("error on Client::get", err)
		}
	}()

	wg.Wait()

	return el, nil
}

func (r *Client) Add(ctx context.Context, e *entity.Client) (*entity.Client, error) {
	err := r.db.Client.
		QueryRow(ctx, "INSERT INTO samples.clients(phone, name, created_at, updated_at) VALUES ($1, $2, now(), now()) RETURNING id", e.Phone, e.Name).
		Scan(&e.Id)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *Client) Update(ctx context.Context, e *entity.Client) (*entity.Client, error) {
	if _, err := r.Get(ctx, e.Id); err != nil {
		return nil, err
	}

	_, err := r.db.Client.
		Query(
			ctx, `UPDATE samples.clients SET phone = $2, name = $3, updated_at = now() WHERE id = $1`, e.Id, e.Phone, e.Name)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *Client) Delete(ctx context.Context, id int64) (bool, error) {
	if _, err := r.Get(ctx, id); err != nil {
		return false, err
	}

	_, err := r.db.Client.
		Query(ctx, "UPDATE samples.clients SET deleted_at = now() WHERE id = $1", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
