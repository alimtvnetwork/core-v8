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

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
//  — corestr remaining ~42 lines (mostly nil-receiver dead code)
// ══════════════════════════════════════════════════════════════════════════════

// ── CharCollectionMap nil-receiver guards ──
// Lines 45, 369, 567, 589, 622, 868, 985, 1105 — documented as dead code

// ── CharHashsetMap nil-receiver guards ──
// Lines 593, 624, 657, 661, 713, 856, 906, 991, 1050 — documented as dead code

// ── CharHashsetMap.efficientAddOfLargeItems (line 748-772) ──

func Test_CharHashsetMap_EfficientAdd_I29(t *testing.T) {
	// Arrange
	chm := corestr.New.CharHashsetMap.Cap(3, 3)

	// Act — add large strings to trigger efficient add path
	chm.AddStrings("alpha", "beta", "gamma", "delta", "epsilon")

	// Assert
	actual := args.Map{
		"hasItems": chm.Length() > 0,
	}
	expected := args.Map{
		"hasItems": true,
	}
	actual.ShouldBeEqual(t, 1, "CharHashsetMap efficient add", expected)
}

// ── Collection.JsonString error path (line 97) ──
// json.Marshal on []string won't fail — dead code

// ── Collection JSON operations (lines 497, 508, 528, 539, 559, 570, 581, 592) ──
// Various nil/error return paths in JSON — mostly dead code on valid data

func Test_Collection_JsonString_I29(t *testing.T) {
	// Arrange
	coll := corestr.New.Collection.Strings([]string{"a", "b", "c"})

	// Act
	jsonStr := coll.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
	}
	expected := args.Map{
		"hasContent": true,
	}
	actual.ShouldBeEqual(t, 1, "Collection JsonString valid", expected)
}

// ── CollectionsOfCollection nil guards (lines 45, 68) ──
// Nil receiver guard — dead code

// ── Hashmap nil guard (line 158) — dead code ──

// ── LinkedCollectionNode nil guards (lines 93, 152, 156, 177, 256) — dead code ──

// ── LinkedCollections various paths ──

func Test_LinkedCollections_SafePointerIndexAt_OutOfRange_I29(t *testing.T) {
	// Arrange
	lc := corestr.New.LinkedCollection.Create()
	lc.Add(corestr.New.Collection.Strings([]string{"x"}))

	// Act
	result := lc.SafePointerIndexAt(99)

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	actual.ShouldBeEqual(t, 1, "LinkedCollections SafePointerIndexAt out of range", expected)
}

// ── LinkedCollections.ToCollection (line 1265) ──

func Test_LinkedCollections_ToCollection_I29(t *testing.T) {
	// Arrange
	lc := corestr.New.LinkedCollection.Create()
	lc.Add(corestr.New.Collection.Strings([]string{"v1", "v2"}))
	lc.Add(corestr.New.Collection.Strings([]string{"v3"}))

	// Act
	merged := lc.ToCollection(0)

	// Assert
	actual := args.Map{"length": merged.Length()}
	expected := args.Map{"length": 3}
	actual.ShouldBeEqual(t, 1, "LinkedCollections ToCollection", expected)
}

// ── LinkedCollections: remaining dead-code paths ──
// Lines 102, 147, 151, 760, 943, 1143, 1182, 1185, 1248 are nil-receiver
// guards or unreachable fallback returns. Documented as accepted dead code.
