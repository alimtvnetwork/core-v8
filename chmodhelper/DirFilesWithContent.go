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

package chmodhelper

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/pathinternal"
)

type DirFilesWithContent struct {
	Dir         string
	Files       []FileWithContent
	DirFileMode os.FileMode
}

func (it *DirFilesWithContent) GetPaths() []string {
	collection := corestr.New.Collection.Cap(
		constants.ArbitraryCapacity50,
	)

	for _, file := range it.Files {
		compiledPath := path.Join(it.Dir, file.RelativePath)
		collection.Add(compiledPath)
	}

	return collection.List()
}

func (it *DirFilesWithContent) GetFilesChmodMap() *corestr.Hashmap {
	files := it.GetPaths()

	hashmap, err := GetFilesChmodRwxFullMap(files)

	errcore.SimpleHandleErr(
		err,
		"GetFilesChmodMap() failed to retrieve hashmap from file paths",
	)

	return hashmap
}

func (it *DirFilesWithContent) Create(
	isRemoveBeforeCreate bool,
) error {
	var sliceErr []string
	fileCreatorFunc := fileWriter{}.Bytes.Chmod

	removeDirErr := removeDirIf(
		isRemoveBeforeCreate,
		it.Dir,
		"Create",
	)

	if removeDirErr != nil {
		sliceErr = append(
			sliceErr,
			removeDirErr.Error(),
		)
	}

	for _, file := range it.Files {
		compiledPath := pathinternal.Join(
			it.Dir,
			file.RelativePath,
		)

		err := fileCreatorFunc(
			isRemoveBeforeCreate,
			it.DirFileMode,
			file.FileMode,
			compiledPath,
			file.ContentToBytes(),
		)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error(),
			)
		}
	}

	if len(sliceErr) == 0 {
		return nil
	}

	// has error
	toString := strings.Join(sliceErr, constants.NewLineUnix)

	return errors.New(toString)
}
