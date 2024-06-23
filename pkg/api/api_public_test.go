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

package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/retr0h/git-url-parse/pkg"
	"github.com/retr0h/git-url-parse/pkg/api"
)

type APIPublicTestSuite struct {
	suite.Suite

	rm pkg.RepositoryManager

	branch    string
	host      string
	href      string
	owner     string
	path      string
	protocol  string
	protocols []string
	provider  string
	repo      string
	resource  string
}

func (suite *APIPublicTestSuite) SetupTest() {
	suite.branch = "branch"
	suite.host = "host"
	suite.href = "href"
	suite.owner = "owner"
	suite.path = "path"
	suite.protocol = "protocol"
	suite.protocols = []string{"protocol"}
	suite.provider = "provider"
	suite.repo = "repo"
	suite.resource = "resource"

	suite.rm = &api.Repository{
		Branch:   suite.branch,
		Host:     suite.host,
		HREF:     suite.href,
		Owner:    suite.owner,
		Path:     suite.path,
		Protocol: suite.protocol,
		Provider: suite.provider,
		Repo:     suite.repo,
		Resource: suite.resource,
	}
}

func (suite *APIPublicTestSuite) TestGetBranchNameOk() {
	got := suite.rm.GetBranchName()

	assert.Equal(suite.T(), suite.branch, got)
}

func (suite *APIPublicTestSuite) TestGetHostNameOk() {
	got := suite.rm.GetHostName()

	assert.Equal(suite.T(), suite.host, got)
}

func (suite *APIPublicTestSuite) TestGetHREFOk() {
	got := suite.rm.GetHREF()

	assert.Equal(suite.T(), suite.href, got)
}

func (suite *APIPublicTestSuite) TestGetOwnerNameOk() {
	got := suite.rm.GetOwnerName()

	assert.Equal(suite.T(), suite.owner, got)
}

func (suite *APIPublicTestSuite) TestGetPathOk() {
	got := suite.rm.GetPath()

	assert.Equal(suite.T(), suite.path, got)
}

func (suite *APIPublicTestSuite) TestGetProtocolOk() {
	got := suite.rm.GetProtocol()

	assert.Equal(suite.T(), suite.protocol, got)
}

func (suite *APIPublicTestSuite) TestGetProtocolsOk() {
	got := suite.rm.GetProtocols()

	assert.Equal(suite.T(), suite.protocols, got)
}

func (suite *APIPublicTestSuite) TestGetProtocolsWithMultiProtocolsOk() {
	suite.rm = &api.Repository{
		Protocol: "git+ssh",
	}
	got := suite.rm.GetProtocols()

	assert.Equal(suite.T(), []string{"git", "ssh"}, got)
}

func (suite *APIPublicTestSuite) TestGetProviderNameOk() {
	got := suite.rm.GetProviderName()

	assert.Equal(suite.T(), suite.provider, got)
}

func (suite *APIPublicTestSuite) TestGetRepoNameOk() {
	got := suite.rm.GetRepoName()

	assert.Equal(suite.T(), suite.repo, got)
}

func (suite *APIPublicTestSuite) TestGetResourceNameOk() {
	got := suite.rm.GetResourceName()

	assert.Equal(suite.T(), suite.resource, got)
}

// In order for `go test` to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestAPIPublicTestSuite(t *testing.T) {
	suite.Run(t, new(APIPublicTestSuite))
}
