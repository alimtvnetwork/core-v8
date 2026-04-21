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

package conditionaltests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v8/conditional"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_AnyFunctionsExecuteResults_False(t *testing.T) {
	// Arrange
	result := conditional.AnyFunctionsExecuteResults(
		false,
		nil,
		[]func() (any, bool, bool){
			func() (any, bool, bool) { return "b", true, false },
		},
	)

	// Act
	actual := args.Map{"result": len(result) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_TypedErrorFunctionsExecuteResults_False(t *testing.T) {
	// Arrange
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](
		false,
		nil,
		[]func() (int, error){
			func() (int, error) { return 42, nil },
		},
	)

	// Act
	actual := args.Map{"result": err != nil || len(results) != 1 || results[0] != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_TypedErrorFunctionsExecuteResults_WithError(t *testing.T) {
	// Arrange
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](
		true,
		[]func() (int, error){
			func() (int, error) { return 0, errors.New("fail") },
			nil,
			func() (int, error) { return 1, nil },
		},
		nil,
	)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual = args.Map{"result": len(results) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_TypedErrorFunctionsExecuteResults_Empty(t *testing.T) {
	// Arrange
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](
		true,
		nil,
		nil,
	)

	// Act
	actual := args.Map{"result": err != nil || results != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}
