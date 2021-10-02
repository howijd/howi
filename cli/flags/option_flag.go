// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import (
	"fmt"
	"strings"

	"github.com/mkungla/vars/v5"
)

// Parse the OptionFlag.
func (f *OptionFlag) Parse(args *[]string) (ok bool, err error) {
	f.isPresent, err = f.parse(args, func(v vars.Value) (err error) {
		f.variable = vars.New(f.name, v.String())
		return err
	})

	if len(f.variable.String()) > 0 {
		opts := strings.Split(f.variable.String(), ",")
		if len(opts) > 0 {
			for _, o := range opts {
				if _, isSet := f.opts[o]; !isSet {
					return f.isPresent, fmt.Errorf("%w: (%s=%q)", ErrInvalidValue, f.name, f.variable.String())
				}
				f.opts[o] = true
			}
		}
	} else {
		return f.isPresent, ErrMissingOption
	}

	return f.isPresent, err
}
