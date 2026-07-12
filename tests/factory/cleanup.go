package factory

import (
	"testing"
)

func (f Factory) Cleanup(t *testing.T) {
	db := f.DB

	db.Exec("TRUNCATE takes CASCADE")
	db.Exec("TRUNCATE classes CASCADE")
	db.Exec("TRUNCATE courses CASCADE")
	db.Exec("TRUNCATE users CASCADE")
}
