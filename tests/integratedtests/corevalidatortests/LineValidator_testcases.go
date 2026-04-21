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
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ==========================================================================
// LineValidator.IsMatch
// ==========================================================================

var lineValidatorIsMatchBothTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true when line and text match",
	ExpectedInput: args.Map{"isMatch": true},
}

var lineValidatorIsMatchLineMismatchTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns false on line number mismatch",
	ExpectedInput: args.Map{"isMatch": false},
}

var lineValidatorIsMatchTextMismatchTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns false on text mismatch",
	ExpectedInput: args.Map{"isMatch": false},
}

var lineValidatorIsMatchSkipLineTestCase = coretestcases.CaseV1{
	Title:         "IsMatch with skip line number passes on text match",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// LineValidator.IsMatchMany
// ==========================================================================

var lineValidatorIsMatchManyAllTestCase = coretestcases.CaseV1{
	Title:         "IsMatchMany returns true when all match",
	ExpectedInput: args.Map{"isMatch": true},
}

var lineValidatorIsMatchManyOneFailsTestCase = coretestcases.CaseV1{
	Title:         "IsMatchMany returns false when one fails",
	ExpectedInput: args.Map{"isMatch": false},
}

var lineValidatorIsMatchManyEmptySkipTestCase = coretestcases.CaseV1{
	Title:         "IsMatchMany empty with skip returns true",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// LineValidator.VerifyError
// ==========================================================================

var lineValidatorVerifyErrorMatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyError returns nil on match",
	ExpectedInput: args.Map{"hasError": false},
}

var lineValidatorVerifyErrorLineMismatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyError returns error on line mismatch",
	ExpectedInput: args.Map{"hasError": true},
}

var lineValidatorVerifyErrorTextMismatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyError returns error on text mismatch",
	ExpectedInput: args.Map{"hasError": true},
}

// ==========================================================================
// LineValidator.VerifyMany
// ==========================================================================

var lineValidatorVerifyManyContinueTestCase = coretestcases.CaseV1{
	Title:         "VerifyMany continueOnError collects errors",
	ExpectedInput: args.Map{"hasError": true},
}

var lineValidatorVerifyManyFirstOnlyTestCase = coretestcases.CaseV1{
	Title:         "VerifyMany firstOnly returns first error",
	ExpectedInput: args.Map{"hasError": true},
}
