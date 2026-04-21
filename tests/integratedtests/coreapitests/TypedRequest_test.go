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

package coreapitests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coreapi"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ==========================================
// Test: NewTypedRequest
// ==========================================

func Test_NewTypedRequest_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}

		// Act
		req := coreapi.NewTypedRequest[string](attr, payload)
		actual := args.Map{
			"payload": req.Request,
			"isValid": fmt.Sprintf("%v", req.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: InvalidTypedRequest
// ==========================================

func Test_InvalidTypedRequest_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestInvalidTestCases {
		// Arrange — nil attribute

		// Act
		req := coreapi.InvalidTypedRequest[string](nil)
		actual := args.Map{
			"payload":      req.Request,
			"isValid":      fmt.Sprintf("%v", req.Attribute.IsValid),
			"hasAttribute": fmt.Sprintf("%v", req.Attribute != nil),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TypedRequest Clone
// ==========================================

func Test_TypedRequest_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}
		req := coreapi.NewTypedRequest[string](attr, payload)

		// Act
		cloned := req.Clone()
		actual := args.Map{
			"payload":       cloned.Request,
			"isValid":       fmt.Sprintf("%v", cloned.Attribute.IsValid),
			"isIndependent": fmt.Sprintf("%v", cloned != req),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedRequest_Clone_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestCloneNilTestCases {
		// Arrange
		var req *coreapi.TypedRequest[string]

		// Act
		cloned := req.Clone()
		actLines := []string{
			fmt.Sprintf("%v", cloned == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponseResult ToTypedResponse
// ==========================================

func Test_TypedResponseResult_ToTypedResponse_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultToTypedResponseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response, _ := input.GetAsString("response")

		attr := &coreapi.ResponseAttribute{IsValid: true}
		result := coreapi.NewTypedResponseResult[string](attr, response)

		// Act
		resp := result.ToTypedResponse()
		actual := args.Map{
			"response": resp.Response,
			"isValid":  fmt.Sprintf("%v", resp.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedResponseResult_ToTypedResponse_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultToTypedResponseNilTestCases {
		// Arrange
		var result *coreapi.TypedResponseResult[string]

		// Act
		resp := result.ToTypedResponse()
		actLines := []string{
			fmt.Sprintf("%v", resp == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponseResult Message
// ==========================================

func Test_TypedResponseResult_Message_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		var result *coreapi.TypedResponseResult[string]

		if response, ok := input.GetAsString("response"); ok {
			message, _ := input.GetAsString("message")
			attr := &coreapi.ResponseAttribute{IsValid: true, Message: message}
			result = coreapi.NewTypedResponseResult[string](attr, response)
		}

		// Act
		var msg string
		if result != nil {
			msg = result.Message()
		}
		actLines := []string{
			msg,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
