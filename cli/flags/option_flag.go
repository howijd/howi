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
func (f *OptionFlag) Parse(args []string) (ok bool, err error) {
	var opts []vars.Variable

	if !f.defval.Empty() {
		defval := strings.Split(f.defval.String(), "|")
		for _, dd := range defval {
			opts = append(opts, vars.New(f.name+":default", dd))
		}
	}

	f.isPresent, err = f.parse(args, func(v []vars.Variable) (err error) {
		opts = v
		return err
	})

	if err != nil && f.defval.Empty() {
		return f.isPresent, err
	}

	if len(opts) > 0 {
		var str []string
		for _, o := range opts {
			if _, isSet := f.opts[o.String()]; !isSet {
				return f.isPresent, fmt.Errorf("%w: (%s=%q)", ErrInvalidValue, f.name, o)
			}
			f.opts[o.String()] = true
			str = append(str, o.String())
		}
		f.variable = vars.New(f.name, strings.Join(str, "|"))
	} else {
		return f.isPresent, ErrMissingValue
	}
	return f.isPresent, err
}

// Option return parsed options.
func (f *OptionFlag) Options() (present []string) {
	for o, set := range f.opts {
		if set {
			present = append(present, o)
		}
	}
	return present
}

// Default sets flag default.
func (f *OptionFlag) Default(def ...interface{}) vars.Variable {
	if len(def) > 0 && def[0] != nil && f.defval.Empty() {
		var defopts = def[0].([]string)
		f.defval = vars.New(f.name+":default", strings.Join(defopts, "|"))
	}
	return f.defval
}
