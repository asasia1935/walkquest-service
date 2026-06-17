package http

import (
	"context"
	"errors"
	nethttp "net/http"
	"strings"
	"time"

	"github.com/asasia1935/walkquest-service/internal/explorer"
)

const userIDHeader = "X-User-Id"

type explorerProfileService interface {
	CreateProfile(ctx context.Context, userID string) (explorer.ExplorerProfile, error)
	GetMyProfile(ctx context.Context, userID string) (explorer.ExplorerProfile, error)
}

type ExplorerProfileHandler struct {
	service explorerProfileService
}

type explorerProfileResponse struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewExplorerProfileHandler(service explorerProfileService) *ExplorerProfileHandler {
	return &ExplorerProfileHandler{
		service: service,
	}
}

func (h *ExplorerProfileHandler) CreateProfile(w nethttp.ResponseWriter, r *nethttp.Request) {
	userID, ok := userIDFromRequest(r)
	if !ok {
		writeError(w, nethttp.StatusUnauthorized, "unauthorized")
		return
	}

	profile, err := h.service.CreateProfile(r.Context(), userID)
	if err != nil {
		writeExplorerProfileError(w, err)
		return
	}

	writeJSON(w, nethttp.StatusCreated, toExplorerProfileResponse(profile))
}

func (h *ExplorerProfileHandler) GetMyProfile(w nethttp.ResponseWriter, r *nethttp.Request) {
	userID, ok := userIDFromRequest(r)
	if !ok {
		writeError(w, nethttp.StatusUnauthorized, "unauthorized")
		return
	}

	profile, err := h.service.GetMyProfile(r.Context(), userID)
	if err != nil {
		writeExplorerProfileError(w, err)
		return
	}

	writeJSON(w, nethttp.StatusOK, toExplorerProfileResponse(profile))
}

func userIDFromRequest(r *nethttp.Request) (string, bool) {
	userID := strings.TrimSpace(r.Header.Get(userIDHeader))
	if userID == "" {
		return "", false
	}

	return userID, true
}

func writeExplorerProfileError(w nethttp.ResponseWriter, err error) {
	switch {
	case errors.Is(err, explorer.ErrProfileAlreadyExists):
		writeError(w, nethttp.StatusConflict, "explorer_profile_already_exists")
	case errors.Is(err, explorer.ErrProfileNotFound):
		writeError(w, nethttp.StatusNotFound, "explorer_profile_not_found")
	default:
		writeError(w, nethttp.StatusInternalServerError, "internal_server_error")
	}
}

func toExplorerProfileResponse(profile explorer.ExplorerProfile) explorerProfileResponse {
	return explorerProfileResponse{
		ID:        profile.ID,
		UserID:    profile.UserID,
		CreatedAt: profile.CreatedAt,
		UpdatedAt: profile.UpdatedAt,
	}
}
