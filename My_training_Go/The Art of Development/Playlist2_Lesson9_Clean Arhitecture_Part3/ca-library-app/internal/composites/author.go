package composites

import (
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson9_Clean Arhitecture_Part3/ca-library-app/internal/adapters/api"
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson9_Clean Arhitecture_Part3/ca-library-app/internal/adapters/api/author"
	author2 "awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson9_Clean Arhitecture_Part3/ca-library-app/internal/domain/author"
)

type AuthorComposite struct {
	Storage author2.Storage
	Service author.Service
	Handler api.Handler
}

func NewAuthorComposite() (*AuthorComposite, error) {
	author3.NewStorage
	service := author2.NewService(storage)
	handler := author.NewHandler(service)
}
