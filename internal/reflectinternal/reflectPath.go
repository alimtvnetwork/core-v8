// MIT License
// 
// Copyright (c) 2020–2026
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package reflectinternal

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/pathinternal"
)

type reflectPath struct{}

func (it reflectPath) RepoDir() string {
	if repoDir != nil {
		return *repoDir
	}

	repoDirLocal := it.repoDirRecursive(
		it.CurDir(),
		20,
	)

	repoDir = &repoDirLocal

	return repoDirLocal
}

func (it reflectPath) repoDirRecursive(curDir string, try int) string {
	if curDir == "" {
		return ""
	}

	modFile := path.Join(curDir, "go.mod")
	gitFolder := path.Join(curDir, ".git")

	if it.isPathsExists(modFile, gitFolder) {
		return curDir
	}

	if it.isPathsExists(gitFolder) {
		return curDir
	}

	if try <= 0 {
		return curDir
	}

	parentDir := it.parentDir(curDir)

	return it.repoDirRecursive(parentDir, try-1)
}

func (it reflectPath) isPathsExists(multiPaths ...string) bool {
	if len(multiPaths) == 0 {
		return false
	}

	for _, singlePath := range multiPaths {
		if !pathinternal.IsPathExists(singlePath) {
			return false
		}
	}

	return true
}

func (it reflectPath) parentDir(curPath string) string {
	return filepath.Dir(curPath)
}

func (it reflectPath) pathName(curPath string) string {
	_, name := filepath.Split(curPath)

	return name
}

func (it reflectPath) CurDir() string {
	return it.CurDirSkipStack(defaultInternalSkip)
}

func (it reflectPath) CurDirSkipStack(skipStack int) string {
	_, filePath, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return filepath.Dir(filePath)
	}

	return constants.EmptyString
}
