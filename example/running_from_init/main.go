package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asccigcc/parrot"
)

func init() {

	token, ok := os.LookupEnv("GITHUB_ACCESS_TOKEN")
	if !ok {
		token, ok = os.LookupEnv("BUNDLE_GITHUB__COM")
		if !ok {
			log.Fatal("Github access token not defined")
		}
	}
	var repo string = "github.com/username/repo"
	var branch string = "master"

	parrot := parrot.NewParrot(repo, branch, token)

	parrot.FetchDir("/path/to/directory")
	parrot.MoveToLocal("dir_destiny")

	parrot.FetchDir("/directory")
	parrot.MoveToLocal("dir_destiny")
}

func main() {
	fmt.Println("My awesome code that will access to my local files")
}
