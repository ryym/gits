package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Dir struct {
	FileInfo os.FileInfo
	Path     string
}

type Options struct {
	FullPath bool
	Help     bool
}

func parseArgs() (Options, string) {
	fullpath := flag.Bool("p", false, "Print full paths")
	help := flag.Bool("h", false, "Show help message")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "gits lists git repositories in the specified path.\n")
		fmt.Fprintf(os.Stderr, "Usage: gits [options] <path>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	opts := Options{
		FullPath: *fullpath,
		Help:     *help,
	}

	if opts.Help || len(flag.Args()) != 1 {
		flag.Usage()
		return Options{}, ""
	}

	return opts, flag.Args()[0]
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	opts, root := parseArgs()
	if root == "" {
		return
	}

	var err error

	root, err = filepath.Abs(root)
	if err != nil {
		fatal(err)
	}

	var dirs []Dir
	dirs, err = ListRepos(root, ".git")
	if err != nil {
		fatal(err)
	}

	for _, d := range dirs {
		repoPath := d.Path
		if !opts.FullPath {
			repoPath, _ = filepath.Rel(root, repoPath)
		}
		fmt.Println(repoPath)
	}
}

func ListRepos(path string, gitDir string) ([]Dir, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	dirs := make([]Dir, 0)

	for _, f := range files {
		if !f.IsDir() || f.Name() == gitDir {
			continue
		}

		dir := filepath.Join(path, f.Name())
		git := filepath.Join(dir, gitDir)

		if _, er := os.Stat(git); os.IsNotExist(er) {
			childs, err := ListRepos(dir, gitDir)
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
