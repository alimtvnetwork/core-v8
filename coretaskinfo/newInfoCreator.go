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

package coretaskinfo

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
)

type newInfoCreator struct {
	Plain  newInfoPlainTextCreator
	Secure newInfoSecureTextCreator
}

func (it newInfoCreator) Deserialized(
	rawBytes []byte,
) (parsedInfo *Info, parsingErr error) {
	emptyInfo := Info{}
	parsingErr = corejson.Deserialize.UsingBytes(
		rawBytes, &emptyInfo)

	return emptyInfo.ToPtr(), parsingErr
}

func (it newInfoCreator) DeserializedUsingJsonResult(
	jsonResult *corejson.Result,
) (parsedInfo *Info, parsingErr error) {
	emptyInfo := Info{}
	parsingErr = jsonResult.Deserialize(
		&emptyInfo)

	return emptyInfo.ToPtr(), parsingErr
}

func (it newInfoCreator) Default(
	name, desc, url string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
	}
}

func (it newInfoCreator) Examples(
	name, desc, url string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		Examples:    examples,
	}
}

func (it newInfoCreator) Create(
	isSecure bool,
	name, desc, url,
	hintUrl, errorUrl,
	exampleUrl,
	chainingExample string,
	examples ...string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		Url:           url,
		HintUrl:       hintUrl,
		ErrorUrl:      errorUrl,
		ExampleUrl:    exampleUrl,
		SingleExample: chainingExample,
		Examples:      examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: isSecure,
		},
	}
}

func (it newInfoCreator) SecureCreate(
	name, desc, url,
	hintUrl, errorUrl,
	exampleUrl,
	chainingExample string,
	examples ...string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		Url:           url,
		HintUrl:       hintUrl,
		ErrorUrl:      errorUrl,
		ExampleUrl:    exampleUrl,
		SingleExample: chainingExample,
		Examples:      examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoCreator) PlainCreate(
	name, desc, url,
	hintUrl, errorUrl,
	exampleUrl,
	chainingExample string,
	examples ...string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		Url:           url,
		HintUrl:       hintUrl,
		ErrorUrl:      errorUrl,
		ExampleUrl:    exampleUrl,
		SingleExample: chainingExample,
		Examples:      examples,
	}
}
