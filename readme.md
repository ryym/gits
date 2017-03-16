# Gits

Gits lists up all git repositories in the specified path.

## Usage

For instance, you can list up git packages in the `$GOPATH/src` using this command.

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
