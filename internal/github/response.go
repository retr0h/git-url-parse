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

// GetBranchName the repo's branch name.
func (u *URL) GetBranchName() string {
	return u.Branch
}

// GetHostName the repo's domain.
func (u *URL) GetHostName() string {
	return u.Host
}

// GetHREF the repo's HREF.
func (u *URL) GetHREF() string {
	return u.HREF
}

// GetOwnerName the repo's owner.
func (u *URL) GetOwnerName() string {
	return u.Owner
}

// GetPath the repo's path.
func (u *URL) GetPath() string {
	return u.Path
}

// GetProtocol the repo's protocol.
func (u *URL) GetProtocol() string {
	return u.Protocol
}

// GetProviderName the repo's provider.
func (u *URL) GetProviderName() string {
	return "github"
}

// GetRepoName the repo's name.
func (u *URL) GetRepoName() string {
	return u.Repo
}

// GetResourceName the repo's resource name aka host name.
func (u *URL) GetResourceName() string {
	return u.Resource
}
