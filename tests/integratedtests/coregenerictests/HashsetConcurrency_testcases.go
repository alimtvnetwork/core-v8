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

package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset Concurrency — test cases
// ==========================================================================

var hashsetAddLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "AddLock concurrent safety -- all items added",
	ExpectedInput: args.Map{"length": 500},
}

var hashsetAddSliceLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "AddSliceLock concurrent safety -- all batches added",
	ExpectedInput: args.Map{"length": 1000},
}

var hashsetContainsLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "ContainsLock concurrent reads during writes",
	ExpectedInput: args.Map{"finalLength": 200},
}

var hashsetRemoveLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "RemoveLock concurrent safety -- all items removed",
	ExpectedInput: args.Map{"length": 0},
}

var hashsetLengthLockConcurrencyTestCase = coretestcases.CaseV1{
	Title: "LengthLock concurrent reads during mutations",
	ExpectedInput: args.Map{
		"finalLength":   100,
		"noNegativeLen": true,
	},
}

var hashsetIsEmptyLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "IsEmptyLock concurrent check with writes",
	ExpectedInput: args.Map{"length": 100},
}
