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
	var opts []string
	f.isPresent, err = f.parseAll(args, func(v vars.Value) (err error) {
		opts = append(opts, v.String())
		return err
	})

	if err != nil {
		return f.isPresent, err
	}

	if len(opts) > 0 {
		for _, o := range opts {
			if _, isSet := f.opts[o]; !isSet {
				return f.isPresent, fmt.Errorf("%w: (%s=%q)", ErrInvalidValue, f.name, f.variable.String())
			}
			f.opts[o] = true
		}
		f.variable = vars.New(f.name, strings.Join(opts, ","))
	} else {
		return f.isPresent, ErrMissingOption
	}
	return f.isPresent, err
}

// Option return parsed options.
func (f *OptionFlag) Options() (present []string) {
	defvals := strings.Split(f.defval.String(), ",")
	for o, set := range f.opts {
		if set {
			present = append(present, o)
		} else if len(defvals) > 0 {
			for _, do := range defvals {
				if o == do {
					present = append(present, o)
				}
			}
		}
	}
	return present
}
