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

package coreinstruction

type BaseSpecification struct {
	Specification *Specification `json:"Specification,omitempty"`
}

func NewBaseSpecification(
	id,
	display,
	typeName string,
	tags []string,
	isGlobal bool,
) *BaseSpecification {
	spec := NewSpecification(
		id, display,
		typeName,
		tags,
		isGlobal)

	return &BaseSpecification{
		Specification: spec,
	}
}

func (b *BaseSpecification) Identifier() BaseIdentifier {
	return b.Specification.BaseIdentifier
}

func (b *BaseSpecification) Display() BaseDisplay {
	return b.Specification.BaseDisplay
}

func (b *BaseSpecification) Type() BaseType {
	return b.Specification.BaseType
}

func (b *BaseSpecification) IsEmptySpec() bool {
	return !b.HasSpec()
}

func (b *BaseSpecification) HasSpec() bool {
	return b != nil && b.Specification != nil
}

func (b *BaseSpecification) Clone() *BaseSpecification {
	if b == nil {
		return nil
	}

	return &BaseSpecification{
		Specification: b.Specification.Clone(),
	}
}
