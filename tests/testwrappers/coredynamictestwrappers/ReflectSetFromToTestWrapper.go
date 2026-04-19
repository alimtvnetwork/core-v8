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

package coredynamictestwrappers

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/errcore"
)

type FromToTestWrapper struct {
	Header                          string
	From, To, ExpectedValue, actual any
	IsUsePointerInFrom              bool
	IsExpectingError                bool
	HasPanic                        bool
	Validator                       corevalidator.TextValidator
}

func (it FromToTestWrapper) CaseTitle() string {
	return it.Header
}

func (it FromToTestWrapper) Input() any {
	return it.From
}

func (it FromToTestWrapper) Expected() any {
	return it.ExpectedValue
}

func (it FromToTestWrapper) ToFieldToDraftType() *coretests.DraftType {
	return coretests.AnyToDraftType(it.To)
}

func (it FromToTestWrapper) ToFieldToBytes() []byte {
	return coretests.AnyToBytes(it.To)
}

func (it FromToTestWrapper) ExpectedFieldToDraftType() *coretests.DraftType {
	return coretests.AnyToDraftType(it.ExpectedValue)
}

func (it FromToTestWrapper) ExpectedFieldToBytes() []byte {
	return coretests.AnyToBytes(it.ExpectedValue)
}

func (it FromToTestWrapper) SetActual(actual any) {
	it.actual = actual
}

func (it FromToTestWrapper) Actual() any {
	return it.actual
}

// ShouldBeEqual asserts actLines match expectedLines using
// the wrapper's Header as the test title.
func (it FromToTestWrapper) ShouldBeEqual(
	t *testing.T,
	caseIndex int,
	actLines []string,
	expectedLines []string,
) {
	t.Helper()

	errcore.AssertDiffOnMismatch(
		t,
		caseIndex,
		it.Header,
		actLines,
		expectedLines,
	)
}

func (it FromToTestWrapper) AsSimpleTestCaseWrapper() coretests.SimpleTestCaseWrapper {
	return &it
}

func (it *FromToTestWrapper) AsSimpleTestCaseWrapperContractsBinder() coretests.SimpleTestCaseWrapperContractsBinder {
	return it
}
