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

package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/enums/versionindexes"
	"github.com/alimtvnetwork/core-v8/errcore"
)

var indexByName = map[string]versionindexes.Index{
	"Major": versionindexes.Major,
	"Minor": versionindexes.Minor,
	"Patch": versionindexes.Patch,
	"Build": versionindexes.Build,
}

func Test_Index_JsonRoundtrip_Verification(t *testing.T) {
	for caseIndex, testCase := range jsonRoundtripTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		indexName, _ := input.GetAsString("index")
		idx := indexByName[indexName]

		// Act
		jsonResult := idx.Json()
		var restored versionindexes.Index
		err := restored.JsonParseSelfInject(&jsonResult)
		errcore.HandleErr(err)

		actual := args.Map{
			"indexName":  restored.Name(),
			"indexValue": restored.ToNumberString(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Index_NameAndNameValue_Verification(t *testing.T) {
	for caseIndex, testCase := range nameAndNameValueTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		indexName, _ := input.GetAsString("index")
		idx := indexByName[indexName]

		// Act
		actual := args.Map{
			"name":      idx.Name(),
			"nameValue": idx.NameValue(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Index_JsonParseSelfInject_Verification(t *testing.T) {
	for caseIndex, testCase := range jsonParseSelfInjectTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sourceName, _ := input.GetAsString("source")
		targetName, _ := input.GetAsString("target")
		source := indexByName[sourceName]
		target := indexByName[targetName]

		// Act
		sourceJson := source.Json()
		err := target.JsonParseSelfInject(&sourceJson)
		errcore.HandleErr(err)

		actual := args.Map{
			"resultName":      target.Name(),
			"resultNameValue": target.NameValue(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
