package db

import (
	"Lesson1_Rest_API/internal/author"
	"Lesson1_Rest_API/pkg/client/postgresql"
	"Lesson1_Rest_API/pkg/logging"
	"context"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, user author.Author) (string, error) {
	q := 'INSERT INTO author (name) VALUES ($1, $2)'
}

func (r *repository) FindAll(ctx context.Context) (u []author.Author, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindOne(ctx context.Context, id string) (author.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, user author.Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
