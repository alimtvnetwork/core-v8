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

package isanytests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// DeepEqual / NotDeepEqual
// ==========================================

var deepEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "DeepEqual true for identical primitives",
		ArrangeInput: args.Map{
			"when": "given same int values",
		},
		ExpectedInput: args.Map{
			"isDeepEqual":    "true",
			"isNotDeepEqual": "false",
		},
	},
	{
		Title: "DeepEqual false for different primitives",
		ArrangeInput: args.Map{
			"when": "given different int values",
		},
		ExpectedInput: args.Map{
			"isDeepEqual":    "false",
			"isNotDeepEqual": "true",
		},
	},
	{
		Title: "DeepEqual true for identical slices",
		ArrangeInput: args.Map{
			"when": "given same string slices",
		},
		ExpectedInput: args.Map{
			"isDeepEqual":    "true",
			"isNotDeepEqual": "false",
		},
	},
	{
		Title: "DeepEqual false for different slices",
		ArrangeInput: args.Map{
			"when": "given different string slices",
		},
		ExpectedInput: args.Map{
			"isDeepEqual":    "false",
			"isNotDeepEqual": "true",
		},
	},
	{
		Title: "DeepEqual true for both nil",
		ArrangeInput: args.Map{
			"when": "given both nil",
		},
		ExpectedInput: args.Map{
			"isDeepEqual":    "true",
			"isNotDeepEqual": "false",
		},
	},
}

// ==========================================
// Zero
// ==========================================

var zeroTestCases = []coretestcases.CaseV1{
	{
		Title: "Zero returns true for zero values and false for non-zero",
		ArrangeInput: args.Map{
			"when": "given various zero and non-zero values",
		},
		ExpectedInput: args.Map{
			"intZero":   "true",
			"int42":     "false",
			"emptyStr":  "true",
			"helloStr":  "false",
			"boolFalse": "true",
		},
	},
}

// ==========================================
// ReflectNull vs Null comparison
// ==========================================

var reflectNullTestCases = []coretestcases.CaseV1{
	{
		Title: "ReflectNull returns true for nil pointer, false for non-nil",
		ArrangeInput: args.Map{
			"when": "given nil and non-nil pointers",
		},
		ExpectedInput: args.Map{
			"nilPtr":    "true",
			"nonNilPtr": "false",
			"nilSlice":  "true",
		},
	},
}

// ==========================================
// NotNull
// ==========================================

var notNullTestCases = []coretestcases.CaseV1{
	{
		Title: "NotNull is inverse of Null",
		ArrangeInput: args.Map{
			"when": "given nil and non-nil values",
		},
		ExpectedInput: args.Map{
			"notNullNil":    "false",
			"notNull42":     "true",
			"inverseEquals": "false",
		},
	},
}

// ==========================================
// StringEqual
// ==========================================

var stringEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "StringEqual compares string representation of values",
		ArrangeInput: args.Map{
			"when": "given values with same/different string representation",
		},
		ExpectedInput: args.Map{
			"sameStrings": "true",
			"diffStrings": "false",
		},
	},
}

// ==========================================
// Pointer
// ==========================================

var pointerTestCases = []coretestcases.CaseV1{
	{
		Title: "Pointer returns true for pointers, false for values",
		ArrangeInput: args.Map{
			"when": "given pointer and non-pointer values",
		},
		ExpectedInput: args.Map{
			"intPtr":    "true",
			"intVal":    "false",
			"stringPtr": "true",
		},
	},
}
