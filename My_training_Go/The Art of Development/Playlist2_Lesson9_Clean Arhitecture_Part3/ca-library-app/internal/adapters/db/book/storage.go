package book

import "awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson9_Clean Arhitecture_Part3/ca-library-app/internal/domain/book"

type bookStorage struct {
}

func NewStorage() book.Storage {
	return &bookStorage{}
}

func (bs *bookStorage) GetOne(uuid string) *book.Book {
	return nil
}
func (bs *bookStorage) GetAll(limit, offset int) []*book.Book {
	return nil
}
func (bs *bookStorage) Create(book *book.Book) *book.Book {
	return nil
}
func (bs *bookStorage) Delete(book *book.Book) error {
	return nil
}
