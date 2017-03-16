package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Dir struct {
	FileInfo os.FileInfo
	Path     string
}

type Options struct {
	Help bool
}

func parseArgs() (Options, string) {
	help := flag.Bool("h", false, "Show help message")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "gits lists up git repositories in the specified path.\n")
		fmt.Fprintf(os.Stderr, "Usage: gits [options] <path>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	opts := Options{
		Help: *help,
	}

	if opts.Help || len(flag.Args()) != 1 {
		flag.Usage()
		return Options{}, ""
	}

	return opts, flag.Args()[0]
}

func main() {
	_, root := parseArgs()
	if root == "" {
		return
	}

	dirs, err := ListRepos(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
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
