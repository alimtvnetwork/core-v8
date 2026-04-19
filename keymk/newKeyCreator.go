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

type newKeyCreator struct{}

func (it *newKeyCreator) Create(
	option *Option,
	main string,
) *Key {
	key := &Key{
		option:   option,
		mainName: main,
		keyChains: make(
			[]string,
			0,
			constants.Capacity7),
	}

	return key
}

func (it *newKeyCreator) PathTemplate(
	root string,
	starterKeyChains ...any,
) *Key {
	return it.All(
		CurlyBracePathJoinerOption,
		root,
		starterKeyChains...)
}

func (it *newKeyCreator) PathTemplateDefault(
	starterKeyChains ...any,
) *Key {
	return it.All(
		CurlyBracePathJoinerOption,
		constants.PathRootTemplate,
		starterKeyChains...)
}

func (it *newKeyCreator) PathTemplatePrefixRelativeIdDefault() *Key {
	return it.All(
		CurlyBracePathJoinerOption,
		root,
		prefix,
		relative,
		id)
}

func (it *newKeyCreator) PathTemplatePrefixRelativeIdFileDefault() *Key {
	return it.All(
		CurlyBracePathJoinerOption,
		root,
		prefix,
		relative,
		id,
		constants.FileKeyword)
}

func (it *newKeyCreator) All(
	option *Option,
	main string,
	starterKeyChains ...any,
) *Key {
	slice := make([]string, 0, len(starterKeyChains)+DefaultCap)

	key := &Key{
		option:    option,
		mainName:  main,
		keyChains: slice,
	}

	if len(starterKeyChains) > 0 {
		key.keyChains = appendAnyItemsWithBaseStrings(
			option.IsSkipEmptyEntry,
			key.keyChains,
			starterKeyChains)
	}

	return key
}

func (it *newKeyCreator) AllStrings(
	option *Option,
	main string,
	starterKeyChains ...string,
) *Key {
	slice := make([]string, 0, len(starterKeyChains)+DefaultCap)

	key := &Key{
		option:    option,
		mainName:  main,
		keyChains: slice,
	}

	if len(starterKeyChains) > 0 {
		key.AppendChainStrings(starterKeyChains...)
	}

	return key
}

func (it *newKeyCreator) StringsWithOptions(
	option *Option,
	main string,
	starterKeyChains ...string,
) *Key {
	slice := make([]string, 0, len(starterKeyChains)+DefaultCap)

	key := &Key{
		option:    option,
		mainName:  main,
		keyChains: slice,
	}

	if len(starterKeyChains) > 0 {
		key.AppendChainStrings(starterKeyChains...)
	}

	return key
}

func (it *newKeyCreator) Parenthesis(
	main string,
	starterKeyChains ...any,
) *Key {
	return it.All(
		ParenthesisJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) ParenthesisStrings(
	main string,
	starterKeyChains ...string,
) *Key {
	return it.AllStrings(
		ParenthesisJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) Curly(
	main string,
	starterKeyChains ...any,
) *Key {
	return it.All(
		CurlyBraceJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) CurlyStrings(
	main string,
	starterKeyChains ...string,
) *Key {
	return it.AllStrings(
		CurlyBraceJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) SquareBrackets(
	main string,
	starterKeyChains ...any,
) *Key {
	return it.All(
		BracketJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) SquareBracketsStrings(
	main string,
	starterKeyChains ...string,
) *Key {
	return it.AllStrings(
		BracketJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) Default(
	main string,
	starterKeyChains ...any,
) *Key {
	return it.All(
		JoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) DefaultStrings(
	main string,
	starterKeyChains ...string,
) *Key {
	return it.Create(
		JoinerOption,
		main).
		AppendChainStrings(starterKeyChains...)
}

func (it *newKeyCreator) DefaultMain(
	main string,
) *Key {
	return it.Create(
		JoinerOption,
		main)
}

func (it *newKeyCreator) OptionMain(
	option *Option,
	main string,
) *Key {
	return it.Create(
		option,
		main)
}
