package main

import (
	"os"
	"testing"
)

func TestListRepos(t *testing.T) {
	// Create an empty directory because it can not be committed to Git.
	err := os.Mkdir("_test/empty-dir", 0755)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove("_test/empty-dir")

	files, err := ListRepos("_test", "git")
	if err != nil {
		t.Error(err)
	}

	actual := make([]string, len(files))
	for i, f := range files {
		actual[i] = f.Path
	}

	expected := []string{
		"_test/repo",
		"_test/x/repo",
		"_test/x/x/repo",
	}

	if !isSamePaths(actual, expected) {
		m := "Listed repositories are wrong.\n expected: %v\n actual: %v\n"
		t.Errorf(m, expected, actual)
	}
}

func isSamePaths(actual []string, expected []string) bool {
	if actual == nil || len(actual) != len(expected) {
		return false
	}
	for i, a := range actual {
		if a != expected[i] {
			return false
		}
	}
	return true
}
