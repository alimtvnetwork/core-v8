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
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
//  — corestr remaining gaps (Iteration 28)
// Split from monolithic file for build isolation.
//
// API Reference (verified from source):
//   - corestr.New.CharCollectionMap.CapSelfCap(cap, selfCap)
//   - corestr.New.CharHashsetMap.Cap(n, n)
//   - corestr.New.LinkedCollection.Create()   (singular, NOT LinkedCollections)
//   - corestr.New.Collection.Strings([]string{...})  (takes slice, NOT variadic)
//   - lc.IsEqualsPtr(other)                   (NOT IsChainEqual)
//   - chm.AddStrings()                        (NOT AddAll)
// ══════════════════════════════════════════════════════════════════════════════

// ---------- CharCollectionMap: AddHashmapsValues nil ----------

func Test_CharCollectionMap_AddHashmapsValues_Nil(t *testing.T) {
	safeTest(t, "Test_I28_CharCollectionMap_AddHashmapsValues_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)

		// Act
		result := ccm.AddHashmapsValues(nil)

		// Assert
		actual := args.Map{"isNil": result == nil}
		expected := args.Map{"isNil": false}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues returns self -- nil input", actual)
	})
}

// ---------- CharCollectionMap: AddHashmapsKeysOrValuesBothUsingFilter nil ----------

func Test_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_I28_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)

		// Act
		result := ccm.AddHashmapsKeysOrValuesBothUsingFilter(nil, nil)

		// Assert
		actual := args.Map{"isNil": result == nil}
		expected := args.Map{"isNil": false}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysOrValuesBothUsingFilter returns self -- nil input", actual)
	})
}

// ---------- CharCollectionMap: AddHashmapsKeysValuesBoth nil ----------

func Test_CharCollectionMap_AddHashmapsKeysValuesBoth_Nil(t *testing.T) {
	safeTest(t, "Test_I28_CharCollectionMap_AddHashmapsKeysValuesBoth_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)

		// Act
		result := ccm.AddHashmapsKeysValuesBoth(nil)

		// Assert
		actual := args.Map{"isNil": result == nil}
		expected := args.Map{"isNil": false}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesBoth returns self -- nil input", actual)
	})
}

// ---------- CharHashsetMap: AddLock triggers nil items init ----------

func Test_CharHashsetMap_AddLock_NilItemsInit(t *testing.T) {
	safeTest(t, "Test_I28_CharHashsetMap_AddLock_NilItemsInit", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)

		// Act
		chm.AddLock("apple")

		// Assert
		actual := args.Map{"has": chm.Has("apple")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddLock initializes items -- nil items", actual)
	})
}

// ---------- CharHashsetMap: Add triggers nil items init ----------

func Test_CharHashsetMap_Add_NilItemsInit(t *testing.T) {
	safeTest(t, "Test_I28_CharHashsetMap_Add_NilItemsInit", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)

		// Act
		chm.Add("banana")

		// Assert
		actual := args.Map{"has": chm.Has("banana")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Add initializes items -- nil items", actual)
	})
}

// ---------- CharHashsetMap: AddStrings empty/nil ----------

func Test_CharHashsetMap_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_I28_CharHashsetMap_AddStrings_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4, 4)

		// Act — AddStrings (NOT AddAll, which doesn't exist)
		result := chm.AddStrings()

		// Assert
		actual := args.Map{"length": result.Length()}
		expected := args.Map{"length": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings returns empty -- nil input", actual)
	})
}

// ---------- SimpleSlice: IsEqualUnorderedLines nil receiver ----------

func Test_SimpleSlice_IsEqualUnorderedLines_NilReceiver(t *testing.T) {
	safeTest(t, "Test_I28_SimpleSlice_IsEqualUnorderedLines_NilReceiver", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		result := ss.IsEqualUnorderedLines([]string{"a"})

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": false}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLines returns false -- nil receiver", actual)
	})
}

// ---------- SimpleSlice: IsEqualUnorderedLinesClone nil receiver ----------

func Test_SimpleSlice_IsEqualUnorderedLinesClone_NilReceiver(t *testing.T) {
	safeTest(t, "Test_I28_SimpleSlice_IsEqualUnorderedLinesClone_NilReceiver", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		result := ss.IsEqualUnorderedLinesClone([]string{"a"})

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": false}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone returns false -- nil receiver", actual)
	})
}

// ---------- SimpleSlice: IsEqualByFuncLinesSplit both empty ----------

func Test_SimpleSlice_IsEqualByFuncLinesSplit_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I28_SimpleSlice_IsEqualByFuncLinesSplit_BothEmpty", func() {
		// Arrange
		ss := corestr.SimpleSlice{}

		// Act
		result := ss.IsEqualByFuncLinesSplit(
			false,
			",",
			"",
			func(index int, left, right string) bool {
				return left == right
			},
		)

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": false}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit returns false -- empty vs single empty string", actual)
	})
}

// ---------- ValidValue: ParseInjectUsingJson error path ----------

func Test_ValidValue_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I28_ValidValue_ParseInjectUsingJson_Error", func() {
		// Arrange
		vv := &corestr.ValidValue{}
		badJson := corejson.NewPtr([]byte("not-valid-json{{{"))

		// Act
		result, err := vv.ParseInjectUsingJson(badJson)

		// Assert
		actual := args.Map{
			"resultNil": result == nil,
			"hasErr":    err != nil,
		}
		expected := args.Map{
			"resultNil": true,
			"hasErr":    true,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns nil,err -- invalid json", actual)
	})
}
