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
	"os"
	"testing"

	"github.com/dgraph-io/badger/v4"
)

func setupDatabase(t *testing.T, storageInMemory bool) *Database {
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
	return NewDatabase(&BadgerDB{DB: db})
}

func TestSetTags(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	tags := []string{"v0", "v1", "v0.1.0", "v0.1", "v1.0.1", "v1.0.0-rc.1"}
	err := db.SetTags(repo, tags)
	if err != nil {
		t.Fatal(err)
	}
	tagsFromDB, err := db.Tags(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) != len(tagsFromDB) {
		t.Fatalf("expected %d tags, got %d", len(tags), len(tagsFromDB))
	}
	for _, tag := range tags {
		found := false
		for _, tagFromDB := range tagsFromDB {
			if tag == tagFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v tags, got %v", tags, tagsFromDB)
		}
	}
}

func TestSetTagsOverwrite(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	tags := []string{"v0", "v1", "v0.1.0", "v0.1", "v1.0.1", "v1.0.0-rc.1"}
	tags2 := []string{"v0.1.1", "v0.1.2", "v0.1.3", "v0.1.4", "v0.1.5", "v0.1.6"}
	err := db.SetTags(repo, tags)
	if err != nil {
		t.Fatal(err)
	}
	err = db.SetTags(repo, tags2)
	if err != nil {
		t.Fatal(err)
	}
	tagsFromDB, err := db.Tags(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags2) != len(tagsFromDB) {
		t.Fatalf("expected %d tags, got %d", len(tags2), len(tagsFromDB))
	}
	for _, tag := range tags2 {
		found := false
		for _, tagFromDB := range tagsFromDB {
			if tag == tagFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v tags, got %v", tags2, tagsFromDB)
		}
	}
}

func TestSetTagsLowercaseKey(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	repoLowerCase := "github.com/kubopscloud/poke"
	tags := []string{"v0", "v1", "v0.1.0", "v0.1", "v1.0.1", "v1.0.0-rc.1"}
	tags2 := []string{"v0.1.1", "v0.1.2", "v0.1.3", "v0.1.4", "v0.1.5", "v0.1.6"}
	err := db.SetTags(repo, tags)
	if err != nil {
		t.Fatal(err)
	}
	err = db.SetTags(repoLowerCase, tags2)
	if err != nil {
		t.Fatal(err)
	}
	// Check if the tags are still there for repo
	tagsFromDB, err := db.Tags(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) != len(tagsFromDB) {
		t.Fatalf("expected %d tags, got %d", len(tags), len(tagsFromDB))
	}
	for _, tag := range tags {
		found := false
		for _, tagFromDB := range tagsFromDB {
			if tag == tagFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v tags, got %v", tags, tagsFromDB)
		}
	}
	// Check if the tags are still there for repoLowerCase
	tagsFromDB, err = db.Tags(repoLowerCase)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags2) != len(tagsFromDB) {
		t.Fatalf("expected %d tags, got %d", len(tags2), len(tagsFromDB))
	}
	for _, tag := range tags2 {
		found := false
		for _, tagFromDB := range tagsFromDB {
			if tag == tagFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v tags, got %v", tags2, tagsFromDB)
		}
	}
}

func TestTags(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	tags := []string{"v0", "v1", "v0.1.0", "v0.1", "v1.0.1", "v1.0.0-rc.1"}
	err := db.SetTags(repo, tags)
	if err != nil {
		t.Fatal(err)
	}
	tagsFromDB, err := db.Tags(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) != len(tagsFromDB) {
		t.Fatalf("expected %d tags, got %d", len(tags), len(tagsFromDB))
	}
	for _, tag := range tags {
		found := false
		for _, tagFromDB := range tagsFromDB {
			if tag == tagFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v tags, got %v", tags, tagsFromDB)
		}
	}
}

func TestTagsLowercaseKey(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	repoLowerCase := "github.com/kubopscloud/poke"
	tags := []string{"v0", "v1", "v0.1.0", "v0.1", "v1.0.1", "v1.0.0-rc.1"}
	err := db.SetTags(repo, tags)
	if err != nil {
		t.Fatal(err)
	}
	tagsFromDB, err := db.Tags(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) != len(tagsFromDB) {
		t.Fatalf("expected %d tags, got %d", len(tags), len(tagsFromDB))
	}
	for _, tag := range tags {
		found := false
		for _, tagFromDB := range tagsFromDB {
			if tag == tagFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v tags, got %v", tags, tagsFromDB)
		}
	}
	tagsFromDB, err = db.Tags(repoLowerCase)
	if err != nil {
		t.Fatal(err)
	}
	if len(tagsFromDB) != 0 {
		t.Fatalf("expected 0 tags, got %d containing %v", len(tagsFromDB), tagsFromDB)
	}

}

func TestTagsEmptyIfKeyNotFound(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	tagsFromDB, err := db.Tags(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(tagsFromDB) != 0 {
		t.Fatalf("expected 0 tags, got %d", len(tagsFromDB))
	}
}

func TestSetBranches(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branches := []string{"main", "develop", "feature/1", "feature/2", "feature/3"}
	err := db.SetBranches(repo, branches)
	if err != nil {
		t.Fatal(err)
	}
	branchesFromDB, err := db.Branches(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(branches) != len(branchesFromDB) {
		t.Fatalf("expected %d branches, got %d", len(branches), len(branchesFromDB))
	}
	for _, branch := range branches {
		found := false
		for _, branchFromDB := range branchesFromDB {
			if branch == branchFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v branches, got %v", branches, branchesFromDB)
		}
	}
}

func TestSetBranchesOverwrite(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branches := []string{"main", "develop", "feature/1", "feature/2", "feature/3"}
	branches2 := []string{"main", "develop", "feature/1", "feature/2", "feature/3", "feature/4"}
	err := db.SetBranches(repo, branches)
	if err != nil {
		t.Fatal(err)
	}
	err = db.SetBranches(repo, branches2)
	if err != nil {
		t.Fatal(err)
	}
	branchesFromDB, err := db.Branches(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(branches2) != len(branchesFromDB) {
		t.Fatalf("expected %d branches, got %d", len(branches2), len(branchesFromDB))
	}
	for _, branch := range branches2 {
		found := false
		for _, branchFromDB := range branchesFromDB {
			if branch == branchFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v branches, got %v", branches2, branchesFromDB)
		}
	}
}

func TestSetBranchesLowercaseKey(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	repoLowerCase := "github.com/kubopscloud/poke"
	branches := []string{"main", "develop", "feature/1", "feature/2", "feature/3"}
	err := db.SetBranches(repo, branches)
	if err != nil {
		t.Fatal(err)
	}
	branchesFromDB, err := db.Branches(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(branches) != len(branchesFromDB) {
		t.Fatalf("expected %d branches, got %d", len(branches), len(branchesFromDB))
	}
	for _, branch := range branches {
		found := false
		for _, branchFromDB := range branchesFromDB {
			if branch == branchFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v branches, got %v", branches, branchesFromDB)
		}
	}
	branchesFromDB, err = db.Branches(repoLowerCase)
	if err != nil {
		t.Fatal(err)
	}
	if len(branchesFromDB) != 0 {
		t.Fatalf("expected 0 branches, got %d containing %v", len(branchesFromDB), branchesFromDB)
	}
}

func TestBranches(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branches := []string{"main", "develop", "feature/1", "feature/2", "feature/3"}
	err := db.SetBranches(repo, branches)
	if err != nil {
		t.Fatal(err)
	}
	branchesFromDB, err := db.Branches(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(branches) != len(branchesFromDB) {
		t.Fatalf("expected %d branches, got %d", len(branches), len(branchesFromDB))
	}
	for _, branch := range branches {
		found := false
		for _, branchFromDB := range branchesFromDB {
			if branch == branchFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v branches, got %v", branches, branchesFromDB)
		}
	}
}

func TestBranchesLowercaseKey(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	repoLowerCase := "github.com/kubopscloud/poke"
	branches := []string{"main", "develop", "feature/1", "feature/2", "feature/3"}
	err := db.SetBranches(repo, branches)
	if err != nil {
		t.Fatal(err)
	}
	branchesFromDB, err := db.Branches(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(branches) != len(branchesFromDB) {
		t.Fatalf("expected %d branches, got %d", len(branches), len(branchesFromDB))
	}
	for _, branch := range branches {
		found := false
		for _, branchFromDB := range branchesFromDB {
			if branch == branchFromDB {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected %v branches, got %v", branches, branchesFromDB)
		}
	}
	branchesFromDB, err = db.Branches(repoLowerCase)
	if err != nil {
		t.Fatal(err)
	}
	if len(branchesFromDB) != 0 {
		t.Fatalf("expected 0 branches, got %d containing %v", len(branchesFromDB), branchesFromDB)
	}
}

func TestBranchesEmptyIfKeyNotFound(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branchesFromDB, err := db.Branches(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(branchesFromDB) != 0 {
		t.Fatalf("expected 0 branches, got %d", len(branchesFromDB))
	}
}

func TestSetLastCommitOfBranch(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branch := "main"
	commit := "66af0f77eaa059852f46a512d15adf09f9f1ea28"
	err := db.SetLastCommitOfBranch(repo, branch, commit)
	if err != nil {
		t.Fatal(err)
	}
	commitFromDB, err := db.LastCommitOfBranch(repo, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commit != commitFromDB {
		t.Fatalf("expected %s, got %s", commit, commitFromDB)
	}
}

func TestSetLastCommitOfBranchOverwrite(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branch := "main"
	commit := "66af0f77eaa059852f46a512d15adf09f9f1ea28"
	commit2 := "66af0f77eaa059852f46a512d15adf09f9f1ea29"
	err := db.SetLastCommitOfBranch(repo, branch, commit)
	if err != nil {
		t.Fatal(err)
	}
	err = db.SetLastCommitOfBranch(repo, branch, commit2)
	if err != nil {
		t.Fatal(err)
	}
	commitFromDB, err := db.LastCommitOfBranch(repo, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commit2 != commitFromDB {
		t.Fatalf("expected %s, got %s", commit2, commitFromDB)
	}
}

func TestSetLastCommitOfBranchLowercaseKey(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	repoLowerCase := "github.com/kubopscloud/poke"
	branch := "main"
	commit := "66af0f77eaa059852f46a512d15adf09f9f1ea28"
	err := db.SetLastCommitOfBranch(repo, branch, commit)
	if err != nil {
		t.Fatal(err)
	}
	commitFromDB, err := db.LastCommitOfBranch(repo, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commit != commitFromDB {
		t.Fatalf("expected %s, got %s", commit, commitFromDB)
	}
	commitFromDB, err = db.LastCommitOfBranch(repoLowerCase, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commitFromDB != "" {
		t.Fatalf("expected empty string, got %s", commitFromDB)
	}
}

func TestLastCommitOfBranch(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branch := "main"
	commit := "66af0f77eaa059852f46a512d15adf09f9f1ea28"
	err := db.SetLastCommitOfBranch(repo, branch, commit)
	if err != nil {
		t.Fatal(err)
	}
	commitFromDB, err := db.LastCommitOfBranch(repo, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commit != commitFromDB {
		t.Fatalf("expected %s, got %s", commit, commitFromDB)
	}
}

func TestCommitLowercaseKey(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	repoLowerCase := "github.com/kubopscloud/poke"
	branch := "main"
	commit := "66af0f77eaa059852f46a512d15adf09f9f1ea28"
	err := db.SetLastCommitOfBranch(repo, branch, commit)
	if err != nil {
		t.Fatal(err)
	}
	commitFromDB, err := db.LastCommitOfBranch(repo, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commit != commitFromDB {
		t.Fatalf("expected %s, got %s", commit, commitFromDB)
	}
	commitFromDB, err = db.LastCommitOfBranch(repoLowerCase, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commitFromDB != "" {
		t.Fatalf("expected empty string, got %s", commitFromDB)
	}
}

func TestLastCommitOfBranchEmptyIfKeyNotFound(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branch := "main"
	commitFromDB, err := db.LastCommitOfBranch(repo, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commitFromDB != "" {
		t.Fatalf("expected empty string, got %s", commitFromDB)
	}
}

func TestDeleteTags(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	tags := []string{"v0", "v1", "v0.1.0", "v0.1", "v1.0.1", "v1.0.0-rc.1"}
	err := db.SetTags(repo, tags)
	if err != nil {
		t.Fatal(err)
	}
	err = db.DeleteTags(repo)
	if err != nil {
		t.Fatal(err)
	}
	tagsFromDB, err := db.Tags(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(tagsFromDB) != 0 {
		t.Fatalf("expected 0 tags, got %d", len(tagsFromDB))
	}
}

func TestDeleteBranches(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branches := []string{"main", "develop", "feature/1", "feature/2", "feature/3"}
	err := db.SetBranches(repo, branches)
	if err != nil {
		t.Fatal(err)
	}
	err = db.DeleteBranches(repo)
	if err != nil {
		t.Fatal(err)
	}
	branchesFromDB, err := db.Branches(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(branchesFromDB) != 0 {
		t.Fatalf("expected 0 branches, got %d", len(branchesFromDB))
	}
}

func TestDeleteLastCommitOfBranch(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	branch := "main"
	commit := "66af0f77eaa059852f46a512d15adf09f9f1ea28"
	err := db.SetLastCommitOfBranch(repo, branch, commit)
	if err != nil {
		t.Fatal(err)
	}
	err = db.DeleteLastCommitOfBranch(repo, branch)
	if err != nil {
		t.Fatal(err)
	}
	commitFromDB, err := db.LastCommitOfBranch(repo, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commitFromDB != "" {
		t.Fatalf("expected empty string, got %s", commitFromDB)
	}
}

func TestDeleteRepository(t *testing.T) {
	db := setupDatabase(t, false)
	repo := "github.com/KubOpsCloud/poke"
	tags := []string{"v0", "v1", "v0.1.0", "v0.1", "v1.0.1", "v1.0.0-rc.1"}
	branches := []string{"main", "develop", "feature/1", "feature/2", "feature/3"}
	branch := "main"
	commit := "66af0f77eaa059852f46a512d15adf09f9f1ea28"
	err := db.SetTags(repo, tags)
	if err != nil {
		t.Fatal(err)
	}
	err = db.SetBranches(repo, branches)
	if err != nil {
		t.Fatal(err)
	}
	err = db.SetLastCommitOfBranch(repo, branch, commit)
	if err != nil {
		t.Fatal(err)
	}
	err = db.DeleteRepo(repo)
	if err != nil {
		t.Fatal(err)
	}
	tagsFromDB, err := db.Tags(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(tagsFromDB) != 0 {
		t.Fatalf("expected 0 tags, got %d", len(tagsFromDB))
	}
	branchesFromDB, err := db.Branches(repo)
	if err != nil {
		t.Fatal(err)
	}
	if len(branchesFromDB) != 0 {
		t.Fatalf("expected 0 branches, got %d", len(branchesFromDB))
	}
	commitFromDB, err := db.LastCommitOfBranch(repo, branch)
	if err != nil {
		t.Fatal(err)
	}
	if commitFromDB != "" {
		t.Fatalf("expected empty string, got %s", commitFromDB)
	}
}
