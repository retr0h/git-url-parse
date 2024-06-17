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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/retr0h/git-url-parse/internal"
	"github.com/retr0h/git-url-parse/internal/github"
)

type ResponsePublicTestSuite struct {
	suite.Suite

	rum internal.RepositoryURLManager

	branch   string
	host     string
	href     string
	owner    string
	path     string
	provider string
	repo     string
	resource string
}

func (suite *ResponsePublicTestSuite) SetupTest() {
	suite.branch = "branch"
	suite.host = "host"
	suite.href = "href"
	suite.owner = "owner"
	suite.path = "path"
	suite.provider = "github"
	suite.repo = "repo"
	suite.resource = "resource"

	suite.rum = &github.URL{
		Branch: suite.branch,
		Host:   suite.host,
		HREF:   suite.href,
		Owner:  suite.owner,
		Path:   suite.path,
		// Provider: suite.provider,
		Repo:     suite.repo,
		Resource: suite.resource,
	}
}

func (suite *ResponsePublicTestSuite) TestGetBranchNameOk() {
	got := suite.rum.GetBranchName()

	assert.Equal(suite.T(), suite.branch, got)
}

func (suite *ResponsePublicTestSuite) TestGetHostNameOk() {
	got := suite.rum.GetHostName()

	assert.Equal(suite.T(), suite.host, got)
}

func (suite *ResponsePublicTestSuite) TestGetHREFOk() {
	got := suite.rum.GetHREF()

	assert.Equal(suite.T(), suite.href, got)
}

func (suite *ResponsePublicTestSuite) TestGetOwnerNameOk() {
	got := suite.rum.GetOwnerName()

	assert.Equal(suite.T(), suite.owner, got)
}

func (suite *ResponsePublicTestSuite) TestGetPathOk() {
	got := suite.rum.GetPath()

	assert.Equal(suite.T(), suite.path, got)
}

func (suite *ResponsePublicTestSuite) TestGetProviderNameOk() {
	got := suite.rum.GetProviderName()

	assert.Equal(suite.T(), suite.provider, got)
}

func (suite *ResponsePublicTestSuite) TestGetRepoNameOk() {
	got := suite.rum.GetRepoName()

	assert.Equal(suite.T(), suite.repo, got)
}

func (suite *ResponsePublicTestSuite) TestGetResourceNameOk() {
	got := suite.rum.GetResourceName()

	assert.Equal(suite.T(), suite.resource, got)
}

// In order for `go test` to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestResponsePublicTestSuite(t *testing.T) {
	suite.Run(t, new(ResponsePublicTestSuite))
}
