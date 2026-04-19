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

package casenilsafetests

// sampleStruct is a test-only struct used to validate CaseNilSafe
// across all edge cases: pointer receivers, value receivers, void
// methods, multi-return, and methods with complex args.
type sampleStruct struct {
	Name  string
	Value int
}

// IsValid — pointer receiver, nil-safe, returns bool.
func (it *sampleStruct) IsValid() bool {
	if it == nil {
		return false
	}

	return it.Name != ""
}

// Length — pointer receiver, nil-safe, returns int.
func (it *sampleStruct) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Name)
}

// String — pointer receiver, nil-safe, returns string.
func (it *sampleStruct) String() string {
	if it == nil {
		return ""
	}

	return it.Name
}

// Reset — pointer receiver, nil-safe, void method.
func (it *sampleStruct) Reset() {
	if it == nil {
		return
	}

	it.Name = ""
	it.Value = 0
}

// Parse — pointer receiver, nil-safe, returns (int, error).
func (it *sampleStruct) Parse(input string) (int, error) {
	if it == nil {
		return 0, nil
	}

	return len(input), nil
}

// Lookup — pointer receiver, nil-safe, returns (string, bool).
func (it *sampleStruct) Lookup(key string) (string, bool) {
	if it == nil {
		return "", false
	}

	if key == it.Name {
		return it.Name, true
	}

	return "", false
}

// UnsafeMethod — pointer receiver, NOT nil-safe (panics on nil).
func (it *sampleStruct) UnsafeMethod() string {
	return it.Name // will panic if it == nil
}

// ValueString — value receiver, always panics on nil pointer.
func (it sampleStruct) ValueString() string {
	return it.Name
}
