package parrot

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asccigcc/parrot/internal/filesystem"
	"github.com/posener/gitfs"
	"github.com/posener/gitfs/fsutil"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type Parrot struct {
	token      string `default:""`
	repo       string
	branch     string
	ctx        context.Context
	client     *http.Client
	fileSystem http.FileSystem
}

// Return a FileSystem function if the path is correct
//
func (r *Parrot) FetchDir(dir string) http.FileSystem {
	log.Println("Starting Fetching Directory")
	path := fmt.Sprintf("%s%s@%s", r.repo, dir, r.branch)
	log.Info(path)

	fs, err := gitfsNew(r, path, r.token)

	log.Println("File system:", fs)
	if err != nil {
		log.Fatalf("Failed initializing git filesystem: %s.", err)
	}

	r.fileSystem = fs

	return fs
}

func (r *Parrot) MoveToLocal(local_dir string) {
	log.Println("Starting Moving Files")

	fs := r.fileSystem
	log.Println(fs)

	walker := fsutil.Walk(fs, "")
	log.Println(walker)

	for walker.Step() {
		if err := walker.Err(); err != nil {
			continue
		}

		folder := walker.Path()
		log.Info(folder)

		if !walker.Stat().IsDir() {
			body := OpenFile(fs, folder)

			filesystem.CreateFile(folder, body, local_dir)
		}
	}
}

func gitfsNew(r *Parrot, repo string, token string) (http.FileSystem, error) {
	var client *http.Client = Auth(token, r.ctx)

	if client != nil {
		return gitfs.New(r.ctx, repo, gitfs.OptClient(client))
	}

	return gitfs.New(r.ctx, repo)
}

func Auth(token string, ctx context.Context) *http.Client {
	// auth is optional
	if len(token) > 0 {
		return oauth2.NewClient(
			ctx,
			oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
		)
	}
	return nil
}

func OpenFile(fs http.FileSystem, folder string) string {
	file, _ := fs.Open(folder)
	log.Info(file)

	fil, _ := ioutil.ReadAll(file)

	body := string(fil)
	log.Info(body)

	return body
}

func NewParrot(repo string, branch string, token string) *Parrot {
	log.Println("Initializing Parrot")
	return &Parrot{
		token:      token,
		repo:       repo,
		branch:     branch,
		ctx:        context.Background(),
		client:     &http.Client{},
		fileSystem: nil,
	}
}
