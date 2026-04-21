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

package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ═══════════════════════════════════════════════════════════════════════
// KeyAnyValuePair — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_01_KeyAnyValuePair_KeyName(t *testing.T) {
	safeTest(t, "Test_01_KeyAnyValuePair_KeyName", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{"result": kv.KeyName() != "k"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected k", actual)
	})
}

func Test_02_KeyAnyValuePair_VariableName(t *testing.T) {
	safeTest(t, "Test_02_KeyAnyValuePair_VariableName", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{"result": kv.VariableName() != "k"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected k", actual)
	})
}

func Test_03_KeyAnyValuePair_ValueAny(t *testing.T) {
	safeTest(t, "Test_03_KeyAnyValuePair_ValueAny", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{"result": kv.ValueAny() != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
	})
}

func Test_04_KeyAnyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_04_KeyAnyValuePair_IsVariableNameEqual", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k"}

		// Act
		actual := args.Map{"result": kv.IsVariableNameEqual("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": kv.IsVariableNameEqual("x")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_05_KeyAnyValuePair_IsValueNull_Nil(t *testing.T) {
	safeTest(t, "Test_05_KeyAnyValuePair_IsValueNull_Nil", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k"}

		// Act
		actual := args.Map{"result": kv.IsValueNull()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_06_KeyAnyValuePair_IsValueNull_NilReceiver(t *testing.T) {
	safeTest(t, "Test_06_KeyAnyValuePair_IsValueNull_NilReceiver", func() {
		// Arrange
		var kv *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{"result": kv.IsValueNull()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_07_KeyAnyValuePair_IsValueNull_NonNil(t *testing.T) {
	safeTest(t, "Test_07_KeyAnyValuePair_IsValueNull_NonNil", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{"result": kv.IsValueNull()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_08_KeyAnyValuePair_HasNonNull(t *testing.T) {
	safeTest(t, "Test_08_KeyAnyValuePair_HasNonNull", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{"result": kv.HasNonNull()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_09_KeyAnyValuePair_HasNonNull_Nil(t *testing.T) {
	safeTest(t, "Test_09_KeyAnyValuePair_HasNonNull_Nil", func() {
		// Arrange
		var kv *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{"result": kv.HasNonNull()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_10_KeyAnyValuePair_HasValue(t *testing.T) {
	safeTest(t, "Test_10_KeyAnyValuePair_HasValue", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.HasValue()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_11_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	safeTest(t, "Test_11_KeyAnyValuePair_IsValueEmptyString", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k"}

		// Act
		actual := args.Map{"result": kv.IsValueEmptyString()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_12_KeyAnyValuePair_IsValueEmptyString_NilReceiver(t *testing.T) {
	safeTest(t, "Test_12_KeyAnyValuePair_IsValueEmptyString_NilReceiver", func() {
		// Arrange
		var kv *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{"result": kv.IsValueEmptyString()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_13_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	safeTest(t, "Test_13_KeyAnyValuePair_IsValueWhitespace", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k"}

		// Act
		actual := args.Map{"result": kv.IsValueWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_14_KeyAnyValuePair_IsValueWhitespace_NilReceiver(t *testing.T) {
	safeTest(t, "Test_14_KeyAnyValuePair_IsValueWhitespace_NilReceiver", func() {
		// Arrange
		var kv *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{"result": kv.IsValueWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_15_KeyAnyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_15_KeyAnyValuePair_ValueString", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s := kv.ValueString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_16_KeyAnyValuePair_ValueString_Cached(t *testing.T) {
	safeTest(t, "Test_16_KeyAnyValuePair_ValueString_Cached", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s1 := kv.ValueString()
		s2 := kv.ValueString()

		// Act
		actual := args.Map{"result": s1 != s2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same", actual)
	})
}

func Test_17_KeyAnyValuePair_ValueString_NilValue(t *testing.T) {
	safeTest(t, "Test_17_KeyAnyValuePair_ValueString_NilValue", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k"}
		s := kv.ValueString()
		_ = s // should not panic
	})
}

func Test_18_KeyAnyValuePair_String(t *testing.T) {
	safeTest(t, "Test_18_KeyAnyValuePair_String", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_19_KeyAnyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_19_KeyAnyValuePair_Compile", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.Compile() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_20_KeyAnyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_20_KeyAnyValuePair_SerializeMust", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()

		// Act
		actual := args.Map{"result": len(b) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_21_KeyAnyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_21_KeyAnyValuePair_Serialize", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b, err := kv.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(b) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_22_KeyAnyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_22_KeyAnyValuePair_Json", func() {
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kv.Json()
		_ = j
	})
}

func Test_23_KeyAnyValuePair_JsonPtr(t *testing.T) {
	safeTest(t, "Test_23_KeyAnyValuePair_JsonPtr", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_24_KeyAnyValuePair_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_24_KeyAnyValuePair_ParseInjectUsingJson", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jp := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		_, err := kv2.ParseInjectUsingJson(jp)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_25_KeyAnyValuePair_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_25_KeyAnyValuePair_ParseInjectUsingJsonMust", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jp := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		result := kv2.ParseInjectUsingJsonMust(jp)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_26_KeyAnyValuePair_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_26_KeyAnyValuePair_AsJsonContractsBinder", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.AsJsonContractsBinder() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_27_KeyAnyValuePair_AsJsoner(t *testing.T) {
	safeTest(t, "Test_27_KeyAnyValuePair_AsJsoner", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.AsJsoner() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_28_KeyAnyValuePair_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_28_KeyAnyValuePair_JsonParseSelfInject", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jp := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		err := kv2.JsonParseSelfInject(jp)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_29_KeyAnyValuePair_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_29_KeyAnyValuePair_AsJsonParseSelfInjector", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.AsJsonParseSelfInjector() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_30_KeyAnyValuePair_Clear(t *testing.T) {
	safeTest(t, "Test_30_KeyAnyValuePair_Clear", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Clear()

		// Act
		actual := args.Map{"result": kv.Key != "" || kv.Value != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cleared", actual)
	})
}

func Test_31_KeyAnyValuePair_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_31_KeyAnyValuePair_Clear_Nil", func() {
		var kv *corestr.KeyAnyValuePair
		kv.Clear() // no panic
	})
}

func Test_32_KeyAnyValuePair_Dispose(t *testing.T) {
	safeTest(t, "Test_32_KeyAnyValuePair_Dispose", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Dispose()
	})
}

func Test_33_KeyAnyValuePair_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_33_KeyAnyValuePair_Dispose_Nil", func() {
		var kv *corestr.KeyAnyValuePair
		kv.Dispose()
	})
}

// ═══════════════════════════════════════════════════════════════════════
// KeyValueCollection — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_34_KeyValueCollection_Add(t *testing.T) {
	safeTest(t, "Test_34_KeyValueCollection_Add", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_35_KeyValueCollection_AddIf_True(t *testing.T) {
	safeTest(t, "Test_35_KeyValueCollection_AddIf_True", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(true, "k", "v")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_36_KeyValueCollection_AddIf_False(t *testing.T) {
	safeTest(t, "Test_36_KeyValueCollection_AddIf_False", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(false, "k", "v")

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_37_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_37_KeyValueCollection_Adds", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_38_KeyValueCollection_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_38_KeyValueCollection_Adds_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_39_KeyValueCollection_Count(t *testing.T) {
	safeTest(t, "Test_39_KeyValueCollection_Count", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.Count() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_40_KeyValueCollection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_40_KeyValueCollection_HasAnyItem", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_41_KeyValueCollection_LastIndex(t *testing.T) {
	safeTest(t, "Test_41_KeyValueCollection_LastIndex", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")

		// Act
		actual := args.Map{"result": kvc.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_42_KeyValueCollection_HasIndex(t *testing.T) {
	safeTest(t, "Test_42_KeyValueCollection_HasIndex", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": kvc.HasIndex(0)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": kvc.HasIndex(5)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_43_KeyValueCollection_First(t *testing.T) {
	safeTest(t, "Test_43_KeyValueCollection_First", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": kvc.First().Key != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_44_KeyValueCollection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_44_KeyValueCollection_FirstOrDefault", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"result": kvc.FirstOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		kvc.Add("a", "1")
		actual = args.Map{"result": kvc.FirstOrDefault().Key != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_45_KeyValueCollection_Last(t *testing.T) {
	safeTest(t, "Test_45_KeyValueCollection_Last", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")

		// Act
		actual := args.Map{"result": kvc.Last().Key != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_46_KeyValueCollection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_46_KeyValueCollection_LastOrDefault", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"result": kvc.LastOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		kvc.Add("a", "1")
		actual = args.Map{"result": kvc.LastOrDefault().Key != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_47_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_47_KeyValueCollection_Find", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "b", false
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_48_KeyValueCollection_Find_Empty(t *testing.T) {
	safeTest(t, "Test_48_KeyValueCollection_Find_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})

		// Act
		actual := args.Map{"result": len(found) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_49_KeyValueCollection_Find_Break(t *testing.T) {
	safeTest(t, "Test_49_KeyValueCollection_Find_Break", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		kvc.Add("c", "3")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, i == 0
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_50_KeyValueCollection_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_50_KeyValueCollection_SafeValueAt", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": kvc.SafeValueAt(0) != "1"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": kvc.SafeValueAt(99) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_51_KeyValueCollection_SafeValueAt_Empty(t *testing.T) {
	safeTest(t, "Test_51_KeyValueCollection_SafeValueAt_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"result": kvc.SafeValueAt(0) != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_52_KeyValueCollection_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_52_KeyValueCollection_SafeValuesAtIndexes", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		vals := kvc.SafeValuesAtIndexes(0, 1)

		// Act
		actual := args.Map{"result": len(vals) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_53_KeyValueCollection_SafeValuesAtIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_53_KeyValueCollection_SafeValuesAtIndexes_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		vals := kvc.SafeValuesAtIndexes()

		// Act
		actual := args.Map{"result": len(vals) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_54_KeyValueCollection_Strings(t *testing.T) {
	safeTest(t, "Test_54_KeyValueCollection_Strings", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		s := kvc.Strings()

		// Act
		actual := args.Map{"result": len(s) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_55_KeyValueCollection_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_55_KeyValueCollection_Strings_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		s := kvc.Strings()

		// Act
		actual := args.Map{"result": len(s) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_56_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_56_KeyValueCollection_StringsUsingFormat", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		s := kvc.StringsUsingFormat("%s=%s")

		// Act
		actual := args.Map{"result": len(s) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_57_KeyValueCollection_StringsUsingFormat_Empty(t *testing.T) {
	safeTest(t, "Test_57_KeyValueCollection_StringsUsingFormat_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		s := kvc.StringsUsingFormat("%s=%s")

		// Act
		actual := args.Map{"result": len(s) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_58_KeyValueCollection_String(t *testing.T) {
	safeTest(t, "Test_58_KeyValueCollection_String", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": kvc.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_59_KeyValueCollection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_59_KeyValueCollection_Length_Nil", func() {
		// Arrange
		var kvc *corestr.KeyValueCollection

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_60_KeyValueCollection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_60_KeyValueCollection_IsEmpty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"result": kvc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_61_KeyValueCollection_Compile(t *testing.T) {
	safeTest(t, "Test_61_KeyValueCollection_Compile", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.Compile() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_62_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_62_KeyValueCollection_AddStringBySplit", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplit("=", "key=value")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_63_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_63_KeyValueCollection_AddStringBySplitTrim", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplitTrim("=", " key = value ")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_64_KeyValueCollection_AddMap(t *testing.T) {
	safeTest(t, "Test_64_KeyValueCollection_AddMap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_65_KeyValueCollection_AddMap_Nil(t *testing.T) {
	safeTest(t, "Test_65_KeyValueCollection_AddMap_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(nil)

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_66_KeyValueCollection_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_66_KeyValueCollection_AddHashsetMap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(map[string]bool{"a": true})

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_67_KeyValueCollection_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_67_KeyValueCollection_AddHashsetMap_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(nil)

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_68_KeyValueCollection_AddHashset(t *testing.T) {
	safeTest(t, "Test_68_KeyValueCollection_AddHashset", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashset(hs)

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_69_KeyValueCollection_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_69_KeyValueCollection_AddHashset_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashset(nil)

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_70_KeyValueCollection_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_70_KeyValueCollection_AddsHashmap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("a", "1")
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmap(hm)

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_71_KeyValueCollection_AddsHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_71_KeyValueCollection_AddsHashmap_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmap(nil)

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_72_KeyValueCollection_Hashmap(t *testing.T) {
	safeTest(t, "Test_72_KeyValueCollection_Hashmap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		hm := kvc.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_73_KeyValueCollection_Hashmap_Empty(t *testing.T) {
	safeTest(t, "Test_73_KeyValueCollection_Hashmap_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		hm := kvc.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_74_KeyValueCollection_IsContains(t *testing.T) {
	safeTest(t, "Test_74_KeyValueCollection_IsContains", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": kvc.IsContains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": kvc.IsContains("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_75_KeyValueCollection_Get(t *testing.T) {
	safeTest(t, "Test_75_KeyValueCollection_Get", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		v, ok := kvc.Get("a")

		// Act
		actual := args.Map{"result": ok || v != "1"}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_, ok2 := kvc.Get("z")
		actual = args.Map{"result": ok2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_76_KeyValueCollection_Map(t *testing.T) {
	safeTest(t, "Test_76_KeyValueCollection_Map", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		m := kvc.Map()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_77_KeyValueCollection_HasKey(t *testing.T) {
	safeTest(t, "Test_77_KeyValueCollection_HasKey", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": kvc.HasKey("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": kvc.HasKey("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_78_KeyValueCollection_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_78_KeyValueCollection_AllKeysSorted", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("b", "2")
		kvc.Add("a", "1")
		keys := kvc.AllKeysSorted()

		// Act
		actual := args.Map{"result": keys[0] != "a" || keys[1] != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted", actual)
	})
}

func Test_79_KeyValueCollection_AllKeys(t *testing.T) {
	safeTest(t, "Test_79_KeyValueCollection_AllKeys", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		keys := kvc.AllKeys()

		// Act
		actual := args.Map{"result": len(keys) != 1 || keys[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_80_KeyValueCollection_AllKeys_Empty(t *testing.T) {
	safeTest(t, "Test_80_KeyValueCollection_AllKeys_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		keys := kvc.AllKeys()

		// Act
		actual := args.Map{"result": len(keys) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_81_KeyValueCollection_AllValues(t *testing.T) {
	safeTest(t, "Test_81_KeyValueCollection_AllValues", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		vals := kvc.AllValues()

		// Act
		actual := args.Map{"result": len(vals) != 1 || vals[0] != "1"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_82_KeyValueCollection_AllValues_Empty(t *testing.T) {
	safeTest(t, "Test_82_KeyValueCollection_AllValues_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		vals := kvc.AllValues()

		// Act
		actual := args.Map{"result": len(vals) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_83_KeyValueCollection_Join(t *testing.T) {
	safeTest(t, "Test_83_KeyValueCollection_Join", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		j := kvc.Join(",")

		// Act
		actual := args.Map{"result": j == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_84_KeyValueCollection_JoinKeys(t *testing.T) {
	safeTest(t, "Test_84_KeyValueCollection_JoinKeys", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")

		// Act
		actual := args.Map{"result": kvc.JoinKeys(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_85_KeyValueCollection_JoinValues(t *testing.T) {
	safeTest(t, "Test_85_KeyValueCollection_JoinValues", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")

		// Act
		actual := args.Map{"result": kvc.JoinValues(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_86_KeyValueCollection_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_86_KeyValueCollection_AddsHashmaps", func() {
		// Arrange
		hm1 := corestr.New.Hashmap.Cap(2)
		hm1.AddOrUpdate("a", "1")
		hm2 := corestr.New.Hashmap.Cap(2)
		hm2.AddOrUpdate("b", "2")
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmaps(hm1, hm2)

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_87_KeyValueCollection_AddsHashmaps_Nil(t *testing.T) {
	safeTest(t, "Test_87_KeyValueCollection_AddsHashmaps_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmaps()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_88_KeyValueCollection_JsonModel(t *testing.T) {
	safeTest(t, "Test_88_KeyValueCollection_JsonModel", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": len(kvc.JsonModel()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_89_KeyValueCollection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_89_KeyValueCollection_JsonModelAny", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": kvc.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_90_KeyValueCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_90_KeyValueCollection_Serialize", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, err := kvc.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(b) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_91_KeyValueCollection_SerializeMust(t *testing.T) {
	safeTest(t, "Test_91_KeyValueCollection_SerializeMust", func() {
		// Arrange
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b := kvc.SerializeMust()

		// Act
		actual := args.Map{"result": len(b) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_92_KeyValueCollection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_92_KeyValueCollection_MarshalJSON", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, err := kvc.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil || len(b) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_93_KeyValueCollection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_93_KeyValueCollection_UnmarshalJSON", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, _ := kvc.MarshalJSON()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"result": err != nil || kvc2.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_94_KeyValueCollection_UnmarshalJSON_Empty(t *testing.T) {
	safeTest(t, "Test_94_KeyValueCollection_UnmarshalJSON_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		err := kvc.UnmarshalJSON([]byte("[]"))

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil err", actual)
	})
}

func Test_95_KeyValueCollection_Json(t *testing.T) {
	safeTest(t, "Test_95_KeyValueCollection_Json", func() {
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		j := kvc.Json()
		_ = j
	})
}

func Test_96_KeyValueCollection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_96_KeyValueCollection_JsonPtr", func() {
		// Arrange
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": kvc.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_97_KeyValueCollection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_97_KeyValueCollection_ParseInjectUsingJson", func() {
		// Arrange
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jp := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		_, err := kvc2.ParseInjectUsingJson(jp)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_98_KeyValueCollection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_98_KeyValueCollection_AsJsonContractsBinder", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"result": kvc.AsJsonContractsBinder() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_99_KeyValueCollection_AsJsoner(t *testing.T) {
	safeTest(t, "Test_99_KeyValueCollection_AsJsoner", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"result": kvc.AsJsoner() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_100_KeyValueCollection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_100_KeyValueCollection_JsonParseSelfInject", func() {
		// Arrange
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jp := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.JsonParseSelfInject(jp)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_101_KeyValueCollection_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_101_KeyValueCollection_AsJsonParseSelfInjector", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"result": kvc.AsJsonParseSelfInjector() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_102_KeyValueCollection_Clear(t *testing.T) {
	safeTest(t, "Test_102_KeyValueCollection_Clear", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Clear()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_103_KeyValueCollection_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_103_KeyValueCollection_Clear_Nil", func() {
		var kvc *corestr.KeyValueCollection
		kvc.Clear()
	})
}

func Test_104_KeyValueCollection_Dispose(t *testing.T) {
	safeTest(t, "Test_104_KeyValueCollection_Dispose", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Dispose()
	})
}

func Test_105_KeyValueCollection_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_105_KeyValueCollection_Dispose_Nil", func() {
		var kvc *corestr.KeyValueCollection
		kvc.Dispose()
	})
}

func Test_106_KeyValueCollection_Deserialize(t *testing.T) {
	safeTest(t, "Test_106_KeyValueCollection_Deserialize", func() {
		// Arrange
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		var target corestr.KeyValueCollection
		err := kvc.Deserialize(&target)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// SimpleStringOnce — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_107_SimpleStringOnce_Value(t *testing.T) {
	safeTest(t, "Test_107_SimpleStringOnce_Value", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": s.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_108_SimpleStringOnce_IsInitialized(t *testing.T) {
	safeTest(t, "Test_108_SimpleStringOnce_IsInitialized", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": s.IsInitialized()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_109_SimpleStringOnce_IsDefined(t *testing.T) {
	safeTest(t, "Test_109_SimpleStringOnce_IsDefined", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": s.IsDefined()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_110_SimpleStringOnce_IsUninitialized(t *testing.T) {
	safeTest(t, "Test_110_SimpleStringOnce_IsUninitialized", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": s.IsUninitialized()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_111_SimpleStringOnce_Invalidate(t *testing.T) {
	safeTest(t, "Test_111_SimpleStringOnce_Invalidate", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")
		s.Invalidate()

		// Act
		actual := args.Map{"result": s.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_112_SimpleStringOnce_Reset(t *testing.T) {
	safeTest(t, "Test_112_SimpleStringOnce_Reset", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")
		s.Reset()

		// Act
		actual := args.Map{"result": s.IsInitialized() || s.Value() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reset", actual)
	})
}

func Test_113_SimpleStringOnce_IsInvalid(t *testing.T) {
	safeTest(t, "Test_113_SimpleStringOnce_IsInvalid", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": s.IsInvalid()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_114_SimpleStringOnce_IsInvalid_Valid(t *testing.T) {
	safeTest(t, "Test_114_SimpleStringOnce_IsInvalid_Valid", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("x")

		// Act
		actual := args.Map{"result": s.IsInvalid()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_115_SimpleStringOnce_ValueBytes(t *testing.T) {
	safeTest(t, "Test_115_SimpleStringOnce_ValueBytes", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": string(s.ValueBytes()) != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_116_SimpleStringOnce_ValueBytesPtr(t *testing.T) {
	safeTest(t, "Test_116_SimpleStringOnce_ValueBytesPtr", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": string(s.ValueBytesPtr()) != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_117_SimpleStringOnce_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_117_SimpleStringOnce_SetOnUninitialized", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}
		err := s.SetOnUninitialized("hello")

		// Act
		actual := args.Map{"result": err != nil || s.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_118_SimpleStringOnce_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_118_SimpleStringOnce_SetOnUninitialized_AlreadyInit", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")
		err := s.SetOnUninitialized("world")

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_119_SimpleStringOnce_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_119_SimpleStringOnce_GetSetOnce", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}
		v := s.GetSetOnce("hello")

		// Act
		actual := args.Map{"result": v != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		v2 := s.GetSetOnce("world")
		actual = args.Map{"result": v2 != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello still", actual)
	})
}

func Test_120_SimpleStringOnce_GetOnce(t *testing.T) {
	safeTest(t, "Test_120_SimpleStringOnce_GetOnce", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}
		v := s.GetOnce()

		// Act
		actual := args.Map{"result": v != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": s.IsInitialized()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)
	})
}

func Test_121_SimpleStringOnce_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_121_SimpleStringOnce_GetOnceFunc", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}
		v := s.GetOnceFunc(func() string { return "hello" })

		// Act
		actual := args.Map{"result": v != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		v2 := s.GetOnceFunc(func() string { return "world" })
		actual = args.Map{"result": v2 != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello still", actual)
	})
}

func Test_122_SimpleStringOnce_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_122_SimpleStringOnce_SetOnceIfUninitialized", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": s.SetOnceIfUninitialized("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": s.SetOnceIfUninitialized("world")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_123_SimpleStringOnce_SetInitialize(t *testing.T) {
	safeTest(t, "Test_123_SimpleStringOnce_SetInitialize", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}
		s.SetInitialize()

		// Act
		actual := args.Map{"result": s.IsInitialized()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_124_SimpleStringOnce_SetUnInit(t *testing.T) {
	safeTest(t, "Test_124_SimpleStringOnce_SetUnInit", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("x")
		s.SetUnInit()

		// Act
		actual := args.Map{"result": s.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_125_SimpleStringOnce_ConcatNew(t *testing.T) {
	safeTest(t, "Test_125_SimpleStringOnce_ConcatNew", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")
		s2 := s.ConcatNew(" world")

		// Act
		actual := args.Map{"result": s2.Value() != "hello world"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello world", actual)
	})
}

func Test_126_SimpleStringOnce_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_126_SimpleStringOnce_ConcatNewUsingStrings", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("a")
		s2 := s.ConcatNewUsingStrings(",", "b", "c")

		// Act
		actual := args.Map{"result": s2.Value() != "a,b,c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_127_SimpleStringOnce_IsEmpty(t *testing.T) {
	safeTest(t, "Test_127_SimpleStringOnce_IsEmpty", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": s.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_128_SimpleStringOnce_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_128_SimpleStringOnce_IsWhitespace", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("   ")

		// Act
		actual := args.Map{"result": s.IsWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_129_SimpleStringOnce_Trim(t *testing.T) {
	safeTest(t, "Test_129_SimpleStringOnce_Trim", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("  hello  ")

		// Act
		actual := args.Map{"result": s.Trim() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_130_SimpleStringOnce_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_130_SimpleStringOnce_HasValidNonEmpty", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("x")

		// Act
		actual := args.Map{"result": s.HasValidNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_131_SimpleStringOnce_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_131_SimpleStringOnce_HasValidNonWhitespace", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("x")

		// Act
		actual := args.Map{"result": s.HasValidNonWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_132_SimpleStringOnce_IsValueBool(t *testing.T) {
	safeTest(t, "Test_132_SimpleStringOnce_IsValueBool", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("true")

		// Act
		actual := args.Map{"result": s.IsValueBool()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_133_SimpleStringOnce_SafeValue(t *testing.T) {
	safeTest(t, "Test_133_SimpleStringOnce_SafeValue", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("x")

		// Act
		actual := args.Map{"result": s.SafeValue() != "x"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_134_SimpleStringOnce_SafeValue_Uninit(t *testing.T) {
	safeTest(t, "Test_134_SimpleStringOnce_SafeValue_Uninit", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": s.SafeValue() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_135_SimpleStringOnce_Uint16(t *testing.T) {
	safeTest(t, "Test_135_SimpleStringOnce_Uint16", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("100")
		v, ok := s.Uint16()

		// Act
		actual := args.Map{"result": ok || v != 100}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
	})
}

func Test_136_SimpleStringOnce_Uint32(t *testing.T) {
	safeTest(t, "Test_136_SimpleStringOnce_Uint32", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("1000")
		v, ok := s.Uint32()

		// Act
		actual := args.Map{"result": ok || v != 1000}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 1000", actual)
	})
}

func Test_137_SimpleStringOnce_WithinRange(t *testing.T) {
	safeTest(t, "Test_137_SimpleStringOnce_WithinRange", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("50")
		v, ok := s.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{"result": ok || v != 50}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
	})
}

func Test_138_SimpleStringOnce_WithinRange_OutOfRange(t *testing.T) {
	safeTest(t, "Test_138_SimpleStringOnce_WithinRange_OutOfRange", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("200")
		v, ok := s.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{"result": ok || v != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100 bounded", actual)
	})
}

func Test_139_SimpleStringOnce_WithinRange_NoBoundary(t *testing.T) {
	safeTest(t, "Test_139_SimpleStringOnce_WithinRange_NoBoundary", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("200")
		v, ok := s.WithinRange(false, 0, 100)

		// Act
		actual := args.Map{"result": ok || v != 200}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200 no boundary", actual)
	})
}

func Test_140_SimpleStringOnce_WithinRange_Below(t *testing.T) {
	safeTest(t, "Test_140_SimpleStringOnce_WithinRange_Below", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("-5")
		v, ok := s.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{"result": ok || v != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 bounded", actual)
	})
}

func Test_141_SimpleStringOnce_WithinRange_ParseErr(t *testing.T) {
	safeTest(t, "Test_141_SimpleStringOnce_WithinRange_ParseErr", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("abc")
		v, ok := s.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{"result": ok || v != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_142_SimpleStringOnce_WithinRangeDefault(t *testing.T) {
	safeTest(t, "Test_142_SimpleStringOnce_WithinRangeDefault", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("50")
		v, ok := s.WithinRangeDefault(0, 100)

		// Act
		actual := args.Map{"result": ok || v != 50}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
	})
}

func Test_143_SimpleStringOnce_Int(t *testing.T) {
	safeTest(t, "Test_143_SimpleStringOnce_Int", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("42")

		// Act
		actual := args.Map{"result": s.Int() != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
	})
}

func Test_144_SimpleStringOnce_Int_Error(t *testing.T) {
	safeTest(t, "Test_144_SimpleStringOnce_Int_Error", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		actual := args.Map{"result": s.Int() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_145_SimpleStringOnce_Byte(t *testing.T) {
	safeTest(t, "Test_145_SimpleStringOnce_Byte", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("200")

		// Act
		actual := args.Map{"result": s.Byte() != 200}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200", actual)
	})
}

func Test_146_SimpleStringOnce_Byte_OutOfRange(t *testing.T) {
	safeTest(t, "Test_146_SimpleStringOnce_Byte_OutOfRange", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("300")

		// Act
		actual := args.Map{"result": s.Byte() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_147_SimpleStringOnce_Byte_Error(t *testing.T) {
	safeTest(t, "Test_147_SimpleStringOnce_Byte_Error", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		actual := args.Map{"result": s.Byte() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_148_SimpleStringOnce_Int16(t *testing.T) {
	safeTest(t, "Test_148_SimpleStringOnce_Int16", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("100")

		// Act
		actual := args.Map{"result": s.Int16() != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
	})
}

func Test_149_SimpleStringOnce_Int16_OutOfRange(t *testing.T) {
	safeTest(t, "Test_149_SimpleStringOnce_Int16_OutOfRange", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("40000")

		// Act
		actual := args.Map{"result": s.Int16() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_150_SimpleStringOnce_Int32(t *testing.T) {
	safeTest(t, "Test_150_SimpleStringOnce_Int32", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("1000")

		// Act
		actual := args.Map{"result": s.Int32() != 1000}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1000", actual)
	})
}

func Test_151_SimpleStringOnce_Int32_OutOfRange(t *testing.T) {
	safeTest(t, "Test_151_SimpleStringOnce_Int32_OutOfRange", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("3000000000")

		// Act
		actual := args.Map{"result": s.Int32() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_152_SimpleStringOnce_BooleanDefault(t *testing.T) {
	safeTest(t, "Test_152_SimpleStringOnce_BooleanDefault", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("true")

		// Act
		actual := args.Map{"result": s.BooleanDefault()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_153_SimpleStringOnce_Boolean_Yes(t *testing.T) {
	safeTest(t, "Test_153_SimpleStringOnce_Boolean_Yes", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("yes")

		// Act
		actual := args.Map{"result": s.Boolean(false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_154_SimpleStringOnce_Boolean_Y(t *testing.T) {
	safeTest(t, "Test_154_SimpleStringOnce_Boolean_Y", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("y")

		// Act
		actual := args.Map{"result": s.Boolean(false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_155_SimpleStringOnce_Boolean_1(t *testing.T) {
	safeTest(t, "Test_155_SimpleStringOnce_Boolean_1", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("1")

		// Act
		actual := args.Map{"result": s.Boolean(false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_156_SimpleStringOnce_Boolean_YES(t *testing.T) {
	safeTest(t, "Test_156_SimpleStringOnce_Boolean_YES", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("YES")

		// Act
		actual := args.Map{"result": s.Boolean(false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_157_SimpleStringOnce_Boolean_CapY(t *testing.T) {
	safeTest(t, "Test_157_SimpleStringOnce_Boolean_CapY", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("Y")

		// Act
		actual := args.Map{"result": s.Boolean(false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_158_SimpleStringOnce_Boolean_Bad(t *testing.T) {
	safeTest(t, "Test_158_SimpleStringOnce_Boolean_Bad", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("xyz")

		// Act
		actual := args.Map{"result": s.Boolean(false)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_159_SimpleStringOnce_Boolean_ConsiderInit_Uninit(t *testing.T) {
	safeTest(t, "Test_159_SimpleStringOnce_Boolean_ConsiderInit_Uninit", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": s.Boolean(true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_160_SimpleStringOnce_IsSetter(t *testing.T) {
	safeTest(t, "Test_160_SimpleStringOnce_IsSetter", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("yes")
		sv := s.IsSetter(false)

		// Act
		actual := args.Map{"result": sv.IsTrue()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_161_SimpleStringOnce_IsSetter_False(t *testing.T) {
	safeTest(t, "Test_161_SimpleStringOnce_IsSetter_False", func() {
		// Arrange
		s := &corestr.SimpleStringOnce{}
		sv := s.IsSetter(true)

		// Act
		actual := args.Map{"result": sv.IsTrue()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_162_SimpleStringOnce_IsSetter_Bad(t *testing.T) {
	safeTest(t, "Test_162_SimpleStringOnce_IsSetter_Bad", func() {
		s := corestr.New.SimpleStringOnce.Init("xyz")
		sv := s.IsSetter(false)
		_ = sv
	})
}

func Test_163_SimpleStringOnce_ValueInt(t *testing.T) {
	safeTest(t, "Test_163_SimpleStringOnce_ValueInt", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("42")

		// Act
		actual := args.Map{"result": s.ValueInt(0) != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
	})
}

func Test_164_SimpleStringOnce_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_164_SimpleStringOnce_ValueDefInt", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("10")

		// Act
		actual := args.Map{"result": s.ValueDefInt() != 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}

func Test_165_SimpleStringOnce_ValueByte(t *testing.T) {
	safeTest(t, "Test_165_SimpleStringOnce_ValueByte", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("200")

		// Act
		actual := args.Map{"result": s.ValueByte(0) != 200}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200", actual)
	})
}

func Test_166_SimpleStringOnce_ValueByte_Overflow(t *testing.T) {
	safeTest(t, "Test_166_SimpleStringOnce_ValueByte_Overflow", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("300")

		// Act
		actual := args.Map{"result": s.ValueByte(99) != 99}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 99", actual)
	})
}

func Test_167_SimpleStringOnce_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_167_SimpleStringOnce_ValueDefByte", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("100")

		// Act
		actual := args.Map{"result": s.ValueDefByte() != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
	})
}

func Test_168_SimpleStringOnce_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_168_SimpleStringOnce_ValueFloat64", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("3.14")

		// Act
		actual := args.Map{"result": s.ValueFloat64(0) != 3.14}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
	})
}

func Test_169_SimpleStringOnce_ValueFloat64_Error(t *testing.T) {
	safeTest(t, "Test_169_SimpleStringOnce_ValueFloat64_Error", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		actual := args.Map{"result": s.ValueFloat64(1.0) != 1.0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1.0", actual)
	})
}

func Test_170_SimpleStringOnce_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_170_SimpleStringOnce_ValueDefFloat64", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("2.5")

		// Act
		actual := args.Map{"result": s.ValueDefFloat64() != 2.5}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_171_SimpleStringOnce_NonPtr(t *testing.T) {
	safeTest(t, "Test_171_SimpleStringOnce_NonPtr", func() {
		s := corestr.New.SimpleStringOnce.Init("x")
		np := s.NonPtr()
		_ = np
	})
}

func Test_172_SimpleStringOnce_Ptr(t *testing.T) {
	safeTest(t, "Test_172_SimpleStringOnce_Ptr", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("x")

		// Act
		actual := args.Map{"result": s.Ptr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same", actual)
	})
}

func Test_173_SimpleStringOnce_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_173_SimpleStringOnce_HasSafeNonEmpty", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("x")

		// Act
		actual := args.Map{"result": s.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_174_SimpleStringOnce_Is(t *testing.T) {
	safeTest(t, "Test_174_SimpleStringOnce_Is", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": s.Is("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_175_SimpleStringOnce_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_175_SimpleStringOnce_IsAnyOf", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("b")

		// Act
		actual := args.Map{"result": s.IsAnyOf("a", "b", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": s.IsAnyOf()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
		actual = args.Map{"result": s.IsAnyOf("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_176_SimpleStringOnce_IsContains(t *testing.T) {
	safeTest(t, "Test_176_SimpleStringOnce_IsContains", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello world")

		// Act
		actual := args.Map{"result": s.IsContains("world")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_177_SimpleStringOnce_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_177_SimpleStringOnce_IsAnyContains", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello world")

		// Act
		actual := args.Map{"result": s.IsAnyContains("xyz", "world")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": s.IsAnyContains()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
	})
}

func Test_178_SimpleStringOnce_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_178_SimpleStringOnce_IsEqualNonSensitive", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("Hello")

		// Act
		actual := args.Map{"result": s.IsEqualNonSensitive("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_179_SimpleStringOnce_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_179_SimpleStringOnce_IsRegexMatches", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": s.IsRegexMatches(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": s.IsRegexMatches(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_180_SimpleStringOnce_RegexFindString(t *testing.T) {
	safeTest(t, "Test_180_SimpleStringOnce_RegexFindString", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": s.RegexFindString(re) != "123"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 123", actual)
		actual = args.Map{"result": s.RegexFindString(nil) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_181_SimpleStringOnce_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_181_SimpleStringOnce_RegexFindAllStringsWithFlag", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, has := s.RegexFindAllStringsWithFlag(re, -1)

		// Act
		actual := args.Map{"result": has || len(items) != 3}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		_, has2 := s.RegexFindAllStringsWithFlag(nil, -1)
		actual = args.Map{"result": has2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_182_SimpleStringOnce_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_182_SimpleStringOnce_RegexFindAllStrings", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("a1b2")
		re := regexp.MustCompile(`\d`)
		items := s.RegexFindAllStrings(re, -1)

		// Act
		actual := args.Map{"result": len(items) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		items2 := s.RegexFindAllStrings(nil, -1)
		actual = args.Map{"result": len(items2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_183_SimpleStringOnce_LinesSimpleSlice(t *testing.T) {
	safeTest(t, "Test_183_SimpleStringOnce_LinesSimpleSlice", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("a\nb")
		sl := s.LinesSimpleSlice()

		// Act
		actual := args.Map{"result": sl.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_184_SimpleStringOnce_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_184_SimpleStringOnce_SimpleSlice", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("a,b,c")
		sl := s.SimpleSlice(",")

		// Act
		actual := args.Map{"result": sl.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_185_SimpleStringOnce_Split(t *testing.T) {
	safeTest(t, "Test_185_SimpleStringOnce_Split", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("a,b")
		sp := s.Split(",")

		// Act
		actual := args.Map{"result": len(sp) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_186_SimpleStringOnce_SplitLeftRight(t *testing.T) {
	safeTest(t, "Test_186_SimpleStringOnce_SplitLeftRight", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("key=value")
		left, right := s.SplitLeftRight("=")

		// Act
		actual := args.Map{"result": left != "key" || right != "value"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected key, value", actual)
	})
}

func Test_187_SimpleStringOnce_SplitLeftRight_NoSep(t *testing.T) {
	safeTest(t, "Test_187_SimpleStringOnce_SplitLeftRight_NoSep", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("nosep")
		left, right := s.SplitLeftRight("=")

		// Act
		actual := args.Map{"result": left != "nosep" || right != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nosep, empty", actual)
	})
}

func Test_188_SimpleStringOnce_SplitLeftRightTrim(t *testing.T) {
	safeTest(t, "Test_188_SimpleStringOnce_SplitLeftRightTrim", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init(" key = value ")
		left, right := s.SplitLeftRightTrim("=")

		// Act
		actual := args.Map{"result": left != "key" || right != "value"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed key, value", actual)
	})
}

func Test_189_SimpleStringOnce_SplitLeftRightTrim_NoSep(t *testing.T) {
	safeTest(t, "Test_189_SimpleStringOnce_SplitLeftRightTrim_NoSep", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init(" nosep ")
		left, right := s.SplitLeftRightTrim("=")

		// Act
		actual := args.Map{"result": left != "nosep" || right != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed nosep, empty", actual)
	})
}

func Test_190_SimpleStringOnce_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_190_SimpleStringOnce_SplitNonEmpty", func() {
		s := corestr.New.SimpleStringOnce.Init("a,,b")
		sp := s.SplitNonEmpty(",")
		_ = sp
	})
}

func Test_191_SimpleStringOnce_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_191_SimpleStringOnce_SplitTrimNonWhitespace", func() {
		s := corestr.New.SimpleStringOnce.Init("a , b , c")
		sp := s.SplitTrimNonWhitespace(",")
		_ = sp
	})
}

func Test_192_SimpleStringOnce_ClonePtr(t *testing.T) {
	safeTest(t, "Test_192_SimpleStringOnce_ClonePtr", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")
		c := s.ClonePtr()

		// Act
		actual := args.Map{"result": c.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_193_SimpleStringOnce_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_193_SimpleStringOnce_ClonePtr_Nil", func() {
		// Arrange
		var s *corestr.SimpleStringOnce

		// Act
		actual := args.Map{"result": s.ClonePtr() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_194_SimpleStringOnce_Clone(t *testing.T) {
	safeTest(t, "Test_194_SimpleStringOnce_Clone", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")
		c := s.Clone()

		// Act
		actual := args.Map{"result": c.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_195_SimpleStringOnce_CloneUsingNewVal(t *testing.T) {
	safeTest(t, "Test_195_SimpleStringOnce_CloneUsingNewVal", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")
		c := s.CloneUsingNewVal("world")

		// Act
		actual := args.Map{"result": c.Value() != "world"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected world", actual)
	})
}

func Test_196_SimpleStringOnce_Dispose(t *testing.T) {
	safeTest(t, "Test_196_SimpleStringOnce_Dispose", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		s.Dispose()
	})
}

func Test_197_SimpleStringOnce_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_197_SimpleStringOnce_Dispose_Nil", func() {
		var s *corestr.SimpleStringOnce
		s.Dispose()
	})
}

func Test_198_SimpleStringOnce_String(t *testing.T) {
	safeTest(t, "Test_198_SimpleStringOnce_String", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hi")

		// Act
		actual := args.Map{"result": s.String() != "hi"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hi", actual)
	})
}

func Test_199_SimpleStringOnce_String_Nil(t *testing.T) {
	safeTest(t, "Test_199_SimpleStringOnce_String_Nil", func() {
		// Arrange
		var s *corestr.SimpleStringOnce

		// Act
		actual := args.Map{"result": s.String() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_200_SimpleStringOnce_StringPtr(t *testing.T) {
	safeTest(t, "Test_200_SimpleStringOnce_StringPtr", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hi")

		// Act
		actual := args.Map{"result": *s.StringPtr() != "hi"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hi", actual)
	})
}
