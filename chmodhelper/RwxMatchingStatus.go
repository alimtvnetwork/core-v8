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
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
)

type RwxMatchingStatus struct {
	RwxMismatchInfos         []*RwxMismatchInfo
	MissingOrPathsWithIssues []string
	IsAllMatching            bool
	Error                    error
}

func InvalidRwxMatchingStatus(err error) *RwxMatchingStatus {
	return &RwxMatchingStatus{
		RwxMismatchInfos:         []*RwxMismatchInfo{},
		MissingOrPathsWithIssues: []string{},
		IsAllMatching:            false,
		Error:                    err,
	}
}

func EmptyRwxMatchingStatus() *RwxMatchingStatus {
	return &RwxMatchingStatus{
		RwxMismatchInfos:         []*RwxMismatchInfo{},
		MissingOrPathsWithIssues: []string{},
		IsAllMatching:            false,
		Error:                    nil,
	}
}

func (it *RwxMatchingStatus) MissingFilesToString() string {
	return strings.Join(it.MissingOrPathsWithIssues, constants.CommaSpace)
}

func (it *RwxMatchingStatus) CreateErrFinalError() error {
	if it.IsAllMatching && it.Error == nil {
		return nil
	}

	length := len(it.RwxMismatchInfos) +
		constants.Capacity2

	sliceErr := make([]string, 0, length)

	for _, info := range it.RwxMismatchInfos {
		expectingMessage := errcore.ExpectingSimpleNoType(
			info.FilePath,
			info.Expecting,
			info.Actual)

		sliceErr = append(
			sliceErr,
			expectingMessage)
	}

	if it.Error != nil {
		sliceErr = append(
			sliceErr,
			it.Error.Error())
	}

	return errcore.SliceToError(sliceErr)
}
