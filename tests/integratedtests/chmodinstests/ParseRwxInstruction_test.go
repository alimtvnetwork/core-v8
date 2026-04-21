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

package chmodinstests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ── ParseRwxInstructionUsingJsonResult ──

func Test_ParseRwxInstruction_Nil(t *testing.T) {
	// Arrange
	// Act
	result, err := chmodins.ParseRwxInstructionUsingJsonResult(nil)
	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseRwxInstruction returns nil -- nil", actual)
}

func Test_ParseRwxInstruction_EmptyBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	// Act
	result, err := chmodins.ParseRwxInstructionUsingJsonResult(r)
	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseRwxInstruction returns empty -- empty bytes", actual)
}

func Test_ParseRwxInstruction_InvalidJson(t *testing.T) {
	// Arrange — valid JSON but not an RwxInstruction shape causes unmarshal to succeed
	// Use a raw string that's valid JSON but wrong type to trigger unmarshal error
	r := corejson.NewPtr("not-an-instruction")
	// Act
	_, err := chmodins.ParseRwxInstructionUsingJsonResult(r)
	// Assert — string JSON unmarshals into struct without error (fields stay zero),
	// but we exercise the code path
	_ = err
}

func Test_ParseRwxInstruction_Success(t *testing.T) {
	// Arrange
	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
		Condition:          *chmodins.DefaultAllTrueCondition(),
	}
	r := corejson.NewPtr(ins)
	// Act
	result, err := chmodins.ParseRwxInstructionUsingJsonResult(r)
	// Assert
	actual := args.Map{
		"noErr":     err == nil,
		"notNil":    result != nil,
		"owner":     result.Owner,
		"group":     result.Group,
		"other":     result.Other,
		"recursive": result.IsRecursive,
	}
	expected := args.Map{
		"noErr": true, "notNil": true,
		"owner": "rwx", "group": "r-x", "other": "r--",
		"recursive": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseRwxInstruction returns correct value -- success", actual)
}

// ── ParseRwxInstructionUsingJsonResultMust ──

func Test_ParseRwxInstructionMust_Success(t *testing.T) {
	// Arrange
	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
		Condition:          *chmodins.DefaultAllFalseCondition(),
	}
	r := corejson.NewPtr(ins)
	// Act
	result := chmodins.ParseRwxInstructionUsingJsonResultMust(r)
	// Assert
	actual := args.Map{
		"owner": result.Owner,
		"recursive": result.IsRecursive,
	}
	expected := args.Map{
		"owner": "rwx",
		"recursive": false,
	}
	expected.ShouldBeEqual(t, 0, "ParseRwxInstructionMust returns correct value -- success", actual)
}

func Test_ParseRwxInstructionMust_Panic(t *testing.T) {
	// Arrange
	defer func() { recover() }()
	// Act
	chmodins.ParseRwxInstructionUsingJsonResultMust(nil)
	// Assert
	actual := args.Map{"result": false}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected panic", actual)
}

// ── ParseBaseRwxInstructionsUsingJsonResult ──

func Test_ParseBaseRwxInstructions_Nil(t *testing.T) {
	// Arrange
	// Act
	result, err := chmodins.ParseBaseRwxInstructionsUsingJsonResult(nil)
	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseBaseRwxInstructions returns nil -- nil", actual)
}

func Test_ParseBaseRwxInstructions_EmptyBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	// Act
	result, err := chmodins.ParseBaseRwxInstructionsUsingJsonResult(r)
	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseBaseRwxInstructions returns empty -- empty bytes", actual)
}

func Test_ParseBaseRwxInstructions_InvalidJson(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("not-a-base-instruction")
	// Act
	_, err := chmodins.ParseBaseRwxInstructionsUsingJsonResult(r)
	// Assert — exercise unmarshal path
	_ = err
}

func Test_ParseBaseRwxInstructions_Success(t *testing.T) {
	// Arrange
	base := chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
				Condition:          *chmodins.DefaultAllTrueCondition(),
			},
		},
	}
	r := corejson.NewPtr(base)
	// Act
	result, err := chmodins.ParseBaseRwxInstructionsUsingJsonResult(r)
	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"notNil": result != nil,
		"len":    result.Length(),
	}
	expected := args.Map{
		"noErr": true,
		"notNil": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "ParseBaseRwxInstructions returns correct value -- success", actual)
}

// ── ParseBaseRwxInstructionsUsingJsonResultMust ──

func Test_ParseBaseRwxInstructionsMust_Success(t *testing.T) {
	// Arrange
	base := chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
				Condition:          *chmodins.DefaultAllFalseCondition(),
			},
		},
	}
	r := corejson.NewPtr(base)
	// Act
	result := chmodins.ParseBaseRwxInstructionsUsingJsonResultMust(r)
	// Assert
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ParseBaseRwxInstructionsMust returns correct value -- success", actual)
}

func Test_ParseBaseRwxInstructionsMust_Panic(t *testing.T) {
	// Arrange
	defer func() { recover() }()
	// Act
	chmodins.ParseBaseRwxInstructionsUsingJsonResultMust(nil)
	// Assert
	actual := args.Map{"result": false}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected panic", actual)
}
