package book

import (
	"awesomeProject/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson8_Clean Arhitecture_Part2/ca-librare-app/internal/adapters/api/book"
	"context"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) user.Service {
	return &service{storage: storage}
}

func (s *service) CreateBook(ctx context.Context, dto *CreateBookDTO) *Book {
	s.storage.Create()
	return nil
}

func (s *service) GetBookByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAllBooks(ctx context.Context, limit, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}
