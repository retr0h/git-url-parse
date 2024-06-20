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

package api

// GetBranchName the repo's branch name.
func (r *Repository) GetBranchName() string {
	return r.Branch
}

// GetHostName the repo's domain.
func (r *Repository) GetHostName() string {
	return r.Host
}

// GetHREF the repo's HREF.
func (r *Repository) GetHREF() string {
	return r.HREF
}

// GetOwnerName the repo's owner.
func (r *Repository) GetOwnerName() string {
	return r.Owner
}

// GetPath the repo's path.
func (r *Repository) GetPath() string {
	return r.Path
}

// GetProtocol the repo's protocol.
func (r *Repository) GetProtocol() string {
	return r.Protocol
}

// GetProviderName the repo's provider.
func (r *Repository) GetProviderName() string {
	return r.Provider
}

// GetRepoName the repo's name.
func (r *Repository) GetRepoName() string {
	return r.Repo
}

// GetResourceName the repo's resource name aka host name.
func (r *Repository) GetResourceName() string {
	return r.Resource
}
