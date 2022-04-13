package inject

import (
	"github.com/go-sql-driver/mysql"
	"github.com/k3forx/coffee_memo/pkg/config"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/reader"
	"github.com/k3forx/coffee_memo/pkg/writer"

	"github.com/k3forx/coffee_memo/pkg/session"
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
			User: reader.NewUserReader(readerClient),
		},
		Writer: Writer{
			User: writer.NewUserWriter(writerClient),
		},
	}

	return injector, fn, nil
}

type Injector struct {
	Reader  Reader
	Writer  Writer
	Session session.Session
}

type Reader struct {
	User reader.User
}

type Writer struct {
	User writer.User
}
