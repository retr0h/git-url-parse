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

package github

import (
	"fmt"
	"log/slog"
	"regexp"
)

// ChatGPT-4 generated regexp
var patterns = []string{
	`^(?P<scheme>https?)://(?P<resource>[^/]+)/(?P<owner>[^/]+)/(?P<repo>[^/]+)(/(?:tree|blob)/(?P<branch>[^/]+)(?P<path>/.*)?)?$`,
	`^(?P<scheme>https?)://(?P<resource>raw\.githubusercontent\.com)/(?P<owner>[^/]+)/(?P<repo>[^/]+)/(?P<branch>[^/]+)(?P<path>/.*)$`,
}

// New factory to create a new GitHub instance.
func New(
	logger *slog.Logger,
) *GitHub {
	return &GitHub{
		logger: logger,
	}
}

// Parse the provided GitHub URL.
func (gh *GitHub) Parse(url string) (*URL, error) {
	response := &URL{}
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(url)
		mm := makeMatchMap(re, matches)

		gh.logger.Debug(
			"matching url",
			slog.String("url", url),
			slog.String("regexp", pattern),
		)

		if matches != nil {
			response = &URL{
				Protocol: mm["scheme"],
				Host:     mm["resource"],
				Resource: mm["resource"],
				Owner:    mm["owner"],
				Repo:     mm["repo"],
				Path:     mm["path"],
				Branch:   mm["branch"],
				HREF:     url,
			}
		}
	}

	if (URL{}) == *response {
		return nil, fmt.Errorf("could match url: %s to any pattern", url)
	}

	return response, nil
}

func makeMatchMap(re *regexp.Regexp, matches []string) map[string]string {
	mm := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" && i < len(matches) {
			mm[name] = matches[i]
		}
	}

	return mm
}
