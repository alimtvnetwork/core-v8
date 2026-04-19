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

package coredynamictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// MapAnyItems — Add and AllKeys
// ==========================================

var mapAnyItemsAddAndKeysTestCase = coretestcases.CaseV1{
	Title: "Add returns stored items and AllKeys returns keys -- 3 items added",
	ArrangeInput: args.Map{
		"when":     "given 3 items added",
		"capacity": 10,
		"keys":     []string{"key1", "key2", "key3"},
	},
	ExpectedInput: args.Map{
		"keyCount": 3,
		"hasAll":   true,
	},
}

// ==========================================
// MapAnyItems — GetPagedCollection
// ==========================================

var mapAnyItemsPagedTestCase = coretestcases.CaseV1{
	Title: "GetPagedCollection returns 5 pages -- 9 items paged by 2",
	ArrangeInput: args.Map{
		"when":      "given 9 items paged by 2",
		"itemCount": 9,
		"pageSize":  2,
	},
	ExpectedInput: args.Map{
		"pageCount": 5,
	},
}

// ==========================================
// MapAnyItems — JSON roundtrip
// ==========================================

var mapAnyItemsJsonRoundtripTestCase = coretestcases.CaseV1{
	Title: "JSON serialize then deserialize returns equal map -- 4 items",
	ArrangeInput: args.Map{
		"when":      "given map serialized and deserialized",
		"itemCount": 4,
	},
	ExpectedInput: args.Map{
		"isEqual": true,
	},
}

// ==========================================
// MapAnyItems — GetItemRef
// ==========================================

var mapAnyItemsGetItemRefTestCase = coretestcases.CaseV1{
	Title: "GetItemRef returns stored value -- existing key 'target-key'",
	ArrangeInput: args.Map{
		"when": "given key exists in map",
		"key":  "target-key",
	},
	ExpectedInput: args.Map{
		"hasItems": true,
	},
}
