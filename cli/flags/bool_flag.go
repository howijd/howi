// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import "github.com/mkungla/vars/v5"

// Bool returns new bool flag. Argument "a" can be any nr of aliases.
func Bool(name string, aliases ...string) (*BoolFlag, error) {
	c, err := newCommon(name, aliases...)
	if err != nil {
		return nil, err
	}
	f := &BoolFlag{*c}
	f.variable, _ = vars.NewTyped(name, "false", vars.TypeBool)
	return f, nil
}

func (b *BoolFlag) Value() bool {
	return b.variable.Bool()
}
