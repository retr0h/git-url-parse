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

package github_test

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/retr0h/git-url-parse/internal"
	"github.com/retr0h/git-url-parse/internal/github"
)

type GitHubPublicTestSuite struct {
	suite.Suite

	rm internal.RepositoryManager

	logger *slog.Logger
}

func (suite *GitHubPublicTestSuite) SetupTest() {
	suite.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	suite.rm = github.New(suite.logger)
}

func (suite *GitHubPublicTestSuite) TestParseOk() {
	type url struct {
		branch   string
		href     string
		owner    string
		path     string
		protocol string
		provider string
		repo     string
		resource string
	}

	type test struct {
		input string
		want  *url
	}

	tests := []test{
		{
			input: "https://github.com/retr0h/git-url-parse",
			want: &url{
				protocol: "https",
				resource: "github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "",
				branch:   "",
				provider: "github",
				href:     "https://github.com/retr0h/git-url-parse",
			},
		},
		{
			input: "https://github.com/retr0h/git-url-parse/blob/main/files/file0.json",
			want: &url{
				protocol: "https",
				resource: "github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "/files/file0.json",
				branch:   "main",
				provider: "github",
				href:     "https://github.com/retr0h/git-url-parse/blob/main/files/file0.json",
			},
		},
		{
			input: "https://github.com/retr0h/git-url-parse/tree/main/files",
			want: &url{
				protocol: "https",
				resource: "github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "/files",
				branch:   "main",
				provider: "github",
				href:     "https://github.com/retr0h/git-url-parse/tree/main/files",
			},
		},
		{
			input: "https://raw.githubusercontent.com/retr0h/git-url-parse/main/files/file0.json",
			want: &url{
				protocol: "https",
				resource: "raw.githubusercontent.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "/files/file0.json",
				branch:   "main",
				provider: "github",
				href:     "https://raw.githubusercontent.com/retr0h/git-url-parse/main/files/file0.json",
			},
		},
		{
			input: "https://www.github.com/retr0h/git-url-parse",
			want: &url{
				protocol: "https",
				resource: "www.github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "",
				branch:   "",
				provider: "github",
				href:     "https://www.github.com/retr0h/git-url-parse",
			},
		},
		{
			input: "git@github.com:retr0h/git-url-parse.git",
			want: &url{
				protocol: "git",
				resource: "github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "",
				branch:   "",
				provider: "github",
				href:     "git@github.com:retr0h/git-url-parse.git",
			},
		},
	}

	for _, tc := range tests {
		var got internal.RepositoryURLManager
		got, err := suite.rm.Parse(tc.input)

		assert.Equal(suite.T(), tc.want.protocol, got.GetProtocol())
		assert.Equal(suite.T(), tc.want.resource, got.GetResourceName())
		assert.Equal(suite.T(), tc.want.href, got.GetHREF())
		assert.Equal(suite.T(), tc.want.owner, got.GetOwnerName())
		assert.Equal(suite.T(), tc.want.repo, got.GetRepoName())
		assert.Equal(suite.T(), tc.want.path, got.GetPath())
		assert.Equal(suite.T(), tc.want.branch, got.GetBranchName())
		assert.Equal(suite.T(), tc.want.provider, got.GetProviderName())
		assert.NoError(suite.T(), err)
	}
}

func (suite *GitHubPublicTestSuite) TestParseFails() {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{
			input: "https://github.com/",
			want:  fmt.Sprintf("could match url: %s to any pattern", "https://github.com/"),
		},
		{
			input: "bogus://url/",
			want:  fmt.Sprintf("could match url: %s to any pattern", "bogus://url/"),
		},
		{
			input: "git@github.com:foobar/retr0h/git-url-parse.git",
			want: fmt.Sprintf(
				"could match url: %s to any pattern",
				"git@github.com:foobar/retr0h/git-url-parse.git",
			),
		},
	}

	for _, tc := range tests {
		_, err := suite.rm.Parse(tc.input)

		assert.EqualError(suite.T(), err, tc.want)
	}
}

// In order for `go test` to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestGitHubPublicTestSuite(t *testing.T) {
	suite.Run(t, new(GitHubPublicTestSuite))
}
