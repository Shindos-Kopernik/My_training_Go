package user

import (
	"awesomeProject/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson8_Clean Arhitecture_Part2/ca-librare-app/internal/domain/book"
	"context"
)

type Service interface { // Этот интерфейс еще можно определить в adapters
	GetBookByUUID(ctx context.Context, uuid string) *book.Book
	GetAllBooks(ctx context.Context, limit, offset int) []*book.Book
	CreateBook(ctx context.Context, dto *book.CreateBookDTO) *book.Book
}
