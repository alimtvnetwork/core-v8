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

package chmodhelpertestwrappers

import (
	"github.com/alimtvnetwork/core-v8/chmodhelper"
	"github.com/alimtvnetwork/core-v8/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core-v8/coretests"
)

type RwxInstructionTestWrapper struct {
	RwxInstructions []chmodins.RwxInstruction
	DefaultRwx      *chmodins.RwxOwnerGroupOther
	IsErrorExpected bool
	CreatePaths     []chmodhelper.DirFilesWithRwxPermission
	TestFuncName    coretests.TestFuncName
	WhatIsExpected  chmodins.RwxOwnerGroupOther
	actual          any
}

func (it *RwxInstructionTestWrapper) Actual() any {
	return it.actual
}

func (it *RwxInstructionTestWrapper) SetActual(actual any) {
	it.actual = actual
}

func (it *RwxInstructionTestWrapper) GetFuncName() string {
	return it.TestFuncName.Value()
}

func (it *RwxInstructionTestWrapper) Value() any {
	return it
}

func (it *RwxInstructionTestWrapper) Expected() any {
	return it.WhatIsExpected
}

func (it *RwxInstructionTestWrapper) ExpectedAsRwxOwnerGroupOtherInstruction() chmodins.RwxOwnerGroupOther {
	return it.WhatIsExpected
}

func (it *RwxInstructionTestWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = it

	return testCaseMessenger
}
