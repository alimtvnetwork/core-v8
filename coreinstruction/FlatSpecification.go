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

import "github.com/alimtvnetwork/core/coredata/stringslice"

type FlatSpecification struct {
	Id       string   `json:"Id"`
	Display  string   `json:"Display"`
	Type     string   `json:"Type"`
	IsGlobal bool     `json:"IsGlobal"`
	Tags     []string `json:"Tags,omitempty"`
	IsValid  bool     `json:"IsValid,omitempty"`
	spec     *Specification
}

func InvalidFlatSpecification() *FlatSpecification {
	return &FlatSpecification{
		Id:       "",
		Display:  "",
		Type:     "",
		IsGlobal: false,
		Tags:     []string{},
		IsValid:  false,
	}
}

func NewFlatSpecificationUsingSpec(spec *Specification, isValid bool) *FlatSpecification {
	clonedTags := make([]string, len(spec.Tags))
	copy(clonedTags, spec.Tags)

	return &FlatSpecification{
		Id:       spec.Id,
		Display:  spec.Display,
		Type:     spec.Type,
		IsGlobal: spec.IsGlobal,
		Tags:     clonedTags,
		IsValid:  isValid,
		spec:     spec,
	}
}

func (it *FlatSpecification) BaseIdentifier() BaseIdentifier {
	return it.Spec().BaseIdentifier
}

func (it *FlatSpecification) BaseTags() BaseTags {
	return it.Spec().BaseTags
}

func (it *FlatSpecification) BaseIsGlobal() BaseIsGlobal {
	return it.Spec().BaseIsGlobal
}

func (it *FlatSpecification) BaseDisplay() BaseDisplay {
	return it.Spec().BaseDisplay
}

func (it *FlatSpecification) BaseType() BaseType {
	return it.Spec().BaseType
}

func (it *FlatSpecification) Spec() *Specification {
	if it == nil {
		return nil
	}

	if it.spec != nil {
		return it.spec
	}

	it.spec = &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{Id: it.Id},
			BaseDisplay:    BaseDisplay{it.Display},
			BaseType:       BaseType{it.Type},
		},
		BaseTags: BaseTags{
			Tags: it.Tags,
		},
		BaseIsGlobal: BaseIsGlobal{IsGlobal: it.IsGlobal},
	}

	return it.spec
}

func (it *FlatSpecification) Clone() *FlatSpecification {
	if it == nil {
		return nil
	}

	return &FlatSpecification{
		Id:       it.Id,
		Display:  it.Display,
		Type:     it.Type,
		IsGlobal: it.IsGlobal,
		Tags:     stringslice.Clone(it.Tags),
		IsValid:  it.IsValid,
	}
}
