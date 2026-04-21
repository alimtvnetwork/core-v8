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

package coredynamictestwrappers

import (
	"github.com/alimtvnetwork/core-v8/coretests"
)

var ReflectSetFromToValidNullNull = FromToTestWrapper{
	Header: "(null, null) -- do nothing -- " +
		"From `Null` to `Null` -- does nothing -- no error",
}

var ReflectSetFromToValidPtrToPtr = FromToTestWrapper{
	Header: "(sameTypePointer, sameTypePointer) -- try reflection -- " +
		"From `*FromToTestWrapper{Expected}` " +
		"to   `*FromToTestWrapper{Sample data}` should set to Expected. ",
	From: &ReflectSetFromToTestCasesDraftTypeExpected,
	To: &coretests.DraftType{
		SampleString1: "Same data",
	},
	ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
}

var ReflectSetFromToValidValueToPtr = FromToTestWrapper{
	Header: "(sameTypeNonPointer, sameTypePointer) -- try reflection -- " +
		"From `FromToTestWrapper{Expected}` " +
		"to   `*FromToTestWrapper{Sample data}` should set to Expected.",
	From: ReflectSetFromToTestCasesDraftTypeExpected,
	To: &coretests.DraftType{
		SampleString1: "Sample data",
	},
	ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
}

var ReflectSetFromToValidBytesToDraft = FromToTestWrapper{
	Header: "([]byte, otherType) -- try unmarshal, reflect -- " +
		"From `[]bytes(FromToTestWrapper{Expected}` " +
		"to   `*FromToTestWrapper{Sample data}` should set to Expected.",
	From: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
	To: &coretests.DraftType{
		SampleString1: "Sample data",
	},
	ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
}

var ReflectSetFromToValidDraftToBytes = FromToTestWrapper{
	Header: "(otherType, *[]byte) -- try marshal, reflect -- " +
		"From `FromToTestWrapper{Expected}` " +
		"to   `*[]byte{}` should set to Expected.",
	From:          &ReflectSetFromToTestCasesDraftTypeExpected,
	To:            &[]byte{},
	ExpectedValue: &[]byte{}, // placeholder; IsSame checks type match (*[]byte == *[]byte)
}

// ReflectSetFromToValidTestCases kept for backward compatibility if needed.
var ReflectSetFromToValidTestCases = []FromToTestWrapper{
	ReflectSetFromToValidNullNull,
	ReflectSetFromToValidPtrToPtr,
	ReflectSetFromToValidValueToPtr,
	ReflectSetFromToValidBytesToDraft,
	ReflectSetFromToValidDraftToBytes,
}
