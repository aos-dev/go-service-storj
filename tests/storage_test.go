package tests

import (
	"os"
	"testing"

	tests "github.com/beyondstorage/go-integration-test/v4"
)

func TestStorager(t *testing.T) {
	if os.Getenv("STORAGE_STORJ_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_STORJ_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestStorager(t, setupTest(t))
}
