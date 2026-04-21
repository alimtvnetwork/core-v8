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
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corerange"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/corevalidator"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
)

func Test_LineNumber_Methods(t *testing.T) {
	// Arrange
	ln := corevalidator.LineNumber{LineNumber: 5}

	// Act
	actual := args.Map{
		"has": ln.HasLineNumber(),
		"isMatch": ln.IsMatch(5),
		"mismatchErr": ln.VerifyError(6) != nil,
	}

	// Assert
	expected := args.Map{
		"has": true,
		"isMatch": true,
		"mismatchErr": true,
	}
	expected.ShouldBeEqual(t, 0, "LineNumber returns correct value -- methods", actual)
}

func Test_Condition_Methods(t *testing.T) {
	// Arrange
	c := corevalidator.Condition{IsUniqueWordOnly: true}

	// Act
	actual := args.Map{"splitByWhitespace": c.IsSplitByWhitespace()}

	// Assert
	expected := args.Map{"splitByWhitespace": true}
	expected.ShouldBeEqual(t, 0, "Condition returns correct value -- methods", actual)
}

func Test_Parameter_Methods(t *testing.T) {
	// Arrange
	p := corevalidator.Parameter{CaseIndex: 1, Header: "h", IsCaseSensitive: false}

	// Act
	actual := args.Map{
		"caseIndex": p.CaseIndex,
		"header": p.Header,
		"ignoreCase": p.IsIgnoreCase(),
	}

	// Assert
	expected := args.Map{
		"caseIndex": 1,
		"header": "h",
		"ignoreCase": true,
	}
	expected.ShouldBeEqual(t, 0, "Parameter returns correct value -- methods", actual)
}

func Test_RangesSegment(t *testing.T) {
	// Arrange
	rs := corevalidator.RangesSegment{RangeInt: corerange.RangeInt{Start: 0, End: 10}, ExpectedLines: []string{"a"}, CompareAs: stringcompareas.Equal}

	// Act
	actual := args.Map{
		"start": rs.Start,
		"end": rs.End,
		"expLen": len(rs.ExpectedLines),
		"compareAs": rs.CompareAs.Name(),
	}

	// Assert
	expected := args.Map{
		"start": 0,
		"end": 10,
		"expLen": 1,
		"compareAs": "Equal",
	}
	expected.ShouldBeEqual(t, 0, "RangesSegment returns correct value -- struct", actual)
}

func Test_BaseValidatorCoreCondition_Default(t *testing.T) {
	// Arrange
	bv := corevalidator.BaseValidatorCoreCondition{}
	cond := bv.ValidatorCoreConditionDefault()

	// Act
	actual := args.Map{
		"notNil": bv.ValidatorCoreCondition != nil,
		"split": cond.IsSplitByWhitespace(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"split": false,
	}
	expected.ShouldBeEqual(t, 0, "BaseValidatorCoreCondition returns non-empty -- default", actual)
}
