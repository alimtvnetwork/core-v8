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
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corepayload"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// =============================================================================
// Attributes.IsEqual — Regression: logic inversion bug in IsSafeValid/HasIssuesOrEmpty
// =============================================================================

func Test_Attributes_IsEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range attributesIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := input.GetAsBoolDefault("left_nil", false)
		rightNil := input.GetAsBoolDefault("right_nil", false)
		samePointer := input.GetAsBoolDefault("same_pointer", false)

		var left, right *corepayload.Attributes

		if !leftNil {
			leftPayload, _ := input.GetAsString("left_payload")
			payload, _ := input.GetAsString("payload")

			if leftPayload == "" {
				leftPayload = payload
			}

			left = &corepayload.Attributes{
				DynamicPayloads: []byte(leftPayload),
			}
		}

		if samePointer {
			right = left
		} else if !rightNil {
			rightPayload, _ := input.GetAsString("right_payload")
			payload, _ := input.GetAsString("payload")

			if rightPayload == "" {
				rightPayload = payload
			}

			right = &corepayload.Attributes{
				DynamicPayloads: []byte(rightPayload),
			}
		}

		// Act
		result := left.IsEqual(right)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"isEqual": result,
		})
	}
}

// =============================================================================
// Attributes.Clone — Regression: deep clone independence
// =============================================================================

func Test_Attributes_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range attributesCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilAttr := input.GetAsBoolDefault("nil_attr", false)
		deep := input.GetAsBoolDefault("deep", false)

		if nilAttr {
			// Act
			var attr *corepayload.Attributes
			clonedPtr, err := attr.ClonePtr(deep)

			// Assert
			testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
				"isNil":    clonedPtr == nil,
				"hasError": err != nil,
			})

			continue
		}

		payload, _ := input.GetAsString("payload")
		attr := &corepayload.Attributes{
			DynamicPayloads: []byte(payload),
		}

		// Act
		cloned, err := attr.Clone(deep)
		if err != nil {
			testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
				"hasError": true,
			})

			continue
		}

		isEqual := attr.IsEqual(&cloned)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"clonedPayload": string(cloned.DynamicPayloads),
			"isEqual":       isEqual,
		})
	}
}

// =============================================================================
// Attributes.IsSafeValid — Regression: was returning HasIssuesOrEmpty() without negation
// =============================================================================

func Test_Attributes_IsSafeValid_Verification(t *testing.T) {
	for caseIndex, testCase := range attributesIsSafeValidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilAttr := input.GetAsBoolDefault("nil_attr", false)
		empty := input.GetAsBoolDefault("empty", false)

		var attr *corepayload.Attributes

		if !nilAttr && !empty {
			payload, _ := input.GetAsString("payload")
			attr = &corepayload.Attributes{
				DynamicPayloads: []byte(payload),
			}
		} else if !nilAttr && empty {
			attr = &corepayload.Attributes{}
		}

		// Act
		result := attr.IsSafeValid()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"isSafeValid": result,
		})
	}
}

// =============================================================================
// AuthInfo.Clone — Regression: was missing Identifier field in clone
// =============================================================================

func Test_AuthInfo_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range authInfoCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilAuth := input.GetAsBoolDefault("nil_auth", false)

		if nilAuth {
			// Act
			var auth *corepayload.AuthInfo
			cloned := auth.ClonePtr()

			// Assert
			testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
				"isNil": cloned == nil,
			})

			continue
		}

		identifier, _ := input.GetAsString("identifier")
		actionType, _ := input.GetAsString("action_type")
		resourceName, _ := input.GetAsString("resource_name")
		newActionType, _ := input.GetAsString("new_action_type")

		auth := &corepayload.AuthInfo{
			Identifier:   identifier,
			ActionType:   actionType,
			ResourceName: resourceName,
		}

		// Act
		cloned := auth.ClonePtr()

		if newActionType != "" {
			// Test independence: mutate clone
			cloned.ActionType = newActionType

			// Assert — original unchanged, clone mutated
			testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
				"originalAction": auth.ActionType,
				"clonedAction":   cloned.ActionType,
			})

			continue
		}

		// Assert — all fields including Identifier are preserved
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"identifier":   cloned.Identifier,
			"actionType":   cloned.ActionType,
			"resourceName": cloned.ResourceName,
		})
	}
}
