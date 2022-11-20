package author

import "context"

type service struct {
	storage Storage
}

func NewService(storage interface{}) interface{} {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *CreateAuthorDTO) *Author {
	s.storage.Create()
	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *Author {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit, offset int) []*Author {
	return s.storage.GetAll(limit, offset)
}
