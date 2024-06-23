---
sidebar_position: 4
---

# Usage

## Package

### Parse a Repository

See example client in `examples/go-client/`.

Currently supports the following parsers:

- GitHub
- GitLab

```go
package main

import (
	"log/slog"
	"os"

	"github.com/retr0h/git-url-parse/pkg"
	"github.com/retr0h/git-url-parse/pkg/repository"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

type repositoryManager interface {
	RegisterParser(url string) error
	Parse() (pkg.RepositoryManager, error)
}

func main() {
	var r repositoryManager = repository.New(logger)

	r.RegisterParser("https://github.com/retr0h/foo")
	repo, err := r.Parse()
	if err != nil {
		panic(err)
	}

	logger.Info(repo.GetProtocol())     // https
	logger.Info(repo.GetProtocols()[0]) // https
	logger.Info(repo.GetResourceName()) // github.com
	logger.Info(repo.GetHREF())         // https://github.com/retr0h/foo
	logger.Info(repo.GetOwnerName())    // retr0h
	logger.Info(repo.GetRepoName())     // foo
	logger.Info(repo.GetPath())
	logger.Info(repo.GetBranchName())
	logger.Info(repo.GetProviderName()) // github
}
```
