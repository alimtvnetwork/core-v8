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

	"github.com/alimtvnetwork/core-v8/internal/pathinternal"
)

type newSimpleFileReaderWriterCreator struct{}

// Create
//
// Arguments:
//   - chmodDir     : applying on parentDir
//   - chmodFile    : applying on file
//   - absParentDir : absolute parentDir ( it can be two level before ),
//     it doesn't have to be relative to absFilePath but can be relative.
//   - absFilePath  : absolute file path
func (it newSimpleFileReaderWriterCreator) Create(
	isRemoveBeforeWrite bool,
	chmodDir,
	chmodFile os.FileMode,
	absParentDir,
	absFilePath string,
) *SimpleFileReaderWriter {
	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              absParentDir,
		FilePath:               absFilePath,
		IsRemoveBeforeWrite:    isRemoveBeforeWrite,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// All
//
// Arguments:
//   - chmodDir     : applying on parentDir
//   - chmodFile    : applying on file
//   - absParentDir : absolute parentDir ( it can be two level before ),
//     it doesn't have to be relative to absFilePath but can be relative.
//   - absFilePath  : absolute file path
func (it newSimpleFileReaderWriterCreator) All(
	chmodDir,
	chmodFile os.FileMode,
	isRemoveBeforeWrite bool,
	isApplyChmodMust bool,
	isApplyOnMismatch bool,
	absParentDir,
	absFilePath string,
) *SimpleFileReaderWriter {
	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              absParentDir,
		FilePath:               absFilePath,
		IsRemoveBeforeWrite:    isRemoveBeforeWrite,
		IsMustChmodApplyOnFile: isApplyChmodMust,
		IsApplyChmodOnMismatch: isApplyOnMismatch,
	}
}

func (it newSimpleFileReaderWriterCreator) Options(
	isRemoveBeforeWrite bool,
	isApplyChmodMust bool,
	isApplyOnMismatch bool,
	absFilePath string,
) *SimpleFileReaderWriter {
	parentDir := pathinternal.ParentDir(absFilePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              fileDefaultChmod,
		ParentDir:              parentDir,
		FilePath:               absFilePath,
		IsRemoveBeforeWrite:    isRemoveBeforeWrite,
		IsMustChmodApplyOnFile: isApplyChmodMust,
		IsApplyChmodOnMismatch: isApplyOnMismatch,
	}
}

// CreateClean
//
//	Applies pathinternal.Clean() to relative actual path from relative(cmd/..) or (.) path
//	then create the reader, writer.
//
// Arguments:
//   - chmodDir  : applying on parentDir
//   - chmodFile : applying on file
//   - parentDir : absolute parentDir ( it can be two level before ) after clean,
//     it doesn't have to be relative to absFilePath but can be relative.
//   - filePath  : absolute file path after clean or else will not work
func (it newSimpleFileReaderWriterCreator) CreateClean(
	isRemoveBeforeWrite bool,
	chmodDir,
	chmodFile os.FileMode,
	parentDir,
	filePath string,
) *SimpleFileReaderWriter {
	parentDir = pathinternal.Clean(parentDir)
	filePath = pathinternal.Clean(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              parentDir,
		FilePath:               filePath,
		IsRemoveBeforeWrite:    isRemoveBeforeWrite,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// Default
//
//	applies default chmod dir - 0755
//	(filemode.DirDefault), file - 0644  (filemode.FileDefault)
//
// Arguments:
//   - chmodDir     : applying on parentDir
//   - chmodFile    : applying on file
//   - absParentDir : absolute parentDir will be from parent of absFilePath
//   - absFilePath  : absolute file path
func (it newSimpleFileReaderWriterCreator) Default(
	isRemoveBeforeWrite bool,
	absFilePath string,
) *SimpleFileReaderWriter {
	parentDir := pathinternal.ParentDir(absFilePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              fileDefaultChmod,
		ParentDir:              parentDir,
		FilePath:               absFilePath,
		IsRemoveBeforeWrite:    isRemoveBeforeWrite,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// DefaultCleanPath
//
//	Applies pathinternal.Clean() to relative actual path from relative(cmd/..) or (.) path
//	then create the reader, writer.
//
//	applies default chmod dir - 0755
//	(filemode.DirDefault), file - 0644  (filemode.FileDefault)
//
// Arguments:
//   - chmodDir     : applying on parentDir
//   - chmodFile    : applying on file
//   - absFilePath  : absolute file path after clean.
//   - absParentDir : absolute parentDir will be from parent of absFilePath after clean
func (it newSimpleFileReaderWriterCreator) DefaultCleanPath(
	isRemoveBeforeWrite bool,
	filePath string,
) *SimpleFileReaderWriter {
	filePath = pathinternal.Clean(filePath)
	parentDir := pathinternal.ParentDir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              fileDefaultChmod,
		ParentDir:              parentDir,
		FilePath:               filePath,
		IsRemoveBeforeWrite:    isRemoveBeforeWrite,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// Path
//
// Arguments:
//   - chmodDir     : applying on parentDir
//   - chmodFile    : applying on file
//   - absFilePath  : absolute file path.
//   - parentDir    : will be extracted from absFilePath.
func (it newSimpleFileReaderWriterCreator) Path(
	isRemoveBeforeWrite bool,
	chmodDir,
	chmodFile os.FileMode,
	absFilePath string,
) *SimpleFileReaderWriter {
	parentDir := pathinternal.ParentDir(absFilePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              parentDir,
		FilePath:               absFilePath,
		IsRemoveBeforeWrite:    isRemoveBeforeWrite,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

func (it newSimpleFileReaderWriterCreator) PathCondition(
	isRemoveBeforeWrite bool,
	isApplyPathClean bool,
	chmodDir,
	chmodFile os.FileMode,
	filePath string,
) *SimpleFileReaderWriter {
	if isApplyPathClean {
		filePath = pathinternal.Clean(filePath)
	}

	parentDir := pathinternal.ParentDir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              parentDir,
		FilePath:               filePath,
		IsRemoveBeforeWrite:    isRemoveBeforeWrite,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// PathDirDefaultChmod
//
//	dir will be applied with default chmod - 0755
func (it newSimpleFileReaderWriterCreator) PathDirDefaultChmod(
	isRemoveBeforeWrite bool,
	chmodFile os.FileMode,
	filePath string,
) *SimpleFileReaderWriter {
	parentDir := pathinternal.ParentDir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              chmodFile,
		ParentDir:              parentDir,
		FilePath:               filePath,
		IsRemoveBeforeWrite:    isRemoveBeforeWrite,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}
