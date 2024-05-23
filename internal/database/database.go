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
	"encoding/json"
	"errors"
)

const (
	tagsSuffix     = "/tags"
	branchesSuffix = "/branches"
	branchSuffix   = "/branch/"
)

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
	return kv.db.Set([]byte(repo+tagsSuffix), v)
}

func (kv *Database) Tags(repo string) ([]string, error) {
	var tags []string
	v, err := kv.db.Get([]byte(repo + tagsSuffix))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(v, &tags)
	if err != nil && len(v) > 0 {
		return nil, err
	}
	return tags, nil
}

func (kv *Database) DeleteTags(repo string) error {
	return kv.db.Delete([]byte(repo + tagsSuffix))
}

func (kv *Database) SetBranches(repo string, branches []string) error {
	v, err := json.Marshal(branches)
	if err != nil {
		return err
	}
	return kv.db.Set([]byte(repo+branchesSuffix), v)
}

func (kv *Database) Branches(repo string) ([]string, error) {
	var branches []string
	v, err := kv.db.Get([]byte(repo + branchesSuffix))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(v, &branches)
	if err != nil && len(v) > 0 {
		return nil, err
	}
	return branches, nil
}

func (kv *Database) DeleteBranches(repo string) error {
	return kv.db.Delete([]byte(repo + branchesSuffix))
}

func (kv *Database) SetLastCommitOfBranch(repo, branch, commit string) error {
	return kv.db.Set([]byte(repo+branchSuffix+branch), []byte(commit))
}

func (kv *Database) LastCommitOfBranch(repo, branch string) (string, error) {
	v, err := kv.db.Get([]byte(repo + branchSuffix + branch))
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func (kv *Database) DeleteLastCommitOfBranch(repo, branch string) error {
	return kv.db.Delete([]byte(repo + branchSuffix + branch))
}

func (kv *Database) DeleteRepo(repo string) error {
	var err error
	branches, err := kv.Branches(repo)
	if err != nil {
		return err
	}
	for _, branch := range branches {
		err = errors.Join(err, kv.DeleteLastCommitOfBranch(repo, branch))
	}
	errBranches := kv.DeleteBranches(repo)
	errTags := kv.DeleteTags(repo)
	return errors.Join(err, errTags, errBranches)
}
