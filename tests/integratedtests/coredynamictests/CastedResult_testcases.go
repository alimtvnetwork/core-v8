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
	"errors"
	"reflect"

	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// Note: Nil receiver test cases migrated to CastedResult_NilReceiver_testcases.go
// using CaseNilSafe pattern with direct method references.

type castedResultTestCase struct {
	Case coretestcases.CaseV1
	CR   *coredynamic.CastedResult
}

// ==========================================
// IsInvalid
// ==========================================

var castedResultIsInvalidTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsInvalid true on nil receiver",
			ExpectedInput: args.Map{"result": true},
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsInvalid false when IsValid=true",
			ExpectedInput: args.Map{"result": false},
		},
		CR: &coredynamic.CastedResult{IsValid: true},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsInvalid true when IsValid=false",
			ExpectedInput: args.Map{"result": true},
		},
		CR: &coredynamic.CastedResult{IsValid: false},
	},
}

// ==========================================
// IsNotNull
// ==========================================

var castedResultIsNotNullTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotNull false on nil receiver",
			ExpectedInput: args.Map{"result": false},
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotNull true when IsNull=false",
			ExpectedInput: args.Map{"result": true},
		},
		CR: &coredynamic.CastedResult{IsNull: false},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotNull false when IsNull=true",
			ExpectedInput: args.Map{"result": false},
		},
		CR: &coredynamic.CastedResult{IsNull: true},
	},
}

// ==========================================
// IsNotPointer
// ==========================================

var castedResultIsNotPointerTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotPointer false on nil receiver",
			ExpectedInput: args.Map{"result": false},
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotPointer true when IsPointer=false",
			ExpectedInput: args.Map{"result": true},
		},
		CR: &coredynamic.CastedResult{IsPointer: false},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotPointer false when IsPointer=true",
			ExpectedInput: args.Map{"result": false},
		},
		CR: &coredynamic.CastedResult{IsPointer: true},
	},
}

// ==========================================
// IsNotMatchingAcceptedType
// ==========================================

var castedResultIsNotMatchingAcceptedTypeTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotMatchingAcceptedType false on nil receiver",
			ExpectedInput: args.Map{"result": false},
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotMatchingAcceptedType true when not matching",
			ExpectedInput: args.Map{"result": true},
		},
		CR: &coredynamic.CastedResult{IsMatchingAcceptedType: false},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsNotMatchingAcceptedType false when matching",
			ExpectedInput: args.Map{"result": false},
		},
		CR: &coredynamic.CastedResult{IsMatchingAcceptedType: true},
	},
}

// ==========================================
// IsSourceKind
// ==========================================

type castedResultIsSourceKindTestCase struct {
	Case      coretestcases.CaseV1
	CR        *coredynamic.CastedResult
	CheckKind reflect.Kind
}

var castedResultIsSourceKindTestCases = []castedResultIsSourceKindTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsSourceKind false on nil receiver",
			ExpectedInput: args.Map{"result": false},
		},
		CR:        nil,
		CheckKind: reflect.String,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsSourceKind true on kind match",
			ExpectedInput: args.Map{"result": true},
		},
		CR:        &coredynamic.CastedResult{SourceKind: reflect.Int},
		CheckKind: reflect.Int,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "IsSourceKind false on mismatch",
			ExpectedInput: args.Map{"result": false},
		},
		CR:        &coredynamic.CastedResult{SourceKind: reflect.Int},
		CheckKind: reflect.String,
	},
}

// ==========================================
// HasError
// ==========================================

var castedResultHasErrorTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "HasError false on nil receiver",
			ExpectedInput: args.Map{"result": false},
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasError true when error present",
			ExpectedInput: args.Map{"result": true},
		},
		CR: &coredynamic.CastedResult{Error: errors.New("fail")},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasError false when no error",
			ExpectedInput: args.Map{"result": false},
		},
		CR: &coredynamic.CastedResult{},
	},
}

// ==========================================
// HasAnyIssues
// ==========================================

var castedResultHasAnyIssuesTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues true on nil receiver",
			ExpectedInput: args.Map{"result": true},
		},
		CR: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues true when invalid",
			ExpectedInput: args.Map{"result": true},
		},
		CR: &coredynamic.CastedResult{IsValid: false, IsMatchingAcceptedType: true},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues true when null",
			ExpectedInput: args.Map{"result": true},
		},
		CR: &coredynamic.CastedResult{IsValid: true, IsNull: true, IsMatchingAcceptedType: true},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues true when type not matching",
			ExpectedInput: args.Map{"result": true},
		},
		CR: &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: false},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "HasAnyIssues false when all good",
			ExpectedInput: args.Map{"result": false},
		},
		CR: &coredynamic.CastedResult{
			IsValid:                true,
			IsNull:                 false,
			IsMatchingAcceptedType: true,
		},
	},
}

// ==========================================
// SourceReflectType
// ==========================================

var castedResultSourceReflectTypeTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "Stores SourceReflectType correctly",
			ExpectedInput: args.Map{
				"typeName":     "string",
				"isStringKind": true,
			},
		},
		CR: &coredynamic.CastedResult{
			SourceReflectType: reflect.TypeOf(""),
			SourceKind:        reflect.String,
		},
	},
}

// ==========================================
// Casted
// ==========================================

var castedResultCastedValueTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "Casted stores value and HasAnyIssues false",
			ExpectedInput: args.Map{
				"castedValue":  42,
				"hasAnyIssues": false,
			},
		},
		CR: &coredynamic.CastedResult{
			Casted:                 42,
			IsValid:                true,
			IsMatchingAcceptedType: true,
		},
	},
}

// ==========================================
// IsSourcePointer
// ==========================================

var castedResultIsSourcePointerTestCases = []castedResultTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "IsSourcePointer field works",
			ExpectedInput: args.Map{"result": true},
		},
		CR: &coredynamic.CastedResult{IsSourcePointer: true},
	},
}
