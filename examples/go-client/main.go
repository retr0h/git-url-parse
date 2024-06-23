package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/retr0h/git-url-parse/pkg"
	"github.com/retr0h/git-url-parse/pkg/repository"
)

type repositoryManager interface {
	RegisterParser(url string) error
	Parse() (pkg.RepositoryManager, error)
}

func getLogger(debug bool) *slog.Logger {
	logLevel := slog.LevelInfo
	if debug {
		logLevel = slog.LevelDebug
	}

	logger := slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      logLevel,
			TimeFormat: time.Kitchen,
		}),
	)

	return logger
}

func main() {
	debug := true
	logger := getLogger(debug)
	var r repositoryManager = repository.New(logger)

	r.RegisterParser("https://github.com/retr0h/foo")
	repo, err := r.Parse()
	if err != nil {
		panic(err)
	}

	logger.Info(repo.GetProtocol())
	logger.Info(repo.GetProtocols()[0])
	logger.Info(repo.GetResourceName())
	logger.Info(repo.GetHREF())
	logger.Info(repo.GetOwnerName())
	logger.Info(repo.GetRepoName())
	logger.Info(repo.GetPath())
	logger.Info(repo.GetBranchName())
	logger.Info(repo.GetProviderName())
}
