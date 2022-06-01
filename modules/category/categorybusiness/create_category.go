package categorybusiness

import (
	"context"
	"github.com/imperiutx/nan_forum/common"
	db "github.com/imperiutx/nan_forum/db/sqlc"
)

type CreateCategoryStorage interface {
	FindCategory(
		ctx context.Context,
		name string,
		moreInfo ...string) (*db.Category, error)
	CreateCategory(ctx context.Context, params db.CreateCategoryParams) error
}

type createCategory struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategory {
	return &createCategory{
		store: store,
	}
}

func (biz *createCategory) CreateNewCategory(ctx context.Context, data db.CreateCategoryParams) error {
	category, err := biz.store.FindCategory(ctx, data.Name)
	if category != nil {
		return common.ErrEntityExisted("category", err)
	}

	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("category", err)
	}

	return nil
}
