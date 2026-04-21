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
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var ifStringTestCases = []coretestcases.CaseV1{
	{
		Title: "If true returns trueValue string",
		ArrangeInput: args.Map{
			"when":       "given true condition",
			"isTrue":     true,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: "yes",
	},
	{
		Title: "If false returns falseValue string",
		ArrangeInput: args.Map{
			"when":       "given false condition",
			"isTrue":     false,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: "no",
	},
	{
		Title: "If true with empty strings returns empty trueValue",
		ArrangeInput: args.Map{
			"when":       "given true with empty trueValue",
			"isTrue":     true,
			"trueValue":  "",
			"falseValue": "fallback",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "If false with empty falseValue returns empty",
		ArrangeInput: args.Map{
			"when":       "given false with empty falseValue",
			"isTrue":     false,
			"trueValue":  "something",
			"falseValue": "",
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "If true with same values returns that value",
		ArrangeInput: args.Map{
			"when":       "given true with identical values",
			"isTrue":     true,
			"trueValue":  "same",
			"falseValue": "same",
		},
		ExpectedInput: "same",
	},
}

var ifIntTestCases = []coretestcases.CaseV1{
	{
		Title: "If true returns trueValue int",
		ArrangeInput: args.Map{
			"when":       "given true condition",
			"isTrue":     true,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: "10",
	},
	{
		Title: "If false returns falseValue int",
		ArrangeInput: args.Map{
			"when":       "given false condition",
			"isTrue":     false,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: "20",
	},
	{
		Title: "If true with zero returns zero",
		ArrangeInput: args.Map{
			"when":       "given true with zero trueValue",
			"isTrue":     true,
			"trueValue":  0,
			"falseValue": 99,
		},
		ExpectedInput: "0",
	},
	{
		Title: "If false with negative returns negative",
		ArrangeInput: args.Map{
			"when":       "given false with negative falseValue",
			"isTrue":     false,
			"trueValue":  100,
			"falseValue": -42,
		},
		ExpectedInput: "-42",
	},
}

var nilDefTestCases = []coretestcases.CaseV1{
	{
		Title: "NilDef with nil pointer returns default",
		ArrangeInput: args.Map{
			"when":   "given nil pointer",
			"isNil":  true,
			"defVal": "default",
		},
		ExpectedInput: "default",
	},
	{
		Title: "NilDef with non-nil pointer returns value",
		ArrangeInput: args.Map{
			"when":   "given non-nil pointer",
			"isNil":  false,
			"value":  "actual",
			"defVal": "default",
		},
		ExpectedInput: "actual",
	},
	{
		Title: "NilDef with empty string pointer returns empty",
		ArrangeInput: args.Map{
			"when":   "given non-nil pointer to empty string",
			"isNil":  false,
			"value":  "",
			"defVal": "fallback",
		},
		ExpectedInput: []string{""},
	},
}

var ifFuncStringTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFunc true evaluates trueFunc only",
		ArrangeInput: args.Map{
			"when":       "given true condition with funcs",
			"isTrue":     true,
			"trueValue":  "from-true-func",
			"falseValue": "from-false-func",
		},
		ExpectedInput: "from-true-func",
	},
	{
		Title: "IfFunc false evaluates falseFunc only",
		ArrangeInput: args.Map{
			"when":       "given false condition with funcs",
			"isTrue":     false,
			"trueValue":  "from-true-func",
			"falseValue": "from-false-func",
		},
		ExpectedInput: "from-false-func",
	},
}

var ifTrueFuncStringTestCases = []coretestcases.CaseV1{
	{
		Title: "IfTrueFunc true returns trueFunc result",
		ArrangeInput: args.Map{
			"when":      "given true condition",
			"isTrue":    true,
			"trueValue": "computed",
		},
		ExpectedInput: "computed",
	},
	{
		Title: "IfTrueFunc false returns zero value",
		ArrangeInput: args.Map{
			"when":      "given false condition",
			"isTrue":    false,
			"trueValue": "computed",
		},
		ExpectedInput: []string{""},
	},
}

var ifSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSlice true returns trueSlice",
		ArrangeInput: args.Map{
			"when":       "given true condition",
			"isTrue":     true,
			"trueValue":  []string{"a", "b"},
			"falseValue": []string{"x", "y"},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "a",
		},
	},
	{
		Title: "IfSlice false returns falseSlice",
		ArrangeInput: args.Map{
			"when":       "given false condition",
			"isTrue":     false,
			"trueValue":  []string{"a", "b"},
			"falseValue": []string{"x", "y", "z"},
		},
		ExpectedInput: args.Map{
			"length": "3",
			"first":  "x",
		},
	},
}

var nilCheckTestCases = []coretestcases.CaseV1{
	{
		Title: "NilCheck returns onNil when input is nil",
		ArrangeInput: args.Map{
			"when":     "given nil input",
			"isNil":    true,
			"onNil":    "nil-result",
			"onNonNil": "non-nil-result",
		},
		ExpectedInput: "nil-result",
	},
	{
		Title: "NilCheck returns onNonNil when input is not nil",
		ArrangeInput: args.Map{
			"when":     "given non-nil input",
			"isNil":    false,
			"onNil":    "nil-result",
			"onNonNil": "non-nil-result",
		},
		ExpectedInput: "non-nil-result",
	},
}

var defOnNilTestCases = []coretestcases.CaseV1{
	{
		Title: "DefOnNil returns default when input is nil",
		ArrangeInput: args.Map{
			"when":     "given nil input",
			"isNil":    true,
			"onNonNil": "default-val",
		},
		ExpectedInput: "default-val",
	},
	{
		Title: "DefOnNil returns input when input is not nil",
		ArrangeInput: args.Map{
			"when":     "given non-nil input",
			"isNil":    false,
			"value":    "actual-val",
			"onNonNil": "default-val",
		},
		ExpectedInput: "actual-val",
	},
}
