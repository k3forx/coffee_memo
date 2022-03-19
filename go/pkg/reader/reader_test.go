package reader_test

import (
	"os"
	"testing"

	"github.com/k3forx/coffee_memo/pkg/ent"

	_ "github.com/go-sql-driver/mysql"
)

var (
	testClient *ent.Client
)

func TestMain(m *testing.M) {
	var err error
	testClient, err = ent.Open("mysql", "root:password@tcp(mysql:3306)/coffee_app_test?parseTime=true")
	if err != nil {
		os.Exit(1)
	}
	defer testClient.Close()

	m.Run()
}
