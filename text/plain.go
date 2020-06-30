// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text // import "github.com/elamre/plot/text"

import "github.com/elamre/plot/vg/draw"

// Plain is a text/plain handler.
type Plain = draw.PlainTextHandler

var _ draw.TextHandler = (*Plain)(nil)
