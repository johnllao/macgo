package kvdb

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGetAndSave(t *testing.T) {
	var err error

	var tempdir = os.TempDir()
	var dbpath = filepath.Join(tempdir, "tempdb_" + time.Now().Format("20060102150405") + ".db")

	var db *DB
	db, err = OpenDB(dbpath)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		db.Close()
		os.Remove(dbpath)
	}()

	var members = "MEMBERS"

	err = db.Save(members, "JL", []byte("John Lao"))
	if err != nil {
		t.Error(err)
		return
	}

	var data []byte
	data, err = db.Get(members, "JL")
	if err != nil {
		t.Error(err)
		return
	}

	if string(data) != "John Lao" {
		t.Error("Invalid value")
		return
	}
}