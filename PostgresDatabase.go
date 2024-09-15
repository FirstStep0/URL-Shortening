package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type PostgresDatabase struct {
	db *sql.DB
}

func (database *PostgresDatabase) GetNextKey() KeyType {
	db := database.db
	var nextId int64 = 0

	row := db.QueryRow("SELECT MAX(id) FROM sites")
	if row != nil {
		row.Scan(&nextId)
	}

	return database.GetKey(nextId)
}

func (database *PostgresDatabase) GetKey(index int64) KeyType {
	base := int64(len(alphabet))
	var build strings.Builder

	for {
		fmt.Fprintf(&build, string(alphabet[index%base]))
		index /= base
		if index == 0 {
			break
		}
	}

	return KeyType(build.String())
}

func (database *PostgresDatabase) Insert(value ValueType) KeyType {
	db := database.db
	key := database.GetNextKey()

	_, err := db.Exec("INSERT INTO sites (Key, Value) VALUES($1, $2)", key, value)
	if err != nil {
		fmt.Println(err)
	}
	key, _ = database.GetByValue(value)
	return key
}

func (database *PostgresDatabase) GetByKey(key KeyType) (ValueType, bool) {
	db := database.db
	var value ValueType

	db.QueryRow("SELECT Value FROM sites WHERE Key = $1", key).Scan(&value)

	return value, value != ""
}

func (database *PostgresDatabase) GetByValue(value ValueType) (KeyType, bool) {
	db := database.db
	var key KeyType
	row := db.QueryRow("SELECT Key FROM sites WHERE Value = $1", value)
	row.Scan(&key)

	return key, key != ""
}
