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

package corepayloadtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/corepayload"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/errcore"
)

func Test_PayloadWrapper_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range payloadWrapperCreateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		id, _ := input.GetAsString("id")
		line := []byte("some payload data")

		// Act
		payload, err := corepayload.New.PayloadWrapper.Create(
			name, id, "task-type", "category", line,
		)
		errcore.HandleErr(err)

		jsonResult := payload.JsonPtr()
		hasJson := !jsonResult.IsEmpty()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"name":       payload.Name,
			"identifier": payload.Identifier,
			"hasJson":    hasJson,
		})
	}
}

func Test_PayloadWrapper_DeserializeRoundtrip_Verification(t *testing.T) {
	for caseIndex, testCase := range payloadWrapperDeserializeRoundtripTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		id, _ := input.GetAsString("id")
		line := []byte("roundtrip payload bytes")

		// Act
		payload, err := corepayload.New.PayloadWrapper.Create(
			name, id, "task-type", "category", line,
		)
		errcore.HandleErr(err)

		jsonResult := payload.JsonPtr()
		restored, restoreErr := corepayload.New.PayloadWrapper.DeserializeUsingJsonResult(jsonResult)
		errcore.HandleErr(restoreErr)

		restoredJson := restored.JsonPtr()
		isEqual := jsonResult.IsEqual(*restoredJson)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"restoredName":       restored.Name,
			"restoredIdentifier": restored.Identifier,
			"jsonIsEqual":        isEqual,
		})
	}
}

func Test_PayloadWrapper_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range payloadWrapperCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		id, _ := input.GetAsString("id")
		newName, _ := input.GetAsString("new_name")
		line := []byte("clone payload")

		// Act
		original, err := corepayload.New.PayloadWrapper.Create(
			name, id, "task-type", "category", line,
		)
		errcore.HandleErr(err)

		cloned, cloneErr := original.ClonePtr(true)
		errcore.HandleErr(cloneErr)

		cloned.Name = newName
		originalUnchanged := original.Name != cloned.Name

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"originalName":  original.Name,
			"clonedName":    cloned.Name,
			"isIndependent": originalUnchanged,
		})
	}
}

func Test_PayloadWrapper_DeserializeToMany_Verification(t *testing.T) {
	for caseIndex, testCase := range payloadWrapperDeserializeToManyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 3)

		// Act
		wrappers := make([]*corepayload.PayloadWrapper, 0, count)
		for i := 0; i < count; i++ {
			payload, createErr := corepayload.New.PayloadWrapper.Create(
				fmt.Sprintf("item-%d", i),
				fmt.Sprintf("id-%d", i),
				"task", "cat",
				[]byte(fmt.Sprintf("data-%d", i)),
			)
			errcore.HandleErr(createErr)
			wrappers = append(wrappers, payload)
		}

		jsonSlice := corejson.Serialize.Apply(wrappers)
		jsonSlice.HandleError()

		deserialized, deserializeErr := corepayload.New.PayloadWrapper.DeserializeToMany(jsonSlice.Bytes)
		errcore.HandleErr(deserializeErr)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"deserializedCount": len(deserialized),
		})
	}
}
