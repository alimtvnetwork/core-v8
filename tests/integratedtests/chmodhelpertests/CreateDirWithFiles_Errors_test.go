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

package chmodhelpertests

import (
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core-v8/chmodhelper"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ── CreateDirWithFiles: file-create error path (line 62) and chmod error (line 75) ──

func Test_CreateDirWithFiles_FileCreateError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	invalidDir := string([]byte{0}) + "/noexist/sub"
	dwf := &chmodhelper.DirWithFiles{
		Dir: invalidDir,
		Files: []string{
			"a.txt",
		},
	}

	// Act
	err := chmodhelper.CreateDirWithFiles(false, 0o644, dwf)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "CreateDirWithFiles file-create error", expected)
}

// ── GetRecursivePathsContinueOnError: walk error appended (line 47-51) ──

func Test_GetRecursivePathsContinueOnError_WalkError(t *testing.T) {
	// Arrange
	invalidPath := string([]byte{0})

	// Act
	paths, err := chmodhelper.GetRecursivePathsContinueOnError(invalidPath)

	// Assert
	actual := args.Map{
		"pathsLen": len(paths),
		"hasError": err != nil,
	}
	expected := args.Map{
		"pathsLen": 0,
		"hasError": true,
	}
	actual.ShouldBeEqual(t, 1, "GetRecursivePathsContinueOnError walk error", expected)
}

// ── MergeRwxWildcardWithFixedRwx: bad existing rwx length (line 38-40) ──

func Test_MergeRwxWildcardWithFixedRwx_BadExistingLength(t *testing.T) {
	// Arrange
	wildcardInput := "r-x"
	existingBadRwx := "rw" // wrong length

	// Act
	result, err := chmodhelper.MergeRwxWildcardWithFixedRwx(wildcardInput, existingBadRwx)

	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasError":  err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasError":  true,
	}
	actual.ShouldBeEqual(t, 1, "MergeRwxWildcardWithFixedRwx bad existing length", expected)
}
