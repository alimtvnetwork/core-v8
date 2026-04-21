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

package defaulterrtests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/defaulterr"
)

func Test_DefaultErr_AllSentinels(t *testing.T) {
	errorMap := map[string]error{
		"Marshalling":                        defaulterr.Marshalling,
		"UnMarshalling":                      defaulterr.UnMarshalling,
		"OutOfRange":                         defaulterr.OutOfRange,
		"CannotProcessNilOrEmpty":            defaulterr.CannotProcessNilOrEmpty,
		"NegativeDataCannotProcess":          defaulterr.NegativeDataCannotProcess,
		"NilResult":                          defaulterr.NilResult,
		"UnexpectedValue":                    defaulterr.UnexpectedValue,
		"CannotRemoveFromEmptyCollection":    defaulterr.CannotRemoveFromEmptyCollection,
		"MarshallingFailedDueToNilOrEmpty":   defaulterr.MarshallingFailedDueToNilOrEmpty,
		"UnmarshallingFailedDueToNilOrEmpty": defaulterr.UnmarshallingFailedDueToNilOrEmpty,
		"KeyNotExistInMap":                   defaulterr.KeyNotExistInMap,
	}

	for caseIndex, tc := range defaultErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		errorName, _ := input.GetAsString("error")
		err := errorMap[errorName]

		// Act
		actual := args.Map{
			"isNotNil":   err != nil,
			"hasMessage": err != nil && err.Error() != "",
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
