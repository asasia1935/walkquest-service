package explorer

import "context"

type profileRepository interface {
	Create(ctx context.Context, userID string) (ExplorerProfile, error)
	FindByUserID(ctx context.Context, userID string) (ExplorerProfile, error)
}

type Service struct {
	repository profileRepository
}

func NewService(repository profileRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateProfile(ctx context.Context, userID string) (ExplorerProfile, error) {
	return s.repository.Create(ctx, userID)
}

func (s *Service) GetMyProfile(ctx context.Context, userID string) (ExplorerProfile, error) {
	return s.repository.FindByUserID(ctx, userID)
}
