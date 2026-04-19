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

package converterstests

type stringToIntegerWithDefaultCase struct {
	name        string
	input       string
	defaultVal  int
	expectedVal int
	expectedOk  bool
}

var stringToIntegerWithDefaultCases = []stringToIntegerWithDefaultCase{
	{"valid integer", "42", 0, 42, true},
	{"empty string", "", 99, 99, false},
	{"non-numeric", "abc", 99, 99, false},
	{"zero", "0", 5, 0, true},
}

type stringToByteCase struct {
	name      string
	input     string
	expected  byte
	expectErr bool
}

var stringToByteCases = []stringToByteCase{
	{"zero", "0", 0, false},
	{"one", "1", 1, false},
	{"valid", "100", 100, false},
	{"max byte", "255", 255, false},
	{"over 255", "256", 0, true},
	{"negative", "-1", 0, true},
	{"empty", "", 0, true},
	{"non-numeric", "abc", 0, true},
}
