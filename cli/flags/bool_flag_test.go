// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import (
	"fmt"
	"testing"

	"github.com/mkungla/vars/v5"
)

func TestBoolFlagPresent(t *testing.T) {
	flag, _ := NewBoolFlag("some-bool-flag")
	if present, err := flag.Parse([]string{"--some-bool-flag"}); !present || err != nil {
		t.Error("expected bool flag parser to return present, nil got ", present, err)
	}

	if !flag.Present() {
		t.Error("expected bool flag to be present")
	}

	if flag.Value().Bool() != true {
		t.Error("expected bool value to be true got ", flag.Value().Bool())
	}

	if flag.Value().String() != "true" {
		t.Error("expected bool value to be \"true\" got ", flag.Value().String())
	}

	flag.Unset()
	if flag.Present() {
		t.Error("expected bool flag not to be present")
	}
}

type booltest struct {
	name  string
	alias string
	arg   string
	str   string
	b     bool
}

func booltests() []booltest {
	return []booltest{
		{"some-true-flag", "t", "true", "true", true},
		{"some-true-flag-on", "t", "on", "true", true},
		{"some-true-flag-1", "t", "1", "true", true},
		{"some-false-flag", "f", "false", "false", false},
		{"some-false-flag-off", "f", "off", "false", false},
		{"some-false-flag-0", "f", "0", "false", false},
	}
}
func TestBoolFlagValues(t *testing.T) {
	for _, tt := range booltests() {
		t.Run(tt.name, func(t *testing.T) {
			flag, _ := NewBoolFlag(tt.name, tt.alias)
			args := fmt.Sprintf("--%s=%s", tt.name, tt.arg)
			if present, err := flag.Parse([]string{args}); !present || err != nil {
				t.Error("expected bool flag parser to return present, nil got ", present, err)
			}
			if !flag.Present() {
				t.Error("expected bool flag to be present")
			}
			if flag.Value().Bool() != tt.b {
				t.Errorf("expected bool value to be %t got %t", tt.b, flag.Value().Bool())
			}
			if flag.String() != tt.str {
				t.Errorf("expected bool value to be %q got %q", tt.str, flag.String())
			}
			if flag.Value().Type() != vars.TypeBool {
				t.Errorf("expected bool value Type to be TypeBool got %v", flag.Value().Type())
			}
			flag.Unset()
			if flag.Present() {
				t.Error("expected bool flag not to be present")
			}
		})
	}
}

func TestBoolFlagAliasValues(t *testing.T) {
	for _, tt := range booltests() {
		t.Run(tt.name, func(t *testing.T) {
			flag, _ := NewBoolFlag(tt.name, tt.alias)
			args := fmt.Sprintf("-%s=%s", tt.alias, tt.str)
			if present, err := flag.Parse([]string{args}); !present || err != nil {
				t.Error("expected bool flag parser to return present, nil got ", present, err)
			}
			if !flag.Present() {
				t.Error("expected bool flag to be present")
			}
			if flag.Value().Bool() != tt.b {
				t.Errorf("expected bool value to be %t got %t", tt.b, flag.Value().Bool())
			}
			if flag.String() != tt.str {
				t.Errorf("expected bool value to be %q got %q", tt.str, flag.String())
			}
		})
	}
}

func TestBoolFlagNotPresent(t *testing.T) {
	flag, _ := NewBoolFlag("some-flag")
	if ok, err := flag.Parse([]string{"--some-flag-2"}); ok {
		t.Error("expected bool flag parser to return not ok, ", ok, err)
	}

	if flag.Present() {
		t.Error("expected bool flag not to be present")
	}

	if flag.Value().Bool() != false {
		t.Error("expected bool value to be false got ", flag.Value().Bool())
	}

	if flag.Value().String() != "false" {
		t.Error("expected bool value to be \"false\" got ", flag.Value().String())
	}
}

func TestBooltName(t *testing.T) {
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
