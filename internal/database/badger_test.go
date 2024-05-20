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
	"bytes"
	"os"
	"testing"

	"github.com/dgraph-io/badger/v4"
)

// setupBadgerDB creates a new BadgerDB instance for testing
func setupBadgerDB(t *testing.T, storageInMemory bool) *BadgerDB {
	t.Helper()
	dir, err := os.MkdirTemp(os.TempDir(), "badger")
	if err != nil {
		t.Fatal(err)
	}
	opt := badger.DefaultOptions(dir)
	if storageInMemory {
		opt = badger.DefaultOptions("").WithInMemory(true)
	}
	db, err := badger.Open(opt)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		db.Close()
		os.RemoveAll(dir)
	})
	return &BadgerDB{db}
}

func TestLowerGet(t *testing.T) {
	db := setupBadgerDB(t, false)
	key := []byte("key")
	value := []byte("value")
	err := db.Set(key, value)
	if err != nil {
		t.Fatal(err)
	}
	v, err := db.get(key)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(v, value) {
		t.Fatalf("expected %s, got %s", string(value), string(v))
	}
}

func TestLowerGetKeyNotFound(t *testing.T) {
	db := setupBadgerDB(t, false)
	key := []byte("key")
	v, err := db.get(key)
	if err != badger.ErrKeyNotFound {
		t.Fatal(err)
	}
	if v != nil {
		t.Fatalf("expected nil, got %s", string(v))
	}
}

func TestGet(t *testing.T) {
	db := setupBadgerDB(t, false)
	key := []byte("key")
	value := []byte("value")
	err := db.Set(key, value)
	if err != nil {
		t.Fatal(err)
	}
	v, err := db.Get(key)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(v, value) {
		t.Fatalf("expected %s, got %s", string(value), string(v))
	}
}

func TestGetKeyNotFound(t *testing.T) {
	db := setupBadgerDB(t, false)
	key := []byte("key")
	value := []byte{}
	v, err := db.Get(key)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(v, value) {
		t.Fatalf("expected %s, got %s", string(value), string(v))
	}
}

func TestSet(t *testing.T) {
	db := setupBadgerDB(t, false)
	key := []byte("key")
	value := []byte("value")
	err := db.Set(key, value)
	if err != nil {
		t.Fatal(err)
	}
	v, err := db.Get(key)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(v, value) {
		t.Fatalf("expected %s, got %s", string(value), string(v))
	}
}

func TestHas(t *testing.T) {
	db := setupBadgerDB(t, false)
	key := []byte("key")
	value := []byte("value")
	err := db.Set(key, value)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := db.Has(key)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatalf("expected key to be present")
	}
}

func TestHasKeyNotFound(t *testing.T) {
	db := setupBadgerDB(t, false)
	key := []byte("key")
	ok, err := db.Has(key)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatalf("expected key to be not present")
	}
}

func TestDelete(t *testing.T) {
	db := setupBadgerDB(t, false)
	key := []byte("key")
	value := []byte("value")
	err := db.Set(key, value)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := db.Has(key)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatalf("expected key to be present")
	}
	err = db.Delete(key)
	if err != nil {
		t.Fatal(err)
	}
	ok, err = db.Has(key)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatalf("expected key to be deleted")
	}
}
