// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import (
	"github.com/mkungla/vars/v5"
)

// Parse the NumFlag.
func (f *NumFlag) Parse(args *[]string) (bool, error) {
	return f.parse(args, func(v vars.Value) (err error) {
		f.variable, err = vars.NewTyped(f.name, v.String(), vars.TypeFloat64)
		return err
	})
}
