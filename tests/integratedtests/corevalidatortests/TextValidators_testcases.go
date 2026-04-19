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

package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// TextValidators — collection operations
// ==========================================================================

var textValidatorsNewEmptyTestCase = coretestcases.CaseV1{
	Title: "TextValidators returns isEmpty true length 0 -- new instance",
	ExpectedInput: args.Map{
		"isEmpty": true,
		"length":  0,
	},
}

var textValidatorsAddTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.Add returns length 2 -- two items added",
	ExpectedInput: args.Map{"length": 2},
}

var textValidatorsAddsTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.Adds returns length 2 -- variadic two items",
	ExpectedInput: args.Map{"length": 2},
}

var textValidatorsAddsEmptyTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.Adds returns length 0 -- no items given",
	ExpectedInput: args.Map{"length": 0},
}

var textValidatorsAddSimpleTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.AddSimple returns length 1 -- one item added",
	ExpectedInput: args.Map{"length": 1},
}

var textValidatorsHasIndexTestCase = coretestcases.CaseV1{
	Title: "TextValidators.HasIndex returns true for 0, false for 1 -- single item",
	ExpectedInput: args.Map{
		"hasIndex0": true,
		"hasIndex1": false,
	},
}

var textValidatorsLastIndexTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.LastIndex returns 1 -- two items",
	ExpectedInput: args.Map{"lastIndex": 1},
}

// ==========================================================================
// TextValidators.IsMatch
// ==========================================================================

var textValidatorsIsMatchEmptyTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.IsMatch returns true -- empty validators",
	ExpectedInput: args.Map{"isMatch": true},
}

var textValidatorsIsMatchAllPassTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.IsMatch returns true -- all validators pass",
	ExpectedInput: args.Map{"isMatch": true},
}

var textValidatorsIsMatchOneFailsTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.IsMatch returns false -- one validator fails",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidators.IsMatchMany
// ==========================================================================

var textValidatorsIsMatchManyEmptyTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.IsMatchMany returns true -- empty validators",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// TextValidators.VerifyFirstError
// ==========================================================================

var textValidatorsVerifyFirstAllPassTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.VerifyFirstError returns nil -- all match",
	ExpectedInput: args.Map{"hasError": false},
}

var textValidatorsVerifyFirstFailTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.VerifyFirstError returns error -- one mismatch",
	ExpectedInput: args.Map{"hasError": true},
}

var textValidatorsVerifyFirstEmptyTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.VerifyFirstError returns nil -- empty validators",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// TextValidators.AllVerifyError
// ==========================================================================

var textValidatorsAllVerifyPassTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.AllVerifyError returns nil -- all match",
	ExpectedInput: args.Map{"hasError": false},
}

var textValidatorsAllVerifyFailTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.AllVerifyError returns error -- multiple mismatches",
	ExpectedInput: args.Map{"hasError": true},
}

// ==========================================================================
// TextValidators.Dispose
// ==========================================================================

var textValidatorsDisposeTestCase = coretestcases.CaseV1{
	Title:         "TextValidators.Dispose returns nil Items -- after dispose",
	ExpectedInput: args.Map{"isNil": true},
}
