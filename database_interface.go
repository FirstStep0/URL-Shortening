package main

type KeyType = string
type ValueType = string

type site struct {
	key   string
	value string
}

type Database interface {
	GetByKey(KeyType) (ValueType, bool)
	GetByValue(ValueType) (KeyType, bool)
	Insert(ValueType) KeyType
}

var alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
