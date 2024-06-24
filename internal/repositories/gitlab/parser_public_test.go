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

package gitlab_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/retr0h/git-url-parse/internal"
	"github.com/retr0h/git-url-parse/internal/repositories/gitlab"
	"github.com/retr0h/git-url-parse/pkg"
)

type ParserPublicTestSuite struct {
	suite.Suite

	rm internal.ParserManager

	logger *slog.Logger
}

func (suite *ParserPublicTestSuite) SetupTest() {
	suite.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	suite.rm = gitlab.New(suite.logger)
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
			input: "https://gitlab.com/owner/repository",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "gitlab.com",
				owner:     "owner",
				repo:      "repository",
				path:      "",
				branch:    "",
				provider:  "gitlab",
				href:      "https://gitlab.com/owner/repository",
			},
			wantErr: false,
		},
		{
			input: "https://gitlab.com/owner/repository/-/blob/main/stable/acs-engine-autoscaler/Chart.yaml",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "gitlab.com",
				owner:     "owner",
				repo:      "repository",
				path:      "stable/acs-engine-autoscaler/Chart.yaml",
				branch:    "main",
				provider:  "gitlab",
				href:      "https://gitlab.com/owner/repository/-/blob/main/stable/acs-engine-autoscaler/Chart.yaml",
			},
			wantErr: false,
		},
		{
			input: "https://gitlab.com/owner/repository/-/blob/dev/README.md",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "gitlab.com",
				owner:     "owner",
				repo:      "repository",
				path:      "README.md",
				branch:    "dev",
				provider:  "gitlab",
				href:      "https://gitlab.com/owner/repository/-/blob/dev/README.md",
			},
			wantErr: false,
		},
		{
			input: "https://gitlab.com/owner/repository/-/tree/dev",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "gitlab.com",
				owner:     "owner",
				repo:      "repository",
				path:      "",
				branch:    "dev",
				provider:  "gitlab",
				href:      "https://gitlab.com/owner/repository/-/tree/dev",
			},
			wantErr: false,
		},
		{
			input: "https://gitlab.com/owner/repository/-/blob/v0.0.0/README.md",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "gitlab.com",
				owner:     "owner",
				repo:      "repository",
				path:      "README.md",
				branch:    "v0.0.0",
				provider:  "gitlab",
				href:      "https://gitlab.com/owner/repository/-/blob/v0.0.0/README.md",
			},
			wantErr: false,
		},
		{
			input: "https://gitlab.com/owner/repository/-/raw/main/stable/acs-engine-autoscaler/Chart.yaml",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "gitlab.com",
				owner:     "owner",
				repo:      "repository",
				path:      "stable/acs-engine-autoscaler/Chart.yaml",
				branch:    "main",
				provider:  "gitlab",
				href:      "https://gitlab.com/owner/repository/-/raw/main/stable/acs-engine-autoscaler/Chart.yaml",
			},
			wantErr: false,
		},
		{
			input: "https://gitlab.com/owner/subgroup/repository.git",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "gitlab.com",
				owner:     "owner",
				repo:      "repository",
				path:      "",
				branch:    "",
				provider:  "gitlab",
				href:      "https://gitlab.com/owner/subgroup/repository.git",
			},
			wantErr: false,
		},
		{
			input: "https://gitlab.com/owner/subgroup/subsubgroup/repository.git",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "gitlab.com",
				owner:     "owner",
				repo:      "repository",
				path:      "",
				branch:    "",
				provider:  "gitlab",
				href:      "https://gitlab.com/owner/subgroup/subsubgroup/repository.git",
			},
			wantErr: false,
		},
		{
			input: "https://gitlab.example.com/owner/repository",
			want: &repository{
				protocol:  "https",
				protocols: []string{"https"},
				resource:  "gitlab.example.com",
				owner:     "owner",
				repo:      "repository",
				path:      "",
				branch:    "",
				provider:  "gitlab",
				href:      "https://gitlab.example.com/owner/repository",
			},
			wantErr: false,
		},

		{
			input: "git@gitlab.com:owner/repository.git",
			want: &repository{
				protocol:  "git",
				protocols: []string{"git"},
				resource:  "gitlab.com",
				owner:     "owner",
				repo:      "repository",
				path:      "",
				branch:    "",
				provider:  "gitlab",
				href:      "git@gitlab.com:owner/repository.git",
			},
			wantErr: false,
		},
		// failure cases
		{
			input:   "https://gitlab.com/",
			want:    &repository{},
			wantErr: true,
		},
		{
			input:   "bogus://url/",
			want:    &repository{},
			wantErr: true,
		},
		{
			input:   "git@gitlab.com:foobar/owner/repository.git",
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
