// Copyright 2016 Charles Banning. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package mxj

import (
	"bytes"
)

var xmlEscapeChars bool

// XMLEscapeChars(true) forces escaping invalid characters in attribute and element values.
// NOTE: this is brute force with NO interrogation of '&' being escaped already; if it is
// then '&amp;' will be re-escaped as '&amp;amp;'.
//  
// This is only for m.Xml() and m.XmlIndent(); m.XmlSeq() and m.XmlSeqIndent() assume that
// values were originally decoded from valid XML.
/*
	The values are:
	"   &quot;
	'   &apos;
	<   &lt;
	>   &gt;
	&   &amp;
*/
func XMLEscapeChars(b bool) {
	xmlEscapeChars = b
}

// order is important - must scan for '&' first
var escapechars = [][2][]byte{
	{[]byte(`&`), []byte(`&amp;`)},
	{[]byte(`<`), []byte(`&lt;`)},
	{[]byte(`>`), []byte(`&gt;`)},
	{[]byte(`"`), []byte(`&quot;`)},
	{[]byte(`'`), []byte(`&apos;`)},
}

func escapeChars(s string) string {
	if len(s) == 0 {
		return s
	}

	b := []byte(s)
	for _, v := range escapechars {
		n := bytes.Count(b, v[0])
		if n == 0 {
			continue
		}
		b = bytes.Replace(b, v[0], v[1], n)
	}
	return string(b)
}
