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

type Specification struct {
	BaseIdDisplayType
	BaseTags
	BaseIsGlobal
	flatSpec *FlatSpecification
}

func NewSpecificationSimple(
	id,
	display,
	typeName string,
) *Specification {
	return &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{Id: id},
			BaseDisplay:    BaseDisplay{display},
			BaseType:       BaseType{typeName},
		},
		BaseTags:     *NewTags(nil),
		BaseIsGlobal: BaseIsGlobal{false},
	}
}

func NewSpecificationSimpleGlobal(
	id,
	display,
	typeName string,
) *Specification {
	return &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{Id: id},
			BaseDisplay:    BaseDisplay{display},
			BaseType:       BaseType{typeName},
		},
		BaseTags:     *NewTags(nil),
		BaseIsGlobal: BaseIsGlobal{true},
	}
}

func NewSpecification(
	id,
	display,
	typeName string,
	tags []string,
	isGlobal bool,
) *Specification {
	return &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{Id: id},
			BaseDisplay:    BaseDisplay{display},
			BaseType:       BaseType{typeName},
		},
		BaseTags:     *NewTags(tags),
		BaseIsGlobal: BaseIsGlobal{isGlobal},
	}
}

func (r *Specification) Clone() *Specification {
	if r == nil {
		return nil
	}

	clonedTags := make([]string, len(r.Tags))
	copy(clonedTags, r.Tags)

	return &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{r.Id},
			BaseDisplay:    BaseDisplay{r.Display},
			BaseType:       BaseType{r.Type},
		},
		BaseTags: BaseTags{
			Tags: clonedTags,
		},
		BaseIsGlobal: BaseIsGlobal{r.IsGlobal},
	}
}

func (r *Specification) FlatSpecification() *FlatSpecification {
	if r == nil {
		return nil
	}

	if r.flatSpec != nil {
		return r.flatSpec
	}

	r.flatSpec = NewFlatSpecificationUsingSpec(
		r,
		true)

	return r.flatSpec
}
