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

package bitbucket_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/retr0h/git-url-parse/internal"
	"github.com/retr0h/git-url-parse/internal/repositories/bitbucket"
	"github.com/retr0h/git-url-parse/pkg"
)

type ParserPublicTestSuite struct {
	suite.Suite

	rm internal.ParserManager

	logger *slog.Logger
}

func (suite *ParserPublicTestSuite) SetupTest() {
	suite.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	suite.rm = bitbucket.New(suite.logger)
}

func (suite *ParserPublicTestSuite) TestParse() {
	type repository struct {
		branch    string
		href      string
		owner     string
		path      string
		protocol  string
		protocols []string
		provider  string
		repo      string
		resource  string
	}

	type test struct {
		input   string
		want    *repository
		wantErr bool
	}

	tests := []test{
		{
			input: "https://bitbucket.org/owner/repository",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "bitbucket.org",
				owner:     "owner",
				repo:      "repository",
				path:      "",
				branch:    "",
				provider:  "bitbucket",
				href:      "https://bitbucket.org/owner/repository",
			},
			wantErr: false,
		},
		{
			input: "https://bitbucket.org/owner/repository/src/main/rules/etcd-encryption-native/raw.rego",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "bitbucket.org",
				owner:     "owner",
				repo:      "repository",
				path:      "rules/etcd-encryption-native/raw.rego",
				branch:    "main",
				provider:  "bitbucket",
				href:      "https://bitbucket.org/owner/repository/src/main/rules/etcd-encryption-native/raw.rego",
			},
			wantErr: false,
		},
		{
			input: "https://bitbucket.org/owner/repository/src/dev/README.md",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "bitbucket.org",
				owner:     "owner",
				repo:      "repository",
				path:      "README.md",
				branch:    "dev",
				provider:  "bitbucket",
				href:      "https://bitbucket.org/owner/repository/src/dev/README.md",
			},
			wantErr: false,
		},
		{
			input: "https://bitbucket.org/owner/repository/src/dev/",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "bitbucket.org",
				owner:     "owner",
				repo:      "repository",
				path:      "",
				branch:    "dev",
				provider:  "bitbucket",
				href:      "https://bitbucket.org/owner/repository/src/dev/",
			},
			wantErr: false,
		},
		{
			input: "https://bitbucket.org/owner/repository/src/v1.0.178/README.md",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "bitbucket.org",
				owner:     "owner",
				repo:      "repository",
				path:      "README.md",
				branch:    "v1.0.178",
				provider:  "bitbucket",
				href:      "https://bitbucket.org/owner/repository/src/v1.0.178/README.md",
			},
			wantErr: false,
		},
		{
			input: "https://bitbucket.org/owner/repository/raw/4502b9b51ee3ac1ea649bacfa0f48ebdeab05f4a/README.md",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "bitbucket.org",
				owner:     "owner",
				repo:      "repository",
				path:      "README.md",
				branch:    "4502b9b51ee3ac1ea649bacfa0f48ebdeab05f4a",
				provider:  "bitbucket",
				href:      "https://bitbucket.org/owner/repository/raw/4502b9b51ee3ac1ea649bacfa0f48ebdeab05f4a/README.md",
			},
			wantErr: false,
		},
		// scp-like syntax supported by git for ssh
		// see: https://mirrors.edge.kernel.org/pub/software/scm/git/docs/git-clone.html#URLS
		// regular form
		{
			input: "git@bitbucket.org:owner/repository.git",
			want: &repository{
				protocol:  "git",
				protocols: []string{"git"},
				resource:  "bitbucket.org",
				owner:     "owner",
				repo:      "repository",
				path:      "",
				branch:    "",
				provider:  "bitbucket",
				href:      "git@bitbucket.org:owner/repository.git",
			},
			wantErr: false,
		},
		// failure cases
		{
			input:   "https://bitbucket.org/",
			want:    &repository{},
			wantErr: true,
		},
		{
			input:   "git@bitbucket.org:owner/to/repository.git",
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
			require.NoError(suite.T(), err)
			assert.Equal(suite.T(), tc.want.protocol, got.GetProtocol())
			assert.Equal(suite.T(), tc.want.protocols, got.GetProtocols())
			assert.Equal(suite.T(), tc.want.resource, got.GetResourceName())
			assert.Equal(suite.T(), tc.want.href, got.GetHREF())
			assert.Equal(suite.T(), tc.want.owner, got.GetOwnerName())
			assert.Equal(suite.T(), tc.want.repo, got.GetRepoName())
			assert.Equal(suite.T(), tc.want.path, got.GetPath())
			assert.Equal(suite.T(), tc.want.branch, got.GetBranchName())
			assert.Equal(suite.T(), tc.want.provider, got.GetProviderName())
		}
	}
}

// In order for `go test` to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestParserPublicTestSuite(t *testing.T) {
	suite.Run(t, new(ParserPublicTestSuite))
}
