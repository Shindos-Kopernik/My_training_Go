package book

import (
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson8_Clean Arhitecture_Part2/ca-librare-app/internal/domain/book"
	"context"
)

type Service interface { // Этот интерфейс еще можно определить в adapters
	GetByUUID(ctx context.Context, uuid string) *book.Book
	GetAll(ctx context.Context, limit, offset int) []*book.Book
	Create(ctx context.Context, dto *book.CreateBookDTO) *book.Book
}
