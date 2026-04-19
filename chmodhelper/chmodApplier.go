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

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/errcore"
)

type chmodApplier struct{}

func (it chmodApplier) Default(
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmod(
		false,
		location)
}

func (it chmodApplier) OnMismatchOption(
	isApply,
	isSkipOnInvalid bool,
	fileMode os.FileMode,
	location string,
) error {
	isSkipApply := !isApply

	if isSkipApply {
		return nil
	}

	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmodOptions(
		isApply,
		true,
		isSkipOnInvalid,
		location)
}

func (it chmodApplier) OnMismatch(
	isSkipOnInvalid bool,
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmodOptions(
		true,
		true,
		isSkipOnInvalid,
		location)
}

func (it chmodApplier) OnMismatchSkipInvalid(
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmodOptions(
		true,
		true,
		true,
		location)
}

func (it chmodApplier) SkipInvalidFile(
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmod(
		true,
		location)
}

func (it chmodApplier) ApplyIf(
	isApply bool,
	fileMode os.FileMode,
	location string,
) error {
	isSkipApply := !isApply

	if isSkipApply {
		return nil
	}

	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmod(
		false,
		location)
}

func (it chmodApplier) Options(
	isSkipInvalidPaths,
	isRecursive bool,
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	if isRecursive {
		return rwx.ApplyRecursive(
			isSkipInvalidPaths,
			location)
	}

	return rwx.ApplyChmod(
		isSkipInvalidPaths,
		location)
}

func (it chmodApplier) RecursivePath(
	isSkipInvalidPaths bool,
	fileMode os.FileMode,
	location string,
) error {
	return it.Options(
		isSkipInvalidPaths,
		true,
		fileMode,
		location)
}

func (it chmodApplier) RecursivePaths(
	isContinueOnError bool,
	isSkipInvalidPaths bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsSkipOnInvalid:   isSkipInvalidPaths,
			IsContinueOnError: isContinueOnError,
			IsRecursive:       true,
		},
		locations...)
}

func (it chmodApplier) RecursivePathsContinueOnError(
	isSkipInvalidPaths bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsSkipOnInvalid:   isSkipInvalidPaths,
			IsContinueOnError: true,
			IsRecursive:       true,
		},
		locations...)
}

func (it chmodApplier) RecursivePathsCaptureInvalids(
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsContinueOnError: false,
			IsRecursive:       true,
		},
		locations...)
}

func (it chmodApplier) PathsUsingFileModeRecursive(
	isContinueOnError bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsContinueOnError: isContinueOnError,
			IsRecursive:       true,
		},
		locations...)
}

func (it chmodApplier) PathsUsingFileModeContinueOnErr(
	isRecursive bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsContinueOnError: true,
			IsRecursive:       isRecursive,
		},
		locations...)
}

func (it chmodApplier) PathsUsingFileModeOptions(
	isSkipOnInvalid,
	isContinueOnError,
	isRecursive bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsSkipOnInvalid:   isSkipOnInvalid,
			IsContinueOnError: isContinueOnError,
			IsRecursive:       isRecursive,
		},
		locations...)
}

func (it chmodApplier) PathsUsingFileModeConditions(
	fileMode os.FileMode,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if condition == nil {
		return errcore.CannotBeNilOrEmptyType.
			ErrorNoRefs("condition")
	}

	rwxWrapper := New.RwxWrapper.UsingFileMode(fileMode)

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}

// RwxPartial
//
// can be any length in
// between 0-10 (rest will be fixed by wildcard)
//
// rwxPartial:
//   - "-rwx" will be "-rwx******"
//   - "-rwxr-x" will be "-rwxr-x***"
//   - "-rwxr-x" will be "-rwxr-x***"
func (it chmodApplier) RwxPartial(
	rwxPartial string,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if condition == nil {
		return errcore.CannotBeNilOrEmptyType.
			ErrorNoRefs("condition")
	}

	if len(locations) == 0 {
		return nil
	}

	rwxInstructionExecutor, err := RwxPartialToInstructionExecutor(
		rwxPartial,
		condition)

	if err != nil {
		return err
	}

	return rwxInstructionExecutor.
		ApplyOnPathsPtr(&locations)
}

// RwxStringApplyChmod
//
//	rwxFullString 10 chars "-rwxrwxrwx"
func RwxStringApplyChmod(
	rwxFullString string, // rwxFullString 10 chars "-rwxrwxrwx"
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	rwxFullStringErr := chmodins.GetRwxFullLengthError(
		rwxFullString)
	if rwxFullStringErr != nil {
		return rwxFullStringErr
	}

	if condition == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("condition")
	}

	rwxWrapper, err := New.RwxWrapper.RwxFullString(
		rwxFullString)

	if err != nil {
		return err
	}

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}

// RwxOwnerGroupOtherApplyChmod rwxFullString 10 chars "-rwxrwxrwx"
func RwxOwnerGroupOtherApplyChmod(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if rwxOwnerGroupOther == nil {
		return errcore.CannotBeNilOrEmptyType.
			ErrorNoRefs("rwxOwnerGroupOther")
	}

	if condition == nil {
		return errcore.CannotBeNilOrEmptyType.
			ErrorNoRefs("condition")
	}

	rwxWrapper, err := New.RwxWrapper.RwxFullString(
		rwxOwnerGroupOther.String())

	if err != nil {
		return err
	}

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}
