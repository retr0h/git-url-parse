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

package repository_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/retr0h/git-url-parse/internal"
	"github.com/retr0h/git-url-parse/internal/repositories/bitbucket"
	"github.com/retr0h/git-url-parse/internal/repositories/github"
	"github.com/retr0h/git-url-parse/internal/repositories/gitlab"
	"github.com/retr0h/git-url-parse/pkg/repository"
)

type RepositoryPublicTestSuite struct {
	suite.Suite

	r *repository.Repository

	logger *slog.Logger
}

func (suite *RepositoryPublicTestSuite) SetupTest() {
	suite.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	suite.r = repository.New(suite.logger)
}

func (suite *RepositoryPublicTestSuite) TestRegisterParser() {
	type test struct {
		input   string
		want    internal.ParserManager
		wantErr bool
	}

	// add additional parser tests
	tests := []test{
		{
			input:   "https://bitbucket.org/retr0h/foo",
			want:    &bitbucket.Bitbucket{},
			wantErr: false,
		},
		{
			input:   "https://github.com/retr0h/foo",
			want:    &github.GitHub{},
			wantErr: false,
		},
		{
			input:   "https://gitlab.com/retr0h/foo",
			want:    &gitlab.GitLab{},
			wantErr: false,
		},
		// failure cases
		{
			input:   "invalid giturls host",
			want:    nil,
			wantErr: true,
		},
		{
			input:   "https://example.com/retr0h/foo",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		err := suite.r.RegisterParser(tc.input)

		if tc.wantErr {
			assert.Error(suite.T(), err)
		} else {
			got := suite.r.GetParser()
			assert.IsType(suite.T(), tc.want, got)
			assert.NoError(suite.T(), err)
		}
	}
}

func (suite *RepositoryPublicTestSuite) TestParse() {
	type test struct {
		input   string
		want    string
		wantErr bool
	}

	// add additional parser tests
	tests := []test{
		{
			input:   "https://github.com/retr0h/foo",
			want:    "github",
			wantErr: false,
		},
		// failure cases
		{
			input:   "https://github.com/",
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		err := suite.r.RegisterParser(tc.input)
		assert.NoError(suite.T(), err)
		got, err := suite.r.Parse()

		if tc.wantErr {
			assert.Error(suite.T(), err)
		} else {
			assert.Equal(suite.T(), tc.want, got.GetProviderName())
			assert.NoError(suite.T(), err)

		}
	}
}

// In order for `go test` to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestRepositoyPublicTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryPublicTestSuite))
}
