package dbs

import (
	"slices"

	badger "github.com/dgraph-io/badger/v4"
)

type BadgerDB struct {
	DB *badger.DB
}

type BadgerKeyValue struct {
	Key   []byte
	Value []byte
}

func Connect(path string) (*badger.DB, error) {
	db, err := badger.Open(badger.DefaultOptions(path))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (b *BadgerDB) UpdateOne(key []byte, value []byte) error {
	err := b.DB.Update(func(transaction *badger.Txn) error {
		err := transaction.Set(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func (b *BadgerDB) UpdateMany(keyValuePairs []BadgerKeyValue) error {
	err := b.DB.Update(func(transaction *badger.Txn) error {
		for _, keyValue := range keyValuePairs {
			err := transaction.Set(keyValue.Key, keyValue.Value)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (b *BadgerDB) FindOne(key []byte) ([]byte, error) {
	transaction := b.DB.NewTransaction(true)
	defer transaction.Discard()

	item, err := transaction.Get(key)

	if err != nil {
		return nil, err
	}

	var value []byte

	item.Value(func(val []byte) error {
		value = slices.Clone(val)
		return nil
	})

	return value, nil
}
