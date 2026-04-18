package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/QianJiuGe/mysite-back/internal/biz"
	"github.com/QianJiuGe/mysite-back/internal/model"
)

type Service struct {
	uc *biz.Usecase
}

func New(uc *biz.Usecase) *Service {
	return &Service{uc: uc}
}

func (s *Service) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "VALIDATION_ERROR", "method not allowed")
		return
	}
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "invalid request body")
		return
	}
	id, status, err := s.uc.Register(r.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, biz.ErrValidation) {
			writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "invalid username or password")
			return
		}
		if errors.Is(err, biz.ErrAlreadyExists) {
			writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "username already exists")
			return
		}
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "register failed")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"userId": id, "status": status})
}

func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "VALIDATION_ERROR", "method not allowed")
		return
	}
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "invalid request body")
		return
	}
	token, role, err := s.uc.Login(r.Context(), req.Username, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, biz.ErrPendingApproval):
			writeError(w, http.StatusForbidden, "FORBIDDEN_PENDING_APPROVAL", "account pending approval")
		case errors.Is(err, biz.ErrUnauthorized):
			writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "invalid credentials")
		default:
			writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "login failed")
		}
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"token": token, "role": role})
}

func (s *Service) PendingUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "VALIDATION_ERROR", "method not allowed")
		return
	}
	current, err := s.currentSession(r.Context(), r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token")
		return
	}
	users, err := s.uc.ListPendingUsers(r.Context(), current)
	if err != nil {
		switch {
		case errors.Is(err, biz.ErrAdminOnly):
			writeError(w, http.StatusForbidden, "FORBIDDEN_ADMIN_ONLY", "admin only")
		default:
			writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "list pending users failed")
		}
		return
	}

	type pendingUser struct {
		UserID   int64  `json:"userId"`
		Username string `json:"username"`
		Status   string `json:"status"`
	}
	resp := make([]pendingUser, 0, len(users))
	for _, u := range users {
		resp = append(resp, pendingUser{UserID: u.ID, Username: u.Username, Status: u.Status})
	}
	writeJSON(w, http.StatusOK, map[string]any{"users": resp})
}

func (s *Service) ApproveUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "VALIDATION_ERROR", "method not allowed")
		return
	}
	current, err := s.currentSession(r.Context(), r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token")
		return
	}

	// expected path: /v1/admin/users/{userId}/approve
	path := strings.TrimPrefix(r.URL.Path, "/v1/admin/users/")
	parts := strings.Split(path, "/")
	if len(parts) != 2 || parts[1] != "approve" {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "invalid path")
		return
	}
	userID, convErr := strconv.ParseInt(parts[0], 10, 64)
	if convErr != nil || userID <= 0 {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "invalid userId")
		return
	}

	if err := s.uc.Approve(r.Context(), current, userID); err != nil {
		switch {
		case errors.Is(err, biz.ErrAdminOnly):
			writeError(w, http.StatusForbidden, "FORBIDDEN_ADMIN_ONLY", "admin only")
		case errors.Is(err, biz.ErrNotFound):
			writeError(w, http.StatusNotFound, "NOT_FOUND", "user not found")
		default:
			writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "approve failed")
		}
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"userId": userID, "status": "approved"})
}

func (s *Service) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "VALIDATION_ERROR", "method not allowed")
		return
	}
	current, err := s.currentSession(r.Context(), r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token")
		return
	}
	if current.Status != "approved" {
		writeError(w, http.StatusForbidden, "FORBIDDEN_PENDING_APPROVAL", "account pending approval")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"welcome": "Welcome, " + current.Username,
		"modules": []string{"profile-card", "feed-placeholder", "announcement"},
	})
}

func (s *Service) Healthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"status": "ok"})
}

func (s *Service) currentSession(ctx context.Context, r *http.Request) (*model.Session, error) {
	auth := strings.TrimSpace(r.Header.Get("Authorization"))
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		return nil, biz.ErrUnauthorized
	}
	token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
	if token == "" {
		return nil, biz.ErrUnauthorized
	}
	return s.uc.GetSessionByToken(ctx, token)
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, code, msg string) {
	writeJSON(w, status, map[string]any{"code": code, "message": msg})
}
