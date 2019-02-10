package db

import (
	"fmt"
	"testing"
)

func TestGetDBInfo(t *testing.T) {
	user, passwd, table, err := getDB()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user, passwd, table)
}

func TestOpenDB(t *testing.T) {
	db, err := OpenDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
}
