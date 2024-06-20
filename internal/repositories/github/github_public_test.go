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
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/retr0h/git-url-parse/internal"
	"github.com/retr0h/git-url-parse/internal/repositories/github"
)

type GitHubPublicTestSuite struct {
	suite.Suite

	rm internal.ParserManager

	logger *slog.Logger
}

func (suite *GitHubPublicTestSuite) SetupTest() {
	suite.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	suite.rm = github.New(suite.logger)
}

func (suite *GitHubPublicTestSuite) TestIsGitHub() {
	type test struct {
		input string
		want  bool
	}

	tests := []test{
		{
			input: "github.com",
			want:  true,
		},
		{
			input: "raw.githubusercontent.com",
			want:  true,
		},
		{
			input: "www.github.com",
			want:  true,
		},
		{
			input: "example.com",
			want:  false,
		},
		{
			input: "example",
			want:  false,
		},
		{
			input: ".com",
			want:  false,
		},
	}

	for _, tc := range tests {
		got := suite.rm.IsGitHub(tc.input)

		assert.Equal(suite.T(), got, tc.want)
	}
}

// In order for `go test` to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestGitHubPublicTestSuite(t *testing.T) {
	suite.Run(t, new(GitHubPublicTestSuite))
}
