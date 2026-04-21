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

package coretests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alimtvnetwork/core-v8/corecmp"
)

// DraftType
//
// Draft type is used for dummy values and reflection testing.
//
// Can be compared and do the verification as well.
type DraftType struct {
	SampleString1, SampleString2 string
	SampleInteger                int
	Lines                        []string
	RawBytes                     []byte
	f1String                     string
	f2Integer                    int
}

func (it *DraftType) SetF2Integer(f2Integer int) {
	it.f2Integer = f2Integer
}

func (it DraftType) F1String() string {
	return it.f1String
}

func (it DraftType) NonPtr() DraftType {
	return it
}

func (it *DraftType) PtrOrNonPtr(isPtr bool) any {
	if it == nil {
		return nil
	}

	if isPtr {
		return it
	}

	return *it
}

func (it *DraftType) VerifyAllNotEqualErr(
	right *DraftType,
) error {
	return it.VerifyNotEqualErr(
		true,
		right)
}

func (it *DraftType) VerifyNotEqualErr(
	isIncludingInnerFields bool,
	right *DraftType,
) error {
	msg := it.VerifyNotEqualMessage(
		isIncludingInnerFields,
		right)

	if msg == "" {
		return nil
	}

	return errors.New(msg)
}

func (it *DraftType) VerifyNotEqualExcludingInnerFieldsErr(
	right *DraftType,
) error {
	msg := it.VerifyNotEqualMessage(
		false,
		right)

	if msg == "" {
		return nil
	}

	return errors.New(msg)
}

func (it *DraftType) VerifyAllNotEqualMessage(
	right *DraftType,
) string {
	return it.VerifyNotEqualMessage(
		true,
		right,
	)
}

func (it *DraftType) VerifyNotEqualMessage(
	isIncludingInnerFields bool,
	right *DraftType,
) string {
	isEqual := it.IsEqual(
		isIncludingInnerFields,
		right)

	if isEqual {
		return ""
	}

	return fmt.Sprintf(
		notEqualComparisonMessageFormat,
		it.JsonString(), it.f1String, it.f2Integer,
		right.JsonString(), right.f1String, right.f2Integer)
}

func (it *DraftType) IsEqualAll(
	right *DraftType,
) bool {
	return it.IsEqual(true, right)
}

func (it *DraftType) IsEqual(
	isIncludingInnerFields bool,
	right *DraftType,
) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	if it == right {
		return true
	}

	if it.SampleString1 != right.SampleString1 {
		return false
	}

	if it.SampleString2 != right.SampleString2 {
		return false
	}

	if it.SampleInteger != right.SampleInteger {
		return false
	}

	if isIncludingInnerFields && it.f1String != right.f1String {
		return false
	}

	if isIncludingInnerFields && it.f2Integer != right.f2Integer {
		return false
	}

	if !bytes.Equal(it.RawBytes, right.RawBytes) {
		return false
	}

	if !corecmp.IsStringsEqual(it.Lines, right.Lines) {
		return false
	}

	return true
}

func (it DraftType) F2Integer() int {
	return it.f2Integer
}

func (it DraftType) JsonString() string {
	rawBytes, err := json.Marshal(it)

	if err != nil {
		panic(err)
	}

	return string(rawBytes)
}

func (it DraftType) JsonBytes() []byte {
	rawBytes, err := json.Marshal(it)

	if err != nil {
		panic(err)
	}

	return rawBytes
}

func (it DraftType) JsonBytesPtr() []byte {
	return it.JsonBytes()
}

func (it DraftType) Clone() DraftType {
	lines := make([]string, it.LinesLength())
	newBytes := make([]byte, it.RawBytesLength())

	copy(lines, it.Lines)
	copy(newBytes, it.RawBytes)

	return DraftType{
		SampleString1: it.SampleString1,
		SampleString2: it.SampleString2,
		SampleInteger: it.SampleInteger,
		Lines:         lines,
		RawBytes:      newBytes,
		f1String:      it.f1String,
		f2Integer:     it.f2Integer,
	}
}

func (it *DraftType) ClonePtr() *DraftType {
	if it == nil {
		return it
	}

	cloned := it.Clone()

	return &cloned
}

func (it *DraftType) RawBytesLength() int {
	return len(it.RawBytes)
}

func (it *DraftType) LinesLength() int {
	return len(it.Lines)
}
