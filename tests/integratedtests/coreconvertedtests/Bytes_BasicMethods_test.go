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

package coreconvertedtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/converters/coreconverted"
)

func Test_Bytes_BasicMethods(t *testing.T) {
	// Arrange
	b := &coreconverted.Bytes{Values: []byte{1, 2, 3}, CombinedError: nil}
	bErr := &coreconverted.Bytes{Values: nil, CombinedError: errors.New("err")}
	var bNil *coreconverted.Bytes

	// Act
	actual := args.Map{
		"hasError":      b.HasError(),
		"length":        b.Length(),
		"hasAny":        b.HasAnyItem(),
		"isEmpty":       b.IsEmpty(),
		"hasIssues":     b.HasIssuesOrEmpty(),
		"errHasError":   bErr.HasError(),
		"errIsEmpty":    bErr.IsEmpty(),
		"errHasIssues":  bErr.HasIssuesOrEmpty(),
		"nilLength":     bNil.Length(),
	}
	expected := args.Map{
		"hasError":      false,
		"length":        3,
		"hasAny":        true,
		"isEmpty":       false,
		"hasIssues":     false,
		"errHasError":   true,
		"errIsEmpty":    true,
		"errHasIssues":  true,
		"nilLength":     0,
	}
	expected.ShouldBeEqual(t, 0, "Bytes_BasicMethods returns correct value -- with args", actual)
}

func Test_Integers_BasicMethods(t *testing.T) {
	// Arrange
	i := &coreconverted.Integers{Values: []int{1, 2, 3}, CombinedError: nil}
	iErr := &coreconverted.Integers{Values: nil, CombinedError: errors.New("err")}
	var iNil *coreconverted.Integers

	// Act
	actual := args.Map{
		"hasError":     i.HasError(),
		"length":       i.Length(),
		"hasAny":       i.HasAnyItem(),
		"isEmpty":      i.IsEmpty(),
		"hasIssues":    i.HasIssuesOrEmpty(),
		"errHasError":  iErr.HasError(),
		"errHasIssues": iErr.HasIssuesOrEmpty(),
		"nilLength":    iNil.Length(),
	}
	expected := args.Map{
		"hasError":     false,
		"length":       3,
		"hasAny":       true,
		"isEmpty":      false,
		"hasIssues":    false,
		"errHasError":  true,
		"errHasIssues": true,
		"nilLength":    0,
	}
	expected.ShouldBeEqual(t, 0, "Integers_BasicMethods returns correct value -- with args", actual)
}
