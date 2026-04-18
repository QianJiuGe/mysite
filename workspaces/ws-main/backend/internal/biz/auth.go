package biz

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/QianJiuGe/mysite-back/internal/data"
	"github.com/QianJiuGe/mysite-back/internal/model"
)

var (
	ErrValidation      = errors.New("VALIDATION_ERROR")
	ErrUnauthorized    = errors.New("UNAUTHORIZED")
	ErrPendingApproval = errors.New("FORBIDDEN_PENDING_APPROVAL")
	ErrAdminOnly       = errors.New("FORBIDDEN_ADMIN_ONLY")
	ErrNotFound        = errors.New("NOT_FOUND")
	ErrInternal        = errors.New("INTERNAL_ERROR")
	ErrAlreadyExists   = errors.New("already exists")
)

type Usecase struct {
	store *data.Store
}

func NewUsecase(store *data.Store) *Usecase {
	return &Usecase{store: store}
}

func (u *Usecase) BootstrapAdmin(ctx context.Context, username, password string) error {
	h, err := hashPassword(password)
	if err != nil {
		return err
	}
	return u.store.EnsureDefaultAdmin(ctx, username, h)
}

func (u *Usecase) Register(ctx context.Context, username, password string) (int64, string, error) {
	username = strings.TrimSpace(username)
	if len(username) < 3 || len(username) > 32 || len(password) < 8 || len(password) > 64 {
		return 0, "", ErrValidation
	}
	if _, err := u.store.GetUserByUsername(ctx, username); err == nil {
		return 0, "", ErrAlreadyExists
	} else if err != nil && !errors.Is(err, data.ErrNotFound) {
		return 0, "", ErrInternal
	}

	h, err := hashPassword(password)
	if err != nil {
		return 0, "", ErrInternal
	}
	id, err := u.store.CreatePendingUser(ctx, username, h)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return 0, "", ErrAlreadyExists
		}
		return 0, "", ErrInternal
	}
	return id, "pending", nil
}

func (u *Usecase) Login(ctx context.Context, username, password string) (string, string, error) {
	usr, err := u.store.GetUserByUsername(ctx, strings.TrimSpace(username))
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return "", "", ErrUnauthorized
		}
		return "", "", ErrInternal
	}
	if err := bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(password)); err != nil {
		return "", "", ErrUnauthorized
	}
	if usr.Status == "pending" {
		return "", "", ErrPendingApproval
	}
	if usr.Status != "approved" {
		return "", "", ErrUnauthorized
	}

	token, err := generateToken()
	if err != nil {
		return "", "", ErrInternal
	}
	if err := u.store.SaveSession(ctx, token, model.Session{
		UserID: usr.ID, Username: usr.Username, Role: usr.Role, Status: usr.Status,
	}, 7*24*time.Hour); err != nil {
		return "", "", ErrInternal
	}
	return token, usr.Role, nil
}

func (u *Usecase) GetSessionByToken(ctx context.Context, token string) (*model.Session, error) {
	sess, err := u.store.GetSession(ctx, token)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return nil, ErrUnauthorized
		}
		return nil, ErrInternal
	}
	return sess, nil
}

func (u *Usecase) ListPendingUsers(ctx context.Context, current *model.Session) ([]model.User, error) {
	if current == nil || current.Role != "admin" {
		return nil, ErrAdminOnly
	}
	users, err := u.store.ListUsersByStatus(ctx, "pending")
	if err != nil {
		return nil, ErrInternal
	}
	return users, nil
}

func (u *Usecase) Approve(ctx context.Context, current *model.Session, targetUserID int64) error {
	if current == nil || current.Role != "admin" {
		return ErrAdminOnly
	}
	if err := u.store.ApproveUser(ctx, targetUserID); err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return ErrNotFound
		}
		return ErrInternal
	}
	return nil
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}
	return string(b), nil
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
