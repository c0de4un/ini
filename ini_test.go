package ini

import "testing"

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Database struct {
	db       string
	name     string
	host     string
	driver   string
	password string
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func newDB() Database {
	var db Database

	return db
}

func (db *Database) OnParam(name string, value string) error {
	if name == "db" {
		db.db = value
		return nil
	}
	if name == "name" {
		db.name = value
		return nil
	}
	if name == "host" {
		db.host = value
		return nil
	}
	if name == "driver" {
		db.driver = value
		return nil
	}
	if name == "password" {
		db.password = value
		return nil
	}

	return nil
}

func (db *Database) readConfig() error {
	reader := NewReader()
	return reader.ReadAll("db.ini", db)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// TESTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func TestIniReader(t *testing.T) {
	db := newDB()
	err := db.readConfig()
	if err != nil {
		t.Fatalf("ini.Reader.ReadAll failed with error: %s", err)
	}

	if db.name != "mysql" {
		t.Fatalf("ini.Reader.ReadAll parse failed, 'mysql' expected as name, got: %s", db.name)
	}

	if len(db.password) > 0 {
		t.Fatalf("ini.Reader.ReadAll parse failed, empty password expected, got: %s", db.password)
	}
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
