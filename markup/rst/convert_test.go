// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rst

import (
	"testing"

	"github.com/gohugoio/hugo/common/hexec"
	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/config/security"

	"github.com/gohugoio/hugo/markup/converter"

	qt "github.com/frankban/quicktest"
)

func TestConvert(t *testing.T) {
	if !Supports() {
		t.Skip("rst not installed")
	}
	c := qt.New(t)
	sc := security.DefaultConfig
	sc.Exec.Allow = security.MustNewWhitelist("rst", "python")

	p, err := Provider.New(
		converter.ProviderConfig{
			Logger: loggers.NewDefault(),
			Exec:   hexec.New(sc),
		})
	c.Assert(err, qt.IsNil)
	conv, err := p.New(converter.DocumentContext{})
	c.Assert(err, qt.IsNil)
	b, err := conv.Convert(converter.RenderContext{Src: []byte("testContent")})
	c.Assert(err, qt.IsNil)
	c.Assert(string(b.Bytes()), qt.Equals, "<div class=\"document\">\n\n\n<p>testContent</p>\n</div>")
}

func TestConvertMathJax(t *testing.T) {
	if !Supports() {
		t.Skip("rst not installed")
	}
	c := qt.New(t)
	p, err := Provider.New(converter.ProviderConfig{Logger: loggers.NewErrorLogger()})
	c.Assert(err, qt.IsNil)
	conv, err := p.New(converter.DocumentContext{})
	c.Assert(err, qt.IsNil)
	b, err := conv.Convert(converter.RenderContext{Src: []byte(":math:`ax^2 + bx + c = 0`")})
	c.Assert(err, qt.IsNil)
	c.Assert(string(b.Bytes()), qt.Equals,
		"<div class=\"document\">\n\n\n<p><span class=\"math\">\\(ax^2 + bx + c = 0\\)</span></p>\n</div>")
}

func TestConvertCodeFormatting(t *testing.T) {
	if !Supports() {
		t.Skip("rst not installed")
	}
	c := qt.New(t)
	p, err := Provider.New(converter.ProviderConfig{Logger: loggers.NewErrorLogger()})
	c.Assert(err, qt.IsNil)
	conv, err := p.New(converter.DocumentContext{})
	c.Assert(err, qt.IsNil)
	b, err := conv.Convert(converter.RenderContext{Src: []byte(".. code:: c\n\n   int i = 0;")})
	c.Assert(err, qt.IsNil)
	c.Assert(string(b.Bytes()), qt.Equals,
		"<div class=\"document\">\n\n\n<pre class=\"code c literal-block\">\n<span class=\"kt\">int</span> <span class=\"n\">i</span> <span class=\"o\">=</span> <span class=\"mi\">0</span><span class=\"p\">;</span>\n</pre>\n</div>")
}
