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
	"github.com/retr0h/git-url-parse/pkg"
)

type ParserPublicTestSuite struct {
	suite.Suite

	rm internal.ParserManager

	logger *slog.Logger
}

func (suite *ParserPublicTestSuite) SetupTest() {
	suite.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	suite.rm = github.New(suite.logger)
}

func (suite *ParserPublicTestSuite) TestParse() {
	type repository struct {
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
		input   string
		want    *repository
		wantErr bool
	}

	tests := []test{
		{
			input: "https://github.com/retr0h/git-url-parse",
			want: &repository{
				protocol: "https",
				resource: "github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "",
				branch:   "",
				provider: "github",
				href:     "https://github.com/retr0h/git-url-parse",
			},
			wantErr: false,
		},
		{
			input: "https://github.com/retr0h/git-url-parse/blob/main/files/file0.json",
			want: &repository{
				protocol: "https",
				resource: "github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "/files/file0.json",
				branch:   "main",
				provider: "github",
				href:     "https://github.com/retr0h/git-url-parse/blob/main/files/file0.json",
			},
			wantErr: false,
		},
		{
			input: "https://github.com/retr0h/git-url-parse/tree/main/files",
			want: &repository{
				protocol: "https",
				resource: "github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "/files",
				branch:   "main",
				provider: "github",
				href:     "https://github.com/retr0h/git-url-parse/tree/main/files",
			},
			wantErr: false,
		},
		{
			input: "https://raw.githubusercontent.com/retr0h/git-url-parse/main/files/file0.json",
			want: &repository{
				protocol: "https",
				resource: "raw.githubusercontent.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "/files/file0.json",
				branch:   "main",
				provider: "github",
				href:     "https://raw.githubusercontent.com/retr0h/git-url-parse/main/files/file0.json",
			},
			wantErr: false,
		},
		{
			input: "https://www.github.com/retr0h/git-url-parse",
			want: &repository{
				protocol: "https",
				resource: "www.github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "",
				branch:   "",
				provider: "github",
				href:     "https://www.github.com/retr0h/git-url-parse",
			},
			wantErr: false,
		},
		{
			input: "git@github.com:retr0h/git-url-parse.git",
			want: &repository{
				protocol: "git",
				resource: "github.com",
				owner:    "retr0h",
				repo:     "git-url-parse",
				path:     "",
				branch:   "",
				provider: "github",
				href:     "git@github.com:retr0h/git-url-parse.git",
			},
			wantErr: false,
		},
		// failure cases
		{
			input:   "https://github.com/",
			want:    &repository{},
			wantErr: true,
		},
		{
			input:   "bogus://url/",
			want:    &repository{},
			wantErr: true,
		},
		{
			input:   "git@github.com:foobar/retr0h/git-url-parse.git",
			want:    &repository{},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		var got pkg.RepositoryManager
		got, err := suite.rm.Parse(tc.input)

		if tc.wantErr {
			assert.Error(suite.T(), err)
		} else {
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
}

// In order for `go test` to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestParserPublicTestSuite(t *testing.T) {
	suite.Run(t, new(ParserPublicTestSuite))
}
