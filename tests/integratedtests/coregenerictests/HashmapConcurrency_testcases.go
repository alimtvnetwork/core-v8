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
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ==========================================================================
// Hashmap Concurrency — test cases
// ==========================================================================

var hashmapSetLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "SetLock concurrent safety -- all entries added",
	ExpectedInput: args.Map{"length": 500},
}

var hashmapGetLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "GetLock concurrent reads with writes",
	ExpectedInput: args.Map{"finalLength": 200},
}

var hashmapContainsLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "ContainsLock concurrent reads",
	ExpectedInput: args.Map{"finalLength": 200},
}

var hashmapRemoveLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "RemoveLock concurrent safety -- all entries removed",
	ExpectedInput: args.Map{"length": 0},
}

var hashmapLengthLockConcurrencyTestCase = coretestcases.CaseV1{
	Title: "LengthLock concurrent reads during mutations",
	ExpectedInput: args.Map{
		"finalLength":   100,
		"noNegativeLen": true,
	},
}

var hashmapIsEmptyLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "IsEmptyLock concurrent check",
	ExpectedInput: args.Map{"length": 100},
}

var hashmapMixedOpsConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "Mixed SetLock+GetLock+RemoveLock concurrent safety",
	ExpectedInput: args.Map{"finalLength": 300},
}
