package main

import (
	"fmt"
	"strings"
)

type MyDatabase struct {
	mKeyToValue map[KeyType]ValueType
	mValueToKey map[ValueType]KeyType
	top         int64
}

func (db *MyDatabase) GetByKey(key KeyType) (value ValueType, ok bool) {
	value, ok = db.mKeyToValue[key]
	if !ok {
		value = ""
	}
	return
}

func (db *MyDatabase) GetByValue(value ValueType) (key KeyType, ok bool) {
	key, ok = db.mValueToKey[value]
	if !ok {
		key = ""
	}
	return
}

func (db *MyDatabase) Insert(value ValueType) KeyType {
	key := db.GetNextKey()
	db.mKeyToValue[key] = value
	db.mValueToKey[value] = key
	return key
}

func (db *MyDatabase) GetNextKey() KeyType {
	index := db.top
	db.top++

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
