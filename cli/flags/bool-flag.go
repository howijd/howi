// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import (
	"github.com/mkungla/vars/v5"
)

// Parse the BoolFlag.
func (f *BoolFlag) Parse(args *[]string) (bool, error) {
	return f.parse(args, func(v vars.Value) (err error) {
		if v.Type() == vars.TypeBool {
			f.variable = vars.New(f.name, v)
			return nil
		}
		f.variable, err = vars.NewTyped(f.name, v.String(), vars.TypeBool)
		return err
	})
}
