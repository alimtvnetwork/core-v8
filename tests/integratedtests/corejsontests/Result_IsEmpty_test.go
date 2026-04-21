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
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var resultIsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmpty - empty bytes returns true",
		ArrangeInput: args.Map{
			"result": func() *corejson.Result {
				r := corejson.NewResult.UsingBytes([]byte{})
				return &r
			}(),
		},
		ExpectedInput: "true", // isEmpty
	},
	{
		Title: "IsEmpty - nil receiver returns true",
		ArrangeInput: args.Map{
			"result": (*corejson.Result)(nil),
		},
		ExpectedInput: "true", // isEmpty
	},
	{
		Title: "IsEmpty - valid bytes returns false",
		ArrangeInput: args.Map{
			"result": func() *corejson.Result {
				r := corejson.New(map[string]string{"key": "value"})
				return &r
			}(),
		},
		ExpectedInput: "false", // isEmpty
	},
}

func Test_Result_IsEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range resultIsEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		result := input["result"].(*corejson.Result)

		// Act
		actual := fmt.Sprintf("%v", result.IsEmpty())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actual)
	}
}
