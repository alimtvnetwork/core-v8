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
	"github.com/alimtvnetwork/core-v8/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core-v8/tests/testwrappers/chmodhelpertestwrappers"
)

// rwxInstructionsUnixApplyRecursivelyTestCase https://ss64.com/bash/chmod.html
var rwxInstructionsUnixApplyRecursivelyTestCase = chmodhelpertestwrappers.RwxInstructionTestWrapper{
	RwxInstructions: []chmodins.RwxInstruction{
		{
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   false,
				IsContinueOnError: false,
				IsRecursive:       true,
			},
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "*-x",
				Group: "**x",
				Other: "-w-",
			},
		},
	},
	DefaultRwx:      &chmodhelpertestwrappers.DefaultRwx,
	IsErrorExpected: false,
	CreatePaths:     pathInstructionsV2,
	TestFuncName:    chmodhelpertestwrappers.RwxApplyOnPath,
	WhatIsExpected:  chmodhelpertestwrappers.DefaultExpected,
}
