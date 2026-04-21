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

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/corepayload"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage27 — corepayload remaining ~35 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── Attributes/AttributesSetters/AttributesGetters/AttributesJson nil guards ──
// Lines 46, 84, 134, 130, 307, 119, 13, 19, 29 — all nil-receiver dead code

// ── PayloadWrapper nil guards ──
// Lines 134, 146, 188, 210, 230, 242, 276, 294, 335, 385 — mostly nil-receiver dead code

func Test_PayloadWrapper_Serialize_Valid_I29(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Identifier = "id1"
	pw.CategoryName = "cat1"
	pw.Name = "name1"

	// Act
	bytes, err := pw.Serialize()

	// Assert
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "PayloadWrapper Serialize valid", expected)
}

func Test_PayloadWrapper_ParseInjectUsingJson_I29(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Identifier = "id1"
	pw.CategoryName = "cat1"
	bytes, _ := pw.Serialize()
	jsonResult := corejson.NewResult.UsingBytes(bytes)

	pw2 := corepayload.New.PayloadWrapper.Empty()

	// Act
	result, err := pw2.ParseInjectUsingJson(&jsonResult)

	// Assert
	actual := args.Map{
		"notNil":   result != nil,
		"hasError": err != nil,
	}
	expected := args.Map{
		"notNil":   true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "PayloadWrapper ParseInjectUsingJson valid", expected)
}

// ── PayloadsCollectionFilter edge paths (lines 52-54, 61, 139) ──

func Test_PayloadsCollectionFilter_Empty_I29(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(0)

	// Act
	filtered := pc.FilterCategoryCollection("nonexistent")

	// Assert
	actual := args.Map{"isEmpty": filtered.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	actual.ShouldBeEqual(t, 1, "PayloadsCollectionFilter empty", expected)
}

// ── TypedPayloadCollection ──

func Test_TypedPayloadCollection_Clone_Empty_I29(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollection[string](0)

	// Act
	cloned, err := tc.Clone()

	// Assert
	actual := args.Map{
		"isEmpty":  cloned.IsEmpty(),
		"hasError": err != nil,
	}
	expected := args.Map{
		"isEmpty":  true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection Clone empty", expected)
}

func Test_TypedPayloadCollection_HasErrors_Empty_I29(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollection[string](0)

	// Act
	hasErrors := tc.HasErrors()
	firstErr := tc.FirstError()
	mergedErr := tc.MergedError()

	// Assert
	actual := args.Map{
		"hasErrors": hasErrors,
		"firstErr":  firstErr,
		"mergedErr": mergedErr,
	}
	expected := args.Map{
		"hasErrors": false,
		"firstErr":  nil,
		"mergedErr": nil,
	}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection error methods empty", expected)
}

func Test_TypedPayloadCollection_IsValid_Empty_I29(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollection[string](0)

	// Act
	isValid := tc.IsValid()

	// Assert
	actual := args.Map{"isValid": isValid}
	expected := args.Map{"isValid": true}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection IsValid empty", expected)
}

// ── TypedPayloadWrapper nil guards (lines 69, 101, 127, 289) — dead code ──

// ── TypedPayloadCollection remaining (lines 460, 474, 487, 574, 591, 616, 633, 648, 664, 668) ──
// Most are nil-guard, error fallback, or internal branch paths — dead code

func Test_TypedPayloadCollection_NewFromData_I29(t *testing.T) {
	// Arrange / Act
	tc, err := corepayload.NewTypedPayloadCollectionFromData[string](
		"test",
		[]string{"data1", "data2"},
	)

	// Assert
	actual := args.Map{
		"count":    tc.Length(),
		"hasError": err != nil,
	}
	expected := args.Map{
		"count":    2,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection NewFromData", expected)
}

func Test_TypedPayloadCollection_NewFromDataMust_I29(t *testing.T) {
	// Arrange / Act
	tc := corepayload.NewTypedPayloadCollectionFromDataMust[string](
		"test",
		[]string{"data1"},
	)

	// Assert
	actual := args.Map{"count": tc.Length()}
	expected := args.Map{"count": 1}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection NewFromDataMust", expected)
}

func Test_TypedPayloadCollection_ConcatNew_I29(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollectionFromDataMust[string](
		"test",
		[]string{"data1"},
	)

	// Act
	concat, err := tc.ConcatNew()

	// Assert
	actual := args.Map{
		"count":    concat.Length(),
		"hasError": err != nil,
	}
	expected := args.Map{
		"count":    1,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection ConcatNew", expected)
}
