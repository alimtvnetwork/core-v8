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

	"github.com/alimtvnetwork/core/internal/pathinternal"
)

type fileBytesWriter struct{}

// WithDirChmodLock
//
// Create dir safely if required.
func (it fileBytesWriter) WithDirChmodLock(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WithDirChmod(
		isRemoveBeforeWrite,
		chmodDir,
		chmodFile,
		writingFilePath,
		contentsBytes,
	)
}

// WithDirChmod
//
// Create dir safely if required.
func (it fileBytesWriter) WithDirChmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	parentDir := pathinternal.ParentDir(writingFilePath)

	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		true,
		true,
		true,
		parentDir,
		writingFilePath,
		contentsBytes,
	)
}

// Chmod
//
// Create dir safely if required.
func (it fileBytesWriter) Chmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	parentDir := pathinternal.ParentDir(writingFilePath)

	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		true,
		true,
		true,
		parentDir,
		writingFilePath,
		contentsBytes,
	)
}

func (it fileBytesWriter) WithDirLock(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WithDir(
		isRemoveBeforeWrite,
		writingFilePath,
		contentsBytes,
	)
}

// WithDir
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileBytesWriter) WithDir(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	contentsBytes []byte,
) error {
	return it.WithDirChmod(
		isRemoveBeforeWrite,
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		contentsBytes,
	)
}

// Default
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileBytesWriter) Default(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	contentsBytes []byte,
) error {
	return it.WithDirChmod(
		isRemoveBeforeWrite,
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		contentsBytes,
	)
}
