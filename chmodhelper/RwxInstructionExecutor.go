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

	"github.com/alimtvnetwork/core-v8/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/internal/messages"
)

type RwxInstructionExecutor struct {
	rwxInstruction *chmodins.RwxInstruction
	varWrapper     *RwxVariableWrapper
}

// IsVarWrapper if it has any wildcard symbol in it
func (it *RwxInstructionExecutor) IsVarWrapper() bool {
	if it.varWrapper == nil {
		return false
	}

	return !it.varWrapper.IsFixedType()
}

// IsFixedWrapper true indicates no wildcard symbol
func (it *RwxInstructionExecutor) IsFixedWrapper() bool {
	if it.varWrapper == nil {
		return false
	}

	return it.varWrapper.IsFixedType()
}

// IsEqualFileInfo fileInfo nil returns false.
func (it *RwxInstructionExecutor) IsEqualFileInfo(fileInfo os.FileInfo) bool {
	return it.varWrapper.IsEqualUsingFileInfo(fileInfo)
}

// IsEqualRwxWrapper nil returns false.
func (it *RwxInstructionExecutor) IsEqualRwxWrapper(rwxWrapper *RwxWrapper) bool {
	return it.varWrapper.IsEqualRwxWrapperPtr(rwxWrapper)
}

// IsEqualRwxPartial
//
// rwxPartial:
//   - "-rwx" will be "-rwx******"
//   - "-rwxr-x" will be "-rwxr-x***"
//   - "-rwxr-x" will be "-rwxr-x***"
func (it *RwxInstructionExecutor) IsEqualRwxPartial(rwxPartial string) bool {
	return it.varWrapper.IsEqualPartialRwxPartial(rwxPartial)
}

func (it *RwxInstructionExecutor) IsEqualFileMode(mode os.FileMode) bool {
	return it.varWrapper.IsEqualUsingFileMode(mode)
}

func (it *RwxInstructionExecutor) CompiledWrapper(mode os.FileMode) (*RwxWrapper, error) {
	if it.IsFixedWrapper() {
		return it.
			varWrapper.
			ToCompileWrapperPtr(nil), nil
	}

	if it.IsVarWrapper() {
		fixedWrapper := New.RwxWrapper.UsingFileMode(mode)

		return it.
			varWrapper.
			ToCompileWrapperPtr(&fixedWrapper), nil
	}

	return nil, failedToCompileVarWrapperToWrapper
}

func (it *RwxInstructionExecutor) CompiledRwxWrapperUsingFixedRwxWrapper(
	wrapper *RwxWrapper,
) (*RwxWrapper, error) {
	if it.IsFixedWrapper() {
		return it.
			varWrapper.
			ToCompileWrapperPtr(nil), nil
	}

	if it.IsVarWrapper() {
		return it.
			varWrapper.
			ToCompileWrapperPtr(wrapper), nil
	}

	return nil, errcore.
		FailedToExecuteType.
		Error(
			messages.FailedToCompileChmodhelperVarWrapperToWrapper,
			wrapper.String())
}

// ApplyOnPath
//
// Warning:
//
//	swallows error if chmodins.RwxInstruction. IsSkipOnInvalid or
//	chmodins.RwxInstruction.IsExitOnInvalid() comes as negative
func (it *RwxInstructionExecutor) ApplyOnPath(location string) error {
	existingRwxFileModWrapper, err := GetExistingChmodRwxWrapperPtr(
		location)

	if it.rwxInstruction.IsExitOnInvalid() && err != nil {
		return errcore.PathErrorType.Error(messages.FailedToGetFileModeRwx, location)
	} else if it.rwxInstruction.IsSkipOnInvalid && err != nil {
		// nothing apply got an error
		return nil
	}

	compiledWrapper, compiledErr := it.CompiledRwxWrapperUsingFixedRwxWrapper(existingRwxFileModWrapper)

	if compiledErr != nil {
		funcWithLoc := "ApplyOnPath" + constants.HyphenAngelRight + location

		return errcore.
			MeaningfulError(
				errcore.PathErrorType, funcWithLoc, compiledErr)
	}

	if it.rwxInstruction.IsRecursive {
		return compiledWrapper.ApplyRecursive(
			it.rwxInstruction.IsSkipOnInvalid,
			location)
	}

	return compiledWrapper.ApplyChmod(
		it.rwxInstruction.IsSkipOnInvalid,
		location,
	)
}

func (it *RwxInstructionExecutor) VerifyRwxModifiersDirect(
	isRecursiveIgnore bool,
	locations ...string,
) error {
	return it.VerifyRwxModifiers(isRecursiveIgnore, locations)
}

func (it *RwxInstructionExecutor) VerifyRwxModifiers(
	isRecursiveIgnore bool,
	locations []string,
) error {
	if len(locations) == 0 {
		return nil
	}

	resultsMap, err := it.
		getVerifyRwxInternalError(
			isRecursiveIgnore,
			locations)

	if err != nil {
		return err
	}

	if it.rwxInstruction.IsContinueOnError {
		return it.verifyChmodLocationsContinueOnError(resultsMap)
	}

	return it.verifyChmodLocationsNoContinue(resultsMap)
}

func (it *RwxInstructionExecutor) getVerifyRwxInternalError(
	isRecursiveIgnore bool,
	locations []string,
) (
	*FilteredPathFileInfoMap, error,
) {
	if !isRecursiveIgnore && it.rwxInstruction.Condition.IsRecursive {
		return nil, errcore.NotSupportedType.Error(
			"IsRecursive is not supported for Verify chmod.",
			locations)
	}

	resultsMap := GetExistsFilteredPathFileInfoMap(
		it.rwxInstruction.IsSkipOnInvalid,
		locations...)

	return resultsMap, nil
}

func (it *RwxInstructionExecutor) verifyChmodLocationsContinueOnError(
	resultsMap *FilteredPathFileInfoMap,
) error {
	var sliceErr []string

	if resultsMap.Error != nil && it.rwxInstruction.IsCollectErrorOnInvalid() {
		sliceErr = append(
			sliceErr,
			resultsMap.Error.Error())
	}

	for filePath, info := range resultsMap.FilesToInfoMap {
		fileMode := info.Mode()
		fixedRwxWrapper, err := it.CompiledWrapper(fileMode)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error()+"- failed to verify rwxInstruction for - "+filePath)
		}

		if fixedRwxWrapper != nil && !fixedRwxWrapper.IsEqualFileMode(fileMode) {
			sliceErr = append(
				sliceErr,
				errcore.ExpectingSimpleNoType(
					"Path:"+filePath,
					fixedRwxWrapper.ToFullRwxValueStringExceptHyphen(),
					fileMode.String()[1:]))
		}
	}

	return errcore.SliceToError(sliceErr)
}

func (it *RwxInstructionExecutor) verifyChmodLocationsNoContinue(
	resultsMap *FilteredPathFileInfoMap,
) error {
	if resultsMap.Error != nil && !it.rwxInstruction.IsSkipOnInvalid {
		return resultsMap.Error
	}

	for filePath, info := range resultsMap.FilesToInfoMap {
		fileMode := info.Mode()
		fixedRwxWrapper, err := it.CompiledWrapper(
			fileMode)

		if err != nil {
			return errcore.MeaningfulErrorWithData(
				errcore.ValidationFailedType,
				"verifyChmodLocationsNoContinue",
				err,
				"failed to verify rwxInstruction for - "+filePath)
		}

		if fixedRwxWrapper.IsDefined() && fixedRwxWrapper.IsNotEqualFileMode(fileMode) {
			expectingMsg := errcore.ExpectingSimpleNoType(
				"Path:"+filePath,
				fixedRwxWrapper.ToFullRwxValueStringExceptHyphen(),
				fileMode.String()[1:])

			return errors.New(expectingMsg)
		}
	}

	return nil
}

func (it *RwxInstructionExecutor) ApplyOnPathsDirect(locations ...string) error {
	if len(locations) == 0 {
		return nil
	}

	return it.ApplyOnPathsPtr(&locations)
}

func (it *RwxInstructionExecutor) ApplyOnPaths(locations []string) error {
	if len(locations) == 0 {
		return nil
	}

	return it.ApplyOnPathsPtr(&locations)
}

func (it *RwxInstructionExecutor) ApplyOnPathsPtr(locations *[]string) error {
	if locations == nil {
		return nil
	}

	isContinueOnError := it.
		rwxInstruction.
		IsContinueOnError

	isStopOnError := !isContinueOnError

	if isStopOnError {
		return it.applyOnPaths(locations)
	}

	return it.applyOnPathsContinueOnError(locations)
}

func (it *RwxInstructionExecutor) applyOnPaths(locations *[]string) error {
	for _, location := range *locations {
		err := it.ApplyOnPath(location)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *RwxInstructionExecutor) applyOnPathsContinueOnError(locations *[]string) error {
	errorSlice := make([]string, constants.Zero)

	for _, location := range *locations {
		err := it.ApplyOnPath(location)

		if err != nil {
			errorSlice = append(errorSlice, err.Error())
		}
	}

	return errcore.SliceToErrorPtr(errorSlice)
}
