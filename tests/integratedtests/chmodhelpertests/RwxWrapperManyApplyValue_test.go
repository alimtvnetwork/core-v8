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

package chmodhelpertests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/errcore"
)

// Test_RwxWrapperManyApplyValue_Unix
//
//	for directory `-` will be placed not `d`
func Test_RwxWrapperManyApplyValue_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	createPathInstructions := pathInstructionsV2
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		createPathInstructions,
	)

	firstCreationIns := createPathInstructions[0]
	paths := firstCreationIns.GetPaths()
	condition := chmodins.DefaultAllTrueCondition()
	existingAppliedRwxFull := firstCreationIns.ApplyRwx.String()

	for caseIndex, testCase := range rwxWrapperManyApplyTestCases {
		// Arrange
		rwxWrapper, err := testCase.SingleRwx.ToDisabledRwxWrapper()
		errcore.SimpleHandleErr(err, "SingleRwx ToDisabledRwxWrapper failed")

		expectation := rwxWrapper.ToFullRwxValueString()

		header := fmt.Sprintf(
			"Existing [%s] Applied by [%s] should result [%s]",
			existingAppliedRwxFull,
			expectation,
			expectation,
		)

		// Act
		applyErr := rwxWrapper.ApplyLinuxChmodOnMany(condition, paths...)
		errcore.SimpleHandleErr(
			applyErr,
			"rwxWrapper.ApplyLinuxChmodOnMany failed",
		)

		fileChmodMap := firstCreationIns.GetFilesChmodMap()
		var actLines []string

		for filePath, chmodValueString := range fileChmodMap.Items() {
			isEqual := chmodValueString == expectation
			actLines = append(actLines, fmt.Sprintf(
				"%s=%v",
				filePath,
				isEqual,
			))

			if !isEqual {
				errcore.PrintDiffOnMismatch(
					caseIndex,
					testCase.Case.Title,
					[]string{chmodValueString},
					[]string{expectation},
					fmt.Sprintf("  File: %s", filePath),
				)
			}
		}

		// Assert
		assertSingleChmod(
			t,
			header,
			firstCreationIns,
			expectation,
		)
	}
}
