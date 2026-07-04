package http

import (
	"context"
	"encoding/json"
	nethttp "net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/asasia1935/walkquest-service/internal/explorer"
)

type fakeExplorerProfileService struct {
	createProfileCalled bool
	createProfileUserID string
	createProfileResult explorer.ExplorerProfile
	createProfileErr    error

	getMyProfileCalled bool
	getMyProfileUserID string
	getMyProfileResult explorer.ExplorerProfile
	getMyProfileErr    error
}

func (f *fakeExplorerProfileService) CreateProfile(ctx context.Context, userID string) (explorer.ExplorerProfile, error) {
	f.createProfileCalled = true
	f.createProfileUserID = userID

	return f.createProfileResult, f.createProfileErr
}

func (f *fakeExplorerProfileService) GetMyProfile(ctx context.Context, userID string) (explorer.ExplorerProfile, error) {
	f.getMyProfileCalled = true
	f.getMyProfileUserID = userID

	return f.getMyProfileResult, f.getMyProfileErr
}

type testErrorResponse struct {
	Error string `json:"error"`
}

type testExplorerProfileResponse struct {
	ID     int64  `json:"id"`
	UserID string `json:"userId"`
}

func TestExplorerProfileHandler_CreateProfile_UnauthorizedWithoutUserID(t *testing.T) {
	handler := NewHandler(NewExplorerProfileHandler(&fakeExplorerProfileService{}))
	request := httptest.NewRequest(nethttp.MethodPost, "/explorer-profiles", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	if recorder.Code != nethttp.StatusUnauthorized {
		t.Fatalf("status code = %d, want %d", recorder.Code, nethttp.StatusUnauthorized)
	}

	var response testErrorResponse
	decodeJSON(t, recorder, &response)

	if response.Error != "unauthorized" {
		t.Fatalf("error = %q, want %q", response.Error, "unauthorized")
	}
}

func TestExplorerProfileHandler_CreateProfile_Success(t *testing.T) {
	userID := "user-123"
	expectedProfile := explorer.ExplorerProfile{
		ID:        10,
		UserID:    userID,
		CreatedAt: time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC),
	}
	service := &fakeExplorerProfileService{createProfileResult: expectedProfile}
	handler := NewHandler(NewExplorerProfileHandler(service))
	request := httptest.NewRequest(nethttp.MethodPost, "/explorer-profiles", nil)
	request.Header.Set(userIDHeader, userID)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	if !service.createProfileCalled {
		t.Fatal("expected service.CreateProfile to be called")
	}

	if service.createProfileUserID != userID {
		t.Fatalf("service.CreateProfile userID = %q, want %q", service.createProfileUserID, userID)
	}

	if recorder.Code != nethttp.StatusCreated {
		t.Fatalf("status code = %d, want %d", recorder.Code, nethttp.StatusCreated)
	}

	var response testExplorerProfileResponse
	decodeJSON(t, recorder, &response)

	if response.ID != expectedProfile.ID {
		t.Fatalf("id = %d, want %d", response.ID, expectedProfile.ID)
	}

	if response.UserID != expectedProfile.UserID {
		t.Fatalf("userId = %q, want %q", response.UserID, expectedProfile.UserID)
	}
}

func TestExplorerProfileHandler_CreateProfile_AlreadyExists(t *testing.T) {
	service := &fakeExplorerProfileService{createProfileErr: explorer.ErrProfileAlreadyExists}
	handler := NewHandler(NewExplorerProfileHandler(service))
	request := httptest.NewRequest(nethttp.MethodPost, "/explorer-profiles", nil)
	request.Header.Set(userIDHeader, "user-123")
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	if recorder.Code != nethttp.StatusConflict {
		t.Fatalf("status code = %d, want %d", recorder.Code, nethttp.StatusConflict)
	}

	var response testErrorResponse
	decodeJSON(t, recorder, &response)

	if response.Error != "explorer_profile_already_exists" {
		t.Fatalf("error = %q, want %q", response.Error, "explorer_profile_already_exists")
	}
}

func TestExplorerProfileHandler_GetMyProfile_Success(t *testing.T) {
	userID := "user-123"
	expectedProfile := explorer.ExplorerProfile{
		ID:        10,
		UserID:    userID,
		CreatedAt: time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC),
	}
	service := &fakeExplorerProfileService{getMyProfileResult: expectedProfile}
	handler := NewHandler(NewExplorerProfileHandler(service))
	request := httptest.NewRequest(nethttp.MethodGet, "/explorer-profiles/me", nil)
	request.Header.Set(userIDHeader, userID)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	if recorder.Code != nethttp.StatusOK {
		t.Fatalf("status code = %d, want %d", recorder.Code, nethttp.StatusOK)
	}

	var response testExplorerProfileResponse
	decodeJSON(t, recorder, &response)

	if response.UserID != expectedProfile.UserID {
		t.Fatalf("userId = %q, want %q", response.UserID, expectedProfile.UserID)
	}
}

func TestExplorerProfileHandler_GetMyProfile_NotFound(t *testing.T) {
	service := &fakeExplorerProfileService{getMyProfileErr: explorer.ErrProfileNotFound}
	handler := NewHandler(NewExplorerProfileHandler(service))
	request := httptest.NewRequest(nethttp.MethodGet, "/explorer-profiles/me", nil)
	request.Header.Set(userIDHeader, "user-123")
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	if recorder.Code != nethttp.StatusNotFound {
		t.Fatalf("status code = %d, want %d", recorder.Code, nethttp.StatusNotFound)
	}

	var response testErrorResponse
	decodeJSON(t, recorder, &response)

	if response.Error != "explorer_profile_not_found" {
		t.Fatalf("error = %q, want %q", response.Error, "explorer_profile_not_found")
	}
}

func decodeJSON(t *testing.T, recorder *httptest.ResponseRecorder, target any) {
	t.Helper()

	if err := json.NewDecoder(recorder.Body).Decode(target); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
}
