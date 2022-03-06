package inject

import (
	"github.com/go-sql-driver/mysql"
	"github.com/k3forx/coffee_memo/pkg/config"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/reader"
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
	dbReader, err := ent.Open("mysql", mc.FormatDSN(), opts...)
	if err != nil {
		return Injector{}, nil, err
	}

	fn := func() {
		_ = dbReader.Close()
	}

	injector := Injector{
		Reader: Reader{
			User: reader.NewUserReader(dbReader),
		},
	}

	return injector, fn, nil
}

type Injector struct {
	Reader Reader
}

type Reader struct {
	User reader.User
}
