package factory

import (
	"testing"
)

func Cleanup(t *testing.T) {
	db := Infra.DB

	db.Exec("TRUNCATE takes CASCADE")
	db.Exec("TRUNCATE classes CASCADE")
	db.Exec("TRUNCATE courses CASCADE")
	db.Exec("TRUNCATE users CASCADE")
}
