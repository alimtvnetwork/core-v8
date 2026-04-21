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

package corevalidator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coreutils/stringutil"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/internal/msgformats"
)

type TextValidator struct {
	Search   string `json:"Search,omitempty"`
	SearchAs stringcompareas.Variant
	Condition
	searchTextFinalized *string
}

func (it *TextValidator) ToString(isSingleLine bool) string {
	if it == nil {
		return constants.EmptyString
	}

	if isSingleLine {
		return fmt.Sprintf(
			msgformats.TextValidatorSingleLineFormat,
			it.Search,
			it.SearchAs.Name(),
			it.IsTrimCompare,
			it.IsSplitByWhitespace(),
			it.IsUniqueWordOnly,
			it.IsNonEmptyWhitespace,
			it.IsSortStringsBySpace,
		)
	}

	return fmt.Sprintf(
		msgformats.TextValidatorMultiLineFormat,
		it.Search,
		it.SearchAs.Name(),
		it.IsTrimCompare,
		it.IsSplitByWhitespace(),
		it.IsUniqueWordOnly,
		it.IsNonEmptyWhitespace,
		it.IsSortStringsBySpace,
	)
}

func (it *TextValidator) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.ToString(true)
}

func (it *TextValidator) SearchTextFinalized() string {
	if it == nil {
		return constants.EmptyString
	}

	return *it.SearchTextFinalizedPtr()
}

func (it *TextValidator) SearchTextFinalizedPtr() *string {
	if it.searchTextFinalized != nil {
		return it.searchTextFinalized
	}

	searchTerm := it.GetCompiledTermBasedOnConditions(
		it.Search,
		it.IsUniqueWordOnly,
	) // for unique word, use lowercase

	it.searchTextFinalized = &searchTerm

	return it.searchTextFinalized
}

func (it *TextValidator) GetCompiledTermBasedOnConditions(
	input string,
	isCaseSensitive bool,
) string {
	searchTerm := input

	if it.IsTrimCompare {
		searchTerm = strings.TrimSpace(searchTerm)
	}

	if it.IsSplitByWhitespace() {
		compiledStringSplits := stringutil.SplitContentsByWhitespaceConditions(
			searchTerm,
			it.IsTrimCompare,
			it.IsNonEmptyWhitespace,
			it.IsSortStringsBySpace,
			it.IsUniqueWordOnly,
			!isCaseSensitive,
		)

		return strings.Join(
			compiledStringSplits,
			constants.Space,
		)
	}

	return searchTerm
}

func (it *TextValidator) IsMatch(
	content string,
	isCaseSensitive bool,
) bool {
	if it == nil {
		return false
	}

	search := it.SearchTextFinalized()
	processedContent := it.GetCompiledTermBasedOnConditions(
		content,
		isCaseSensitive,
	)

	isIgnoreCase := !isCaseSensitive

	return it.SearchAs.IsCompareSuccess(
		isIgnoreCase,
		processedContent,
		search,
	)
}

func (it *TextValidator) IsMatchMany(
	isSkipOnContentsEmpty,
	isCaseSensitive bool,
	contents ...string,
) bool {
	if it == nil {
		return true
	}

	if len(contents) == 0 && isSkipOnContentsEmpty {
		return true
	}

	for _, content := range contents {
		if !it.IsMatch(content, isCaseSensitive) {
			return false
		}
	}

	return true
}

func (it *TextValidator) VerifyDetailError(
	params *Parameter,
	content string,
) error {
	if it == nil {
		return nil
	}

	return it.verifyDetailErrorUsingLineProcessing(
		constants.InvalidValue,
		params,
		content,
	)
}

func (it *TextValidator) verifyDetailErrorUsingLineProcessing(
	lineProcessingIndex int,
	params *Parameter,
	content string,
) error {
	if it == nil {
		return nil
	}

	processedSearch := it.SearchTextFinalized()
	processedContent := it.GetCompiledTermBasedOnConditions(
		content,
		params.IsCaseSensitive,
	)

	isMatch := it.SearchAs.IsCompareSuccess(
		params.IsIgnoreCase(),
		processedContent,
		processedSearch,
	)

	if isMatch {
		return nil
	}

	expectationMethod := it.SearchAs.Name()

	msg := errcore.GetSearchTermExpectationMessage(
		params.CaseIndex,
		params.Header,
		expectationMethod,
		lineProcessingIndex,
		processedContent,
		processedSearch,
		it.String(),
	)

	return errors.New(msg)
}

func (it *TextValidator) MethodName() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.SearchAs.Name()
}

func (it *TextValidator) VerifySimpleError(
	processingIndex int,
	params *Parameter,
	content string,
) error {
	if it == nil {
		return nil
	}

	processedSearch := it.SearchTextFinalized()
	processedContent := it.GetCompiledTermBasedOnConditions(
		content,
		params.IsCaseSensitive,
	)

	isMatch := it.SearchAs.IsCompareSuccess(
		params.IsIgnoreCase(),
		processedContent,
		processedSearch,
	)

	if isMatch {
		return nil
	}

	method := it.SearchAs.Name()

	msg := errcore.GetSearchTermExpectationSimpleMessage(
		params.CaseIndex,
		method,
		processingIndex,
		processedContent,
		processedSearch,
	)

	return errors.New(msg)
}

func (it *TextValidator) VerifyMany(
	isContinueOnError bool,
	params *Parameter,
	contents ...string,
) error {
	if isContinueOnError {
		return it.AllVerifyError(
			params,
			contents...,
		)
	}

	return it.VerifyFirstError(
		params,
		contents...,
	)
}

func (it *TextValidator) VerifyFirstError(
	params *Parameter,
	contents ...string,
) error {
	if it == nil {
		return nil
	}

	length := len(contents)
	if length == 0 && params.IsSkipCompareOnActualEmpty {
		return nil
	}

	for i, content := range contents {
		err := it.verifyDetailErrorUsingLineProcessing(
			i,
			params,
			content,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *TextValidator) AllVerifyError(
	params *Parameter,
	contents ...string,
) error {
	if it == nil {
		return nil
	}

	length := len(contents)
	if length == 0 && params.IsSkipCompareOnActualEmpty {
		return nil
	}

	var sliceErr []string

	for i, content := range contents {
		err := it.verifyDetailErrorUsingLineProcessing(
			i,
			params,
			content,
		)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error(),
			)
		}
	}

	return errcore.SliceToError(
		sliceErr,
	)
}
