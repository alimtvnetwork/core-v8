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

package chmodhelper

// AttrVariant
//
// 1 - Execute true
// 2 - Write true
// 3 - Write + Execute true
// 4 - Read true
// 5 - Read + Execute true
// 6 - Read + Write true
// 7 - Read + Write + Execute all true
type AttrVariant byte

//goland:noinspection ALL
const (
	Execute          AttrVariant = 1
	Write            AttrVariant = 2
	WriteExecute     AttrVariant = 3
	Read             AttrVariant = 4
	ReadExecute      AttrVariant = 5
	ReadWrite        AttrVariant = 6
	ReadWriteExecute AttrVariant = 7
)

// IsGreaterThan v > byte(attrVariant)
func (attrVariant AttrVariant) IsGreaterThan(v byte) bool {
	return v > byte(attrVariant)
}

func (attrVariant AttrVariant) String() string {
	return string(attrVariant)
}

func (attrVariant AttrVariant) Value() byte {
	return byte(attrVariant)
}

func (attrVariant AttrVariant) ToAttribute() Attribute {
	return New.Attribute.UsingVariantMust(attrVariant)
}
