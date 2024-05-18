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

import "encoding/json"

type IDatabase interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
	Has(key []byte) (bool, error)
	Delete(key []byte) error
}

type Database struct {
	db IDatabase
}

func NewDatabase(db IDatabase) *Database {
	return &Database{
		db: db,
	}
}

func (kv *Database) SetTags(repo string, tags []string) error {
	v, err := json.Marshal(tags)
	if err != nil {
		return err
	}
	return kv.db.Set([]byte(repo), v)
}

func (kv *Database) Tags(repo string) ([]string, error) {
	var tags []string
	v, err := kv.db.Get([]byte(repo))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(v, &tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (kv *Database) DeleteRepoTags(repo string) error {
	return kv.db.Delete([]byte(repo))
}

func (kv *Database) DeleteRepo(repo string) error {
	return kv.DeleteRepoTags(repo)
}
