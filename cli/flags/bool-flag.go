// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import (
	"testing"

	"github.com/mkungla/vars/v5"
)

// Parse the BoolFlag.
func (f *BoolFlag) Parse(args *[]string) (bool, error) {
	return f.parse(args, func(v vars.Value) (err error) {
		f.variable, err = vars.NewTyped(f.name, v.String(), vars.TypeBool)
		return err
	})
}

func TesBooltName(t *testing.T) {
	for _, tt := range testflags() {
		t.Run(tt.name, func(t *testing.T) {
			flag, err := NewBoolFlag(tt.name)
			if !tt.valid {
				if err == nil {
					t.Errorf("invalid flag %q expected error got <nil>", tt.name)
				}
				if flag != nil {
					t.Errorf("invalid flag %q should be <nil> got %#v", tt.name, flag)
				}
			}
		})
	}
}
