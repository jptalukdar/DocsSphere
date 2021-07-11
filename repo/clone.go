package repo

import (
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/jptalukdar/DocsSphere/response"
)

func GetRepository(repoUrl string, path string) response.Response {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		panic(err)
	}
	return response.Response{}
}
