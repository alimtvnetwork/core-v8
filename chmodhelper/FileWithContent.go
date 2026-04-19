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
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/pathinternal"
)

type FileWithContent struct {
	RelativePath string
	FileMode     os.FileMode // default for file fileDefaultChmod
	Content      []string
}

func (it FileWithContent) AbsPath(parentDir string) string {
	return pathinternal.Join(parentDir, it.RelativePath)
}

func (it FileWithContent) ContentToString() string {
	return strings.Join(it.Content, constants.NewLineUnix)
}

func (it FileWithContent) ContentToBytes() []byte {
	return []byte(it.ContentToString())
}

func (it FileWithContent) Read(parentDir string) ([]byte, error) {
	return SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              it.FileMode,
		ParentDir:              parentDir,
		FilePath:               it.AbsPath(parentDir),
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}.Read()
}

func (it FileWithContent) ReadString(parentDir string) (string, error) {
	return SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              it.FileMode,
		ParentDir:              parentDir,
		FilePath:               it.AbsPath(parentDir),
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}.ReadString()
}

func (it FileWithContent) ReadLines(parentDir string) ([]string, error) {
	toStr, err := it.ReadString(parentDir)

	if err != nil {
		return []string{}, err
	}

	// safe

	return strings.Split(toStr, constants.NewLineUnix), nil
}

func (it FileWithContent) Write(parentDir string) error {
	return SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              it.FileMode,
		ParentDir:              parentDir,
		FilePath:               it.AbsPath(parentDir),
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}.Write(it.ContentToBytes())
}
