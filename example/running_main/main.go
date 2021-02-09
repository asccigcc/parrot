package main

import (
	"github.com/asccigcc/parrot"
)

func main() {
	var token string = "TOKEN_HERE"
	var repo string = "github.com/username/repo"
	var branch string = "master"

	parrot := parrot.NewParrot(repo, branch, token)

	parrot.FetchDir("/path/to/directory")
	parrot.MoveToLocal("dir_destiny")

	parrot.FetchDir("/directory")
	parrot.MoveToLocal("dir_destiny")
}
