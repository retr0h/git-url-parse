// Copyright (c) 2024 John Dewey

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package repository

import (
	"fmt"
	"log/slog"

	giturls "github.com/chainguard-dev/git-urls"

	"github.com/retr0h/git-url-parse/internal"
	"github.com/retr0h/git-url-parse/internal/repositories/bitbucket"
	"github.com/retr0h/git-url-parse/internal/repositories/github"
	"github.com/retr0h/git-url-parse/internal/repositories/gitlab"
	"github.com/retr0h/git-url-parse/pkg"
)

// New factory to create a new Repository instance.
func New(
	// url string,
	logger *slog.Logger,
) *Repository {
	return &Repository{
		logger: logger,
	}
}

// RegisterParser register the parser to be used.
func (r *Repository) RegisterParser(url string) error {
	host, err := getHost(url)
	if err != nil {
		return err
	}

	var bb internal.ParserManager = bitbucket.New(r.logger)
	var gh internal.ParserManager = github.New(r.logger)
	var gl internal.ParserManager = gitlab.New(r.logger)
	r.SetURL(url)

	// add additional parsers
	if bb.ShouldParse(host) {
		r.SetParser(bb)
	} else if gh.ShouldParse(host) {
		r.SetParser(gh)
	} else if gl.ShouldParse(host) {
		r.SetParser(gl)
	} else {
		return fmt.Errorf("could not find parser for host: %s", host)
	}

	return nil
}

// Parse the URL via the delegated parser.
func (r *Repository) Parse() (pkg.RepositoryManager, error) {
	url := r.GetURL()

	return r.parser.Parse(url)
}

// SetParser set the parser to be used.
func (r *Repository) SetParser(parser internal.ParserManager) { r.parser = parser }

// GetParser getthe parser to be used.
func (r *Repository) GetParser() internal.ParserManager { return r.parser }

// SetURL set the URL to be parsed.
func (r *Repository) SetURL(url string) { r.url = url }

// GetURL get the URL to be parsed.
func (r *Repository) GetURL() string { return r.url }

// getHost take the provided URL and return the `Host` property; delegated to
// the `git-urls` package.
func getHost(url string) (string, error) {
	// This function will never return an error as each parser's error is dropped
	// for the next parser, and the final parser is simply net/url's `URL` type.
	parsedURL, _ := giturls.Parse(url)
	if parsedURL.Host == "" {
		return "", fmt.Errorf("could parse url for host: %s", url)
	}

	return parsedURL.Host, nil
}
