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

import "github.com/alimtvnetwork/core/constants"

type newKeyWithLegendCreator struct{}

// All
//
// Chain Sequence (Root-Package-Group-State-User-Item)
func (it *newKeyWithLegendCreator) All(
	option *Option,
	legendName LegendName,
	isAttachLegendNames bool,
	rootName,
	packageName,
	group,
	stateName string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          legendName,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		stateName:           stateName,
		isAttachLegendNames: isAttachLegendNames,
	}

	return keyWithLegend
}

// Create
//
// Chain Sequence (Root-Package-Group-State-User-Item)
func (it *newKeyWithLegendCreator) Create(
	option *Option,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          FullLegends,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: true,
	}

	return keyWithLegend
}

// NoLegend
//
// Chain Sequence (Root-Package-Group-State-User-Item)
func (it *newKeyWithLegendCreator) NoLegend(
	option *Option,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: false,
	}

	return keyWithLegend
}

// NoLegendPackage
//
// Chain Sequence (Root-Group-State-User-Item)
func (it *newKeyWithLegendCreator) NoLegendPackage(
	isAttachLegend bool,
	option *Option,
	rootName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		rootName:            rootName,
		packageName:         constants.EmptyString,
		groupName:           group,
		isAttachLegendNames: isAttachLegend,
	}

	return keyWithLegend
}

func (it *newKeyWithLegendCreator) ShortLegend(
	option *Option,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          ShortLegends,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: true,
	}

	return keyWithLegend
}
