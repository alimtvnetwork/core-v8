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
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/alimtvnetwork/core-v8/constants"
)

type anyItemWriter struct{}

func (it anyItemWriter) ChmodLock(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	parentDir,
	writingFilePath string,
	anyItem any,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Chmod(
		isRemoveBeforeWrite,
		chmodDir,
		chmodFile,
		parentDir,
		writingFilePath,
		anyItem,
	)
}

// Chmod
//
//	Writes contents to file system by serializing using JSON.
//
// parentDirPath:
//   - is a full path to the parent dir for checking
//     if parent dir exist if not then created
//
// writingFilePath:
//   - is a full path to the actual file where to write contents
func (it anyItemWriter) Chmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	parentDir,
	writingFilePath string,
	anyItem any,
) error {
	jsonBytes, err := json.Marshal(anyItem)

	if err == nil {
		return fileWriter{}.All(
			chmodDir,
			chmodFile,
			isRemoveBeforeWrite,
			true,
			true,
			true,
			parentDir,
			writingFilePath,
			jsonBytes,
		)
	}

	var typeName, anyString string
	if anyItem != nil {
		// fine if var type not detected as nil
		// we want to avoid interface nil only
		typeName = reflect.TypeOf(anyItem).String()
		anyString = fmt.Sprintf(
			constants.SprintValueFormat,
			anyItem,
		)
	}

	// has err
	return errors.New(
		"json convert failed," +
			", filePath : " + writingFilePath +
			", AnyType : " + typeName +
			", AnyItem(String) : " + anyString +
			", chmodFile :" + chmodFile.String() + ", " +
			", chmodDir :" + chmodDir.String() + ", " +
			err.Error(),
	)
}

// DefaultLock
//
//	Writes contents to file system by serializing using JSON.
//	Applies default chmod (for dir - 0755, for file - 0644)
//
// writingFilePath:
//   - is a full path to the actual file where to write contents
func (it anyItemWriter) DefaultLock(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	anyItem any,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Default(
		isRemoveBeforeWrite,
		writingFilePath,
		anyItem,
	)
}

// Default
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it anyItemWriter) Default(
	isRemoveBeforeWrite bool,
	writingFilePath string,
	anyItem any,
) error {
	parentDir := filepath.Dir(writingFilePath)

	return it.Chmod(
		isRemoveBeforeWrite,
		dirDefaultChmod,
		fileDefaultChmod,
		parentDir,
		writingFilePath,
		anyItem,
	)
}
