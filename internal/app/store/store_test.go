package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = `host=pbx-1.hl.ua dbname=restapi_test user = restapi_master password = jscjkklqeuiw38d`
	}

	os.Exit(m.Run())
}
