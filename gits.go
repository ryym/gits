package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Dir struct {
	FileInfo os.FileInfo
	Path     string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("gits lists up git repositories in the specified path.")
		fmt.Println("Usage: gits <path>")
		os.Exit(1)
	}

	root := os.Args[1]

	dirs, err := ListRepos(root)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, d := range dirs {
		relpath, _ := filepath.Rel(root, d.Path)
		fmt.Println(relpath)
	}
}

func ListRepos(path string) ([]Dir, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	dirs := make([]Dir, 0)

	for _, f := range files {
		if !f.IsDir() || f.Name() == ".git" {
			continue
		}

		dir := filepath.Join(path, f.Name())
		git := filepath.Join(dir, ".git")

		if _, er := os.Stat(git); os.IsNotExist(er) {
			childs, err := ListRepos(dir)
			if err != nil {
				return nil, err
			}
			dirs = append(dirs, childs...)
		} else {
			dirs = append(dirs, Dir{FileInfo: f, Path: dir})
		}
	}
	return dirs, nil
}
