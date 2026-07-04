package explorer

import (
	"context"
	"errors"
	"testing"
	"time"
)

type fakeProfileRepository struct {
	createCalled bool
	createUserID string
	createResult ExplorerProfile
	createErr    error

	findByUserIDCalled bool
	findByUserID       string
	findByUserIDResult ExplorerProfile
	findByUserIDErr    error
}

func (f *fakeProfileRepository) Create(ctx context.Context, userID string) (ExplorerProfile, error) {
	f.createCalled = true
	f.createUserID = userID

	return f.createResult, f.createErr
}

func (f *fakeProfileRepository) FindByUserID(ctx context.Context, userID string) (ExplorerProfile, error) {
	f.findByUserIDCalled = true
	f.findByUserID = userID

	return f.findByUserIDResult, f.findByUserIDErr
}

func TestService_CreateProfile_Success(t *testing.T) {
	ctx := context.Background()
	userID := "user-123"
	expectedProfile := ExplorerProfile{
		ID:        1,
		UserID:    userID,
		CreatedAt: time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC),
	}
	repository := &fakeProfileRepository{createResult: expectedProfile}
	service := NewService(repository)

	profile, err := service.CreateProfile(ctx, userID)
	if err != nil {
		t.Fatalf("CreateProfile returned unexpected error: %v", err)
	}

	if !repository.createCalled {
		t.Fatal("expected repository.Create to be called")
	}

	if repository.createUserID != userID {
		t.Fatalf("repository.Create userID = %q, want %q", repository.createUserID, userID)
	}

	if profile != expectedProfile {
		t.Fatalf("CreateProfile returned profile = %+v, want %+v", profile, expectedProfile)
	}
}

func TestService_CreateProfile_AlreadyExists(t *testing.T) {
	ctx := context.Background()
	userID := "user-123"
	repository := &fakeProfileRepository{createErr: ErrProfileAlreadyExists}
	service := NewService(repository)

	_, err := service.CreateProfile(ctx, userID)
	if !errors.Is(err, ErrProfileAlreadyExists) {
		t.Fatalf("CreateProfile error = %v, want %v", err, ErrProfileAlreadyExists)
	}
}

func TestService_GetMyProfile_Success(t *testing.T) {
	ctx := context.Background()
	userID := "user-123"
	expectedProfile := ExplorerProfile{
		ID:        1,
		UserID:    userID,
		CreatedAt: time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC),
	}
	repository := &fakeProfileRepository{findByUserIDResult: expectedProfile}
	service := NewService(repository)

	profile, err := service.GetMyProfile(ctx, userID)
	if err != nil {
		t.Fatalf("GetMyProfile returned unexpected error: %v", err)
	}

	if !repository.findByUserIDCalled {
		t.Fatal("expected repository.FindByUserID to be called")
	}

	if repository.findByUserID != userID {
		t.Fatalf("repository.FindByUserID userID = %q, want %q", repository.findByUserID, userID)
	}

	if profile != expectedProfile {
		t.Fatalf("GetMyProfile returned profile = %+v, want %+v", profile, expectedProfile)
	}
}

func TestService_GetMyProfile_NotFound(t *testing.T) {
	ctx := context.Background()
	userID := "user-123"
	repository := &fakeProfileRepository{findByUserIDErr: ErrProfileNotFound}
	service := NewService(repository)

	_, err := service.GetMyProfile(ctx, userID)
	if !errors.Is(err, ErrProfileNotFound) {
		t.Fatalf("GetMyProfile error = %v, want %v", err, ErrProfileNotFound)
	}
}
