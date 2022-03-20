package reader

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/user"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func NewUserReader(db *ent.Client) *UserReader {
	return &UserReader{
		db: db,
	}
}

//go:generate mockgen -source=./user.go -destination=./mock/user_mock.go -package=reader
type User interface {
	GetByID(ctx context.Context, userID int) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
}

type UserReader struct {
	db *ent.Client
}

var _ User = (*UserReader)(nil)

func (impl *UserReader) GetByID(ctx context.Context, userID int) (model.User, error) {
	user, err := impl.db.User.Query().Where(user.IDEQ(int32(userID))).Only(ctx)
	if err != nil {
		return model.User{}, ent.MaskNotFound(err)
	}
	return model.NewUser(user), nil
}

func (impl *UserReader) GetByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := impl.db.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return model.User{}, ent.MaskNotFound(err)
	}
	return model.NewUser(user), nil
}
