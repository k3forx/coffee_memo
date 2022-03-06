package user

import (
	"context"
	"fmt"
	"log"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/user"
)

func NewUsecase() *UserUsecase {
	return &UserUsecase{}
}

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=user
type Usecase interface {
	GetByID(ctx context.Context)
}

type UserUsecase struct{}

var _ Usecase = (*UserUsecase)(nil)

func (u *UserUsecase) GetByID(ctx context.Context) {
	client, err := ent.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	user, err := client.User.Query().Where(user.ID(int32(1))).Only(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("user: %+v\n", user)
}
