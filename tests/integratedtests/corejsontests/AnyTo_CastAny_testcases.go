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

package corejsontests

import (
	"errors"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// anyTo — SerializedJsonResult branches
// =============================================================================

// Branch: nil input → error result
var anyToSerializedJsonResultNilTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedJsonResult returns error -- nil input",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

// Branch: Result type → direct return
var anyToSerializedJsonResultResultTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedJsonResult returns no error -- Result type input",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// Branch: *Result type → direct return
var anyToSerializedJsonResultResultPtrTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedJsonResult returns no error -- *Result type input",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// Branch: []byte → NewResult.UsingBytesTypePtr
var anyToSerializedJsonResultBytesTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedJsonResult returns no error -- byte slice input",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// Branch: string → NewResult.UsingBytesTypePtr
var anyToSerializedJsonResultStringTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedJsonResult returns no error -- string input",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// Branch: error with message → NewResult.UsingTypePlusString
var anyToSerializedJsonResultErrorTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedJsonResult returns no error -- error with message",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// Branch: error with empty message → NewResult.UsingBytesTypePtr
var anyToSerializedJsonResultErrorEmptyTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedJsonResult returns no error -- error with empty message",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// Branch: default → Serialize.Apply
var anyToSerializedJsonResultDefaultTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedJsonResult returns no error -- integer fallback to default",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// =============================================================================
// anyTo — SerializedString / SerializedSafeString
// =============================================================================

var anyToSerializedStringErrorTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedString returns error -- nil input",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var anyToSerializedSafeStringNilTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedSafeString returns empty -- nil input swallowed",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

// =============================================================================
// anyTo — JsonString / JsonStringWithErr branches
// =============================================================================

// Branch: string → direct return
var anyToJsonStringStringTestCase = coretestcases.CaseV1{
	Title: "AnyTo JsonString returns raw -- string type passthrough",
	ExpectedInput: args.Map{
		"result": "raw",
	},
}

// Branch: []byte → BytesToString
var anyToJsonStringBytesTestCase = coretestcases.CaseV1{
	Title: "AnyTo JsonString returns non-empty -- byte slice input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

// Branch: Result → .JsonString()
var anyToJsonStringResultTestCase = coretestcases.CaseV1{
	Title: "AnyTo JsonString returns non-empty -- Result type input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

// Branch: *Result → .JsonString()
var anyToJsonStringResultPtrTestCase = coretestcases.CaseV1{
	Title: "AnyTo JsonString returns non-empty -- *Result type input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

// Branch: default → New(anyItem).JsonString()
var anyToJsonStringDefaultTestCase = coretestcases.CaseV1{
	Title: "AnyTo JsonString returns non-empty -- integer fallback default",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

// Branch: JsonStringWithErr Result with error → return error
var anyToJsonStringWithErrResultErrorTestCase = coretestcases.CaseV1{
	Title: "AnyTo JsonStringWithErr returns error -- Result with error",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

// Branch: JsonStringWithErr *Result with error → return error
var anyToJsonStringWithErrResultPtrErrorTestCase = coretestcases.CaseV1{
	Title: "AnyTo JsonStringWithErr returns error -- *Result with error",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

// =============================================================================
// anyTo — PrettyStringWithError branches
// =============================================================================

// Branch: string → direct return
var anyToPrettyStringWithErrorStringTestCase = coretestcases.CaseV1{
	Title: "AnyTo PrettyStringWithError returns hello -- string passthrough",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// Branch: []byte → BytesToPrettyString
var anyToPrettyStringWithErrorBytesTestCase = coretestcases.CaseV1{
	Title: "AnyTo PrettyStringWithError returns non-empty -- byte slice input",
	ExpectedInput: args.Map{
		"hasError":   false,
		"hasContent": true,
	},
}

// Branch: Result with error → return pretty + error
var anyToPrettyStringWithErrorResultErrTestCase = coretestcases.CaseV1{
	Title: "AnyTo PrettyStringWithError returns error -- Result with error",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

// Branch: *Result with error → return pretty + error
var anyToPrettyStringWithErrorResultPtrErrTestCase = coretestcases.CaseV1{
	Title: "AnyTo PrettyStringWithError returns error -- *Result with error",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

// =============================================================================
// anyTo — SafeJsonPrettyString branches
// =============================================================================

var anyToSafeJsonPrettyStringStringTestCase = coretestcases.CaseV1{
	Title: "AnyTo SafeJsonPrettyString returns hello -- string passthrough",
	ExpectedInput: args.Map{
		"result": "hello",
	},
}

var anyToSafeJsonPrettyStringBytesTestCase = coretestcases.CaseV1{
	Title: "AnyTo SafeJsonPrettyString returns non-empty -- byte slice input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

var anyToSafeJsonPrettyStringResultTestCase = coretestcases.CaseV1{
	Title: "AnyTo SafeJsonPrettyString returns non-empty -- Result type input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

var anyToSafeJsonPrettyStringResultPtrTestCase = coretestcases.CaseV1{
	Title: "AnyTo SafeJsonPrettyString returns non-empty -- *Result type input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

var anyToSafeJsonPrettyStringDefaultTestCase = coretestcases.CaseV1{
	Title: "AnyTo SafeJsonPrettyString returns non-empty -- integer default",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

// =============================================================================
// anyTo — SerializedFieldsMap
// =============================================================================

var anyToSerializedFieldsMapTestCase = coretestcases.CaseV1{
	Title: "AnyTo SerializedFieldsMap returns map -- with valid input",
	ExpectedInput: args.Map{
		"hasError": false,
		"hasName":  true,
	},
}

// =============================================================================
// castingAny — FromToOption branches
// =============================================================================

// Branch: bytes → Deserialize.UsingBytes
var castAnyFromToBytesTestCase = coretestcases.CaseV1{
	Title: "CastAny FromToOption returns hello -- byte slice to string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// Branch: string → Deserialize.UsingBytes
var castAnyFromToStringTestCase = coretestcases.CaseV1{
	Title: "CastAny FromToOption returns hello -- json string to string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// Branch: Result → .Deserialize
var castAnyFromToResultTestCase = coretestcases.CaseV1{
	Title: "CastAny FromToOption returns hello -- Result to string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// Branch: *Result → .Deserialize
var castAnyFromToResultPtrTestCase = coretestcases.CaseV1{
	Title: "CastAny FromToOption returns hello -- *Result to string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// Branch: func() ([]byte, error) → serializerFunc
var castAnyFromToSerializerFuncTestCase = coretestcases.CaseV1{
	Title: "CastAny FromToOption returns hello -- serializer func to string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// Branch: error → .Error() as json bytes
var castAnyFromToErrorTestCase = coretestcases.CaseV1{
	Title: "CastAny FromToOption returns hello -- error message as json string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// Branch: default → Serialize.Apply then Deserialize
var castAnyFromToDefaultTestCase = coretestcases.CaseV1{
	Title: "CastAny FromToOption returns 42 -- struct to struct via json",
	ExpectedInput: args.Map{
		"hasError": false,
		"value":    42,
	},
}

// Branch: reflection same type both ptr → direct set
var castAnyFromToReflectionTestCase = coretestcases.CaseV1{
	Title: "CastAny FromToDefault returns hello -- same ptr type reflection",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// Branch: nil from → error
var castAnyFromToNilFromTestCase = coretestcases.CaseV1{
	Title: "CastAny FromToOption returns error -- nil from input",
	ExpectedInput: args.Map{
		"hasIssue": true,
	},
}

// suppress unused import
var _ = errors.New
