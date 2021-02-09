# Parrot

Parrot is a Go module that retrieves and loads a file from a remote git repository and copies it into a defined directory.

## Installation
To install the package run

```
go get -u github.com/asccigcc/parrot
```

## Usage

This packages is oriented to just dowload files from a specific directory.


```
package main

import (
    "github.com/asccigcc/parrot"
)

func main() {
    var token string = "GITHUB_TOKEN"
    var repo string = "github.com/username/repository"
    var branch string = "master"

    parrot := parrot.NewParrot(repo, branch, token)

	parrot.FetchDir("/path/to/directory")
	parrot.MoveToLocal("dir_destiny")

	parrot.FetchDir("/directory")
	parrot.MoveToLocal("dir_destiny"))
}
```