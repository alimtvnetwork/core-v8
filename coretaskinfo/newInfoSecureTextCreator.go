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

type newInfoSecureTextCreator struct{}

func (it newInfoSecureTextCreator) Default(
	name, desc, url string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NameDescUrl(
	name, desc, url string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NameDescUrlExamples(
	name, desc, url string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		Examples:    examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NewNameDescUrlErrorUrl(
	name, desc,
	url, errUrl string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		ErrorUrl:    errUrl,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NameDescUrlErrUrlExamples(
	name, desc,
	url, errUrl string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		ErrorUrl:    errUrl,
		Examples:    examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NameDescExamples(
	name, desc string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Examples:    examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) Examples(
	name, desc string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Examples:    examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) ExamplesOnly(
	examples ...string,
) *Info {
	return &Info{
		Examples: examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) UrlOnly(
	url string,
) *Info {
	return &Info{
		Url: url,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) ErrorUrlOnly(
	errUrl string,
) *Info {
	return &Info{
		ErrorUrl: errUrl,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) HintUrlOnly(
	hintUrl string,
) *Info {
	return &Info{
		HintUrl: hintUrl,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) DescHintUrlOnly(
	desc, hintUrl string,
) *Info {
	return &Info{
		Description: desc,
		HintUrl:     hintUrl,
	}
}

func (it newInfoSecureTextCreator) NameHintUrlOnly(
	name, hintUrl string,
) *Info {
	return &Info{
		RootName: name,
		HintUrl:  hintUrl,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) SingleExampleOnly(
	singleExample string,
) *Info {
	return &Info{
		SingleExample: singleExample,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) AllUrlExamples(
	name, desc string,
	url, hintUrl, errUrl string,
	examples ...string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		HintUrl:     hintUrl,
		ErrorUrl:    errUrl,
		Examples:    examples,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) AllUrl(
	name, desc string,
	url, hintUrl, errUrl string,
) *Info {
	return &Info{
		RootName:    name,
		Description: desc,
		Url:         url,
		HintUrl:     hintUrl,
		ErrorUrl:    errUrl,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) UrlSingleExample(
	name, desc string,
	url string,
	chainingExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		Url:           url,
		SingleExample: chainingExample,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) SingleExample(
	name, desc string,
	singleExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		SingleExample: singleExample,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) ExampleUrl(
	name, desc string,
	exampleUrl string,
	singleExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		ExampleUrl:    exampleUrl,
		SingleExample: singleExample,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) ExampleUrlSingleExample(
	name, desc string,
	exampleUrl string,
	singleExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		ExampleUrl:    exampleUrl,
		SingleExample: singleExample,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}

func (it newInfoSecureTextCreator) NewExampleUrlSecure(
	name, desc string,
	exampleUrl string,
	chainingExample string,
) *Info {
	return &Info{
		RootName:      name,
		Description:   desc,
		ExampleUrl:    exampleUrl,
		SingleExample: chainingExample,
		Examples:      nil,
		ExcludeOptions: &ExcludingOptions{
			IsSecureText: true,
		},
	}
}
