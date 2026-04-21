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

package coredynamictests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ==========================================
// Test: IsInvalid
// ==========================================

func Test_CastedResult_IsInvalid_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsInvalidTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsInvalid()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsNotNull
// ==========================================

func Test_CastedResult_IsNotNull_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotNullTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsNotNull()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsNotPointer
// ==========================================

func Test_CastedResult_IsNotPointer_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotPointerTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsNotPointer()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsNotMatchingAcceptedType
// ==========================================

func Test_CastedResult_IsNotMatchingAcceptedType_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotMatchingAcceptedTypeTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsNotMatchingAcceptedType()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsSourceKind
// ==========================================

func Test_CastedResult_IsSourceKind_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsSourceKindTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsSourceKind(tc.CheckKind)}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: HasError
// ==========================================

func Test_CastedResult_HasError_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultHasErrorTestCases {
		// Act
		actual := args.Map{"result": tc.CR.HasError()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: HasAnyIssues
// ==========================================

func Test_CastedResult_HasAnyIssues_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultHasAnyIssuesTestCases {
		// Act
		actual := args.Map{"result": tc.CR.HasAnyIssues()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: SourceReflectType
// ==========================================

func Test_CastedResult_SourceReflectType_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultSourceReflectTypeTestCases {
		// Act
		actual := args.Map{
			"typeName":     tc.CR.SourceReflectType.Name(),
			"isStringKind": tc.CR.IsSourceKind(reflect.String),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Casted
// ==========================================

func Test_CastedResult_CastedValue_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultCastedValueTestCases {
		// Act
		actual := args.Map{
			"castedValue":  fmt.Sprintf("%v", tc.CR.Casted),
			"hasAnyIssues": tc.CR.HasAnyIssues(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsSourcePointer
// ==========================================

func Test_CastedResult_IsSourcePointer_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsSourcePointerTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsSourcePointer}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
