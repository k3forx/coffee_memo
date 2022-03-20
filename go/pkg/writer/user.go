package writer

import (
	"context"
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func NewUserWriter(db *ent.Client) *UserWriter {
	return &UserWriter{
		db: db,
	}
}

//go:generate mockgen -source=./user.go -destination=./mock/user_mock.go -package=writer
type User interface {
	Create(ctx context.Context, user model.User) error
}

type UserWriter struct {
	db *ent.Client
}

var _ User = (*UserWriter)(nil)

func (impl *UserWriter) Create(ctx context.Context, user model.User) error {
	now := time.Now().UTC()
	if _, err := impl.db.User.Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetFlags(user.Flags.Int()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx); err != nil {
		return err
	}
	return nil
}
