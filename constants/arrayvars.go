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

package constants

//goland:noinspection ALL
var (
	// Copied from golang strings
	AsciiSpace = [256]uint8{
		TabByte:            One,
		LineFeedUnixByte:   One,
		TabVByte:           One,
		FormFeedByte:       One,
		CarriageReturnByte: One,
		SpaceByte:          One,
		0x85:               One, // reference : https://bit.ly/2JWdIoj
		0xA0:               One, // reference : https://bit.ly/2JWdIoj
	}

	// FormFeed \f is also marked as newline here.
	AsciiNewLinesChars = [256]uint8{
		LineFeedUnix:   One,
		FormFeed:       One,
		CarriageReturn: One,
	}

	SpecialChars = [256]uint8{
		'!': One,
		'@': One,
		'#': One,
		'$': One,
		'%': One,
		'^': One,
		'&': One,
		'*': One,
		'(': One,
		')': One,
	}

	BracketChars = [256]uint8{
		'[': One,
		']': One,
		'{': One,
		'}': One,
		'(': One,
		')': One,
		'<': One,
		'>': One,
	}

	EmptyStrings       []string
	EmptyPtrStrings    []*string
	EmptyInts          []int
	EmptyBytes         []byte
	EmptyFloats        []float32
	EmptyFloat64s      []float64
	EmptyInterfaces    []any
	EmptyIntToIntsMap  map[int][]int
	EmptyIntToBytesMap map[int][]byte
	EmptyStringMap     map[string]string
	EmptyStrToIntsMap  map[string][]int
	EmptyStrToBytesMap map[string][]byte
	EmptyStringsMap    map[string][]string

	EmptyIntToPtrIntsMap  map[int]*[]int
	EmptyIntToPtrBytesMap map[int]*[]byte
	EmptyStrToPtrIntsMap  map[string]*[]int
	EmptyStrToPtrBytesMap map[string]*[]byte
	EmptyPtrStringsMap    map[string]*[]string
)
