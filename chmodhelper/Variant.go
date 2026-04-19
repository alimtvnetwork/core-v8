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

type Variant string

//goland:noinspection ALL
const (
	AllRead                             Variant = "444"
	AllWrite                            Variant = "222"
	AllExecute                          Variant = "111"
	AllReadWrite                        Variant = "666"
	AllReadExecute                      Variant = "555"
	AllWriteExecute                     Variant = "333"
	OwnerAllReadWriteGroupOther         Variant = "755"
	ReadWriteOwnerReadGroupOther        Variant = "644"
	ReadWriteOwnerReadExecuteGroupOther Variant = "655"
	X111                                Variant = "111"
	X222                                Variant = "222"
	X333                                Variant = "333"
	X444                                Variant = "444"
	X555                                Variant = "555"
	X644                                Variant = "644"
	X655                                Variant = "655"
	X666                                Variant = "666"
	X677                                Variant = "677"
	X711                                Variant = "711"
	X722                                Variant = "722"
	X733                                Variant = "733"
	X744                                Variant = "744"
	X755                                Variant = "755"
	X766                                Variant = "766"
	X777                                Variant = "777"
)

func (it Variant) String() string {
	return string(it)
}

// ExpandOctalByte
//
//	returns byte values at most 7 for each.
func (it Variant) ExpandOctalByte() (r7, w7, x7 byte) {
	return ExpandCharRwx(string(it))
}

func (it Variant) ToWrapper() (RwxWrapper, error) {
	return New.RwxWrapper.UsingVariant(it)
}

func (it Variant) ToWrapperPtr() (*RwxWrapper, error) {
	return New.RwxWrapper.UsingVariantPtr(it)
}
