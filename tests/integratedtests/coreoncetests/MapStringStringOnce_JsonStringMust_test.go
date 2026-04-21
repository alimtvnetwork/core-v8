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

package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coreonce"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// MapStringStringOnce.JsonStringMust() — success path
// Covers MapStringStringOnce.go L306-317
// Note: Error panic path (L309-314) is unreachable — MarshalJSON on
// map[string]string never returns an error in standard Go.
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapStringStringOnce_JsonStringMust_FromMapStringStringOnceJ_FromMapStringStringOnceJ(t *testing.T) {
	// Arrange
	tc := cov13MapStringStringOnceJsonStringMustTestCase
	m := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{
			"key1": "val1",
		}
	})

	// Act
	noPanic := true
	var jsonStr string
	func() {
		defer func() {
			if r := recover(); r != nil {
				noPanic = false
			}
		}()
		jsonStr = m.JsonStringMust()
	}()

	actual := args.Map{
		"nonEmpty": jsonStr != "",
		"noPanic":  noPanic,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// StringsOnce.JsonStringMust() — success path
// Covers StringsOnce.go L248-258
// Note: Error panic path (L251-256) is unreachable — MarshalJSON on
// []string never returns an error in standard Go.
// ══════════════════════════════════════════════════════════════════════════════

func Test_StringsOnce_JsonStringMust_FromMapStringStringOnceJ_FromMapStringStringOnceJ(t *testing.T) {
	// Arrange
	tc := cov13StringsOnceJsonStringMustTestCase
	s := coreonce.NewStringsOnce(func() []string {
		return []string{"a", "b", "c"}
	})

	// Act
	noPanic := true
	var jsonStr string
	func() {
		defer func() {
			if r := recover(); r != nil {
				noPanic = false
			}
		}()
		jsonStr = s.JsonStringMust()
	}()

	actual := args.Map{
		"nonEmpty": jsonStr != "",
		"noPanic":  noPanic,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
