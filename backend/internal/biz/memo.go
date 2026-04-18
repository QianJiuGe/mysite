package biz

import (
	"context"

	"github.com/QianJiuGe/mysite/backend/internal/data"
	"github.com/QianJiuGe/mysite/backend/internal/model"
)

type MemoUsecase struct {
	store *data.Store
}

func NewMemoUsecase(store *data.Store) *MemoUsecase {
	return &MemoUsecase{store: store}
}

func (m *MemoUsecase) List(ctx context.Context, userID int64) ([]model.Memo, error) {
	return m.store.ListMemos(ctx, userID)
}

func (m *MemoUsecase) ListAll(ctx context.Context) ([]model.Memo, error) {
	return m.store.ListAllMemos(ctx)
}

func (m *MemoUsecase) Create(ctx context.Context, userID int64, content string, priority int) (*model.Memo, error) {
	if content == "" {
		return nil, ErrValidation
	}
	return m.store.CreateMemo(ctx, userID, content, priority)
}

func (m *MemoUsecase) Update(ctx context.Context, userID, memoID int64, isAdmin bool, content *string, done *bool, priority *int) error {
	memo, err := m.store.GetMemo(ctx, memoID)
	if err != nil {
		return err
	}
	if !isAdmin && memo.UserID != userID {
		return ErrAdminOnly
	}
	if content != nil {
		memo.Content = *content
	}
	if done != nil {
		memo.Done = *done
	}
	if priority != nil {
		memo.Priority = *priority
	}
	return m.store.UpdateMemo(ctx, memo)
}

func (m *MemoUsecase) Delete(ctx context.Context, userID, memoID int64, isAdmin bool) error {
	memo, err := m.store.GetMemo(ctx, memoID)
	if err != nil {
		return err
	}
	if !isAdmin && memo.UserID != userID {
		return ErrAdminOnly
	}
	return m.store.DeleteMemo(ctx, memoID)
}
