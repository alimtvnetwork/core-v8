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
// SliceValidator.IsValid
// ==========================================================================

var svIsValidExactMatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- exact match",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidMismatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns false -- content mismatch",
	ExpectedInput: args.Map{"isValid": false},
}

var svIsValidLengthMismatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns false -- length mismatch",
	ExpectedInput: args.Map{"isValid": false},
}

var svIsValidBothNilTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- both nil",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidOneNilTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns false -- one nil",
	ExpectedInput: args.Map{"isValid": false},
}

var svIsValidBothEmptyTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- both empty",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidTrimMatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- trimmed match",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidContainsTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- contains substrings",
	ExpectedInput: args.Map{"isValid": true},
}

// ==========================================================================
// SliceValidator — helper methods
// ==========================================================================

var svActualLinesLengthTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.ActualLinesLength returns 2 -- two actual lines",
	ExpectedInput: args.Map{"length": 2},
}

var svExpectingLinesLengthTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.ExpectingLinesLength returns 3 -- three expected lines",
	ExpectedInput: args.Map{"length": 3},
}

var svIsUsedAlreadyFalseTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsUsedAlready returns false -- fresh instance",
	ExpectedInput: args.Map{"isUsed": false},
}

var svIsUsedAlreadyTrueTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsUsedAlready returns false -- ComparingValidators does not set isUsed",
	ExpectedInput: args.Map{"isUsed": false},
}

var svMethodNameTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.MethodName returns 'IsContains' -- Contains compare mode",
	ExpectedInput: args.Map{"name": "IsContains"},
}

// ==========================================================================
// SliceValidator.SetActual / SetActualVsExpected
// ==========================================================================

var svSetActualTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.SetActual returns length 1 -- one line set",
	ExpectedInput: args.Map{"length": 1},
}

var svSetActualVsExpectedTestCase = coretestcases.CaseV1{
	Title: "SliceValidator.SetActualVsExpected returns both set -- one actual one expected",
	ExpectedInput: args.Map{
		"actualLen":   1,
		"expectedLen": 1,
	},
}

// ==========================================================================
// SliceValidator.IsValidOtherLines
// ==========================================================================

var svIsValidOtherLinesMatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValidOtherLines returns true -- matching lines",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidOtherLinesMismatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValidOtherLines returns false -- mismatching lines",
	ExpectedInput: args.Map{"isValid": false},
}

// ==========================================================================
// SliceValidator.AllVerifyError
// ==========================================================================

var svAllVerifyErrorPassTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.AllVerifyError returns nil -- matching lines",
	ExpectedInput: args.Map{"hasError": false},
}

var svAllVerifyErrorFailTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.AllVerifyError returns error -- mismatched lines",
	ExpectedInput: args.Map{"hasError": true},
}

var svAllVerifyErrorSkipEmptyTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.AllVerifyError returns nil -- skip when actual empty",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// SliceValidator.VerifyFirstError
// ==========================================================================

var svVerifyFirstErrorPassTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.VerifyFirstError returns nil -- matching lines",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// SliceValidator.Dispose
// ==========================================================================

var svDisposeTestCase = coretestcases.CaseV1{
	Title: "SliceValidator.Dispose returns nil lines -- after dispose",
	ExpectedInput: args.Map{
		"actualNil":   true,
		"expectedNil": true,
	},
}

// ==========================================================================
// SliceValidator — case sensitivity
// ==========================================================================

var svCaseInsensitiveTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- case-insensitive match",
	ExpectedInput: args.Map{"isValid": true},
}

var svCaseSensitiveFailTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns false -- case-sensitive different case",
	ExpectedInput: args.Map{"isValid": false},
}

// ==========================================================================
// NewSliceValidatorUsingErr
// ==========================================================================

var svNewUsingErrNilTestCase = coretestcases.CaseV1{
	Title: "NewSliceValidatorUsingErr returns non-nil with 0 actual -- nil error input",
	ExpectedInput: args.Map{
		"isNotNil":  true,
		"actualLen": 0,
	},
}
