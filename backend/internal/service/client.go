package service

import (
	"context"
	"git.sample.ru/sample/internal/entity"
	"git.sample.ru/sample/internal/repository"
)

type ClientService struct {
	r *repository.Client
}

type ClientInterface interface {
	GetClient(ctx context.Context, id int64) (*entity.Client, error)
	FindClient(ctx context.Context, phone string, name string) (*entity.Client, error)
	GetClientList(ctx context.Context) (*[]entity.Client, error)
	AddClient(ctx context.Context, e *entity.Client) (*entity.Client, error)
	UpdateClient(ctx context.Context, e *entity.Client) (*entity.Client, error)
	DeleteClient(ctx context.Context, id int64) (bool, error)
}

func NewClientService(r *repository.Client) *ClientService {
	return &ClientService{r}
}

func (s *ClientService) GetClient(ctx context.Context, id int64) (*entity.Client, error) {
	return s.r.Get(ctx, id)
}

func (s *ClientService) FindClient(ctx context.Context, phone string, name string) (*entity.Client, error) {
	return s.r.Find(ctx, phone, name)
}

func (s *ClientService) GetClientList(ctx context.Context) ([]*entity.Client, error) {
	return s.r.List(ctx)
}

func (s *ClientService) AddClient(ctx context.Context, e *entity.Client) (*entity.Client, error) {
	return s.r.Add(ctx, e)
}

func (s *ClientService) UpdateClient(ctx context.Context, e *entity.Client) (*entity.Client, error) {
	return s.r.Update(ctx, e)
}

func (s *ClientService) DeleteClient(ctx context.Context, id int64) (bool, error) {
	return s.r.Delete(ctx, id)
}
