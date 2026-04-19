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

import "os"

type fileStringWriter struct{}

// All
//
//	Writes contents to file system.
//
// parentDirPath:
//   - is a full path to the parent dir for checking
//     if parent dir exist if not then created
//
// writingFilePath:
//   - is a full path to the actual file where to write contents
func (it fileStringWriter) All(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	content string,
) error {
	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		isApplyChmodMust,
		isApplyChmodOnMismatch,
		isCreateDirOnRequired,
		parentDirPath,
		writingFilePath,
		[]byte(content),
	)
}

func (it fileStringWriter) DefaultLock(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	content string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Default(
		isRemoveBeforeWrite,
		writingFilePath,
		content,
	)
}

// Default
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileStringWriter) Default(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	content string,
) error {
	return fileWriter{}.Bytes.Default(
		isRemoveBeforeWrite,
		writingFilePath,
		[]byte(content),
	)
}

func (it fileStringWriter) Chmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	content string,
) error {
	return fileWriter{}.Bytes.WithDirChmod(
		isRemoveBeforeWrite,
		chmodDir,
		chmodFile,
		writingFilePath,
		[]byte(content),
	)
}

func (it fileStringWriter) ChmodLock(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	content string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return fileWriter{}.Bytes.WithDirChmod(
		isRemoveBeforeWrite,
		chmodDir,
		chmodFile,
		writingFilePath,
		[]byte(content),
	)
}
