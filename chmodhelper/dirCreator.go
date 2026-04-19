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
)

type dirCreator struct{}

// If
//
// if isCreate + is missing director then only create dir.
func (it dirCreator) If(
	isCreate bool,
	chmod os.FileMode,
	dirPath string,
) error {
	if !isCreate {
		return nil
	}

	return it.IfMissing(
		chmod,
		dirPath,
	)
}

func (it dirCreator) IfMissingLock(
	applyChmod os.FileMode,
	dirPath string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.IfMissing(
		applyChmod,
		dirPath,
	)
}

// IfMissing
//
// Only create dir if missing.
func (it dirCreator) IfMissing(
	applyChmod os.FileMode,
	dirPath string,
) error {
	if IsPathExists(dirPath) {
		return nil
	}

	// Use permissive mode for MkdirAll to allow nested dir creation,
	// then apply the requested chmod afterward.
	err := os.MkdirAll(
		dirPath,
		dirDefaultChmod,
	)

	if err != nil {
		// has err
		return newError.pathErrorWithDirValidate(
			"dir creation failed",
			applyChmod,
			dirPath,
			err,
		)
	}

	// Apply the requested chmod if it differs from the default
	if applyChmod != dirDefaultChmod {
		chmodErr := os.Chmod(dirPath, applyChmod)
		if chmodErr != nil {
			return newError.chmodApplyFailed(
				applyChmod,
				dirPath,
				chmodErr,
			)
		}
	}

	return nil
}

// ByChecking
//
// Check only if the dir is missing and apply chmod.
func (it dirCreator) ByChecking(
	applyChmod os.FileMode,
	dirPath string,
) error {
	isExists := IsPathExists(dirPath)
	isDir := IsDirectory(dirPath)

	if isExists && isDir {
		return os.Chmod(dirPath, applyChmod)
	}

	if isExists && !isDir {
		return newError.notDirError(dirPath)
	}

	// Use permissive mode for MkdirAll to allow nested dir creation,
	// then apply the requested chmod afterward.
	err := os.MkdirAll(
		dirPath,
		dirDefaultChmod,
	)

	if err != nil {
		return newError.pathErrorWithDirValidate(
			"dir creation failed",
			applyChmod,
			dirPath,
			err,
		)
	}

	chmodErr := os.Chmod(
		dirPath,
		applyChmod,
	)

	if chmodErr != nil {
		return newError.chmodApplyFailed(
			applyChmod,
			dirPath,
			chmodErr,
		)
	}

	return nil
}

func (it dirCreator) DefaultLock(
	applyChmod os.FileMode,
	dirPath string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Default(
		applyChmod,
		dirPath,
	)
}

// Default
//
// Direct try to create without checking if directory exists.
func (it dirCreator) Default(
	applyChmod os.FileMode,
	dirPath string,
) error {
	err := os.MkdirAll(
		dirPath,
		applyChmod,
	)

	if err == nil {
		return nil
	}

	// has err
	return newError.dirError(dirPath, err)
}

func (it dirCreator) DirectLock(
	dirPath string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Direct(
		dirPath,
	)
}

// Direct
//
// Dir default chmod 0755
func (it dirCreator) Direct(
	dirPath string,
) error {
	err := os.MkdirAll(
		dirPath,
		dirDefaultChmod,
	)

	if err == nil {
		return nil
	}

	return newError.dirError(dirPath, err)
}
