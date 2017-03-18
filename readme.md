# Gits

[![travis](https://travis-ci.org/ryym/gits.svg?branch=master)](https://travis-ci.org/ryym/gits)

Gits lists all git repositories in the specified path.

## Usage

For instance, you can list git packages in the `$GOPATH/src` using this command.

```sh
$ gits $GOPATH/src
hello-world
github.com/peco/peco
github.com/ryym/bar
github.com/ryym/baz
golang.org/x/net
golang.org/x/text
...
```

## Install

```
go get -u github.com/ryym/gits
```
