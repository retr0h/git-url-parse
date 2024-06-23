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

package repositories

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoriesTestSuite struct {
	suite.Suite
}

func (suite *RepositoriesTestSuite) SetupTest() {}

func (suite *RepositoriesTestSuite) TestMakeMatchMap() {
	type input struct {
		pattern string
		str     string
	}

	type test struct {
		input *input
		want  map[string]string
	}

	tests := []test{
		{
			input: &input{
				pattern: `^(?P<alpha>[a-zA-Z]*)(?P<numeric>[0-9]*)(?P<special>[^a-zA-Z0-9]*)$`,
				str:     "abc123!@#",
			},
			want: map[string]string{
				"alpha":   "abc",
				"numeric": "123",
				"special": "!@#",
			},
		},
		{
			input: &input{
				pattern: `^(?P<alpha>[a-zA-Z]*)(?P<numeric>[0-9]*)(?P<special>[^a-zA-Z0-9]*)?$`,
				str:     "abc123",
			},
			want: map[string]string{
				"alpha":   "abc",
				"numeric": "123",
				"special": "",
			},
		},
		// failure cases
		{
			input: &input{
				pattern: `^(?P<alpha>[a-zA-Z]*)(?P<numeric>[0-9]*)(?P<special>[^a-zA-Z0-9]*)$`,
				str:     "abc123abc",
			},
			want: map[string]string{},
		},
	}

	for _, tc := range tests {
		re := regexp.MustCompile(tc.input.pattern)
		matches := re.FindStringSubmatch(tc.input.str)
		got := MakeMatchMap(re, matches)

		assert.Equal(suite.T(), tc.want, got)
	}
}

// In order for `go test` to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestRepositoriesTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoriesTestSuite))
}
