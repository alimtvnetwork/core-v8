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

package errcoretests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ── RawErrCollection Basic ──

var rawErrCollBasicTestCases = []coretestcases.CaseV1{
	{Title: "RawErrCollection returns empty state -- new collection", ExpectedInput: args.Map{
		"isEmpty": true, "hasError": false, "hasAnyError": false,
		"len": 0, "hasAnyIssues": false, "isValid": true,
		"isSuccess": true, "isFailed": false, "isInvalid": false,
		"isCollectionType": true,
	}},
}

// ── RawErrCollection Add methods ──

var rawErrCollAddTestCases = []coretestcases.CaseV1{
	{Title: "RawErrCollection.Add returns 1 -- nil then error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddError returns 1 -- nil then error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.Adds returns 2 -- mixed nil and errors", ExpectedInput: args.Map{"len": 2}},
	{Title: "RawErrCollection.AddErrors returns 1 -- one error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddString returns 1 -- empty then string", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddMsg returns 1 -- one msg", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddIf returns 1 -- false then true", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddFunc returns 1 -- nil, nil-return, error-return", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddFuncIf returns 1 -- various combos", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.ConditionalAddError returns 1 -- false then true", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddMsgStackTrace returns 1 -- empty then msg", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddStackTrace returns 1 -- nil then error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddMsgErrStackTrace returns 1 -- nil then error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddMethodName returns 1 -- empty then msg", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddMessages returns 1 -- empty then messages", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddErrorWithMessage returns 1 -- nil then error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddErrorWithMessageRef returns 2 -- various combos", ExpectedInput: args.Map{"len": 2}},
	{Title: "RawErrCollection.AddFmt returns 1 -- nil then error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.Fmt returns 1 -- empty then formatted", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.FmtIf returns 1 -- false then true", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.References returns 1 -- msg and ref", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddWithRef returns 1 -- nil then error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddWithTraceRef returns 1 -- nil then error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddWithCompiledTraceRef returns 1 -- nil then error", ExpectedInput: args.Map{"len": 1}},
	{Title: "RawErrCollection.AddStringSliceAsErr returns 2 -- empty and non-empty strings", ExpectedInput: args.Map{"len": 2}},
	{Title: "RawErrCollection.AddErrorGetters does not panic -- empty", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "RawErrCollection.AddCompiledErrorGetters does not panic -- empty", ExpectedInput: args.Map{"noPanic": true}},
}

// ── RawErrCollection String/Error methods ──

var rawErrCollStringTestCases = []coretestcases.CaseV1{
	{Title: "RawErrCollection.Strings returns correct -- empty then one", ExpectedInput: args.Map{
		"emptyLen": 0,
		"oneLen": 1,
	}},
	{Title: "RawErrCollection.String returns correct -- empty then one", ExpectedInput: args.Map{
		"emptyStr": "",
		"nonEmpty": true,
	}},
	{Title: "RawErrCollection.StringUsingJoiner returns correct -- empty then one", ExpectedInput: args.Map{
		"emptyStr": "",
		"nonEmpty": true,
	}},
	{Title: "RawErrCollection.StringUsingJoinerAdditional returns correct -- empty then one", ExpectedInput: args.Map{
		"emptyStr": "",
		"nonEmpty": true,
	}},
	{Title: "RawErrCollection.StringWithAdditionalMessage returns correct -- empty then one", ExpectedInput: args.Map{
		"emptyStr": "",
		"nonEmpty": true,
	}},
}

var rawErrCollCompiledTestCases = []coretestcases.CaseV1{
	{Title: "RawErrCollection.CompiledError returns correct -- empty then one", ExpectedInput: args.Map{
		"emptyNil": true,
		"nonNil": true,
	}},
	{Title: "RawErrCollection.CompiledErrorUsingJoiner returns correct -- empty then one", ExpectedInput: args.Map{
		"emptyNil": true,
		"nonNil": true,
	}},
	{Title: "RawErrCollection.CompiledErrorUsingJoinerAdditionalMessage returns correct -- empty then one", ExpectedInput: args.Map{
		"emptyNil": true,
		"nonNil": true,
	}},
	{Title: "RawErrCollection.CompiledErrorWithStackTraces returns correct -- empty then one", ExpectedInput: args.Map{
		"emptyNil": true,
		"nonNil": true,
	}},
}

var rawErrCollMiscTestCases = []coretestcases.CaseV1{
	{Title: "RawErrCollection.Serialize returns nil -- empty collection", ExpectedInput: args.Map{
		"bytesNil": true,
		"errNil": true,
	}},
	{Title: "RawErrCollection.SerializeWithoutTraces returns nil -- empty collection", ExpectedInput: args.Map{
		"bytesNil": true,
		"errNil": true,
	}},
	{Title: "RawErrCollection.MarshalJSON returns nil -- empty collection", ExpectedInput: args.Map{
		"bytesNil": true,
		"errNil": true,
	}},
	{Title: "RawErrCollection.UnmarshalJSON does not panic -- empty array", ExpectedInput: args.Map{"noPanic": true}},
	{Title: "RawErrCollection.Value returns nil -- empty", ExpectedInput: args.Map{"isNil": true}},
	{Title: "RawErrCollection.Clear returns 0 -- after add", ExpectedInput: args.Map{"len": 0}},
	{Title: "RawErrCollection.IsErrorsCollected returns correct -- nil and error", ExpectedInput: args.Map{
		"nilFalse": true,
		"errTrue": true,
	}},
	{Title: "RawErrCollection.CountStateChangeTracker returns same -- no changes", ExpectedInput: args.Map{"isSame": true}},
	{Title: "RawErrCollection.ToRawErrCollection returns non-nil -- valid", ExpectedInput: args.Map{"nonNil": true}},
	{Title: "RawErrCollection.ReflectSetTo returns error -- value type", ExpectedInput: args.Map{
		"valueErr": true,
		"nilPtrErr": true,
		"validNoErr": true,
		"otherErr": true,
	}},
}
