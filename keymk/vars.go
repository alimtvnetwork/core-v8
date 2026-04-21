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

package keymk

import "github.com/alimtvnetwork/core-v8/constants"

var (
	PipeJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.Pipe,
		EndBracket:       constants.Pipe,
	}

	PipeCurlyJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.Pipe + constants.CurlyStart,
		EndBracket:       constants.CurlyEnd + constants.Pipe,
	}

	PipeSquareJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.Pipe + constants.SquareStart,
		EndBracket:       constants.SquareEnd + constants.Pipe,
	}

	CurlyBracePathJoinerOption = &Option{
		Joiner:           constants.PathSeparator,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.CurlyStart,
		EndBracket:       constants.CurlyEnd,
	}

	CurlyBraceJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.CurlyStart,
		EndBracket:       constants.CurlyEnd,
	}

	BracketJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.SquareStart,
		EndBracket:       constants.SquareEnd,
	}

	ParenthesisJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.ParenthesisStart,
		EndBracket:       constants.ParenthesisEnd,
	}

	JoinerOption = &Option{
		Joiner:           constants.Hyphen,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    false,
		StartBracket:     constants.EmptyString,
		EndBracket:       constants.EmptyString,
	}

	FullLegends = LegendName{
		Root:    "root",
		Package: "package",
		Group:   "group",
		User:    "user",
		Item:    "item",
	}

	FullCategoryLegends = LegendName{
		Root:    "root",
		Package: "package",
		Group:   "category",
		User:    "user",
		Item:    "item",
	}

	FullEventLegends = LegendName{
		Root:    "root",
		Package: "package",
		Group:   "event",
		User:    "user",
		Item:    "item",
	}

	ShortLegends = LegendName{
		Root:    "r",
		Package: "p",
		Group:   "g",
		User:    "u",
		Item:    "i",
	}

	ShortEventLegends = LegendName{
		Root:    "r",
		Group:   "e",
		Package: "p",
		Item:    "i",
	}

	NewKey           = &newKeyCreator{}
	NewKeyWithLegend = &newKeyWithLegendCreator{}
	FixedLegend      = fixedLegend{}
)
