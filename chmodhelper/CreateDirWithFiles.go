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
	"os"
	"path"

	"github.com/alimtvnetwork/core/errcore"
)

func CreateDirWithFiles(
	isRemoveAllDirBeforeCreate bool,
	fileChmod os.FileMode,
	dirWithFile *DirWithFiles,
) error {
	const funcName = "CreateDirWithFiles"
	dir := dirWithFile.Dir

	removeDirErr := removeDirIf(
		isRemoveAllDirBeforeCreate,
		dir,
		funcName,
	)

	if removeDirErr != nil {
		return removeDirErr
	}

	mkDirErr := os.MkdirAll(
		dir, dirDefaultChmod,
	)

	if mkDirErr != nil {
		return errcore.PathMeaningfulError(
			errcore.PathCreateFailedType,
			mkDirErr,
			dir,
		)
	}

	var fileManipulateErr error

	if len(dirWithFile.Files) == 0 {
		return nil
	}

	for _, filePath := range dirWithFile.Files {
		compiledPath := path.Join(dir, filePath)
		osFile, err := os.Create(compiledPath)

		if err != nil {
			return errcore.PathMeaningfulError(
				errcore.PathCreateFailedType,
				err,
				dir,
			)
		}

		if osFile != nil {
			fileManipulateErr = osFile.Close()
		}

		if fileManipulateErr != nil {
			return errcore.PathMeaningfulError(
				errcore.FileCloseFailedType,
				fileManipulateErr,
				compiledPath,
			)
		}

		chmodErr := os.Chmod(
			compiledPath,
			fileChmod,
		)

		if chmodErr != nil {
			return errcore.PathMeaningfulError(
				errcore.PathChmodApplyType,
				chmodErr,
				compiledPath,
			)
		}
	}

	return nil
}
