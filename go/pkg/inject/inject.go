package inject

import (
	"github.com/go-sql-driver/mysql"
	"github.com/k3forx/coffee_memo/pkg/config"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/reader"
	"github.com/k3forx/coffee_memo/pkg/writer"
)

func New() (Injector, func(), error) {
	var opts []ent.Option
	if !config.IsProduction() {
		opts = append(opts, ent.Debug())
	}

	mc := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "mysql" + ":" + "3306",
		DBName:               "coffee_app",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	readerClient, err := ent.Open("mysql", mc.FormatDSN(), opts...)
	if err != nil {
		return Injector{}, nil, err
	}

	// TODO: change config
	writerClient, err := ent.Open("mysql", mc.FormatDSN(), opts...)
	if err != nil {
		return Injector{}, nil, err
	}

	fn := func() {
		_ = readerClient.Close()
		_ = writerClient.Close()
	}

	injector := Injector{
		Reader: Reader{
			UserCoffeeBean: reader.NewUserCoffeeBeanReader(readerClient),
			User:           reader.NewUserReader(readerClient),
		},
		Writer: Writer{
			UserBrewRecipe: writer.NewUserBrewRecipeWriter(writerClient),
			UserCoffeeBean: writer.NewUserCoffeeBeanWriter(writerClient),
			User:           writer.NewUserWriter(writerClient),
		},
	}

	return injector, fn, nil
}

type Injector struct {
	Reader Reader
	Writer Writer
}

type Reader struct {
	UserCoffeeBean reader.UserCoffeeBean
	User           reader.User
}

type Writer struct {
	UserBrewRecipe writer.UserBrewRecipe
	UserCoffeeBean writer.UserCoffeeBean
	User           writer.User
}
