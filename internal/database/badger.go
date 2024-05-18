/*
Copyright (C) 2024  KubOps Technology

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package database

import (
	"github.com/dgraph-io/badger/v4"
)

type BadgerDB struct {
	DB *badger.DB
}

func (kv *BadgerDB) get(key []byte) ([]byte, error) {
	var value []byte

	err := kv.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			value = append([]byte{}, val...)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (kv *BadgerDB) Get(key []byte) ([]byte, error) {
	v, err := kv.get(key)
	if err == badger.ErrKeyNotFound {
		return []byte{}, nil
	}
	return v, err
}

func (kv *BadgerDB) Set(key, value []byte) error {
	return kv.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})
}

func (kv *BadgerDB) Has(key []byte) (bool, error) {
	_, err := kv.get(key)
	if err == badger.ErrKeyNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, err
}

func (kv *BadgerDB) Delete(key []byte) error {
	return kv.DB.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		return err
	})
}
