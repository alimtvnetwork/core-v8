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
)

type errorCreator struct{}

func (it errorCreator) dirError(dirPath string, err error) error {
	notDirErr := it.notDirError(dirPath)
	if notDirErr != nil {
		return notDirErr
	}

	// has err
	return errors.New(
		"dir : " + dirPath +
			", applyChmod :" + dirDefaultChmod.String() +
			", " + err.Error(),
	)
}

func (it errorCreator) notDirError(dirPath string) error {
	if IsPathInvalid(dirPath) {
		return nil
	}

	// exist

	if !IsDirectory(dirPath) {
		return errors.New(
			"dir : " + dirPath +
				", applyChmod :" + dirDefaultChmod.String() +
				", path exist but it is not a dir.",
		)
	}

	return nil
}

func (it errorCreator) pathError(
	message string,
	applyChmod os.FileMode,
	location string,
	err error,
) error {
	if err == nil {
		return nil
	}

	compiledMessage := pathErrorMessage(
		message,
		applyChmod,
		location,
		err,
	)

	return errors.New(compiledMessage)
}

func (it errorCreator) pathErrorWithDirValidate(
	message string,
	applyChmod os.FileMode,
	location string,
	err error,
) error {
	notDirErr := it.notDirError(location)

	if notDirErr != nil {
		return notDirErr
	}

	if err == nil {
		return nil
	}

	compiledMessage := pathErrorMessage(
		message,
		applyChmod,
		location,
		err,
	)

	return errors.New(compiledMessage)
}

func (it errorCreator) chmodApplyFailed(
	applyChmod os.FileMode,
	location string,
	err error,
) error {
	if err == nil {
		return nil
	}

	compiledMessage := pathErrorMessage(
		"chmod apply failed",
		applyChmod,
		location,
		err,
	)

	return errors.New(compiledMessage)
}
