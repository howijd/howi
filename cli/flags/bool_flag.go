// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

func (b *BoolFlag) Value() bool {
	return b.variable.Bool()
}
