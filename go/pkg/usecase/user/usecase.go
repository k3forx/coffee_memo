package user

import (
	"context"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/user"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func NewUsecase() *UserUsecase {
	return &UserUsecase{}
}

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=user
type Usecase interface {
	GetByID(ctx context.Context, in GetByIDInput) *GetByIDOutput
}

type UserUsecase struct{}

var _ Usecase = (*UserUsecase)(nil)

func (u *UserUsecase) GetByID(ctx context.Context, in GetByIDInput) *GetByIDOutput {
	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())

	mc := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "mysql" + ":" + "3306",
		DBName:               "coffee_app",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("open error: %+v\n", err)
		return nil
	}
	defer client.Close()

	user, err := client.User.Query().Where(user.IDEQ(int32(in.UserID))).Only(ctx)
	if err != nil {
		log.Fatalf("select err: %+v\n", err)
		return nil
	}

	return &GetByIDOutput{
		User: model.NewUser(user),
	}
}
